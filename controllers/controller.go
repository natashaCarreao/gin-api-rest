package controllers

import (
	"net/http"

	"github.com/gin-api-rest/database"
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

func NovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}
