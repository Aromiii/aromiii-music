package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func HandleRoot(c *gin.Context) {
	/*sessionTokenBytes, ok := c.MustGet("session").([]byte)
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
	}*/

	// Set the response headers to indicate audio streaming.
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Transfer-Encoding", "chunked")

	// Replace "your-audio-file.mp3" with the path to your MP3 file.
	filePath := "./music/rickroll.mp3"

	// Open the MP3 file.
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	// Create a buffer for reading chunks of the file.
	chunkSize := 1024
	buffer := make([]byte, chunkSize)

	// Read and send chunks of the file to the client.
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			c.Status(http.StatusInternalServerError)
			return
		}

		_, err = c.Writer.Write(buffer[:n])
		if err != nil {
			break
		}
		c.Writer.Flush()
	}

	c.Status(http.StatusOK)
}
