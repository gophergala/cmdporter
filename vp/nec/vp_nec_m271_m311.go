package nec

// Note : user manual advises to lower baud rate to 9600 for long cables
<<<<<<< HEAD
import ()

type nec_m271_m311 struct {
	Json_FilePath string
	ModelName     string
=======
import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type nec_m271_m311 struct {
	ModelName string
>>>>>>> e7bc7230bb514cdc1c29caab20ed9cf7fcd0aada

	Commands map[string][]byte //PowerOn, PowerOff, SoundMuteOn, SoundMuteOff, ...

}

func (o nec_m271_m311) GetName() string {
	return o.ModelName
}

<<<<<<< HEAD
func (o nec_m271_m311) SetName(name string) {
	o.ModelName = name
}

func (o nec_m271_m311) RegisterCmd(sCmdName string, Bytes []byte) {
	o.Commands[sCmdName] = Bytes
}

func (o nec_m271_m311) GetNumCommands() int {
	return len(o.Commands)
}

func (o nec_m271_m311) DoCmd(sCmdName string) []byte {
	return o.Commands[sCmdName]
}

func (o nec_m271_m311) GetJsonPath() string {
	return o.Json_FilePath
}

var Nec_m271_m311 nec_m271_m311

// Load json file containing device commands into device global var, eg. Nec_m271_m311 here
func init() {

	Nec_m271_m311.Commands = make(map[string][]byte)

	Nec_m271_m311.Json_FilePath = "vp/nec/vp_nec_m271_m311.json"
=======
func (o nec_m271_m311) RegisterCmd(sCmdName string, Bytes []byte) {
	o.Commands[sCmdName] = Bytes
}

func (o nec_m271_m311) DoCmd(sCmdName string) {
	// TODO
}

var Nec_m271_m311 nec_m271_m311

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
>>>>>>> e7bc7230bb514cdc1c29caab20ed9cf7fcd0aada

}
