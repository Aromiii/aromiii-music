package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRoot(c *gin.Context) {
	sessionTokenBytes, ok := c.MustGet("session").([]byte)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Session data not found",
		})
		return
	}

	var session Session
	if err := json.Unmarshal(sessionTokenBytes, &session); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to deserialize session data",
		})
		return
	}

	c.File("./music/rickroll.mp3")
}
