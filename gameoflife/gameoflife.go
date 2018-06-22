package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"github.com/worldeater/go/gameoflife/grid"
	"github.com/worldeater/go/terminal"
	"hash/crc32"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var interrupted bool = false
var delay = 50

func initSigHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for sig := range sigChan {
			if sig == os.Interrupt {
				fmt.Println()
				terminal.ShowCursor()
				os.Exit(1)
			}
		}
	}()
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
	var (
		width, height, delay int
		seed                 int64
	)

	//defer profile.Start().Stop() // XXX XXX XXX

	flag.IntVar(&width, "x", 160, "Grid width, rouned to the next multiple of 8")
	flag.IntVar(&height, "y", 92, "Grid height, rouned to the next multiple of 8")
	flag.IntVar(&delay, "d", 25, "Delay between generations, in milliseconds")
	flag.Int64Var(&seed, "s", 1, "PRNG seed")
	flag.Parse()

	if width%8 != 0 {
		width += 8 - width%8
	}
	if height%8 != 0 {
		height += 8 - height%8
	}

	genCount := 0
	curGen := grid.New(width, height)
	newGen := grid.New(width, height)
	lastChksum := crc32.ChecksumIEEE(*curGen.RawData())

	rand.Seed(seed) //time.Now().UnixNano())
	curGen.Randomize()

	initSigHandling()
	terminal.EraseDisplay()
	terminal.HideCursor()
	defer terminal.ShowCursor()
	defer fmt.Println()

	for {
		curGen.NextGen(&newGen)

		// check for progress
		if genCount%2 == 0 {
			curChksum := crc32.ChecksumIEEE(*curGen.RawData())
			if curChksum == lastChksum {
				return
			}
			lastChksum = curChksum
		}

		terminal.CursorHome()
		fmt.Printf("%vGeneration #%v", curGen.ToString(), genCount)
		genCount++
		sleep(delay)
		curGen, newGen = newGen, curGen
	}
}
