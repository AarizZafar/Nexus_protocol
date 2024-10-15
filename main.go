package main

import (
	"fmt"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	"github.com/AarizZafar/Nexus_protocol.git/router"
	"github.com/AarizZafar/Nexus_protocol.git/verification"
)

func main() {
	port := ":8080"
	router := router.Router()
	
	fmt.Println(">>>>>>>>>>>>>>>>>> Starting server <<<<<<<<<<<<<<<<<<")
	fmt.Printf(">>>>>>>>>>>>>>>>>> Listening at port %s <<<<<<<<<<\n", port)

	fmt.Println("\n*****************************************************")
	fmt.Println("-------NETWORK BIOMETRICS EXTRACTION PROCESS----------")
	bioMetrics.ExtractNetBioMetrix()

	fmt.Println("-------SYSTEM BIOMETRICS EXTRACTION PROCESS----------")
	bioMetrics.ExtractSysBioMetrix()

	fmt.Println("---------VERIFICATION PROCESS HAS STARTED------------")
	fmt.Println("Verification step has started")
	verification.StartVerification()

	router.Run(port)
}
