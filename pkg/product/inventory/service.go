package inventory

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

// IsDBActive gets user data by her userID
func (s *Service) UpsertInventory(ctx context.Context, inventory *Inventory) error {
	return s.Repo.upsertInventory(ctx, inventory)
}
func (s *Service) FetchInventoryByFilter(ctx context.Context, filter model.Filter) ([]Inventory, error) {
	return s.Repo.fetchInventoryByFilter(ctx, filter)
}
