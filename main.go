package main

import (
	"github.com/anselmotaccola/api-go-gin/database"

	"github.com/anselmotaccola/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Ana Clara", Id: "1", RG: "12.345.678-9", CPF: "123.456.789-00"},
	// 	{Nome: "Bruno Souza", Id: "2", RG: "98.765.432-1", CPF: "987.654.321-00"},
	// 	{Nome: "Carla Mendes", Id: "3", RG: "11.223.344-5", CPF: "111.222.333-44"},
	// }
	routes.HundleRequest()
}