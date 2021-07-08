package hash

import "golang.org/x/crypto/bcrypt"

func ToBcrypt(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}
