package main

import (
	"github.com/gin-api-rest/models"
	"github.com/gin-api-rest/routes"
)

func main() {
	models.Aluno = []models.Aluno{
		{Nome: "Aluno de teste", CPF: "123.123.123-55", RG: "53.03030-0"},
	}
	routes.HandleRequests()
}
