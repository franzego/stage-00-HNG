package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func CatApi() (string, error) {

	var ApiFact struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	catUrl := os.Getenv("catFactUrl")
	// catUrl := "https.//catfact."
	resp, err := client.Get(catUrl)
	if err != nil {
		return "Cannot get CatFact at the moment", fmt.Errorf("[CatApi] is unreachable")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//log.Printf("Invalid status code for [CatApi]: %v", err)
		return "Invalid Status Code", fmt.Errorf("[CatAPI] returned status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&ApiFact); err != nil {
		log.Printf("[CatApi] JSON decode error: %v", err)
		return "Error in getting to CatFact Url", err
	}
	return ApiFact.Fact, err

	// newMsg := string(msg)
	// return newMsg, nil
}
