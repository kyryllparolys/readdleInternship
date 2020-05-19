// Package main provides starting the application
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// holiday structure
type Holiday struct {
	Name        string      `json:"name"`
	CountryCode string      `json:"countryCode"`
	Date       	string     `json:"date"`
	Fixed       bool        `json:"fixed"`
	Global      bool        `json:"global"`
	Counties    interface{} `json:"counties"`
	LaunchYear  interface{} `json:"launchYear"`
	LocalName   string      `json:"localName"`
	Type        string      `json:"type"`
}

func isWeekend(date time.Time) bool {
	if string(date.Weekday()) == "Sunday" || string(date.Weekday()) == "Saturday" {
		return true
	}
	return false
}

func isAdjacent(date time.Time, adjacent bool) (time.Time, bool) {
	if isWeekend(date) {
		isAdjacent(date.AddDate(0, 0, 1), true)
	}
	return date, adjacent
}

func main() {
	var arr []Holiday
	t := time.Now()
	url := fmt.Sprintf("https://date.nager.at/api/v2/publicholidays/%d/UA", t.Year())
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(contents), &arr)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	for _, foo := range arr {
		holiday, err := time.Parse("2006-01-02", foo.Date)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		if holiday.After(t) {
			nextDay := holiday.AddDate( 0, 0, 1)

			var lastDay time.Time
			var adjacent bool
			if isWeekend(nextDay) {
				lastDay, adjacent = isAdjacent(nextDay, false)
			}

			if adjacent {
				//TODO: Test adjacent
				daysBetween := lastDay.Sub(holiday).Hours() / 24
				fmt.Printf(
					"The next holiday is %v. It will last for %v days. From %v %v till %v %v",
					foo.Name,
					daysBetween,
					holiday.Month(),
					holiday.Day(),
					lastDay.Month(),
					lastDay.Day())
			} else {
				fmt.Printf("The next holiday %v is on: %v %v", foo.Name, holiday.Month(), holiday.Day())
			}
			break
		}
	}
}
