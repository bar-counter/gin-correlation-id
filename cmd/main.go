//go:build !test

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/bar-counter/gin-correlation-id/ginid_uuidv4"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const cliVersion = "0.1.2"

var serverPort = flag.String("serverPort", "49002", "http service address")

func main() {
	log.Printf("-> env:ENV_WEB_AUTO_HOST %s", os.Getenv("ENV_WEB_AUTO_HOST"))
	flag.Parse()
	log.Printf("-> run serverPort %v", *serverPort)
	log.Printf("=> now version %v", cliVersion)

	// Create the Gin engine.
	g := gin.New()

	middlewareList := []gin.HandlerFunc{
		ginid_uuidv4.CorrelationIDUUidV4Middleware(),
	}
	// Routes.
	Load(
		// Cores.
		g,

		// middlewareList.
		middlewareList...,
	)
	log.Printf("=> now server access as: http://127.0.0.1:%s", *serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *serverPort), g)
	if err != nil {
		fmt.Printf("run gin server err %v\n", err)
		return
	}
}

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found.")
	})

	g.Use(Options)
	g.Use(mw...)

	g.GET("/ping", ping)
	log.Printf("-> ping api as: http://127.0.0.1:%s/ping", *serverPort)

	return g
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong as correlation ID %s", ginid_uuidv4.GetCorrelationIDUUidV4(c))
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}
