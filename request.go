package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	var compactJSON bytes.Buffer
	if err := json.Compact(&compactJSON, body); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("workfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(compactJSON.String() + "\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", compactJSON.String())
}
