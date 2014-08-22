package main

import (
	"flag"
	"os"

	"github.com/nabeken/go-ltsview"
)

var flagKeys = flag.String("k", "", "Specify comma-sparated keys to show")
var flagIkeys = flag.String("i", "", "Specify comma-spareted keys to ignore")

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
