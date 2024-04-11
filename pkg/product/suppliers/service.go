package suppliers

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

func (s *Service) UpsertSuppliers(ctx context.Context, suppliers *Supplier) error {
	return s.Repo.upsertSuppliers(ctx, suppliers)
}

func (s *Service) FetchSuppliersByFilter(ctx context.Context, Filter model.Filter) (suppliers []Supplier, err error) {
	return s.Repo.fetchSuppliers(ctx, Filter)
}
