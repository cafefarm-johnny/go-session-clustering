package user

import (
	"Go-Session-Clustering/src/domain"
	"Go-Session-Clustering/src/domain/user/model"
	"Go-Session-Clustering/src/hash"
)

type UserService struct {
	ur *userRepository
}

func NewUserService() *UserService {
	return &UserService{
		ur: NewUserRepository(),
	}
}

func (us *UserService) Signup(dto *model.UserDTO) error {
	hashPwd, err := hash.ToBcrypt(dto.Password)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return us.ur.Save(dto.Username, hashPwd)
}

func (us *UserService) SelfAuthenticate(dto *model.UserDTO) error {
	u := us.ur.Find(dto.Username)
	if u == nil {
		return domain.ErrNotFoundUser
	}

	if !hash.CompareBcrypt(u.Password, []byte(dto.Password)) {
		return domain.ErrInvalidPassword
	}

	return nil
}
