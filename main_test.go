package main

import (
	"testing"
	"time"
)

func TestDepartureTime(t *testing.T) {
	hour, minutes := departureTime("10:30")
	expected_hour := 10
	expected_minutes := 30
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime2(t *testing.T) {
	hour, minutes := departureTime("0:0")
	expected_hour := 0
	expected_minutes := 0
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime3(t *testing.T) {
	hour, minutes := departureTime("00:00")
	expected_hour := 0
	expected_minutes := 0
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime4(t *testing.T) {
	hour, minutes := departureTime("23:59")
	expected_hour := 23
	expected_minutes := 59
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime5(t *testing.T) {
	hour, minutes := departureTime("44:30")
	now := time.Now()
	expected_hour := now.Hour()
	expected_minutes := now.Minute()
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime6(t *testing.T) {
	hour, minutes := departureTime("24:30")
	now := time.Now()
	expected_hour := now.Hour()
	expected_minutes := now.Minute()
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestDepartureTime7(t *testing.T) {
	hour, minutes := departureTime("12:60")
	now := time.Now()
	expected_hour := now.Hour()
	expected_minutes := now.Minute()
	if hour != expected_hour {
		t.Errorf("got %v\nwant %v", hour, expected_hour)
	}
	if minutes != expected_minutes {
		t.Errorf("got %v\nwant %v", minutes, expected_minutes)
	}
}

func TestGetSelector(t *testing.T) {
	selector := getSelector()
	if time.Now().Weekday().String() == "Saturday" {
		if selector != "#tab-2 .standard2" {
			t.Errorf("got %v\nwant #tab-2 .standard2", selector)
		}
	} else if time.Now().Weekday().String() == "Sunday" {
		if selector != "#tab-3 .standard2" {
			t.Errorf("got %v\nwant #tab-3 .standard2", selector)
		}
	} else {
		if selector != "#tab-1 .standard2" {
			t.Errorf("got %v\nwant #tab-1 .standard2", selector)
		}
	}
}

func TestToString(t *testing.T) {
	value := Time{8, 8}
	ret := value.toString()
	if ret != "8:08 " {
		t.Error("got %v, wrong Format", ret)
	}
}

func TestToString2(t *testing.T) {
	value := Time{10, 10}
	ret := value.toString()
	if ret != "10:10 " {
		t.Error("got %v, wrong Format", ret)
	}
}
