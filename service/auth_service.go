package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/taco-tortilla/jwt-go/initializers"
	"github.com/taco-tortilla/jwt-go/models"
	"github.com/taco-tortilla/jwt-go/request"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(body request.SingUpBody) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func Login(body request.SingUpBody) (string, error) {
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokenString, nil
}
