package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
		grp1.GET("mesa", Controllers.GetMesas)
		grp1.POST("mesa", Controllers.CreateMesa)
		grp1.GET("mesa/:id", Controllers.GetMesaByID)
		grp1.PUT("mesa/:id", Controllers.UpdateMesa)
		grp1.DELETE("mesa/:id", Controllers.DeleteMesa)
	}
	return r
}
