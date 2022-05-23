package routes

import (
	"github.com/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/ben-vindo/:nome", controllers.BemVindo)
	r.POST("/alunos", controllers.NovoAluno)
	r.Run(":8081")
}
