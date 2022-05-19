package controllers

import (
	"github.com/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func BemVindo(c *gin.Context) {

	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Bem vindo: ": nome,
	})

}
