package productdetails

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"ecommerce_backend_project/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf     *viper.Viper
	log      *logrus.Logger
	Repo     Repository
	s3Config *AWSS3Config
}

type AWSS3Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository) *Service {
	s3Config := AWSS3Config{
		AccessKeyID:     conf.GetString(utils.AccessKeyEnv),
		SecretAccessKey: conf.GetString(utils.SecretAccessKey),
		Region:          conf.GetString(utils.Region),
		Bucket:          conf.GetString(utils.BucketName),
	}
	return &Service{
		conf:     conf,
		log:      log,
		Repo:     Repo,
		s3Config: &s3Config,
	}
}

func (s *Service) UpsertProductDetails(dCtx context.Context, req *Product) error {
	err := s.Repo.upsertProductDetails(dCtx, req)
	if err != nil {
		return err
	}
	req.Dimensions.ProductID = req.ID
	err = s.Repo.upsertDimentions(dCtx, req.Dimensions)
	if err != nil {
		return err
	}
	for i := range req.Variants {
		req.Variants[i].ProductID = req.ID
		req.Variants[i].CreatedAt = time.Now()
		req.Variants[i].UpdatedAt = time.Now()
	}
	err = s.Repo.upsertProductVariants(dCtx, req.Variants)
	return err
}

func (s *Service) GetProductByID(dCtx context.Context, productID int) (*Product, error) {
	return s.Repo.getProductByID(dCtx, productID)
}

func (s *Service) GetProductListByCategory(dCtx context.Context, categoryID int) ([]Product, error) {
	return s.Repo.getProductListByCategory(dCtx, categoryID)
}

func (s *Service) GetProductCategoryList(dCtx context.Context) ([]Category, error) {
	return s.Repo.getProductCategoryList(dCtx)
}

func (s *Service) DisableProductByID(dCtx context.Context, req Product) error {
	return s.Repo.disableProductByID(dCtx, req)
}

func (s *Service) UploadProductImage(r *ProductImage, file multipart.File, fileName, contentType string, sess *session.Session, dCtx context.Context) error {
	var err error
	defer func() {
	}()

	//upload to the s3 bucket
	uploader := s3manager.NewUploader(sess)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.s3Config.Bucket),
		ACL:         aws.String("public-read"),
		Key:         aws.String(fileName),
		Body:        file,
		ContentType: &contentType,
	})
	if err != nil {
		s.log.Error("Failed to upload file to S3: " + err.Error())
		s.log.WithContext(dCtx).Info(up.UploadID)
		err = errors.New("failed to upload file")
		return err
	}
	filepath := "https://" + s.s3Config.Bucket + "." + "s3-" + s.s3Config.Region + ".amazonaws.com/" + fileName
	r.ImageURL = filepath
	return s.Repo.productImageDetails(dCtx, r)
}
