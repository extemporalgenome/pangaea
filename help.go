package main

import (
	"fmt"
	"sort"
)

func CmdHelp(args []string) {
	fmt.Println("Available commands:")
	cmds := make(sort.StringSlice, 0, len(Commands))
	for cmd := range Commands {
		cmds = append(cmds, cmd)
	}
	cmds.Sort()
	for _, cmd := range cmds {
		fmt.Println("\t" + cmd)
	}
}

func init() {
	// adding this here to prevent an initialization loop
	Commands["help"] = CmdHelp
}
