package service

import (
	"fmt"
	"strings"
)

var months = map[string]string{
	"01": "Jan",
	"02": "Feb",
	"03": "Mar",
	"04": "Apr",
	"05": "May",
	"06": "Jun",
	"07": "July",
	"08": "Aug",
	"09": "Sep",
	"10": "Oct",
	"11": "Nov",
	"12": "Dec",
}

func ConvertDate(date string) string {
	from := 0
	if date[0] == '0' {
		from = 1
	}
	return fmt.Sprintf("%s %s %s", date[from:2], months[date[3:5]], date[6:])
}

func ConvertLocation(location string) string {
	locTemp := strings.Split(strings.ReplaceAll(location, "_", " "), "-")
	city := strings.Title(locTemp[0])
	var country string
	switch locTemp[1] {
	case "usa", "uk":
		country = strings.ToUpper(locTemp[1])
	default:
		country = strings.Title(locTemp[1])
	}
	return fmt.Sprintf("%s, %s", city, country)
}
