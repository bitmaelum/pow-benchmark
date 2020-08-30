package main

import (
	"fmt"
	"github.com/bitmaelum/bitmaelum-suite/pkg/proofofwork"
	"github.com/jaypipes/ghw"
	"runtime"
	"time"
)

var (
	message string = "This is the message we are going to provide proof-of-work on"
)

func main() {
	fmt.Println(`Hi there! This program benchmarks your proof-of-work capabilities. This is basically a test to see
how fast your system is in hashing SHA256 numbers. With this information, I can find a balance 
between effort (time) it takes for doing proof-of-work in BitMaelum, and workability. My aim is to 
find a bit-size that would result in a proof-of-work between 2-3 minutes on an average machine.

Hit CTRL-C to stop. The more data, the better though. Please share the results on slack 
(DM to @jaytaph in PHPNL slack)`)

	fmt.Println("---------- START CPU INFO -------------")
	fmt.Println("(Don't send this part if you do not feel comfortable)")
	fmt.Println("")
	displayCpuInfo()
	fmt.Println("---------- END CPU INFO -------------")


	fmt.Println("---------- START WORK INFO -------------")
	doProofOfWork()
	fmt.Println("---------- END WORK INFO -------------")

	fmt.Println("Thanks for participating. Please share the results on slack (phpnl) in a DM to @jaytaph")
}

func displayCpuInfo() {
	fmt.Printf("OS     : %s\n", runtime.GOOS)
	fmt.Printf("Arch   : %s\n", runtime.GOARCH)


	info, err := ghw.CPU()
	if err == nil {
		fmt.Printf("CPUs   : %d (cores: %d)\n", runtime.NumCPU(), info.TotalCores)
		fmt.Printf("Vendor : %s\n", info.Processors[0].Vendor)
		fmt.Printf("Model  : %s\n", info.Processors[0].Model)
	}
}

func doProofOfWork() {
	// Start with 10 bits and gradually increase the number (until 64 or users CTRL-C's)
	bits := 1
	for {
		startRun := time.Now()
		runUntil := startRun.Add(5 * time.Second)

		// Do 100 runs or until time expires
		var cnt, totalTime int64
		for cnt < 1000 {
			cnt++

			startTime := time.Now().UnixNano()
			pow := proofofwork.New(bits, string(message), 0)
			pow.Work()
			endTime := time.Now().UnixNano()

			totalTime += endTime - startTime

			// Check if the run has run for 5 seconds
			if time.Now().After(runUntil) {
				break;
			}
		}

		runTime := time.Now().Sub(startRun)
		avgTime := time.Duration(runTime.Nanoseconds() / cnt)
		fmt.Printf("Bits: %02d   Cnt: %04d   Avg: %-15s Total: %-15s\n", bits, cnt, avgTime, runTime)

		bits++
		if bits > 64 {
			fmt.Printf("oh wow!")
			break
		}
	}

}
