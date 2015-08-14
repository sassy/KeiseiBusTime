package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"time"
)

func main() {
	var timetable = make([][]int, 24)
	now := time.Now()

	var departure = flag.String("t", "", "specify departure time.")
	flag.Parse()
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
			hour = now.Minute()
			minute = now.Minute()
		}
	} else {
		hour = now.Hour()
		minute = now.Minute()
	}
	weekday := now.Weekday().String()

	doc, _ := goquery.NewDocument("http://www.keiseibus.co.jp/jikoku/bs_tt.php?key=04159_01a")
	var selector string
	if weekday == "Saturday" {
		selector = "#tab-2 .standard2"
	} else if weekday == "Sunday" {
		selector = "#tab-3 .standard2"
	} else {
		selector = "#tab-1 .standard2"
	}
	//selector := "#tab-2 .standard2"
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

	for _, v := range result {
		if v < 10 {
			fmt.Println(fmt.Sprintf("%d:0%d ", hour, v))
		} else {
			fmt.Println(fmt.Sprintf("%d:%d ", hour, v))
		}
	}

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
		for _, v := range result2 {
			if v < 10 {
				fmt.Println(fmt.Sprintf("%d:0%d ", hour+1, v))
			} else {
				fmt.Println(fmt.Sprintf("%d:%d ", hour+1, v))
			}
		}
	}

}
