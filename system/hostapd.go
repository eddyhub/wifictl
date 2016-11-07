package system

import (
	"github.com/coreos/go-systemd/dbus"
	"fmt"
)

var (
	conn *dbus.Conn
	ch chan string
)

const (
	serviceName = "hostapd.service"
)

func init() {
	var err error
	conn, err = dbus.New()

	if err != nil {
		panic("Couldn't establish connection to systemd!")
	}

	ch = make(chan string)
}

func StartHostapd() {
	conn.StartUnit(serviceName, "replace", ch)
	outcome := <-ch

	if (outcome == "done") {
		fmt.Println("Starting ", serviceName, " [success]")
	} else {
		fmt.Println("Starting ", serviceName, " [failed]")
	}
}

func StopHostapd() {
	conn.StopUnit(serviceName, "replace", ch)
	outcome := <-ch

	if (outcome == "done") {
		fmt.Println("Stopping ", serviceName, " [success]")
	} else {
		fmt.Println("Stopping ", serviceName, " [failed]")
	}
}

func IsHostapdRunning() bool {
	units, err := conn.ListUnits()
	if err != nil {
		panic("Couldn't get unit " + serviceName)
	}
	for _, unit := range units {
		if unit.Name == serviceName {
			return "active" == unit.ActiveState
		}
	}
	return false
}