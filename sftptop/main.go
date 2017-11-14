package main

import (
	"fmt"
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
		if interval, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println("Usage: sftptop [refresh interval]\nValue of refresh interval must be an integer. If no value is specified then the interval defaults to", interval, "seconds")
			os.Exit(1)
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for range time.NewTicker(time.Second * time.Duration(interval)).C {
		exec.Command("clear")

		proc, err := exec.Command("ps", "aux").Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
