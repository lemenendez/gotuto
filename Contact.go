package main

type NameType int
type IdType int

const (
	First NameType = iota()
	Middle,
	Last,
	SecondLast
)

type AutoIdShort uint8
type AutoIdMiddle uint16
type AutoIdLarge uint32
type AutoIdHuge uint64


type IID interface {
	Validate
}

/*
type ID struct {
	Id int
	Short string
	Trans
}
*/

type Name struct {
	NType NameType
	Value string	
}

type Contact struct {
	Id int
	names []Name
}




