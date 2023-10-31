test:
	go test ./...
scrape:
	./bin/ip-scraper
build:
	go build cmd/ip-scraper/main.go && mv main ./bin/ip-scraper