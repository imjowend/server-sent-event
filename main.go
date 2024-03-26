package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imjowend/server-sent-event/http"
)

func main() {
	ch := make(chan string)

	router := gin.Default()
	router.POST("/event-stream", func(c *gin.Context) {
		http.HandleEventStreamPost(c, ch)
	})
	router.GET("/event-stream", func(c *gin.Context) {
		http.HandleEventStreamGet(c, ch)
	})

	log.Fatalf("error running HTTP server: %s\n", router.Run(":9990"))
}
