package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/database"
	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/models"
)

func AllScores(c *gin.Context) {
	var scores []models.Score
	database.DB.Find(&scores)
	c.JSON(http.StatusOK, scores)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"APP": "Welcome " + name + " Tudo Tranquilo?",
	})
}

func CreatePoints(c *gin.Context) {
	var s models.Score
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(404, gin.H{
			"erro": err.Error()})
	}
	database.DB.Create(&s)
	c.JSON(http.StatusOK, gin.H{
		"status": "Score lançado com sucesso!"})
}

func ReturnScore(c *gin.Context) {
	var score models.Score
	id := c.Params.ByName("id")
	database.DB.First(&score, id)
	if score.ID == 0 {
		c.JSON(http.StatusFound, gin.H{
			"Status": "Score não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, score)
}

func DeleteScore(c *gin.Context) {
	var score models.Score
	id := c.Params.ByName("id")
	database.DB.First(&score, id)
	if score.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Score não foi deletado, pois não foi encontrado!"})
		return
	}
	database.DB.Delete(&score, id)
	c.JSON(http.StatusOK, score)

}

/*func EditScore(c *gin.Context) {
	var score models.Score
	id := c.Params.ByName("id")
	database.DB.First(&score, id)
	if err := c.ShouldBindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&score).UpdateColumns(score)
	c.JSON(http.StatusOK, gin.H{
		"status": "Score editado com sucesso!"})
	c.JSON(http.StatusOK, score)
}*/

func ShowCpf(c *gin.Context) {
	var score models.Score
	cpf := c.Param("name")
	database.DB.Where(&models.Score{Cpf: cpf}).First(&score)
	if score.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Usuário não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, score)

}
