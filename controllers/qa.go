package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"shamilabs.com/qa-api/config"
	"shamilabs.com/qa-api/models"
)

func CreateQuestion(c *gin.Context) {

	//TODO add service layer and add validation
	var question models.Question
	err := c.BindJSON(&question)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = models.CreateAQuestion(&question)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, question)
	}
}

func CreateAnswer(c *gin.Context) {

	//TODO add service layer and add validation
	var answer models.Answer
	err := c.BindJSON(&answer)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	questionId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	answer.QuestionID = uint(questionId)
	err = models.CreateAnswer(&answer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, answer)
	}
}

func UpdateQuestionWithTag(c *gin.Context) {
	questionId, err := strconv.Atoi(c.Param("id"))
	var question models.Question
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// get Question
	if err = config.GetDB().First(&question, questionId).Error; err != nil {
		log.Error().Msg("could not fetch question with")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var tag models.Tag
	err = c.BindJSON(&tag)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	models.TagQuestion(&question, &tag)
}
func GetQuestion(c *gin.Context) {

}

func DeleteQuestion(c *gin.Context) {

}

func IsHealthy(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func IsReady(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
