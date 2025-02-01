package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type HealthCheck struct {
    Status string `json:"status"`
}

func main() {
	res, err := http.Get("https://simpledebit.gocardless.io/health_check")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	var health HealthCheck
	if err := json.Unmarshal(body, &health); err != nil {
		log.Fatal(err)
	}
	fmt.Println(health.Status) // This will print "ok"

	if res.StatusCode == 200 {
		fmt.Println("SUCCESS")
	} else if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
