package middleware

import (
	"net/http"

	"Training/Assignment1/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifikasi token (misalnya, cocokkan dengan token yang diharapkan)
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Butuh Authorization basic coy!!"})
			c.Abort()
			return
		}

		//log.Printf("Received username: %s, password: %s", username, password)

		isValid := (username == config.AuthBasicUsername) && (password == config.AuthBasicPassword)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization tidak ada atau salah ya"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}
