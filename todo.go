package todo

import "time"

type todo struct {
	guid        uint64
	description string
	created     time.Time
	updated     []time.Time
	completed   time.Time
}

type todoCollection struct {
	guid     uint64
	name     string
	elements []uint64
}

type app struct {
	defaultCollection uint64
	otherCollections  []uint64
}

func (a *app) newCollection(name, collection string) error {}

func (a *app) newTodoInDefaultCollection(description string) error {}

func (a *app) listAllTodos() {}

func (a *app) markDone(guid uint64) error {}
