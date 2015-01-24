package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type JSONCommands struct {
	Commands []JSONCommand
}

type JSONCommand struct {
	CommandName      string
	StringCodedBytes []string `json:"bytes"`
	Bytes            []byte
}

// type BytesContainer struct {
// 	StringCodedBytes []string `json:"bytes"`
// 	Bytes            []byte
// }

func main() {

	json_string :=
		`{
			"commands": [
			{
				"CommandName":"PowerOn",
		 		"bytes": [ "0x31", "0x02","0x02", "0x00", "0x00", "0x00", "0x00", "0x02" ]
			},
			{
				"CommandName":"PowerOff",
		 		"bytes": ["0x02", "0x01", "0x00", "0x00", "0x00", "0x03"]
			}
			]
		}`

	var err error

	//res := &JSONCommands{}

	var res = JSONCommands{}

	err = json.Unmarshal([]byte(json_string), &res)
	if err != nil {
		fmt.Println("err :", err)
		os.Exit(-1)
	}

	for key, value := range res.Commands {
		command := value
		for _, cvalue := range command.StringCodedBytes {
			chex, err := hex.DecodeString(cvalue[2:])
			if err != nil {
				panic(err)
			}
			res.Commands[key].Bytes = append(res.Commands[key].Bytes, chex[0])
		}
	}

}
