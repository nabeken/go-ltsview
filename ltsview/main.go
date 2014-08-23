package main

import (
	"flag"
	"os"

	"github.com/nabeken/go-ltsview"
)

var flagKeys = flag.String("k", "", "Specify comma-separated keys to show")
var flagIkeys = flag.String("i", "", "Specify comma-separated keys to ignore")

func main() {
	flag.Parse()
	v := &ltsview.LTSView{
		os.Stdin,
		os.Stdout,
		ltsview.ParseKeysByFlag(*flagKeys),
		ltsview.ParseKeysByFlag(*flagIkeys),
	}
	v.Start()
}
