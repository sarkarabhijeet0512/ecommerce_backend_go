package dummy

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	IsActive() (ok bool, err error)
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

//IsActive checks if DB is connected
func (r *PGRepo) IsActive() (ok bool, err error) {

	ctx := context.Background()
	err = r.db.Ping(ctx)
	if err == nil {
		ok = true
	}
	return
}
