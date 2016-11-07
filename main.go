package main

import (
	"time"
	"fmt"
	"sync"
	"github.com/eddyhub/wifictl/system"
)

type loginTime struct {
	from, till time.Time
}

const (
	stop = iota  // stop == 0
)

var (
	ticker *time.Ticker
	wg sync.WaitGroup
	timeLayout = "15:4:5"
	loginTimeMap = map[time.Weekday][]loginTime{
		time.Monday: {
			{
				from:onlyTime(timeLayout, "17:00:00"),
				till:onlyTime(timeLayout, "18:00:00"),
			},
			{
				from:onlyTime(timeLayout, "8:59:10"),
				till:onlyTime(timeLayout, "8:59:50"),
			},
			{
				from:onlyTime(timeLayout, "19:00:00"),
				till:onlyTime(timeLayout, "19:30:00"),
			},
			{
				from:onlyTime(timeLayout, "20:30:00"),
				till:onlyTime(timeLayout, "21:00:00"),
			},
		},
		time.Tuesday: {
			{
				from:onlyTime(timeLayout, "17:00:00"),
				till:onlyTime(timeLayout, "18:00:00"),
			},
			{
				from:onlyTime(timeLayout, "19:00:00"),
				till:onlyTime(timeLayout, "19:30:00"),
			},
			{
				from:onlyTime(timeLayout, "20:30:00"),
				till:onlyTime(timeLayout, "21:00:00"),
			},
		},
		time.Wednesday: {
			{
				from:onlyTime(timeLayout, "14:00:00"),
				till:onlyTime(timeLayout, "15:30:00"),
			},
			{
				from:onlyTime(timeLayout, "20:30:00"),
				till:onlyTime(timeLayout, "21:00:00"),
			},
		},
		time.Thursday: {
			{
				from:onlyTime(timeLayout, "17:00:00"),
				till:onlyTime(timeLayout, "18:00:00"),
			},
			{
				from:onlyTime(timeLayout, "19:00:00"),
				till:onlyTime(timeLayout, "19:30:00"),
			},
			{
				from:onlyTime(timeLayout, "20:30:00"),
				till:onlyTime(timeLayout, "21:00:00"),
			},
		},
		time.Saturday: {
			{
				from:onlyTime(timeLayout, "10:00:00"),
				till:onlyTime(timeLayout, "11:00:00"),
			},
			{
				from:onlyTime(timeLayout, "18:00:00"),
				till:onlyTime(timeLayout, "19:00:00"),
			},
		},
		time.Sunday: {
			{
				from:onlyTime(timeLayout, "10:00:00"),
				till:onlyTime(timeLayout, "11:00:00"),
			},
			{
				from:onlyTime(timeLayout, "18:00:00"),
				till:onlyTime(timeLayout, "19:00:00"),
			},
		},
	}
)

func onlyTime(timeLayout string, timeString string) (t time.Time) {
	t, err := time.Parse(timeLayout, timeString)
	if err != nil {
		panic("Couldn't parse time string!")
	}
	return t
}

func isAfterHourMinute(t1, t2 time.Time) bool {
	return (t1.Hour() * 10000 + t1.Minute() * 100 + t1.Second()) >= (t2.Hour() * 10000 + t2.Minute() * 100 + t2.Second())
}

func isBeforeHourMinute(t1, t2 time.Time) bool {
	return (t1.Hour() * 10000 + t1.Minute() * 100 + t1.Second()) < (t2.Hour() * 10000 + t2.Minute() * 100 + t2.Second())
}

func isInBetween(current, from, till time.Time) bool {
	return isAfterHourMinute(current, from) && isBeforeHourMinute(current, till)
}

func isTimeValid(currentTime time.Time) bool {
	currentWeekday := currentTime.Weekday()
	for _, lt := range loginTimeMap[currentWeekday] {
		if isInBetween(currentTime, lt.from, lt.till) {
			return true;
		}
	}
	return false
}

//func run() {
//	ticker = time.NewTicker(1 * time.Second)
//	go func() {
//		for {
//			select {
//			case tick := <-ticker.C:
//				{
//					fmt.Println(tick)
//					if (isCurrentTimeValid()) {
//						fmt.Println("Valid time...")
//					}
//				}
//			default:
//				time.Sleep(500 * time.Millisecond)
//				fmt.Println("=== default ===")
//			}
//		}
//	}()
//}

func run() {
	ticker = time.NewTicker(1 * time.Second)
	wg.Add(1)
	go func() {
		for {
			currentTime := <-ticker.C
			if (isTimeValid(currentTime) && !system.IsHostapdRunning()) {
				fmt.Println("Starting hostapd!")
				//system.StartHostapd()
			} else if (!isTimeValid(currentTime) && system.IsHostapdRunning()) {
				fmt.Println("Stopping hostapd!")
				//system.StopHostapd()
			}
		}
	}()
}

func main() {

	run()
	wg.Wait()


	//r := mux.NewRouter()
	//api.SetRoutes(r)
	//server := &http.Server{
	//	Addr:    ":8080",
	//	Handler: r,
	//}
	//log.Println("Listining...")
	//server.ListenAndServe()
}
