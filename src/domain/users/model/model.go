package model

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/go-uuid"
)

type User struct {
	UUID     string
	Username string
	Password []byte
}

func CreateUser(username string, password []byte) User {
	return User{
		UUID:     generateUUID(),
		Username: username,
		Password: password,
	}
}

func generateUUID() string {
	v, err := uuid.GenerateUUID()
	if err != nil {
		return fmt.Sprintf("custom%08d", rand.Int())
	}
	return v
}
