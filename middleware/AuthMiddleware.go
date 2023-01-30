package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/response"
	"mini-douyin/utils/jwt"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.ErrorMessage{
				Response: response.Response{
					StatusMsg:  "unauthorized",
					StatusCode: 1,
				},
			})
			return
		}

		_, e := c.Get("userId")
		if e {
			c.Next()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			log.Printf("AuthMiddleware|token解析错误|%v", err)
			c.JSON(http.StatusUnauthorized, response.ErrorMessage{
				Response: response.Response{
					StatusMsg:  "unauthorized",
					StatusCode: 1,
				},
			})
			return
		}
		userID := claims.UserID
		c.Set("userId", userID)
		c.Next()
		return
	}
}
