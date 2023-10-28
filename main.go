package botdetector

import "github.com/felihenrique/go-botdetector/internal/detector"

var detectorInstance = detector.New()

func IsBotIp(ip string) (bool, error) {
	return detectorInstance.IsBotIp(ip)
}

func IsBotAgent(agent string) bool {
	return detectorInstance.IsBotAgent(agent)
}
