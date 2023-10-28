package tools

import (
	"log"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
)

const GOOGLE_IPS_URL = "https://www.gstatic.com/ipranges/cloud.json"

type googleData struct {
	Prefixes []googlePrefix `json:"prefixes"`
}

type googlePrefix struct {
	IpV4Prefix string `json:"ipv4Prefix"`
}

func FetchGoogleIps() ([]string, error) {
	log.Println("Fetching google cloud data")

	data, err := httputils.GetJSON[googleData](GOOGLE_IPS_URL)

	if err != nil {
		return nil, err
	}

	var ranges = make([]string, 0, len(data.Prefixes))

	for _, p := range data.Prefixes {
		ranges = append(ranges, p.IpV4Prefix)
	}

	log.Printf("Fetched %d IPS for google cloud", len(ranges))

	return ranges, nil
}
