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
	"github.com/gophergala/cmdporter/vp/nec"
	"github.com/tarm/goserial"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"text/template"
)

var SerialPortStatus bool = false

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

	//On Linux
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	//On Macos
	//c := &serial.Config{Name: "/dev/cu.PL2303-00002014", Baud: 9600}

	s, err := serial.OpenPort(c)

	if err != nil {
		log.Println(err)

	} else {
		SerialPortStatus = true
		n, err := s.Write(nec.Nec_m271_m311.PowerOn)

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(n)
	}

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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
