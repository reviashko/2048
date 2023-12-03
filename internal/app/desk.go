package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/reviashko/2048/internal/text"
	"github.com/reviashko/2048/model"
)

// DeskInterface interface
type DeskInterface interface {
	Print()
	Reset()
	AddDigit(num int) bool
	ShiftLeft()
	ShiftRight()
	ShiftTop()
	ShiftDown()
}

// Desk struct
type Desk struct {
	Desk        [][]int
	Size        int
	ColorWriter text.ColorDigitInterface
}

// NewDesk func
func NewDesk(size int, colorWriter text.ColorDigitInterface) Desk {

	if size < 4 || size > 64 {
		panic("4 <= size <= 64")
	}

	desk := make([][]int, size)
	row := 0

	for row < size {
		desk[row] = make([]int, size)
		row++
	}

	return Desk{Desk: desk, Size: size, ColorWriter: colorWriter}
}

// Reset func
func (d *Desk) Reset() {
	row := 0
	for row < d.Size {
		d.Desk[row] = make([]int, d.Size)
		row++
	}
}

// Print func
func (d *Desk) Print() {

	row := 0
	fmt.Println("Esc - exit. Use ↑ → ↓ ← to shift")
	for row < d.Size {
		fmt.Print("[ ")

		col := 0
		for col < d.Size {
			d.ColorWriter.PrintDigit(d.Desk[row][col])
			col++
		}

		fmt.Print("]\n")

		row++
	}
}

// AddDigit func
func (d *Desk) AddDigit(num int) bool {

	row := 0

	avl := make([]model.Cell, 0)

	for row < d.Size {
		col := 0
		for col < d.Size {

			if d.Desk[row][col] == 0 {
				avl = append(avl, model.Cell{Row: row, Col: col})
			}

			col++
		}
		row++
	}

	if len(avl) == 0 {
		return false
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	cell := r1.Intn(len(avl))

	d.Desk[avl[cell].Row][avl[cell].Col] = num

	return true
}

// shiftLeftRow func
func (d *Desk) shiftLeftRow(row int) {
	col := 0
	next := 0
	for col < d.Size {

		if d.Desk[row][col] != 0 {
			d.Desk[row][next] = d.Desk[row][col]
			if col != next {
				d.Desk[row][col] = 0
			}
			next++
		}
		col++
	}
}

// ShiftLeft func
func (d *Desk) ShiftLeft() {

	row := 0

	for row < d.Size {

		d.shiftLeftRow(row)

		col := 0
		for col < d.Size {

			if col < (d.Size-1) && d.Desk[row][col] == d.Desk[row][col+1] {
				d.Desk[row][col] = d.Desk[row][col] * 2
				d.Desk[row][col+1] = 0
			}

			col++
		}

		d.shiftLeftRow(row)

		row++
	}

}

// shiftRightRow func
func (d *Desk) shiftRightRow(row int) {
	col := d.Size - 1
	next := d.Size - 1
	for col >= 0 {

		if d.Desk[row][col] != 0 {
			d.Desk[row][next] = d.Desk[row][col]
			if col != next {
				d.Desk[row][col] = 0
			}
			next--
		}
		col--
	}
}

// ShiftRight func
func (d *Desk) ShiftRight() {

	row := 0

	for row < d.Size {

		d.shiftRightRow(row)

		col := d.Size - 1
		for col >= 0 {

			if col > 0 && d.Desk[row][col] == d.Desk[row][col-1] {
				d.Desk[row][col] = d.Desk[row][col] * 2
				d.Desk[row][col-1] = 0
			}

			col--
		}

		d.shiftRightRow(row)

		row++
	}

}

// shiftDownCol func
func (d *Desk) shiftDownCol(col int) {
	row := d.Size - 1
	next := d.Size - 1
	for row >= 0 {

		if d.Desk[row][col] != 0 {
			d.Desk[next][col] = d.Desk[row][col]
			if row != next {
				d.Desk[row][col] = 0
			}
			next--
		}
		row--
	}
}

// ShiftDown func
func (d *Desk) ShiftDown() {

	col := 0

	for col < d.Size {

		d.shiftDownCol(col)

		row := d.Size - 1
		for row >= 0 {

			if row > 0 && d.Desk[row][col] == d.Desk[row-1][col] {
				d.Desk[row][col] = d.Desk[row][col] * 2
				d.Desk[row-1][col] = 0
			}

			row--
		}

		d.shiftDownCol(col)

		col++
	}

}

// shiftTopCol func
func (d *Desk) shiftTopCol(col int) {
	row := 0
	next := 0
	for row < d.Size {

		if d.Desk[row][col] != 0 {
			d.Desk[next][col] = d.Desk[row][col]
			if row != next {
				d.Desk[row][col] = 0
			}
			next++
		}
		row++
	}
}

// ShiftTop func
func (d *Desk) ShiftTop() {

	col := 0

	for col < d.Size {

		d.shiftTopCol(col)

		row := 0
		for row < d.Size {

			if row < (d.Size-1) && d.Desk[row][col] == d.Desk[row+1][col] {
				d.Desk[row][col] = d.Desk[row][col] * 2
				d.Desk[row+1][col] = 0
			}

			row++
		}

		d.shiftTopCol(col)

		col++
	}

}
