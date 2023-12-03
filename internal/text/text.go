package text

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// ColorDigitInterface interface
type ColorDigitInterface interface {
	PrintDigit(digit int)
}

// ColorDigit struct
type ColorDigit struct {
	Colors       map[int]func(format string, a ...interface{})
	RenderLength int
}

// NewColorDigit func
func NewColorDigit(renderLength int) ColorDigit {
	tmp := ColorDigit{Colors: make(map[int]func(format string, a ...interface{})), RenderLength: renderLength}
	tmp.Colors[0] = color.New(color.FgRed).PrintfFunc()
	tmp.Colors[2] = color.New(color.FgGreen).PrintfFunc()
	tmp.Colors[4] = color.New(color.FgBlue).PrintfFunc()
	tmp.Colors[8] = color.New(color.FgCyan).PrintfFunc()
	tmp.Colors[16] = color.New(color.FgHiBlue).PrintfFunc()
	tmp.Colors[32] = color.New(color.FgHiGreen).PrintfFunc()
	tmp.Colors[64] = color.New(color.FgHiRed).PrintfFunc()
	tmp.Colors[128] = color.New(color.FgHiCyan).PrintfFunc()
	tmp.Colors[256] = color.New(color.FgHiYellow).PrintfFunc()
	tmp.Colors[512] = color.New(color.FgYellow).PrintfFunc()
	tmp.Colors[1024] = color.New(color.FgMagenta).PrintfFunc()
	tmp.Colors[2048] = color.New(color.FgHiMagenta).PrintfFunc()

	return tmp
}

// PrintDigit func
func (c *ColorDigit) PrintDigit(digit int) {

	text := fmt.Sprintf("%d%s", digit, strings.Repeat(" ", c.RenderLength-len(fmt.Sprint(digit))))
	if digit == 0 {
		text = strings.Repeat(" ", c.RenderLength)
	}

	fn, isExists := c.Colors[digit]
	if !isExists {
		c.Colors[0](text)
		return
	}

	fn(text)
}
