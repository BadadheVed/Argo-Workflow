package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/routes")
	setupRoutes(api)
	return r
}
func setupRoutes(r *gin.RouterGroup) {
	r.GET("/hello", Hello)
	r.GET("/hello/:name", HelloUser)

}
