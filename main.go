package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Fadhelbulloh/api-with-redis/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if e := gotenv.Load(env); e != nil {
		log.Println(e)
	}
}
func main() {
	errs := make(chan error)

	go func() {
		// go routine for signalling if app terminated
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {
		// setting gin mode
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.Use(cors.Default())

		// default response if page not found
		router.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found or wrong path"})
		})

		// accsessing api
		route.BasicService(router)

		// running on port
		if e := router.Run(":7200"); e != nil {
			errs <- e
		}
	}()

	log.Fatal(<-errs)
}
