package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/cookiejar"
)

var Url = "http://localhost:3000"

func Authenticate(c *gin.Context) {
	// Create a cookie jar to manage cookies
	cookieJar, _ := cookiejar.New(nil)

	// Create an HTTP client with the cookie jar
	client := &http.Client{
		Jar: cookieJar,
	}

	// Create a GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/auth/session", Url), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retrieve cookies from the Gin context and add them to the request
	cookies := c.Request.Cookies()
	for _, cookie := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	// Send the GET request
	resp, err := client.Do(req)

	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "Authentication service is unavailable"})
		return
	}
	defer resp.Body.Close()

	// Read the response body into a []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Unable to read session data",
		})
		return
	}

	if string(body) == "{}" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Authentication service didn't return user data",
		})
		return
	}

	// Now you can work with the 'body' []byte as needed
	fmt.Println("Response Body:", string(body))

	c.Next()
}
