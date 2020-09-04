package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/kakty/taskmgr/models"
	"gorm.io/gorm"
)

//SignIn :
func SignIn(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var jwtKey = []byte(os.Getenv("JWTSECRETKEY"))
	fmt.Println(jwtKey)
	var creds models.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.Table("users").Where("username = ?", creds.Username).First(&user)

	expectedPassword := user.Password

	if result.Error != nil || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("token error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
