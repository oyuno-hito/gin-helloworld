oapi-codegen -generate "types" -package openapi schema/gin-helloworld.yml > ./src/openapi/types.gen.go
oapi-codegen -generate "gin" -package openapi schema/gin-helloworld.yml > ./src/openapi/gin.gen.go
oapi-codegen -generate "spec" -package openapi schema/gin-helloworld.yml > ./src/openapi/spec.gen.go