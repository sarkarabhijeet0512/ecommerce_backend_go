package payment

import (
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
func (s *Service) IsActive() (bool, error) {
	return s.Repo.IsActive()
}
