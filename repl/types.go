package repl

import "github.com/Frank-gh/go-cmd-repl"

type ActionFunc func(...cmd_repl.Argument) interface{}

type ReplCommand struct {
	Name   string
	Action ActionFunc
}

type Repl interface {
	RegisterCommand(name string, action ActionFunc)
	Start() error
}
