package detector

import (
	"log"
	"strings"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
	"github.com/felihenrique/go-botdetector/pkg/iptree"
)

const JSON_URL = "https://raw.githubusercontent.com/felihenrique/go-botdetector/master/assets/ips.json"

type detector struct {
	tree *iptree.IpTree
}

func New() detector {
	d := detector{}
	d.tree = iptree.New(0)

	data, err := httputils.GetJSON[[]string](JSON_URL)

	if err != nil {
		log.Printf("Error loading Ips JSON: %s", err.Error())
	}

	for _, v := range *data {
		d.AddRange(v)
	}

	return d
}

func (d *detector) AddRange(ipRange string) (bool, error) {
	return d.tree.Add(ipRange)
}

func (d *detector) IsBotIp(ip string) (bool, error) {
	return d.tree.Has(ip)
}

func (d *detector) IsBotAgent(agent string) bool {
	ag := strings.ToLower(agent)
	for _, crawler := range botAgents {
		if strings.Contains(ag, crawler) {
			return true
		}
	}
	return false
}
