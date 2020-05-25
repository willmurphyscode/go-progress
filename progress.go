package progress

import (
	"errors"
	"io"
)

var ErrCompleted = io.EOF

type Sizable interface {
	Size() int64
}

type Monitorable interface {
	Current() int64
	Error() error
}

type Progressable interface {
	Monitorable
	Sizable
}

type Progressor interface {
	Progress() Progress
}

type Progress struct {
	current int64
	size    int64
	err     error
}

func (p Progress) Current() int64 {
	return int64(p.current)
}

func (p Progress) Size() int64 {
	return int64(p.size)
}

func (p Progress) Error() error {
	return p.err
}

func (p Progress) Complete() bool {
	return isCompleted(&p)
}

func (p Progress) Ratio() float64 {
	if p.current == 0 || p.size < 0 {
		return 0
	}
	if p.current >= p.size {
		return 1
	}
	return float64(p.current) / float64(p.size)
}

func (p Progress) Percent() float64 {
	if p.current == 0 || p.size < 0 {
		return 0
	}
	if p.current >= p.size {
		return 100
	}
	return 100 / (float64(p.size) / float64(p.current))
}

func isCompleted(p Progressable) bool {
	if errors.Is(p.Error(), io.EOF) || errors.Is(p.Error(), ErrCompleted) {
		return true
	}
	if p.Size() < 0 {
		return false
	}
	return p.Current() >= p.Size()
}
