package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/oklog/ulid"
)

type IdStrategy string

const (
	UUID IdStrategy "uuid"
	ULID IdStrategy "ulid"
)

type ID interface {		
	MarshalText() ([] byte, error)
}

type State struct {
	Id ID
	Name string
}

func NewState(name string) State {
	var state State
	state.Id = uuid.New()
	state.Name = name
	return state
}


/*

type ULID ulid.ULID

func (id UUID) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func (id UUID) ToString() string {
	return id.UUID.MarshalText()
}

func (id ULID) Parse(s string) (ulid.ULID, error) {
	return ulid.Parse(s)
}
*/

/*
func NewUUID() ID {
	return uuid.New()
}
*/





func main() {

	fmt.Printf("Machine State in Go\n")
	
	// fmt.Printf("My id:%v\n", myId)

	/*
	openState :=NewState("open")
	closeState :=NewState("close")

	fmt.Printf("Open State %v\n", openState)
	fmt.Printf("Close State %v\n", closeState)
	*/
}

/*

// crand "crypto/rand"
// ulid "github.com/oklog/ulid"

type ULID struct {
	Id ulid.ULID
}

type State struct {
	Id UID
	Name string
}

type FSM struct {
	Id UID
	states []State
}

func NewStateUUID(n string) State {
	var state State
	state.UUID = uuid.New()
	state.Name = n
	return state
}

func NewStateUlid(n string) State {
	var state State
	state.UUID = uuid.New()
	state.Name = n
	return state
}
*/