package logic

import (
	"os"
	"time"

	"github.com/Sahilb315/trello_clone/models"
	JWT "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	return err == nil
}

func CreateToken(user models.User) (string, error) {
	token := JWT.New(JWT.SigningMethodHS256)
	claims := token.Claims.(JWT.MapClaims)
	claims["id"] = user.ID
	claims["fullName"] = user.FullName
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}
