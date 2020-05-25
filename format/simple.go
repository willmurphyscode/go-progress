package format

import (
	"strings"

	"github.com/gookit/color"
	"github.com/wagoodman/go-progress"
)

const (
	liteTheme SimpleTheme = iota
	liteSquashTheme
	heavyTheme
	heavySquashTheme
	reallyHeavySquashTheme
)

const (
	emptyPosition SimplePosition = iota
	fullPosition
	leftCapPosition
	rightCapPosition
)

type SimpleTheme int
type SimplePosition int

var lookup = map[SimpleTheme][]string{
	liteTheme:              {" ", "─", "├", "┤"},
	liteSquashTheme:        {" ", "─", "▕", "▏"},
	heavyTheme:             {"━", "━", "┝", "┥"},
	heavySquashTheme:       {"━", "━", "▕", "▏"},
	reallyHeavySquashTheme: {"━", "━", "▐", "▌"},
}

var (
	doneColor = color.HEX("#ff8700")
	todoColor = color.HEX("#c6c6c6")
)

type Simple struct {
	width   int
	theme   SimpleTheme
	charSet []string
}

func NewSimple(width int) Simple {
	theme := heavySquashTheme
	return Simple{
		width:   width,
		theme:   theme,
		charSet: lookup[theme],
	}
}

func (s Simple) Format(p progress.Progress) (string, error) {

	completedRatio := p.Ratio()
	completedCount := int(completedRatio * float64(s.width))
	todoCount := s.width - completedCount

	completedSection := doneColor.Sprint(strings.Repeat(string(s.charSet[fullPosition]), completedCount))
	todoSection := todoColor.Sprint(strings.Repeat(string(s.charSet[fullPosition]), todoCount))

	return s.charSet[leftCapPosition] + completedSection + todoSection + s.charSet[rightCapPosition], nil
}