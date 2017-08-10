package lib

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/0xAX/notificator"
)

func Set(mode string) {
	if user, err := user.Current(); err == nil {
		fileName := user.HomeDir + "/.gomato_timer"
		if file, err := os.Create(fileName); err == nil {
			defer file.Close()
			n := time.Now()
			timestr := fmt.Sprintf("%s\n%d", mode, n.Unix())
			file.Write([]byte(timestr))
		}
	}
}

func Get(remaining bool) (string, error) {
	if user, err := user.Current(); err == nil {
		fileName := user.HomeDir + "/.gomato_timer"
		if file, err := os.Open(fileName); err == nil {
			var starttime int64 = 0
			var mode = "N"
			fmt.Fscanln(file, &mode)
			fmt.Fscanln(file, &starttime)

			now := time.Now()
			diff := int(now.Unix() - starttime)
			var min int = diff / 60
			var sec int = diff - (min * 60)

			if mode == "N" {
				return "", nil
			}

			maxTime := 25
			desc := "Working"
			if mode == "R" {
				maxTime = 5
				desc = "Resting"
			}

			if mode == "X" {
				return "", nil
			}

			if min < maxTime {
				if remaining {
					desc = fmt.Sprintf("%s %d:%02.f", desc, (maxTime - min - 1), (60 - float32(sec)))
				} else {
					desc = fmt.Sprintf("%s %d:%02.f", desc, min, float32(sec))
				}
			} else if min >= maxTime {
				desc = desc + " done"
				Alert(desc, "Good Job!")
				Rm()
			}

			return desc, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func Rm() {
	if user, err := user.Current(); err == nil {
		fileName := user.HomeDir + "/.gomato_timer"
		if file, err := os.Create(fileName); err == nil {
			defer file.Close()
		}
	}
}

func Alert(title string, text string) {
	var notify *notificator.Notificator

	notify = notificator.New(notificator.Options{
		AppName: "Gomato",
	})

	notify.Push(title, text, "", notificator.UR_CRITICAL)
}
