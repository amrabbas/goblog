package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthResponse struct {
	TokenString string `json:"tokenString"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// Decode credentials
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &AuthResponse{
		TokenString: tokenString,
	}

	json.NewEncoder(w).Encode(response)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Extract authorization token
	authorizationHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authorizationHeader, "Bearer ")
	tokenString := splitToken[1]

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
