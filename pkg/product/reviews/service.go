package reviews

import (
	"context"
	model "ecommerce_backend_project/utils/models"

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

func (s *Service) UpdateReviewByProductID(ctx context.Context, review *Review) error {
	return s.Repo.updateReviewByProductID(ctx, review)
}
func (s *Service) FetchReviewByFilter(ctx context.Context, filter *model.Filter) ([]Review, error) {
	return s.Repo.fetchReviewByFilter(ctx, filter)
}
