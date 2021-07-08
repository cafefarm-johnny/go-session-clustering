package db

import "Go-Session-Clustering/src/domain/user/model"

var users map[string]model.User

func init() {
	users = map[string]model.User{}
}

func Save(u model.User) {
	users[u.Username] = u
}

func Exist(u model.User) bool {
	_, has := users[u.Username]
	return has
}
