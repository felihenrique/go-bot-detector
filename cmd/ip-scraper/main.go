package main

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/felihenrique/go-botdetector/internal/tools"
)

func main() {
	data, errs := tools.FetchAllIps()

	if len(errs) > 0 {
		log.Fatalln(errs)
	}

	slices.Sort(data)

	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("assets/ips.json")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(jsonData)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("JSON data written to 'assets/ips.json'")
}
