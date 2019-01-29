package analysis

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func Geo_search(connect_ip string) (string, string) {
	var city string
	var country string

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(connect_ip)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	city = record.City.Names["en"]
	country = record.Country.Names["en"]

	if city == "" && country == "" {

		city = "Unknown"
		country = "Unknown"
	} else if city == "" {
		city = "Unknown"

	} else if country == "" {

		country = "Unknown"
	} else {
		return city, country
	}

	return city, country

}
