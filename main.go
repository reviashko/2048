package main

import (
	"flag"

	"github.com/reviashko/2048/internal/app"
	"github.com/reviashko/2048/internal/text"
)

func main() {

	var size, digitTextLength int
	flag.IntVar(&size, "size", 5, "size")
	flag.IntVar(&digitTextLength, "dlen", 5, "digit text length")
	flag.Parse()

	colorwriter := text.NewColorDigit(digitTextLength)
	desk := app.NewDesk(size, &colorwriter)
	cntrl := app.NewCntrl(&desk)

	cntrl.Run()
}
