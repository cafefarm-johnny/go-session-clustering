package user

import (
	"Go-Session-Clustering/src/domain"
	"Go-Session-Clustering/src/domain/user/db"
	"Go-Session-Clustering/src/domain/user/model"
)

func saveUser(username string, password []byte) error {
	u := model.CreateUser(username, password)
	if db.Exist(u) {
		return domain.ErrDuplicatedUser
	}

	db.Save(u)
	return nil
}
