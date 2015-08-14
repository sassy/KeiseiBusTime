package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"time"
)

func departureTime(departure *string) (int, int) {
	ret, _ := regexp.MatchString("^[0-9]{1,2}:[0-9]{1,2}$", *departure)
	var hour int
	var minute int
	if ret {
		re := regexp.MustCompile("^([0-9]{1,2}):([0-9]{1,2})$")
		bs := []byte(*departure)
		group := re.FindSubmatch(bs)
		h, _ := strconv.Atoi(string(group[1]))
		m, _ := strconv.Atoi(string(group[2]))
		if h >= 0 && h < 24 && m >= 0 && m < 60 {
			hour = h
			minute = m
		} else {
			now := time.Now()
			hour = now.Minute()
			minute = now.Minute()
		}
	} else {
		now := time.Now()
		hour = now.Hour()
		minute = now.Minute()
	}
	return hour, minute
}

func getSelector() string {
	weekday := time.Now().Weekday().String()
	if weekday == "Saturday" {
		return "#tab-2 .standard2"
	} else if weekday == "Sunday" {
		return "#tab-3 .standard2"
	} else {
		return "#tab-1 .standard2"
	}
}

func createTimetable(selector string) [][]int {
	var timetable = make([][]int, 24)
	doc, _ := goquery.NewDocument("http://www.keiseibus.co.jp/jikoku/bs_tt.php?key=04159_01a")
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
			key, _ := strconv.Atoi(s.Find("th").Text())
			s.Find("td>span").Each(func(_ int, s *goquery.Selection) {
				s.Find(".notes").Remove()
				s.Find("br").Remove()
				if s.Text() != "" {
					value, _ := strconv.Atoi(s.Text())
					timetable[key] = append(timetable[key], value)
				}
			})
		})
	})
	return timetable
}

func printTimes(hour int, minuteses []int) {
	for _, v := range minuteses {
		if v < 10 {
			fmt.Println(fmt.Sprintf("%d:0%d ", hour, v))
		} else {
			fmt.Println(fmt.Sprintf("%d:%d ", hour, v))
		}
	}
}

func main() {
	var departure = flag.String("t", "", "specify departure time.")
	flag.Parse()
	hour, minute := departureTime(departure)
	timetable := createTimetable(getSelector())

	arrivals := timetable[hour]
	result := make([]int, 0, 3)
	for _, v := range arrivals {
		if v > minute {
			result = append(result, v)
			if len(result) >= 3 {
				break
			}
		}
	}
	printTimes(hour, result)

	if hour != 23 && len(result) < 3 {
		max := 3 - len(result)
		arrivals = timetable[hour+1]
		result2 := make([]int, 0, max)
		for _, v := range arrivals {
			result2 = append(result2, v)
			if len(result2) >= max {
				break
			}
		}
		printTimes(hour+1, result2)
	}
}
