package main

import (
	"ip-info/handlers/root"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", root.Index)

	if err := r.Run("lan-ip.ast.moe:80"); err != nil {
		log.Fatal(err)
	}
}
