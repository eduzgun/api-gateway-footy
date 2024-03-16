package routes

import (
	"github.com/eduzgun/api-gateway-footy/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/api/fixtures/statistics/fixture/:fixtureId", controllers.Fixture)
}
