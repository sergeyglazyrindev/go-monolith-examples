package models

import "github.com/sergeyglazyrindev/go-monolith/core"

type Todo struct {
	core.Model
	TaskAlias       string
	TaskDescription string
}

func (t *Todo) String() string {
	return t.TaskAlias
}
