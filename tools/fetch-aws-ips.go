package tools

import httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"

const AWS_IPS_URL = "https://ip-ranges.amazonaws.com/ip-ranges.json"

type awsPrefix struct {
	IpPrefix string `json:"ip_prefix"`
}

type awsData struct {
	Prefixes []awsPrefix `json:"prefixes"`
}

func FetchAWSIps() ([]string, error) {
	data, err := httputils.GetJSON[awsData](AWS_IPS_URL)

	if err != nil {
		return nil, err
	}

	var ranges []string = make([]string, 0, len(data.Prefixes))

	for _, p := range data.Prefixes {
		ranges = append(ranges, p.IpPrefix)
	}

	return ranges, nil
}
