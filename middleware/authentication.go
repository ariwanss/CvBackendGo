package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ariwanss/CvBackendGo/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthPayload struct {
	UserID    primitive.ObjectID
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (p AuthPayload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("expired token")
	}
	return nil
}

func generateToken(userId primitive.ObjectID) (string, error) {
	p := AuthPayload{
		UserID:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24 * 30),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func verifyToken(tokenStr string) (*AuthPayload, error) {
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	}

	var p AuthPayload

	_, err := jwt.ParseWithClaims(tokenStr, &p, keyfunc)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func Authorize(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("No authorization"))
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer") {
		c.AbortWithError(http.StatusBadRequest, errors.New("No bearer token"))
		return
	}

	tokenStr := strings.Split(authHeader, " ")[1]

	p, err := verifyToken(tokenStr)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// c.Set("authPayload", p)
	c.Set("userId", p.UserID)
	c.Next()
}

func AttachToken(c *gin.Context) {
	user := c.Value("user").(*entity.User)
	tokenStr, err := generateToken(user.ID)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenStr,
	})
}
