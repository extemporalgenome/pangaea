package main

var Prefixes = []string{
	"pn",
	"pangaea",
}

var Commands = map[string]CmdHandler{
	"cat":  CmdCat,
	"echo": CmdEcho,
}
