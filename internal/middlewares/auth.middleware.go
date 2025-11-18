package middlewares

import (
	"go-ecommerce-backend-api/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c * gin.Context){
		token := c.GetHeader("Authorization")
		if token != "valid-token"{
			response.ErrorReponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}