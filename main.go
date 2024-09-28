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
	/* when the browser loads an html file it reads the <link > tag that points to a CSS file
	   The browser will send a GET request to the server for the css file.
	   If the server has been set up to serve static files
	   router.Static("webpages/css", "./webpages/css")
	   the CSS file is served, and the page is styled accordingly.
	   the <link> automatically request the css file from the local file path  
	*/

	router.Static("webPages/css", "./webPages/css")             // Serve the static files from here, arg1 - url we specify, 
	                                                            //                                   arg2 - local system path from where the files get fetched
    router.Static("webPages/JS", "./webPages/JS/")   
																

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login_page.html", nil) /* in the nil parameter we can pass some default parameters like usename, passwd */
	})




	router.Run(port)
}