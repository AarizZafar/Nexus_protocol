package Vmcontrols

import (
	"fmt"
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
)

const baseURL = "http://13.90.73.228:8080/verify"

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendBioMetx() {
	// the device data that we have assigned to the struct in our local code
	jsonData, err := json.Marshal(bioMetrics.BMstruct)
	handleErr(err)

	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonData))
	handleErr(err)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response : ", string(body))
}
