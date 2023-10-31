package botdetector

import "github.com/felihenrique/go-botdetector/internal/detector"

var detectorSingleton = detector.New()

func IsBotIp(ip string) (bool, error) {
	return detectorSingleton.IsBotIp(ip)
}

func IsBotAgent(agent string) bool {
	return detectorSingleton.IsBotAgent(agent)
}
