package tools

import httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"

const ORACLE_IPS_URL = "https://docs.oracle.com/iaas/tools/public_ip_ranges.json"

type oracleCidr struct {
	Cidr string `json:"cidr"`
}

type oracleRegion struct {
	Cidrs []oracleCidr `json:"cidrs"`
}

type oracleData struct {
	Regions []oracleRegion `json:"regions"`
}

func FetchOracleIps() ([]string, error) {
	data, err := httputils.GetJSON[oracleData](ORACLE_IPS_URL)

	if err != nil {
		return nil, err
	}

	var ranges = make([]string, 0)
	for _, re := range data.Regions {
		for _, ci := range re.Cidrs {
			ranges = append(ranges, ci.Cidr)
		}
	}

	return ranges, err
}