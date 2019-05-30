package apple

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}

func TestGetAppById(t *testing.T) {
	input := 1222530780
	want := &Response{}

	err := json.Unmarshal(loadFixture("result.txt"), want)
	if err != nil {
		panic(err)
	}
	got, err := GetAppByID(input)
	if err != nil {
		t.Errorf("GetAppById returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAppById= %+v, want %+v", got, want)
	}
}

func TestGetAppByInvalidId(t *testing.T) {
	input := 0
	want := &Response{
		ResultCount: 0,
		Results:     []Result{},
	}
	got, err := GetAppByID(input)
	if err != nil {
		t.Errorf("GetAppById returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAppById= %+v, want %+v", got, want)
	}

}
