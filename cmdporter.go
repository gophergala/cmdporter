package main

/* ====================================================================================================

cmdporter : a wifi intercom to talk to various devices

By Fred Ménez & Gaël Reyrol

==================================================================================================== */

/* TODO Serial

x looks for serial device depending on OS (Macos, Linux)
x discover serial device or read configuration
x load commands params from file

*/

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gophergala/cmdporter/vp/nec"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

var (
	SerialPortStatus bool = false
	g_Device         Device
)

func Render(w http.ResponseWriter, view string, content interface{}) {
	layout, err := ioutil.ReadFile(path.Join("views", "layout.html"))
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadFile(path.Join("views", view))
	if err != nil {
		log.Fatal(err)
	}

	layoutTemplate := template.New("layout")
	pageTemplate := template.New("page")

	template.Must(layoutTemplate.Parse(string(layout)))
	template.Must(pageTemplate.Parse(string(page)))

	pageBuffer := new(bytes.Buffer)
	pageTemplate.Execute(pageBuffer, content)

	layoutContent := map[string]interface{}{"View": string(pageBuffer.Bytes())}
	layoutTemplate.Execute(w, layoutContent)

}

func main() {

	g_Device = nec.Nec_m271_m311
	LoadCommands(g_Device)

	devices := []string{
		"Nec mg271wg",
		"Arduino One",
	}

	// Start Http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		content := map[string]interface{}{
			"SerialPortStatus": SerialPortStatus,
			"Devices":          devices,
		}

		Render(w, "index.html", content)
	})

	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		content := map[string]interface{}{
			"Devices": devices,
		}

		Render(w, "devices.html", content)
	})

	http.HandleFunc("/device/", func(w http.ResponseWriter, r *http.Request) {
		content := map[string]interface{}{
			"Name": r.URL.Path[8:],
		}

		Render(w, "device.html", content)
	})

	http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

		}
		w.WriteHeader(http.StatusNotFound)
	})

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("Running for device", g_Device.GetName())
	log.Println("Waiting for http connections on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func LoadCommands(d Device) {
	var err error

	// Load json file into string
	jsonbytes, err := ioutil.ReadFile(d.GetJsonPath())
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

	//CREATE A MAPPING FOR THE nec_m271_m311 COMMANDS
	for _, IntermediateCmd := range IntermediateStruct.Commands {
		d.RegisterCmd(IntermediateCmd.CommandName, IntermediateCmd.Bytes)
	}
	log.Println("test n :", IntermediateStruct.Name)
	d.SetName(IntermediateStruct.Name)

	log.Printf("Loaded %d commands for %s\n", d.GetNumCommands(), d.GetName())
}

type JSONCommands struct {
	Name     string
	Commands []JSONCommand
}

type JSONCommand struct {
	CommandName      string
	StringCodedBytes []string `json:"bytes"`
	Bytes            []byte
}
