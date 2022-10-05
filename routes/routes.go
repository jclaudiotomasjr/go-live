package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/scores/", controllers.AllScores)
	//r.GET(":name", controllers.Welcome)
	r.POST("/scores/", controllers.CreatePoints)
	r.GET("/scores/:id", controllers.ReturnScore)
	r.DELETE("/score/:id", controllers.DeleteScore)
	//r.PATCH("/scores/:id", controllers.EditScore)
	r.GET("/score/cpf/:cpf", controllers.ShowCpf)
	r.Run()

}
