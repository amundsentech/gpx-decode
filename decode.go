package gpxdecode

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type GPX struct {
	XMLName  xml.Name   `xml:"gpx" json:"gpx"`
	Version  string     `xml:"version,attr" json:"version,attr"`
	Creator  string     `xml:"creator,attr" json"creator,attr"`
	Time     string     `xml:"time" json:"time"`
	Metadata Metadata   `xml:"metadata" json:"metadata"`
	Waypoint []Waypoint `xml:"wpt,omitempty" json:"wpt,omitempty"`
	Track    []Track    `xml:"trk,omitempty" json:"trk,omitempty"`
	Route    []Route    `xml:"rte,omitempty" json:"rte,omitempty"`
}

type Metadata struct {
	Boundary []Bounds `xml:"bounds" json:"bounds"`
	Keywords string   `xml:"keywords,omitempty"`
}

type Bounds struct {
	Minlat float64 `xml:"minlat,attr" json:"minlat,attr"`
	Minlon float64 `xml:"minlon,attr" json:"minlon,attr"`
	Maxlat float64 `xml:"maxlat,attr" json:"maxlat,attr"`
	Maxlon float64 `xml:"maxlon,attr" json:"maxlon,attr"`
}

type TrackPoint struct {
	Lat  float64   `xml:"lat,attr" json:"lat,attr"`
	Lon  float64   `xml:"lon,attr" json:"lon,attr"`
	Ele  float64   `xml:"ele" json:"ele"`
	Time time.Time `xml:"time" json:"time"`
}

type TrackSegment struct {
	TrackPoint []TrackPoint `xml:"trkpt" json:"trkpt"`
}

type Track struct {
	Extensions   Extensions	    `xml:"extensions" json:"extensions"`
	Name         string         `xml:"name,omitempty" json:"name,omitempty"`
	Description  string         `xml:"desc,omitempty" json:"desc,omitempty"`
	TrackSegment []TrackSegment `xml:"trkseg" json:"trksed"`
}

type Route struct {
	Extensions  Extensions   `xml:"extensions" json:"extensions"`
	Name        string       `xml:"name,omitempty" json:"name,omitempty"`
	Description string       `xml:"desc,omitempty" json:"desc,omitempty"`
	RoutePoints []RoutePoint `xml:"rtept" json:"rtept"`
}

type RoutePoint struct {
	Lat float64 `xml:"lat,attr" json:"lat,attr"`
	Lon float64 `xml:"lon,attr" json:"lon,attr"`
	Ele float64 `xml:"ele,omitempty" json:"ele,omitempty"`
}

type Waypoint struct {
	Extensions  Extensions `xml:"extensions" json:"extensions"`
	Name        string     `xml:"name,omitempty" json:"name,omitempty"`
	Description string     `xml:"desc,omitempty" json:"desc,omitempty"`
	Lat         float64    `xml:"lat,attr" json:"lat,attr"`
	Lon         float64    `xml:"lon,attr" json:"lon,attr"`
	Ele         float64    `xml:"ele,omitempty" json:"ele,omitempty"`
}

type Extensions struct {
	//XMLName  xml.Name   `xml:"extensions" json:"extensions"`
	OGR []OGR `xml:",any" json:",any"`
}

type OGR struct {
	//XMLName  xml.Name   `xml:"ogr,any" json:"ogr,any"`
	Key   string `xml:"Key" json:"Key"`
	Value string `xml:",chardata" json:",chardata"`
}

func GPXDecode (f *bytes.Buffer, gpx *GPX) {

	xmlraw, _ := ioutil.ReadAll(f)

	// First, build as much as we can of the GPX Struct
	xmlbytes := bytes.NewBuffer(xmlraw)
	d := xml.NewDecoder(xmlbytes)
	d.Decode(gpx)

	// Second, rebuild a buffer to just parse for extensions
	xmlbytes = bytes.NewBuffer(xmlraw)
	d = xml.NewDecoder(xmlbytes)

	// set a boolean extension tracker
	isExt := 0

	// set counter for tracking # of features
	elementNum := -1
	extNum := 0

	// parse each token, looking for 'extensions'
	for {
		t, _ := d.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {

		case xml.StartElement:
			// the next feature coming after this one will be an extension
			if strings.Contains(se.Name.Local,"extensions") {
				isExt = 1
				elementNum  ++
				extNum = 0
				continue
			}

			// Reached a new feature, so Reset extension counter and isExt boolean
                        nonExt := []string{"metadata","wpt","trkpt","rtept","trkseg","lon","lat","ele","name","desc"}
                        for _, substr := range nonExt {
                                if strings.Contains(se.Name.Local,substr) {
                                        isExt = 0
                                }
                        }

			// this feature is flagged as an extension (falls after prior)
			if isExt == 1 {
				var ogr OGR
				ogr.Key = se.Name.Local

				var frag []byte
				if err := d.DecodeElement(&frag, &se); err != nil {
					fmt.Printf("Error decoding token: %v\n", err.Error())
				}
				ogr.Value = string(frag)

				// is Point only
				if gpx.Waypoint != nil && len(gpx.Waypoint) >= elementNum {
					gpx.Waypoint[elementNum].Extensions.OGR[extNum].Key = ogr.Key
				}

				// is Track
				if gpx.Track != nil && len(gpx.Track[elementNum].Extensions.OGR) >= extNum {
					gpx.Track[elementNum].Extensions.OGR[extNum].Key = ogr.Key
				}

				// is Route
				if gpx.Route != nil && len(gpx.Route[elementNum].Extensions.OGR) >= extNum {
					gpx.Route[elementNum].Extensions.OGR[extNum].Key = ogr.Key
				}

				// increase the counter
				extNum ++
			}

		}

	}
}
