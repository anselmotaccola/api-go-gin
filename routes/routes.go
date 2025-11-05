package routes

import (
	"github.com/anselmotaccola/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)


func HundleRequest() {
	// Lógica para lidar com requisições
	r := gin.Default()
	r.GET("/", controllers.Exibeindex)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	
	r.Run(":5000")
}