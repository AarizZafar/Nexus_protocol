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

	fmt.Println("\033[97;46m>>>>>>>>>>>>>>>>>> STARTING SERVER <<<<<<<<<<<<<<<<<<\033[0m")
    fmt.Printf("\033[97;46m>>>>>>>>>>>>>>>>>> LISTENING AT PORT %s <<<<<<<<<<\033[0m\n\n", port)


	fmt.Println("\033[97;46m       NETWORK BIOMETRICS EXTRACTION PROCESS         \033[0m")
	bioMetrics.ExtractNetBioMetrix()
	fmt.Println("\033[97;42m           NET BIOMETRICS EXTRACTION DONE ✔          \033[0m\n")
	
	fmt.Println("\033[97;46m       SYSTEM BIOMETRICS EXTRACTION PROCESS          \033[0m")
	bioMetrics.ExtractSysBioMetrix()
	fmt.Println("\033[97;42m     SYSTEM BIOMETRICS EXTRACTION PROCESS DONE ✔     \033[0m\n")

	fmt.Println("\033[97;46m           VERIFICATION PROCESS HAS STARTED          \033[0m")
	verification.StartVerification()

	router.Run(port)
}
