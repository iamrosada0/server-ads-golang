package main

import (
	"log"
	"simple-ads-server/internal/ads"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	reader, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	s := ads.NewServer(reader)

	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
