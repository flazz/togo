package the

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestThePdf(t *testing.T) {
	b, err := ioutil.ReadFile("The.pdf")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ThePdf, b) {
		t.Fail()
	}
}
