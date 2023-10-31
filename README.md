# Go bot detector

This library was created to provide ways to detect if a requisition was made from a bot or not.
You can use the library to check if an IP is probably from a bot (checking if the requisition was made from a cloud provider like AWS or Azure) or to check if the user agent is from a list of common agents used in scraping. The check is based on an IP tree to provide fast returns.

## Up to date ip ranges list

This repository has a scraper that do daily scraping of public ip ranges from major clouds (you can see the list at <a href="https://github.com/felihenrique/go-botdetector/tree/master/internal/tools">internal/tools</a>). You are welcome to contribute with more open data.

## Using the library

- First install the library on your project:
```
go get github.com/felihenrique/go-botdetector
```

- Using the library to check an ip:
```go
package main

import "github.com/felihenrique/go-botdetector"

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

import "github.com/felihenrique/go-botdetector"

func main() {
	isBotAgent := botdetector.IsBotAgent("googlebot")

	println(isBotAgent)
}
```

