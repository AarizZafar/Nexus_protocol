package main

import (
	"fmt"
	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	
	"github.com/AarizZafar/Nexus_protocol.git/router"
)

func main() {
	port := ":8000"
	router := router.Router()
	
	fmt.Println(">>>>>>>>>>>>>>>>>> Starting server <<<<<<<<<<<<<<<<<<")
	fmt.Printf(">>>>>>>>>>>>>>>>>> Listening at port %s <<<<<<<<<<\n", port)

	fmt.Println("\n*****************************************************")
	fmt.Println("-------SYSTEM BIOMETRICS EXTRACTION PROCESS----------")
	bioMetrics.ExtSysBioMetrix()

	router.Run(port)
}
