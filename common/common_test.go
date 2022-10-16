package common

import (
	"testing"
)

func Test_JsonStringToStruct(t *testing.T) {
	s := `{"Name": "Platypus"}`
	type person struct {
		Name string `json:"Name"`
	}
	var p person
	JsonStringToStruct(s, &p)
	if p.Name != "Platypus" {
		t.Errorf("Wanted Platypus but got %s", p.Name)
	}
}
