package requel_test

import (
	"reflect"
	"testing"

	"github.com/s1ntaxe770r/requel"
)

func TestRequestTimer(T *testing.T) {
	expectedtime := requel.TimeReq()
	var expectedtype float64
	if reflect.TypeOf(expectedtime) != reflect.TypeOf(expectedtype) {
		T.Errorf("time retruned the wrong value expected %v got %v", expectedtype, expectedtime)
	}
}


