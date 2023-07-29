package helpers

import (
	"os"
	"strings"
	"time"

	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")
var tokenError = errs.NewUnauthenticatedError("Invalid token")

func claimToken(userID int, name string, email string, role string) jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": userID,
		"name":    name,
		"email":   email,
		"role":    role,
		"exp":     3600 * time.Second,
	}
}

func signToken(signingMethod jwt.SigningMethod, token jwt.MapClaims) (string, errs.ErrMessage) {
	claim := jwt.NewWithClaims(signingMethod, token)

	signedToken, err := claim.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", tokenError
	}

	return signedToken, nil
}

func GenerateToken(userID int, name string, email string, role string) (string, errs.ErrMessage) {
	claim := claimToken(userID, name, email, role)

	signedToken, err := signToken(jwt.SigningMethodHS256, claim)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetUserData(bearerToken string) (*entity.User, errs.ErrMessage) {
	//validate token
	mapClaims, err := validateToken(bearerToken)

	if err != nil {
		return nil, err
	}

	user, err := bindTokenToUserEntity(*mapClaims)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func bindTokenToUserEntity(mapClaims jwt.MapClaims) (*entity.User, errs.ErrMessage) {
	var user entity.User

	if id, ok := mapClaims["user_id"].(float64); !ok {
		return nil, tokenError
	} else {
		user.ID = uint(id)
	}

	if name, ok := mapClaims["name"].(string); !ok {
		return nil, tokenError
	} else {
		user.Name = name
	}

	if email, ok := mapClaims["email"].(string); !ok {
		return nil, tokenError
	} else {
		user.Email = email
	}

	if role, ok := mapClaims["role"].(string); !ok {
		return nil, tokenError
	} else {
		user.Role = role
	}

	return &user, nil
}

func validateToken(bearerToken string) (*jwt.MapClaims, errs.ErrMessage) {
	if !strings.HasPrefix(bearerToken, "Bearer") || len(strings.Split(bearerToken, " ")) != 2 {
		return nil, tokenError
	}

	tokenized := strings.Split(bearerToken, " ")

	stringToken := tokenized[1]

	token, err := parseToken(stringToken)

	if err != nil {
		return nil, err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, tokenError
	} else {
		mapClaims = claims
	}

	return &mapClaims, nil
}

func parseToken(tokenString string) (*jwt.Token, errs.ErrMessage) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, tokenError
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, tokenError
	}

	return token, nil
}
