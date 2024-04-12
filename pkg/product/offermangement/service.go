package offermangement

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

func (s *Service) UpsertOfferDetails(ctx context.Context, discount *Discount) error {
	return s.Repo.upsertOfferDetails(ctx, discount)
}

func (s *Service) FetchOfferByFilter(ctx context.Context, filter model.Filter) ([]Discount, error) {
	return s.Repo.fetchOfferByFilter(ctx, filter)
}
