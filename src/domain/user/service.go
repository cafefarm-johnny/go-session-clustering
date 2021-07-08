package user

import "Go-Session-Clustering/src/domain/user/model"

func signup(dto *model.UserDTO) error {
	return saveUser(dto.Username, dto.Password)
}
