package utils

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
	"log"
	"slices"
)

/*
GetOpenApiTags 引数で受け取ったタグがopenapi上に含まれているかをチェックするutil関数
*/
func GetOpenApiTags(swagger *openapi3.T, c *gin.Context, tags []string) bool {

	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(err)
	}

	route, _, err := router.FindRoute(c.Request)

	if err != nil {
		log.Fatal(err)
		return false
	}

	currentTags := route.PathItem.GetOperation(c.Request.Method).Tags
	for _, tag := range tags {
		if slices.Contains(currentTags, tag) {
			return true
		}
	}
	return false
}
