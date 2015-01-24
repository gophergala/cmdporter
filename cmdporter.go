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
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/gophergala/cmdporter/vp/nec"
	"github.com/tarm/goserial"
	"log"
)

func main() {

	//On Linux
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	//On Macos
	//c := &serial.Config{Name: "/dev/cu.PL2303-00002014", Baud: 9600}

	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write(nec.Nec_m271_m311.PowerOn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(n)

	// Start Http Server
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Extensions: []string{".html"},
	}))

	m.Get("/", func(r render.Render) {
		content := map[string]interface{}{"slogan": "Gopher is coming"}

		r.HTML(200, "index", content)
	})

	m.Run()
}
