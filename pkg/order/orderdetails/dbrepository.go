package orderdetails

import (
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"deliveryRiderDB"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {
	Repo = &PGRepo{
		log: i.Log,
		db:  i.DB,
	}
	return
}
