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
	interval    = 2
	usage       = fmt.Sprintf("Usage: sftptop [refresh interval]\nValue of refresh interval must be an integer. If no value is specified then the interval defaults to %d second(s).", interval)
	sftpSession = "notty"
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
		load, err := exec.Command("uptime").Output()
		fatal(err)

		proc, err := exec.Command("ps", "aux").Output()
		fatal(err)

		clear()
		fmt.Println(string(load))
		fmt.Fprintln(w, "USER\tPID\tCPU\tMEM\tSTATUS\tSTART")
		for _, line := range strings.Split(string(proc), "\n") {
			if !strings.Contains(line, sftpSession) {
				continue
			}

			f := strings.Fields(line)
			if len(f) < 12 {
				continue
			}

			var (
				user   = strings.Replace(f[11], "@"+sftpSession, "", -1)
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

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
