package ltsview

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

type writer struct {
	io.Writer
}

func (w *writer) WriteSeparator(sep string) (int, error) {
	return fmt.Fprintln(w.Writer, sep)
}

func NewWriter(w io.Writer) Writer {
	return &writer{w}
}

type Writer interface {
	io.Writer
	SeparatorWriter
}

type SeparatorWriter interface {
	WriteSeparator(sep string) (int, error)
}

type ColorableWriter struct {
	io.Writer
}

func (w *ColorableWriter) WriteSeparator(sep string) (int, error) {
	return fmt.Fprintln(w.Writer, sep)
}

func (w *ColorableWriter) Write(p []byte) (int, error) {
	f := color.New(color.FgMagenta).SprintFunc()
	keyPrinted := false
	for i := range p {
		str := string(p[i])
		if !keyPrinted && p[i] == ':' {
			fmt.Fprint(w.Writer, str)
			keyPrinted = true
			f = color.New(color.FgGreen).SprintFunc()
			continue
		}
		fmt.Fprint(w.Writer, f(str))
	}
	return len(p), nil
}
