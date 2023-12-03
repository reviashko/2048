package app

import (
	"fmt"
	"os"

	term "github.com/nsf/termbox-go"
)

// Cntrl struct
type Cntrl struct {
	Desk DeskInterface
}

// NewCntrl func
func NewCntrl(desk DeskInterface) Cntrl {
	return Cntrl{Desk: desk}
}

// Run func
func (c *Cntrl) Run() {

	c.Desk.AddDigit(2)
	c.Desk.AddDigit(2)

	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	c.Desk.Print()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowUp:
				term.Sync()
				c.Desk.ShiftTop()
			case term.KeyArrowDown:
				term.Sync()
				c.Desk.ShiftDown()
			case term.KeyArrowLeft:
				term.Sync()
				c.Desk.ShiftLeft()
			case term.KeyArrowRight:
				term.Sync()
				c.Desk.ShiftRight()
			case term.KeyEsc:
				term.Sync()
				os.Exit(0)
			case term.KeyEnter:
				term.Sync()
				c.Desk.Reset()

			}

			if !c.Desk.AddDigit(2) {
				fmt.Println("Game over!")
			}
			c.Desk.Print()

		case term.EventError:
			panic(ev.Err)
		}
	}

}
