package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"meilian/constants"
	"time"
)

func GenerateToken(data map[string]interface{}, expiration time.Duration) (string, error) {
	secretKeyBytes := []byte(constants.JWTSECREETKEY)
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(expiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用HS256算法进行签名
	return token.SignedString(secretKeyBytes)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, nil
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTSECREETKEY), nil // 使用常量的别名
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func GetTokenData(tokenString string) (map[string]interface{}, error) {

	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	for key, value := range claims {
		data[key] = value
	}

	return data, nil
}
