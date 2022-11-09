package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/parser", parser)

	log.Println("Staring server")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Println("Error when staring server")
		panic(err)
	}
	log.Println("Terminating server")
}

type Payload struct {
	Key        string   `json:"key"`
	Properties []string `json:"properties"`
}

func parser(w http.ResponseWriter, r *http.Request) {
	payload := getPayload(r)
	fmt.Println(payload)
	xml := toXml(payload)

	io.WriteString(w, xml)
}

func getPayload(r *http.Request) Payload {
	var payload Payload

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Error when reading body from request")
		panic(err)
	}

	err = json.Unmarshal(body, &payload)

	if err != nil {
		log.Println("Error when reading unmarshaing body")
		panic(err)
	}

	return payload
}

func toXml(payload Payload) string {
	valuesXml, err := xml.Marshal(&payload)

	if err != nil {
		log.Println("Error when marshaing to xml")
		panic(err)
	}

	xml := buildXml(valuesXml)

	return string(xml)
}

func buildXml(values []byte) []byte {
	writeToFile(values)
	xml := applyValuesOnTemplate()

	return xml
}

func writeToFile(content []byte) {
	err := os.WriteFile("values.xml", content, 0644)

	if err != nil {
		log.Println("Error when writing to file")
		panic(err)
	}
}

func applyValuesOnTemplate() []byte {
	//xsltproc template.xsl values.xml
	xsltproc := exec.Command("xsltproc", "template.xsl", "values.xml")
	xml, err := xsltproc.Output()

	if err != nil {
		log.Println("Error when appling template")
		panic(err)
	}

	return xml
}
