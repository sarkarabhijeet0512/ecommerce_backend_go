package productdetails

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	Repo Repository
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository) *Service {
	return &Service{
		conf: conf,
		log:  log,
		Repo: Repo,
	}
}

func (s *Service) UpsertProductDetails(dCtx context.Context, req Product) error {
	return s.Repo.upsertProductDetails(dCtx, req)
}

func (s *Service) GetProductByID(dCtx context.Context, req Product) error {
	return s.Repo.getProductByID(dCtx, req)
}

func (s *Service) GetProductListByCategory(dCtx context.Context, req Product) error {
	return s.Repo.getProductListByCategory(dCtx, req)
}

func (s *Service) GetProductList(dCtx context.Context, req Product) error {
	return s.Repo.getProductList(dCtx, req)
}

func (s *Service) DisableProductByID(dCtx context.Context, req Product) error {
	return s.Repo.disableProductByID(dCtx, req)
}
