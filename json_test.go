package objecttext_test

import (
	"encoding/json"
	"testing"

	objecttext "github.com/go-marshaltemabu/go-objecttext"
)

type dataInChecked struct {
	TargetValue objecttext.CheckedObjectText `json:"target_v"`
}

type dataInUncheck struct {
	TargetValue objecttext.UncheckObjectText `json:"target_v"`
}

type dataOutString struct {
	TargetString string `json:"target_v"`
}

type dataOutMap struct {
	TargetMap map[string]interface{} `json:"target_v"`
}

type dataOutABCDEF struct {
	Target struct {
		ABC int32 `json:"abc"`
		DEF int32 `json:"def"`
	} `json:"target_v"`
}

const (
	sampleJSONabcdef       = "{\"abc\":123,\"def\":789}"
	sampleTargetJSONabcdef = "{\"target_v\":{\"abc\":123,\"def\":789}}"
)

func verifyTargetOutputJSONabcdef(buf []byte, t *testing.T) {
	t.Logf("JSON: %s", string(buf))
	var d1 dataOutABCDEF
	if err := json.Unmarshal(buf, &d1); nil != err {
		t.Fatalf("json.Unmarshal (string) failed: %v", err)
		return
	}
	if (d1.Target.ABC != 123) || (d1.Target.DEF != 789) {
		t.Errorf("unexpect result: %#v", &d1)
	}
	var d2 dataOutMap
	if err := json.Unmarshal(buf, &d2); nil != err {
		t.Fatalf("json.Unmarshal (map) failed: %v", err)
		return
	}
	if vABC, ok := d2.TargetMap["abc"].(float64); (!ok) || (vABC != 123) {
		t.Errorf("unexpected result: ok=%v, vABC=%v, map=%v",
			ok, vABC, d2.TargetMap)
	}
	if vDEF, ok := d2.TargetMap["def"].(float64); (!ok) || (vDEF != 789) {
		t.Errorf("unexpected result: ok=%v, vDEF=%v, map=%v",
			ok, vDEF, d2.TargetMap)
	}
}

func verifyTargetSampleJSONabcdef(buf []byte, t *testing.T) {
	t.Logf("JSON: %s", string(buf))
	var d1 struct {
		ABC int32 `json:"abc"`
		DEF int32 `json:"def"`
	}
	if err := json.Unmarshal(buf, &d1); nil != err {
		t.Fatalf("json.Unmarshal (string) failed: %v", err)
		return
	}
	if (d1.ABC != 123) || (d1.DEF != 789) {
		t.Errorf("unexpect result: %#v", &d1)
	}
}

const sampleJSONabcdefxyz = "{\"abc\":123,\"def\":789}xyz"

const sampleTargetJSONempty = "{\"target_v\":{}}"

func TestJSONChecked_01(t *testing.T) {
	aux1 := dataInChecked{
		TargetValue: sampleJSONabcdef,
	}
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	verifyTargetOutputJSONabcdef(buf, t)
}

func TestJSONChecked_02(t *testing.T) {
	aux1 := dataInChecked{
		TargetValue: sampleJSONabcdefxyz,
	}
	_, err := json.Marshal(&aux1)
	if nil == err {
		t.Fatal("expect json.Marshal fail")
		return
	}
}

func TestJSONChecked_03(t *testing.T) {
	var aux1 dataInChecked
	if err := json.Unmarshal([]byte(sampleTargetJSONabcdef), &aux1); nil != err {
		t.Fatalf("json.Unmarshal failed: %v", err)
		return
	}
	t.Logf("JSONText: [%v]", aux1.TargetValue)
	verifyTargetSampleJSONabcdef(([]byte)(aux1.TargetValue), t)
}

func TestJSONChecked_04(t *testing.T) {
	aux1 := dataInChecked{
		TargetValue: "",
	}
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	s := string(buf)
	t.Logf("JSON: [%v]", s)
	if s != sampleTargetJSONempty {
		t.Errorf("unexpect result: [%s]", s)
	}
}

func TestJSONUncheck_01(t *testing.T) {
	aux1 := dataInUncheck{
		TargetValue: sampleJSONabcdef,
	}
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	verifyTargetOutputJSONabcdef(buf, t)
}

func TestJSONUncheck_02(t *testing.T) {
	aux1 := dataInUncheck{
		TargetValue: sampleJSONabcdefxyz,
	}
	_, err := json.Marshal(&aux1)
	if nil == err {
		t.Fatal("expect json.Marshal fail")
		return
	}
	t.Logf("fail as expect: %v", err)
}

func TestJSONUncheck_03(t *testing.T) {
	var aux1 dataInUncheck
	if err := json.Unmarshal([]byte(sampleTargetJSONabcdef), &aux1); nil != err {
		t.Fatalf("json.Unmarshal failed: %v", err)
		return
	}
	t.Logf("JSONText: [%v]", aux1.TargetValue)
	verifyTargetSampleJSONabcdef(([]byte)(aux1.TargetValue), t)
}

func TestJSONUncheck_04(t *testing.T) {
	aux1 := dataInUncheck{
		TargetValue: "",
	}
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	s := string(buf)
	t.Logf("JSON: [%v]", s)
	if s != sampleTargetJSONempty {
		t.Errorf("unexpect result: [%s]", s)
	}
}
