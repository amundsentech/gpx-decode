package gpxdecode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	//testing datasets
	points3D = "tests/testpoints.gpx"
	lines    = "tests/testlines.gpx"
	lines3D  = "tests/testlines with elevation.gpx"
)

func TestGPX2Struct(t *testing.T) {

	// build a map of the testing data and inputs
	data := make(map[string]string)
	data[points3D] = points3D
	//data[lines] = lines
	//data[lines3D] = lines3D

	for item, fileloc := range data {

		// psa
		fmt.Printf("Decoder is starting in on: %v\n", item)

		// prase the inputDetails from origin
		gpxfile, err := os.Open(fileloc)
		if err != nil {
			t.Errorf(err.Error())
		}

		// load file as byte
		gpxbyte, _ := ioutil.ReadAll(gpxfile)
		gpxbuf := bytes.NewBuffer(gpxbyte)

		// this is the test!!!
		var gpx GPX

		GPXDecode(gpxbuf, &gpx)

		gpxjson, err := json.Marshal(gpx)
		if err != nil {
			fmt.Println("error:", err)
		}

		//fmt.Printf("Inbound reads like:\n%v\n", string(gpxbyte))
		// what does it look like?
		fmt.Printf("GPX struct as json reads:\n%v\n", string(gpxjson))
	}
}
