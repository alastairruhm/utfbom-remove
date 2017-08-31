package main

import (
	"reflect"
	"testing"
)

func TestRemoveUtfBom(t *testing.T) {
	var testCases = []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("\xEF\xBB\xBFhello"), []byte("hello")},
		{[]byte("\xEFhello"), []byte("\xEFhello")},
		{[]byte("hello"), []byte("hello")},
	}

	for _, tt := range testCases {
		result, err := RemoveUtfBom(tt.input)
		if err != nil {
			t.Errorf("%v", err)
		}

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v, but is:%v\n", tt.expected, result)
		}
	}

}
