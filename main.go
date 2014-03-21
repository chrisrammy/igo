package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DAddYE/igo/cmd"
)

type Cmd int

const (
	INVALID Cmd = iota
	COMPILE
	PARSE
	BUILD
)

var commands = []string{
	COMPILE: "compile",
	PARSE: "parse",
	BUILD: "build",
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: igo [" + strings.Join(commands[1:], "|") + "] [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func toCmd(c string) Cmd {
	for i, command := range commands {
		if c == command {
			return Cmd(i)
		}
	}
	return Cmd(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	var (
		command Cmd
		paths []string
		exitCode = 0
	)

	for i := 0; i < flag.NArg(); i++ {
		s := flag.Arg(i)
		if cmd := toCmd(s); cmd > 0 {
			command = cmd
		} else {
			// Could be a path
			paths = append(paths, s)
		}
	}

	switch command {
	case PARSE:
		exitCode = cmd.To(cmd.IGO, paths)
	case COMPILE:
		exitCode = cmd.To(cmd.GO, paths)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		usage()
	}

	os.Exit(exitCode)
}