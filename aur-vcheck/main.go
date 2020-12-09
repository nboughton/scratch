package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
	regex "github.com/nboughton/go-utils/regex/common"
)

type config struct {
	Versions []string `json:"versions,omitempty"`
}

func main() {
	confPath := "config.json"
	// Check for config path
	if len(os.Args) > 1 {
		confPath = os.Args[1]
	}

	// Read config
	log.Println("Reading config")
	f, err := os.Open(confPath)
	if err != nil {
		log.Fatalf("cannot open config: %s", err)
	}

	c := config{}
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatalf("failed to decode config: %s", err)
	}

	// Check Versions
	log.Println("Checking versions")
	out := []string{}
	for _, appVer := range c.Versions {
		app := fmt.Sprintf("@%s", strings.Split(appVer, "@")[1])
		o, err := exec.Command("npm", "info", app).CombinedOutput()
		if err != nil {
			log.Fatalf("npm exec failed: %s", err)
		}

		v := regex.ANSI.ReplaceAllString(strings.Fields(strings.Split(string(o), "\n")[1])[0], "")
		if appVer != v {
			out = append(out, fmt.Sprintf("%s -> %s", appVer, v))
		}
	}

	if len(out) > 0 {
		log.Println("Connecting to DBUS")
		conn, err := dbus.SessionBusPrivate()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		if err = conn.Auth(nil); err != nil {
			log.Fatal(err)
		}

		if err = conn.Hello(); err != nil {
			log.Fatal(err)
		}

		// Send notification
		log.Println("Sending notification")
		notify.SendNotification(conn, notify.Notification{
			AppName:       "AUR VCHECK",
			ReplacesID:    uint32(0),
			AppIcon:       "mail-message-new",
			Summary:       "AUR NPM Packages Need Updating",
			Body:          strings.Join(out, "\n"),
			Hints:         map[string]dbus.Variant{},
			ExpireTimeout: 10000,
		})
	}
}
