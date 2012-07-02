package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type CmdHandler func([]string)

func Exit(code int, args ...interface{}) {
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(code)
}

func NormalizeName(name string, sub bool) string {
	name = path.Base(name)
	norm := name
	for _, p := range Prefixes {
		if strings.HasPrefix(norm, p) {
			norm = norm[len(p):]
			if strings.HasPrefix(norm, "-") {
				norm = norm[1:]
			} else if norm != "" {
				norm = name
			}
			if sub && norm == "" {
				norm = name
			}
			break
		}
	}
	return norm
}

// Currently doesn't support external commands
func Dispatch(args []string) (CmdHandler, []string) {
	i := 0
start:
	if i <= 1 && len(args) > 0 {
		name := args[0]
		norm := NormalizeName(name, i > 0)
		if norm == "" {
			args = args[1:]
			i++
			goto start
		}
		if cmd, ok := Commands[norm]; ok {
			return cmd, args
		}
		args[0] = norm
	}
	return nil, args
}

func main() {
	cmd, args := Dispatch(os.Args)
	if cmd == nil {
		CmdHelp(args)
		if len(args) > 0 {
			Exit(2, "Unrecognized command:", args[0])
		}
	} else {
		cmd(args)
	}
}
