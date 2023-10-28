package tools

import (
	"sync"
)

type FetchFunc func() ([]string, error)

func fetchData(ff FetchFunc, rc chan []string, ec chan error, wg *sync.WaitGroup) {
	data, err := ff()
	defer wg.Done()

	if err != nil {
		ec <- err
		return
	}

	rc <- data
}

func FetchAllIps() ([]string, []error) {
	rgChan := make(chan []string)
	errChan := make(chan error)
	wg := &sync.WaitGroup{}

	funcs := []FetchFunc{
		FetchAWSIps, FetchCloudflareIps, FetchDigitalOceanIps, FetchGoogleIps, FetchLinodeIps,
		FetchAzureIps,
	}
	wg.Add(len(funcs))
	for _, ff := range funcs {
		go fetchData(ff, rgChan, errChan, wg)
	}

	go func() {
		wg.Wait()
		close(rgChan)
		close(errChan)
	}()

	unqRanges := make(map[string]byte)
	allErrors := make([]error, 0, len(funcs))

	for data := range rgChan {
		for _, v := range data {
			if v == "" {
				continue
			}
			unqRanges[v] = 0
		}
	}

	ranges := make([]string, 0, len(unqRanges))

	for ip := range unqRanges {
		ranges = append(ranges, ip)
	}

	for err := range errChan {
		allErrors = append(allErrors, err)
	}

	return ranges, allErrors
}
