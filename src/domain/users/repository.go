package users

import (
	"Go-Session-Clustering/src/domain"
	"Go-Session-Clustering/src/domain/db"
	"Go-Session-Clustering/src/domain/users/model"
)

type userRepository struct {
	userDB *db.UserStore
}

func NewUserRepository() *userRepository {
	return &userRepository{
		userDB: db.GetUserStore(),
	}
}

func (ur *userRepository) Save(username string, password []byte) error {
	u := model.CreateUser(username, password)
	if ur.userDB.Exist(u) {
		return domain.ErrDuplicatedUser
	}

	ur.userDB.Save(u)
	return nil
}

func (ur *userRepository) Find(username string) *model.User {
	u := model.CreateUser(username, nil)
	dbu, has := ur.userDB.Find(u)
	if !has {
		return nil
	}

	return dbu
}
