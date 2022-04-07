package middlewares

import (
	"net/http"
	"strings"

	e "github.com/akkien/learngo-explorer-api/err"
	"github.com/akkien/learngo-explorer-api/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		authHeaderArr := c.Request.Header["Authorization"]
		if len(authHeaderArr) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid token",
				"data":    "",
			})
			c.Abort()
			return
		}

		splittedAuthHeader := strings.Split(authHeaderArr[0], " ")
		if len(splittedAuthHeader) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid token",
				"data":    "",
			})
			c.Abort()
			return
		}

		token := splittedAuthHeader[1]
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data":    data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
