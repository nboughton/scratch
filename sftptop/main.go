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

var (
	interval = 1
	usage    = fmt.Sprintf("Usage: sftptop [refresh interval]\nValue of refresh interval must be an integer. If no value is specified then the interval defaults to %d second(s).", interval)
)

func main() {
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(usage)
			os.Exit(1)
		}

		interval = i
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for range time.NewTicker(time.Second * time.Duration(interval)).C {
		exec.Command("clear")

		load, err := exec.Command("uptime").Output()
		fatal(err)
		fmt.Println("LOAD:", string(load))

		fmt.Fprintln(w, "USER\tPID\tCPU\tMEM\tSTATUS\tSTART")
		proc, err := exec.Command("ps", "aux").Output()
		fatal(err)

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

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
