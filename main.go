package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "time"
)

func main() {
  var timetable = make([][]int, 24)
  now := time.Now()
  hour := now.Hour()
  minute := now.Minute()

  doc, _ := goquery.NewDocument("http://www.keiseibus.co.jp/jikoku/bs_tt.php?key=04159_01a")
  doc.Find("#tab-1 .standard2").Each(func(_ int, s *goquery.Selection) {
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
  
  for _, v:= range result {
    if v < 10 {
      fmt.Println(fmt.Sprintf("%d:0%d ", hour, v))
    } else {
      fmt.Println(fmt.Sprintf("%d:%d ", hour, v))
    }

  }

}
