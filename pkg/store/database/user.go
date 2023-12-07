package database

import (
	"context"
	"time"

	"git-devops.opencsg.com/product/community/starhub-server/pkg/model"
	"github.com/uptrace/bun"
)

type UserStore struct {
	db *model.DB
}

func NewUserStore(db *model.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

type User struct {
	ID           int64         `bun:",pk,autoincrement" json:"id"`
	GitID        int64         `bun:",pk" json:"git_id"`
	Name         string        `bun:",notnull" json:"name"`
	Username     string        `bun:",notnull,unique" json:"username"`
	Email        string        `bun:",notnull,unique" json:"email"`
	Password     string        `bun:",notnull" json:"-"`
	AccessTokens []AccessToken `bun:"rel:has-many,join:id=user_id"`
	Namespaces   []Namespace   `bun:"rel:has-many,join:id=user_id" json:"namespace"`
	times
}

func (s *UserStore) FindByUsername(ctx context.Context, username string) (user User, err error) {
	user.Username = username
	err = s.db.Operator.Core.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	return
}

func (s *UserStore) FindByID(ctx context.Context, id int) (user User, err error) {
	user.ID = int64(id)
	err = s.db.Operator.Core.NewSelect().Model(&user).WherePK().Scan(ctx)
	return
}

func (s *UserStore) UpdateByUsername(ctx context.Context, u *User) (err error) {
	u.UpdatedAt = time.Now()
	err = assertAffectedOneRow(s.db.Operator.Core.NewUpdate().
		Model(u).
		Where("username = ?", u.Username).
		Exec(ctx),
	)
	return
}

func (s *UserStore) Create(ctx context.Context, user *User, namespace *Namespace) (err error) {
	err = s.db.Operator.Core.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err = assertAffectedOneRow(tx.NewInsert().Model(user).Exec(ctx)); err != nil {
			return err
		}
		namespace.UserID = user.ID
		namespace.NamespaceType = UserNamespace
		if err = assertAffectedOneRow(tx.NewInsert().Model(namespace).Exec(ctx)); err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *UserStore) IsExist(ctx context.Context, username string) (exists bool, err error) {
	var user User
	exists, err = s.db.Operator.Core.
		NewSelect().
		Model(&user).
		Where("username =?", username).
		Exists(ctx)
	if err != nil {
		return
	}
	return
}
