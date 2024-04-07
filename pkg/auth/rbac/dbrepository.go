package rbac

import (
	"context"

	"ecommerce_backend_project/utils"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertRole(context.Context, *Role) error
	upsertRolePermission(context.Context, *RolePermission) error
	upsertPermission(context.Context, *Permission) error
	upsertAssignRole(dCtx context.Context, r *UserRole) error
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"userdb"`
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

func (r *PGRepo) upsertRole(dCtx context.Context, role *Role) error {
	utils.SetGenericFieldValue(role)
	_, err := r.db.ModelContext(dCtx, role).OnConflict("(role_name) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) upsertRolePermission(dCtx context.Context, rolePermissions *RolePermission) error {
	utils.SetGenericFieldValue(rolePermissions)
	_, err := r.db.ModelContext(dCtx, rolePermissions).OnConflict("(permission_id,resourse_id) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) upsertPermission(dCtx context.Context, permission *Permission) error {
	utils.SetGenericFieldValue(permission)
	_, err := r.db.ModelContext(dCtx, permission).OnConflict("(read,write,edit,remove) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) upsertAssignRole(dCtx context.Context, userRole *UserRole) error {
	utils.SetGenericFieldValue(userRole)
	_, err := r.db.ModelContext(dCtx, userRole).OnConflict("(user_id) DO UPDATE").Insert()
	return err
}
