package main

import (
	"testing"
	"time"
)

func TestDepartureTime(t *testing.T) {
	depTime := departureTime("10:30")
	expected_time := Time{10, 30}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime2(t *testing.T) {
	depTime := departureTime("0:0")
	expected_time := Time{0, 0}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime3(t *testing.T) {
	depTime := departureTime("00:00")
	expected_time := Time{0, 0}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime4(t *testing.T) {
	depTime := departureTime("23:59")
	expected_time := Time{23, 59}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime5(t *testing.T) {
	depTime := departureTime("44:30")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime6(t *testing.T) {
	depTime := departureTime("24:30")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
	}
}

func TestDepartureTime7(t *testing.T) {
	depTime := departureTime("12:60")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if depTime != expected_time {
		t.Errorf("got %v\nwant %v", depTime, expected_time)
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

func TestIsLaterThan1(t *testing.T) {
	t1 := Time{11, 00}
	t2 := Time{10, 30}
	if !t1.isLaterThan(t2) {
		t.Error("%v should be later than %v", t1, t2)
	}
}

func TestIsLaterThan2(t *testing.T) {
	t1 := Time{10, 10}
	t2 := Time{10, 00}
	if !t1.isLaterThan(t2) {
		t.Errorf("%v should be later than %v", t1, t2)
	}
}
