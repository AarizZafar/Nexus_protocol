package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

// Create an item
func createItem() {
    item := Item{
        Name:  "Smartphone",
        Price: 800,
    }
    jsonData, _ := json.Marshal(item)
    resp, err := http.Post("http://13.90.73.228:8080/items", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

// Get items
func getItems() {
    resp, err := http.Get("http://13.90.73.228:8080/items")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

// Update an item
func updateItem(id string) {
    item := Item{
        Name:  "Updated Smartphone",
        Price: 1000,
    }
    jsonData, _ := json.Marshal(item)
    req, _ := http.NewRequest(http.MethodPut, "http://13.90.73.228:8080/items/"+id, bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

// Delete an item
func deleteItem(id string) {
    req, _ := http.NewRequest(http.MethodDelete, "http://13.90.73.228:8080/items/"+id, nil)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

func main() {
    // Replace <VM-IP> with your Azure VM's IP address
    // Replace <ITEM_ID> with an actual item ID after testing the getItems() call

    createItem()       // Test Create
    getItems()         // Test Read

    // Replace with a valid ID from getItems() response
    updateItem("item-id")   // Test Update
    deleteItem("item-id")   // Test Delete
}
