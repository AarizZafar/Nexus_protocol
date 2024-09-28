package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/AarizZafar/Nexus_protocol.git/router"
)

func main() {
	port := ":8000"
	router := router.Router()

	fmt.Println(">>>>>>>>>>>>>>>>>> Starting server <<<<<<<<<<<<<<<<<<")
	fmt.Printf(">>>>>>>>>>>>>>>>>> Listening at port %s <<<<<<<<<<\n", port)

	router.LoadHTMLGlob("webPages/html/*")                      // Load all the html files from here
	router.Static("webPages/css", "./webPages/css")             // Serve the static files from here, arg1 - url we specify, 
	                                                            //                                   arg2 - local system path from where the files get fetched
	
    router.Static("webPages/JS", "./webPages/JS/")   
																

	router.GET("/", func(ctx *gin.Context) {
											// in the nil parameter we can pass some default parameters like usename, passwd 
		ctx.HTML(http.StatusOK, "login_page.html", nil)
	})

	router.Run(port)
}