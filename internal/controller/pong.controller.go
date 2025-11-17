package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController{
	return &PongController{}
}


//uc user controller
//us user service
func (uc *PongController) GetPongById(c *gin.Context) {
	name := c.DefaultQuery("name", "Giau")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H {
		"message": " pong..ping" + name,
		"uid": uid,
		"users": []string{"cr7", "m10", "Giau"},
	})
}