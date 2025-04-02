package main

import (
	"flag"
	"fmt"
	"ip-info/handlers/root"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("host", "", "host")
	port := flag.String("port", "8080", "port")
	withSSL := flag.Bool("ssl", false, "use ssl")
	cert := flag.String("cert", "", "ssl cert")
	key := flag.String("key", "", "ssl key")

	withPid := flag.Bool("pid", false, "write pid to file")
	pidFile := flag.String("pid-file", "/run/ip-info.pid", "pid file")

	if *withPid {
		pid := os.Getpid()
		if err := os.WriteFile(*pidFile, []byte(fmt.Sprintf("%d", pid)), 0644); err != nil {
			log.Fatalf("failed to write pid file: %v", err)
		}
		defer func() {
			if err := os.Remove(*pidFile); err != nil {
				log.Printf("failed to remove pid file: %v", err)
			}
		}()
	}

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
