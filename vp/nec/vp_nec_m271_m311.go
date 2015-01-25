package nec

// Note : user manual advises to lower baud rate to 9600 for long cables
import ()

type nec_m271_m311 struct {
	Json_FilePath string
	ModelName     string

	Commands map[string][]byte //PowerOn, PowerOff, SoundMuteOn, SoundMuteOff, ...
}

func (o nec_m271_m311) GetName() string {
	return o.ModelName
}

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

}
