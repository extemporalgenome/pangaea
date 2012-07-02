package main

import (
	"io"
	"os"
)

func CmdCat(args []string) {
	if len(args) == 1 {
		args[0] = "-"
	} else {
		args = args[1:]
	}
	var (
		stdindone bool
		err       error
	)
	for _, arg := range args {
		var r io.ReadCloser
		if arg == "-" {
			if stdindone {
				continue
			}
			r = os.Stdin
			stdindone = true
		} else if r, err = os.Open(arg); err != nil {
			Exit(1, err)
		}
		if _, err = io.Copy(os.Stdout, r); err != nil {
			if arg == "-" {
				arg = "(stdin)"
			}
			Exit(1, arg, "error:", err)
		}
	}
}
