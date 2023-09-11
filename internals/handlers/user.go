package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func SignupHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
