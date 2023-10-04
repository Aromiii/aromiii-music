package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var Url = "http://localhost:3000"

type Session struct {
	User struct {
		Email       string `json:"email"`
		Image       any    `json:"image"`
		ID          string `json:"id"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
	Expires time.Time `json:"expires"`
}

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
		c.AbortWithStatusJSON(500, gin.H{"error": "Error creating request to authentication service"})
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Unable to read session data",
		})
		return
	}

	if string(body) == "{}" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Authentication service didn't return user data",
		})
		return
	}

	println(string(body))

	c.Set("session", body)

	c.Next()
}
