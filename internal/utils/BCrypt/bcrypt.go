package BCrypt

import (
	"caipiaotong/configs/constant"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), constant.EncodeCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Check 判断password加密后是否相同
func Check(password string, hashedPassword string) (bool, error) {
	encode, err := Encode(password)
	if err != nil {
		return false, err
	}
	return encode == hashedPassword, nil
}
