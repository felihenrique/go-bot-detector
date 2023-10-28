package tools

import (
	"log"
	"strings"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
)

const DIGITALOCEAN_IPS_URL = "https://digitalocean.com/geo/google.csv"

func FetchDigitalOceanIps() ([]string, error) {
	log.Println("Fetching Digital ocean data")

	data, err := httputils.GetData(DIGITALOCEAN_IPS_URL)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var ranges = make([]string, 0, len(lines))

	for _, v := range lines {
		values := strings.Split(v, ",")
		ranges = append(ranges, values[0])
	}

	return ranges, nil
}
