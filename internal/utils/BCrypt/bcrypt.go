package BCrypt

import (
	"caipiaotong/configs/constant"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), constant.EncodeCost)
	fmt.Println("原密码:", password, "加密后:", string(hashedPassword))
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Check 判断password加密后是否相同,不相等返回err
func Check(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
