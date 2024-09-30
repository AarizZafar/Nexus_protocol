package router

import (
	"github.com/gin-gonic/gin"
	"github.com/AarizZafar/Nexus_protocol.git/controls"
)


func Router() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("webPages/html/*")                      // Load all the html files from here

    // Serve the static files from here, 
	// arg1 - url we specify, 
	// arg2 - local system path from where the files get fetched
	router.Static("webPages/css", "./webPages/css")             
    router.Static("webPages/js", "./webPages/js/")   

	// GET used to display -> web page, when we load a login page the browser will send a get request to server
	router.GET("/",          controllers.LoginPage)             
	// Post used to send data from (browser) to like netName, netPass          
	router.POST("/login",    controllers.NetAuthentication)      

	return router
}

/*
+-------------------+          +-------------------------+
|    User Browser   |          |       Web Server        |
+-------------------+          +-------------------------+
        |                                 |
        | 1. GET Request "/"              |    <-- User opens login page.
        |-------------------------------> |
        |                                 |
        | 2. HTML Form (login_page.html)  |    <-- Server sends login page.
        |<------------------------------- |
        |                                 |
User enters username/password, clicks "Login" (Form data sent via POST)
        |                                 |
        | 3. POST Request "/login"        |    <-- User submits login data.
        |-------------------------------> |
        |                                 |
        | 4. Validate Credentials         |    <-- Server checks login info.
        |                                 |
        | 5. Success or Error Response    |    <-- Server responds with result.
        |<------------------------------- |

*/
