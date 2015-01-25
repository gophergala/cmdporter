package main

<<<<<<< HEAD
type Device interface {
	GetJsonPath() string
	GetName() string
	SetName(name string)
	RegisterCmd(sCmdName string, Bytes []byte)
	GetNumCommands() int
=======
// import (
// 	"log"
// )

type Device interface {
	GetName() string
	RegisterCmd(sCmdName string, Bytes []byte)
>>>>>>> e7bc7230bb514cdc1c29caab20ed9cf7fcd0aada
	DoCmd(sCmdName string)
}
