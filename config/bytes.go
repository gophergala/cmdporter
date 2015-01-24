package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type BytesContainer struct {
	StringCodedBytes []string `json:"bytes"`
	Bytes            []byte
}

func main() {

	json_string :=
		`{
		 	"bytes" : [ "0x31", "0x02","0x02", "0x00", "0x00", "0x00", "0x00", "0x02" ]
		 }`

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

	//var v []byte

	//v, err = hex.DecodeString(res.Bytes[0][2:])
	//if err != nil {
	//	fmt.Println("err :", err)
	//		os.Exit(-1)
	//}

	//fmt.Println("Byte ", res.Bytes[0])
	//fmt.Printf("Byte %d\n", v)
	//fmt.Printf("%x", res.Bytes[0])

	fmt.Println(res)

}
