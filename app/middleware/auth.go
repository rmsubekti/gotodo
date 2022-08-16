package middlewares

import (
	"gotodo/app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var user any

		autH := c.GetHeader("Authorization")
		token := strings.SplitN(autH, " ", 2)
		if (len(token) != 2) && (token[0] != "Bearer") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request header incorrect format"})
		}
		if user, err = utils.ParseAndVerify(token[1]); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
