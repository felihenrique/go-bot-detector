package tools

import (
	"log"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
)

const AZURE_IPS_URL = "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20231023.json"

type azureProperty struct {
	AddressPrefixes []string `json:"addressPrefixes"`
}

type azureValue struct {
	Properties azureProperty `json:"properties"`
}

type azureData struct {
	Values []azureValue `json:"values"`
}

func FetchAzureIps() ([]string, error) {
	log.Println("Fetching AWS data")

	data, err := httputils.GetJSON[azureData](AZURE_IPS_URL)

	if err != nil {
		return nil, err
	}

	ranges := make([]string, 0)
	for _, dv := range data.Values {
		ranges = append(ranges, dv.Properties.AddressPrefixes...)
	}

	return ranges, nil
}
