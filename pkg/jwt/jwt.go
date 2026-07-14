package jwt

import (
	"time"
	// "fmt"
	golangjwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, email string, role string) (string, error) {

	expiry := time.Now().Add(24 * time.Hour)

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,

		RegisteredClaims: golangjwt.RegisteredClaims{
			ExpiresAt: golangjwt.NewNumericDate(expiry),
			IssuedAt:  golangjwt.NewNumericDate(time.Now()),
		},
	}

	// fmt.Println(claims);
	token := golangjwt.NewWithClaims(
		golangjwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte("your-super-secret-key"))
}

func ValidateToken(tokenString string) (*Claims, error) {

	token, err := golangjwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *golangjwt.Token) (interface{}, error) {
			return []byte("your-super-secret-key"), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
