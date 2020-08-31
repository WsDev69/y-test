package persistence

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"y-test/internal/constant/model"
	"y-test/internal/constant/query"
	"y-test/internal/module/user"
	"y-test/platform/sqlite"
)

type UserPersistence struct {
	db *sqlite.Sqlite
}

func NewUserPersistence(db *sqlite.Sqlite) user.Persistence {
	return &UserPersistence{db: db}
}

func (up *UserPersistence) Create(ctx context.Context, user *model.User) (*model.User, error) {
	log := logrus.WithContext(ctx).WithField("user", user)
	tx := up.db.DB.MustBegin()
	_, err := tx.NamedExec(query.InsertNewUser, user)
	if err != nil {
		log.Error("couldn't create user :", err)
		return nil, err
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, err
	}

	return user, nil
}

func (up *UserPersistence) FindByID(ctx context.Context, userId string) (*model.User, error) {
	log := logrus.WithContext(ctx).WithField("userId", userId)
	u := model.User{}
	err := up.db.DB.Get(&u, query.SelectUserById, userId)
	if err != nil {
		log.Error("couldn't find user :", err)
		return nil, err
	}

	return &u, nil
}

func (up *UserPersistence) Find(ctx context.Context, email, password string) (*model.User, error) {
	log := logrus.WithContext(ctx).WithField("email", email)
	u := model.User{}
	err := up.db.DB.Get(&u, query.SelectUserByPassAndEmail, email, password)
	if err != nil {
		log.Error("couldn't find user by email and password :", err)
		return nil, err
	}
	return &u, nil
}

func (up *UserPersistence) Update(ctx context.Context, user *model.User) (*model.User, error) {
	if user.UserId == "" {
		return nil, fmt.Errorf("user id is empty")
	}

	log := logrus.WithContext(ctx).WithField("user", user)
	tx := up.db.DB.MustBegin()

	var s string
	var values = make(map[string]interface{})
	values["userId"] = user.UserId

	if user.FirstName != "" {
		s += "first_name=:firstName,"
		values["firstName"] = user.FirstName
	}
	if user.LastName != "" {
		s += "last_name=:lastName"
		values["lastName"] = user.LastName
	}
	if user.AvatarLink != "" {
		s += "avatar_l=:avatarLink"
		values["avatarLink"] = user.AvatarLink
	}

	q := fmt.Sprintf(query.UpdateUser, s)
	res, err := tx.NamedExec(q, values)
	if err != nil {
		log.Error("couldn't update user :", err)
		return nil, err
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, err
	}

	if rows, errRows := res.RowsAffected(); rows < 1 || errRows != nil {
		log.WithField("user", user).Debug("couldn't update, not found")
		return nil, fmt.Errorf("not found")
	}

	return user, nil
}

func (up UserPersistence) FindPageable(ctx context.Context, pageSize, pageNumber int) ([]*model.User, error) {
	log := logrus.WithContext(ctx)
	tx := up.db.DB.MustBegin()
	var u []*model.User
	err := tx.Select(&u, query.SelectPageable, pageSize, pageNumber-1)
	if err != nil {
		log.Error("couldn't find user by pageSize and pageNumber : ", err)
		return nil, err
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, err
	}
	return u, nil
}
