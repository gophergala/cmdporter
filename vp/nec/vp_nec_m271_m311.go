package nec

//package main

// Note : user manual advises to lower baud rate to 9600 for long cables
import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Nec_m271_m311 nec_m271_m311

var SerialPortStatus bool = false

type nec_m271_m311 struct {
	ModelName string

	Commands map[string][]byte //PowerOn, PowerOff, SoundMuteOn, SoundMuteOff, ...

}

type JSONCommand struct {
	CommandName      string
	StringCodedBytes []string `json:"bytes"`
	Bytes            []byte
}

type JSONCommands struct {
	Name     string
	Commands []JSONCommand
}

// Load json file containing device commands into device global var, eg. Nec_m271_m311 here
func init() {

	Nec_m271_m311.Commands = make(map[string][]byte)

	var err error

	// Load json file into string
	jsonbytes, err := ioutil.ReadFile("vp/nec/vp_nec_m271_m311.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(-1)
	}
	json_string := string(jsonbytes)

	// Load it into our intermediate struct containing string encoded commands either in base 10 or hexa
	var IntermediateStruct = JSONCommands{}
	err = json.Unmarshal([]byte(json_string), &IntermediateStruct)
	if err != nil {
		fmt.Println("err :", err)
		os.Exit(-1)
	}

	// Convert these string encoded commands into bytes
	for key, value := range IntermediateStruct.Commands {
		command := value
		for _, cvalue := range command.StringCodedBytes {
			// TODO check whether string encoded commands actually begins with 0x, if not then it's base 10
			cmd_bytes, err := hex.DecodeString(cvalue[2:])
			if err != nil {
				fmt.Println("err :", err)
				os.Exit(-1)
			}
			// FIX this for commands containing more than one byte
			IntermediateStruct.Commands[key].Bytes = append(IntermediateStruct.Commands[key].Bytes, cmd_bytes[0])
		}
	}

	for _, IntermediateCmd := range IntermediateStruct.Commands {
		Nec_m271_m311.Commands[IntermediateCmd.CommandName] = IntermediateCmd.Bytes
	}

	Nec_m271_m311.ModelName = IntermediateStruct.Name

	fmt.Printf("Loaded %d commands for %s\n", len(Nec_m271_m311.Commands), Nec_m271_m311.ModelName)

}

func (o nec_m271_m311) Load() {

}
