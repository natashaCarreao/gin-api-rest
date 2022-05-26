package routes

import (
	"github.com/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/ben-vindo/:nome", controllers.BemVindo)
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.POST("/alunos", controllers.NovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.Run(":8081")
}
