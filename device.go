package main

// import (
// 	"log"
// )

type Device interface {
	GetName() string
	DoCmd(sCmdName string)
}
