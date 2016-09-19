package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
)

func init() {
	verbose := flag.Bool("v", false, "Verbose mode will display all debug informations")
	flag.Parse()

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	g := newGame()
	err := g.run()
	if nil != err {
		panic(err)
	}
}
