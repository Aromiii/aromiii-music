package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRoot(c *gin.Context) {
	sessionTokenBytes, ok := c.MustGet("session").([]byte)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Session data not found",
		})
		return
	}

	var session Session
	if err := json.Unmarshal(sessionTokenBytes, &session); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to deserialize session data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": session,
	})
}
