// Copyright 2020 NGR Softlab
//
package ngr_math

import (
	"testing"
)

func TestDoMathFunc(t *testing.T) {

	for k, _ := range Functions {
		res, err := DoMathFunc(FuncParams{
			Start:    -1,
			End:      1,
			Func:     k,
			Addition: 1,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	}

}

/////////////////////////////////////////////////
/////////////////////////////////////////////////
/////////////////////////////////////////////////

/////////////////////////////////////////////////
func BenchmarkTest(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := makeSqrtX(FuncParams{
			Start:    -1,
			End:      1,
			Addition: 1,
		})
		if err != nil {
			b.Error(err)
		}
	}

}

/////////////////////////////////////////////////
func BenchmarkTest1(b *testing.B) {
	//create upgraded copy of math func from BenchmarkTest and test it here
}
