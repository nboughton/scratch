package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for range time.NewTicker(time.Second * 5).C {
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
