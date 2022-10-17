package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Id  string `json:"id"`
	Txt string `json:"txt"`
}

var entries = []Entry{
	{Id: "1", Txt: "Hellow"},
}

func data(c *gin.Context) {
	var entry Entry

	if c.Request.Method == "POST" {
		if err := c.BindJSON(&entry); err != nil {
			return
		}
		entries = append(entries, entry)
		c.IndentedJSON(http.StatusCreated, entry)
	}

	if c.Request.Method == "GET" {
		c.IndentedJSON(http.StatusOK, entries)
	}
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.GET("/", ping)
	r.GET("data/", data)
	r.POST("data/", data)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
