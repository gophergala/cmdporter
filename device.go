package main

// import (
// 	"log"
// )

type Device interface {
	GetName() string
	RegisterCmd(sCmdName string, Bytes []byte)
	DoCmd(sCmdName string)
}
