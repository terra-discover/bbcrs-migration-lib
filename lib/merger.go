package lib

import (
	"encoding/json"
	"encoding/xml"
)

// Merge a struct to another struct
func Merge(from interface{}, to interface{}) error {
	j, e := json.Marshal(from)
	if nil == e {
		e = json.Unmarshal(j, to)
	}

	return e
}

func MergeXML(from interface{}, to interface{}) error {
	x, e := xml.Marshal(from)
	if nil == e {
		e = xml.Unmarshal(x, to)
	}
	return e
}
