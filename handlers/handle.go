package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnParameterMissingError(c *gin.Context, parameter string) {
	var err = fmt.Sprintf("Required parameter %s missing.", parameter)
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}
