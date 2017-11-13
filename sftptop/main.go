package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var interval = 5

func main() {
	if len(os.Args) > 1 {
		var err error
		interval, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Invalid interval. Usage: sftptop [refresh interval]. Value of refresh interval must be an integer.")
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for range time.NewTicker(time.Second * time.Duration(interval)).C {
		exec.Command("clear")

		proc, err := exec.Command("ps", "aux").Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, "USER\tPID\tCPU\tMEM\tSTATUS\tSTART")
		for _, line := range strings.Split(string(proc), "\n") {
			if !strings.Contains(line, "internal-sftp") {
				continue
			}

			f := strings.Fields(line)
			if len(f) < 12 {
				continue
			}

			var (
				user   = strings.Replace(f[11], "@internal-sftp", "", -1)
				pid    = f[1]
				cpu    = f[2]
				mem    = f[3]
				status = f[7]
				start  = f[8]
			)

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", user, pid, cpu, mem, status, start)
		}
		w.Flush()
	}
}
