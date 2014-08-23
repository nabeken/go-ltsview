package main

import (
	"flag"
	"os"

	"github.com/nabeken/go-ltsview"
)

var flagKeys = flag.String("k", "", "Specify comma-separated keys to show")
var flagIkeys = flag.String("i", "", "Specify comma-separated keys to ignore")
var flagNocolor = flag.Bool("n", false, "Set to true to disable colorized outputs")

func main() {
	flag.Parse()
	var w ltsview.Writer
	if *flagNocolor {
		w = ltsview.NewWriter(os.Stdout)
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
