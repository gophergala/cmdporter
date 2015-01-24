package nec

// Note : user manual advises to lower baud rate to 9600 for long cables
import "encoding/json"

type Command []byte

type nec_m271_m311 struct {
	Commands map[string]Command 'json:commands'    	//PowerOn, PowerOff, SoundMuteOn, SoundMuteOff 
}

var Nec_m271_m311 nec_m271_m311

func init() {

	//MAKE MAP
	Nec_m271_m311.Commands = make(map[string]Command)


	//ADD COMMAND TO THE MAP
	//Nec_m271_m311.Commands["PowerOn"] = []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x02}
	//Nec_m271_m311.Commands["PowerOff"] = []byte{0x02, 0x01, 0x00, 0x00, 0x00, 0x03}
	//Nec_m271_m311.Commands["SoundMuteOn"] = []byte{0x02, 0x12, 0x00, 0x00, 0x00, 0x14}
	//Nec_m271_m311.Commands["SoundMuteOff"] = []byte{0x02, 0x13, 0x00, 0x00, 0x00, 0x15}

	//CHECKING
	//fmt.Println(Nec_m271_m311)

	//TRANSFORM TO JSON
	//temp_json_Nec_m271_m311 := &nec_m271_m311{
	//	Commands: Nec_m271_m311.Commands}
	//json_Nec_m271_m311, _ := json.Marshal(temp_json_Nec_m271_m311)
	//fmt.Println(string(json_Nec_m271_m311))



	//FROM JSON FILE, WHICH INCLUDED ALL SCENARIOS TO
	type BytesContainer struct {
	StringCodedBytes []string `json:"bytes"`
	Bytes []byte
	}

	var err error
	res := &BytesContainer{}
	err = json.Unmarshal([]byte(json_string), &res)
	if err != nil {
		fmt.Println("err :", err)
		os.Exit(-1)
	}

	for _, StringCodedByte := range res.StringCodedBytes {
		var v []byte
		v, err = hex.DecodeString(StringCodedByte[2:])
		fmt.Printf("Byte %d\n", v)

		if err != nil {
			fmt.Println("err :", err)
			os.Exit(-1)
		}

		res.Bytes = append(res.Bytes, v[0])
	}
	fmt.Println(res.Bytes) //v[len(v)-2:])
	fmt.Println(res)



}


