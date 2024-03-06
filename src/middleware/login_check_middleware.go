package middleware

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/utils"
	"gorm.io/gorm"
	"net/http"
)

func LoginCheckMiddleware(db *gorm.DB, swagger *openapi3.T) gin.HandlerFunc {
	return func(c *gin.Context) {
		tags := []string{"user", "admin", "developer"}
		isTarget := utils.GetOpenApiTags(swagger, c, tags)
		if !isTarget {
			c.Next()
			return
		}
		// TODO: セッション情報として格納する情報を検討する
		session := sessions.Default(c)
		id := session.Get("id")

		if id == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			c.Set("session", id)
			c.Next()
		}
	}

}
