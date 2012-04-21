// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/knieriem/peg"
	"io/ioutil"
	"os"
	"runtime"
)

var inline = flag.Bool("inline", false, "parse rule inlining")
var _switch = flag.Bool("switch", false, "replace if-else if-else like blocks with switch blocks")

func main() {
	runtime.GOMAXPROCS(2)
	flag.BoolVar(&peg.Verbose, "verbose", false, "enable additional output, like statistics")
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "  FILE: the leg file to compile\n")
		os.Exit(1)
	}
	file := flag.Arg(0)

	buffer, error := ioutil.ReadFile(file)
	if error != nil {
		fmt.Printf("%v\n", error)
		return
	}
	p := &Leg{Tree: peg.New(*inline, *_switch), Buffer: string(buffer)}
	p.Init()
	if p.Parse(0) {
		p.Compile(file + ".go")
	} else {
		p.PrintError()
	}
}
