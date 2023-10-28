package tools

import (
	"log"
	"strings"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
)

const LINODE_IPS_URL = "https://geoip.linode.com/"

func FetchLinodeIps() ([]string, error) {
	log.Println("Fetching linode data")

	data, err := httputils.GetData(LINODE_IPS_URL)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	var ranges = make([]string, 0, len(lines))
	for i, v := range lines {
		if i < 3 {
			continue
		}

		values := strings.Split(v, ",")

		ranges = append(ranges, values[0])
	}

	log.Printf("Fetched %d IPS for linode", len(ranges))

	return ranges, err
}
