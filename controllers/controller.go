package controllers

import (
	"net/http"

	"github.com/anselmotaccola/api-go-gin/database"
	"github.com/anselmotaccola/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	// Lógica para exibir todos os alunos
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Exibeindex(c *gin.Context) {
	// Lógica para exibir todos os alunos
	c.JSON(200, gin.H{
		"Menssagem": "Estamos e um api Go GIN!!!",
		
	})
}

func Saudacao(c *gin.Context) {
	// Lógica para exibir todos os alunos
	nome:= c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API Diz": "Seja bem vindo " + nome,
	})
}


func CriaNovoAluno(c *gin.Context) {
	var novoAluno models.Aluno

	
	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&novoAluno)
	c.JSON(http.StatusOK, novoAluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	// Lógica para buscar aluno por ID
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)

}

func DeletaAluno(c *gin.Context) {
	// Lógica para deletar aluno por ID
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Deleted": "Aluno deletado com sucesso",
	})
}

func AtualizaAluno(c *gin.Context) {
	// Lógica para atualizar aluno por ID
	id := c.Params.ByName("id")
	var aluno models.Aluno


	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&aluno).Where("id = ?", id).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	// Lógica para buscar aluno por CPF
	cpf := c.Param("cpf")
	var aluno models.Aluno

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0  {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
	
}