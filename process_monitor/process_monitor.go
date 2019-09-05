package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	humanize "github.com/dustin/go-humanize"
	ps "github.com/shirou/gopsutil/process"
)

var pid int32
var process ps.Process
var statusString string

func main() {
	if len(os.Args) == 1 {
		log.Fatal("no pid provided")
	}

	pid, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	process, err := ps.NewProcess(int32(pid))
	if err != nil {
		log.Fatal(err)
	}

	for {
		mem, err := process.MemoryInfo()
        if err != nil {
            log.Fatal(err)
        }

		times, err := process.Times()
		if err != nil {
			log.Fatal(err)
		}

		rss := humanize.Ftoa(float64(mem.RSS / 1024.0 / 1024.0))
		vms := humanize.Ftoa(float64(mem.VMS / 1024.0 / 1024.0))

		fmt.Printf("\rRSS: %vMB\tVMS: %vMB\tSYS: %v\tUSR: %v", rss, vms, times.System, times.User)
		time.Sleep(time.Second)
		fmt.Printf("\r%50v", ' ') // clear status string
	}
}
