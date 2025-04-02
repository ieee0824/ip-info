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
	withSSL := flag.Bool("ssl", false, "use ssl")
	cert := flag.String("cert", "", "ssl cert")
	key := flag.String("key", "", "ssl key")
	flag.Parse()

	r := gin.Default()

	r.GET("/", root.Index)

	if *withSSL {
		if err := r.RunTLS(fmt.Sprintf("%s:%s", *host, *port), *cert, *key); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := r.Run(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
			log.Fatal(err)
		}
	}

}
