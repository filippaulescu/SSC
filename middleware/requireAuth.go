package middleware

import (
	"ic-project/initializers"
	"ic-project/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// 1. Obține token-ul din cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Autentificare necesară",
		})
		return
	}

	// 2. Parsează token-ul
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid",
		})
		return
	}

	// 3. Procesează claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token format",
		})
		return
	}

	// 4. Verifică expirarea
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token expirat",
		})
		return
	}

	// 5. Găsește utilizatorul
	var user models.User
	if err := initializers.DB.First(&user, claims["sub"]).Error; err != nil || user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Utilizatorul nu există",
		})
		return
	}

	// 6. Adaugă user-ul în context
	c.Set("user", user)
	c.Next()
}