package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"strings"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName = "GOLANG CLI REMINDER"
	markValue = "1"
)

func main() {
	// If no arguments is passed
	if len(os.Args) < 3 {
		// os.Args[0] is the location of the project
		fmt.Printf("Usage %s <hh:mm> <text mesage\n>", os.Args[0]) 
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Println("Unable to parse time")
		os.Exit(2)
	}

	if now.After(t.Time) {
		fmt.Println("invalid! Set a future time")
		os.Exit(3)
	}

	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err := beeep.Alert("Reminder: ", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		fmt.Printf("Reminder will be displayed after %s\n", diff.Round(time.Second))
		os.Exit(0)
	}
}