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

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", strings.Replace(f[11], "@internal-sftp", "", -1), f[1], f[2], f[3], f[7], f[8])
		}
		w.Flush()
	}
}
