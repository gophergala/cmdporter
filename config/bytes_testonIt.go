package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {

	var err error

	//IMPORT FROM A JSON FILE
	file, err := ioutil.ReadFile("./commands.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(-1)
	}
	json_string := string(file)

	//TRANSFORM THE JSON TO STUCT DATA
	var res = JSONCommands{}
	err = json.Unmarshal([]byte(json_string), &res)
	if err != nil {
		fmt.Println("err :", err)
		os.Exit(-1)
	}

	//CONVERT THE STRING DATA TO BYTE DATA, USING FOR SEND INSTRUCTIONS TO HARDWARE DEVICES
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

	//CREATE A MAPPING FOR THE COMMANDS
	nec_m271_m311_Commands := make(map[string][]byte)

	for _, value2 := range res.Commands {
		command := value2
		nec_m271_m311_Commands[command.CommandName] = command.Bytes
	}
	fmt.Println(nec_m271_m311_Commands["SoundMuteOn"])

}
