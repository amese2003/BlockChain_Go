package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"

	s := struct{ Test string }{Test: "test"}

	t.Run("해시는 항상 같습니다.", func(t *testing.T) {
		x := Hash(s)
		if x != hash {
			t.Errorf("예상값 : %s, 실제값 : %s", hash, x)
		}
	})

	t.Run("해시는 hex 인코딩", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Errorf("해쉬값이 hex 인코딩값이 아닙니다.")
		}
	})
}

func ExampleHash() {
	s := struct{ Test string }{Test: "test"}
	x := Hash(s)
	fmt.Println(x)
	// Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("ToBytes should return a slice of bytes got %s", k)
	}
}

func TestSpliter(t *testing.T) {
	type test struct {
		input  string
		sep    string
		index  int
		output string
	}

	tests := []test{
		{input: "0:16:0", sep: ":", index: 1, output: "16"},
		{input: "0:6:0", sep: ":", index: 10, output: ""},
		{input: "0:6:0", sep: "/", index: 0, output: "0:6:0"},
	}

	for _, tc := range tests {
		got := Splitter(tc.input, tc.sep, tc.index)
		if got != tc.output {
			t.Errorf("Expected %s and got %s", tc.output, got)
		}
	}
}

func TestHandleError(t *testing.T) {
	oldLogFn := logfn
	defer func() {
		logfn = oldLogFn
	}()

	called := false
	logfn = func(v ...interface{}) {
		called = true
	}

	err := errors.New("test")
	HandleError(err)
	if !called {
		t.Error("HandleError should call fn")
	}
}
