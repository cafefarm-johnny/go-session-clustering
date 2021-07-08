package user

import (
	"Go-Session-Clustering/src/domain"
	"Go-Session-Clustering/src/domain/user/model"
	"Go-Session-Clustering/src/hash"
)

func signup(dto *model.UserDTO) error {
	hashPwd, err := hash.ToBcrypt(dto.Password)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return saveUser(dto.Username, hashPwd)
}
