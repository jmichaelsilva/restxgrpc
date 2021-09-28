package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Read the entire file into a byte slice
	imgBytes, err := ioutil.ReadFile("./img/IMG_2416.png")
	if err != nil {
		log.Printf("Load image error")
		log.Fatal(err)
	}
	
	values := map[string]string{"Data": base64.StdEncoding.EncodeToString(imgBytes) }
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatalf("json error: %v", err)
    }    

	resp, err := http.Post("http://localhost:3000/image", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
        log.Fatal(err)
    }

	var res string
    json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
}
