# Go bot detector

This library was created to provide a means of detecting whether a request was made by a bot or a human user. You can use this library to check if an IP address is likely associated with a bot, such as when the request originates from a cloud provider like AWS or Azure. This check is based on an IP tree structure for fast and efficient results. Additionally, you can verify if the user agent matches a list of common agents commonly used in web scraping.

## Up to date ip ranges list

This repository includes a scraper that performs daily scraping of public IP address ranges from major cloud providers. You can find the list [here](https://github.com/felihenrique/go-bot-detector/tree/master/internal/tools). We welcome contributions of additional open data.

## Using the library

- First install the library on your project:
```
go get github.com/felihenrique/go-bot-detector
```

- Using the library to check an ip:
```go
package main

import "github.com/felihenrique/go-bot-detector"

func main() {
	// Using a specific IP
	isBotIp, err := botdetector.IsBotIp("65.123.44.11")

	if err != nil {
		println(err.Error())
	}

	println(isBotIp)

	// Using an IP range
	isBotRange, err := botdetector.IsBotIp("65.123.44.11/21")

	if err != nil {
		println(err.Error())
	}

	println(isBotRange)
}
```

- Using the library to check a user agent:
```go
package main

import "github.com/felihenrique/go-bot-detector"

func main() {
	isBotAgent := botdetector.IsBotAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	println(isBotAgent)
}
```

