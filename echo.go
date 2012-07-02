package main

import "os"

func CmdEcho(args []string) {
	f := os.Stdout
	if len(args) > 1 {
		f.WriteString(args[1])
		for _, arg := range args[2:] {
			// ignoring possible write errors -- don't care yet
			f.Write([]byte{' '})
			f.WriteString(arg)
		}
	}
	f.Write([]byte{'\n'})
}
