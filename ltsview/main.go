package main

import (
	"flag"
	"io"
	"os"

	"github.com/nabeken/go-ltsview"
)

var flagKeys = flag.String("k", "", "Specify comma-separated keys to show")
var flagIkeys = flag.String("i", "", "Specify comma-separated keys to ignore")
var flagNocolor = flag.Bool("n", false, "Set to true to disable colorize outputs")

func main() {
	flag.Parse()
	var w io.Writer
	if *flagNocolor {
		w = os.Stdout
	} else {
		w = &ltsview.ColorableWriter{os.Stdout}
	}
	v := &ltsview.LTSView{
		os.Stdin,
		w,
		ltsview.ParseKeysByFlag(*flagKeys),
		ltsview.ParseKeysByFlag(*flagIkeys),
	}
	v.Start()
}
