package objecttext

import (
	"encoding/json"
)

const emptyJSONObjectText = "{}"

// RemarshalJSON decode and encode given string v as string map.
func RemarshalJSON(v string) (result string, err error) {
	if v == "" {
		return
	}
	var aux map[string]interface{}
	if err = json.Unmarshal([]byte(v), &aux); nil != err {
		return
	}
	if len(aux) == 0 {
		return
	}
	buf, err := json.Marshal(aux)
	if nil != err {
		return
	}
	result = string(buf)
	return
}

// MarshalJSON implements the json.Marshaler interface.
func (v UncheckObjectText) MarshalJSON() ([]byte, error) {
	if v == "" {
		return ([]byte)(emptyJSONObjectText), nil
	}
	return ([]byte)(v), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *UncheckObjectText) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == emptyJSONObjectText {
		*v = ""
		return nil
	} else if s == "null" {
		return nil
	}
	*v = (UncheckObjectText)(s)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (v CheckedObjectText) MarshalJSON() ([]byte, error) {
	if v == "" {
		return ([]byte)(emptyJSONObjectText), nil
	}
	aux, err := RemarshalJSON(string(v))
	return ([]byte)(aux), err
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *CheckedObjectText) UnmarshalJSON(data []byte) error {
	result, err := RemarshalJSON(string(data))
	if nil != err {
		return err
	}
	*v = (CheckedObjectText)(result)
	return nil
}
