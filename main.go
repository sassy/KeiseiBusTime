package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	DefaultNumOfResultToShow = 3
)

type Time struct {
	hour, minute int
}

func (time *Time) toString() string {
	if time.minute < 10 {
		return fmt.Sprintf("%d:0%d ", time.hour, time.minute)
	} else {
		return fmt.Sprintf("%d:%d ", time.hour, time.minute)
	}
}

func (t1 *Time) isLaterThan(t2 Time) bool {
	return t1.hour > t2.hour || (t1.hour == t2.hour && t1.minute > t2.minute)
}

func departureTime(departure string) Time {
	ret, _ := regexp.MatchString("^[0-9]{1,2}:[0-9]{1,2}$", departure)
	var hour int
	var minute int
	if ret {
		re := regexp.MustCompile("^([0-9]{1,2}):([0-9]{1,2})$")
		bs := []byte(departure)
		group := re.FindSubmatch(bs)
		h, _ := strconv.Atoi(string(group[1]))
		m, _ := strconv.Atoi(string(group[2]))
		if h >= 0 && h < 24 && m >= 0 && m < 60 {
			hour = h
			minute = m
		} else {
			now := time.Now()
			hour = now.Hour()
			minute = now.Minute()
		}
	} else {
		now := time.Now()
		hour = now.Hour()
		minute = now.Minute()
	}
	return Time{hour, minute}
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

func createTimetable(selector string) []Time {
	var timetable = make([]Time, 0)
	doc, _ := goquery.NewDocument("http://www.keiseibus.co.jp/jikoku/bs_tt.php?key=04159_01a")
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
			key, _ := strconv.Atoi(s.Find("th").Text())
			s.Find("td>span").Each(func(_ int, s *goquery.Selection) {
				s.Find(".notes").Remove()
				s.Find("br").Remove()
				if s.Text() != "" {
					value, _ := strconv.Atoi(s.Text())
					timetable = append(timetable, Time{key, value})
				}
			})
		})
	})
	return timetable
}

func printTimes(times []string) {
	for _, v := range times {
		fmt.Println(v)
	}
}

func main() {
	var (
		departure   string
		numOfResult int
		isLast      bool
	)
	flag.StringVar(&departure, "t", "", "specify departure time.")
	flag.BoolVar(&isLast, "l", false, "show last bus of the day.")
	flag.IntVar(&numOfResult, "n", DefaultNumOfResultToShow, "specify amount of result.")
	flag.Parse()

	if numOfResult < 0 {
		fmt.Fprintf(os.Stderr, "parameter for -n must be greater than 0.\n")
		os.Exit(2)
	}

	depTime := departureTime(departure)
	timetable := createTimetable(getSelector())

	if isLast {
		result := []string{timetable[len(timetable)-1].toString()}
		printTimes(result)
		return
	}

	result := make([]string, 0, numOfResult)
	for i := 0; i < len(timetable); i++ {
		v := timetable[i]
		if v.isLaterThan(depTime) {
			timeStr := v.toString()
			if i < len(timetable)-1 && v == timetable[i+1] {
				timeStr += "(2)"
				i++
			}
			result = append(result, timeStr)
			if len(result) >= numOfResult {
				break
			}
		}
	}
	printTimes(result)
}
