package middleware

import (
	"learn-go/internal/constant"
	"learn-go/internal/utils"
	pkg "learn-go/pkg/panic"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer pkg.PanicHandler(c)
		tokenString := c.GetHeader("Authorization")
		log.Debug("tokenString: ", tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Token is required"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		claims, err := utils.GetClaimsFromToken(tokenString)
		if err != nil {
			pkg.PanicException(constant.Unauthorized, "Unauthorized")
		}
		// You can access claims.UserID or other claims in your handlers
		c.Set("userId", claims.Id)
		c.Next()
	}
}
