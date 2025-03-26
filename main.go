package main

import (
	"ip-info/handlers/root"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", root.Index)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
