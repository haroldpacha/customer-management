package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/haroldpacha/customer-management/helper"
	"github.com/haroldpacha/customer-management/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("failed to process request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(strings.Replace(authHeader, "Bearer ", "", 1))
		if !token.Valid {
			log.Println(err)
			response := helper.BuildErrorResponse("token invalidd", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		log.Println("claim[user_id]: ", claims["user_id"])
		log.Println("claim[issuer]: ", claims["issuer"])
	}
}
