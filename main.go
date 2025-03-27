package main

import (
	"flag"
	"fmt"
	"ip-info/handlers/root"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("h", "", "host")
	port := flag.String("p", "8080", "port")
	flag.Parse()

	r := gin.Default()

	r.GET("/", root.Index)

	if err := r.Run(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
		log.Fatal(err)
	}
}
