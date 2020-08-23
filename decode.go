package gpxdecode

import (
	"bytes"
	"encoding/xml"
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
	//Extensions   []Ogr	    `xml:"ogr,attr" json:"type,attr"`
	Name         string         `xml:"name" json:"name"`
	Description  string         `xml:"desc" json:"desc"`
	TrackSegment []TrackSegment `xml:"trkseg" json:"trksed"`
}

type Route struct {
	//Extensions  Extensions   `xml:"extensions" json:"extensions"`
	Name        string       `xml:"name" json:"name"`
	Description string       `xml:"desc" json:"desc"`
	RoutePoints []RoutePoint `xml:"rtept" json:"rtept"`
}

type RoutePoint struct {
	Lat float64 `xml:"lat,attr" json:"lat,attr"`
	Lon float64 `xml:"lon,attr" json:"lon,attr"`
	Ele float64 `xml:"ele" json:"ele"`
}

type Waypoint struct {
	Extensions  Extensions `xml:"extensions" json:"extensions"`
	Name        string     `xml:"name" json:"name"`
	Description string     `xml:"desc" json:"desc"`
	Lat         float64    `xml:"lat,attr" json:"lat,attr"`
	Lon         float64    `xml:"lon,attr" json:"lon,attr"`
	Ele         float64    `xml:"ele" json:"ele"`
}

type Extensions struct {
	OGR string `xml:"ogr distance,attr" json:"ogr distance,attr"`
}

type OGR struct {
	//Key   string `xml:" ,attr" json:" ,attr"`
	Value string `xml:"ogr distance,chardata" json:"ogr distance,chardata"`
}

func GPXDecode(f *bytes.Buffer, gpx *GPX) {

	// xml.Unmarshal(f, kml)
	d := xml.NewDecoder(f)

	// gpx encodes some data as namespace eg. <ogr:key>value</ogr:key>
	d.DefaultSpace = "_"

	// place all the xml where it belongs
	d.Decode(gpx)
}
