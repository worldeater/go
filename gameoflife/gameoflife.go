package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
	"github.com/worldeater/go/gameoflife/grid"
	"github.com/worldeater/go/terminal"
)

var interrupted bool = false
var delay = 50

func initSigHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for sig := range sigChan {
			if sig == os.Interrupt {
				terminal.ShowCursor()
				os.Exit(1)
			}
		}
	}()
}

func sleep(ms uint) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
	var (
		width, height, delay uint
		grid                 [2]grid.Grid
		stepMode             bool
		reader               *bufio.Reader
	)

	flag.UintVar(&width, "x", 160, "Grid width, rouned to the next multiple of 8")
	flag.UintVar(&height, "y", 92, "Grid height, rouned to the next multiple of 8")
	flag.UintVar(&delay, "d", 25, "Delay between generations, in milliseconds")
	flag.BoolVar(&stepMode, "s", false, "Activate step mode, hit Return to get the next generation")
	flag.Parse()

	if stepMode {
		reader = bufio.NewReader(os.Stdin)
	}

	if width%8 != 0 {
		width += 8 - width%8
	}
	if height%8 != 0 {
		height += 8 - height%8
	}

	grid[0].Init(width, height)
	grid[1].Init(width, height)

	genCount := 0
	curGen := &grid[0]
	newGen := &grid[1]
	lastChksum := curGen.Checksum()

	curGen.Randomize()
	initSigHandling()
	terminal.HideCursor()
	terminal.Clear()

	for {
		curGen.NextGen(newGen)

		// check for progress
		if genCount%2 == 0 {
			curChksum := curGen.Checksum()
			if lastChksum == curChksum {
				terminal.ShowCursor()
				os.Exit(0)
			}
			lastChksum = curChksum
		}

		terminal.Home()
		fmt.Printf("%vGeneration #%v", curGen.ToString(), genCount)
		genCount++
		sleep(delay)
		curGen, newGen = newGen, curGen
		if stepMode {
			reader.ReadString('\n')
		}
	}
}
