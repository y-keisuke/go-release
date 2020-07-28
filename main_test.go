package main

import "testing"

func TestReturnStr(t *testing.T) {
	testStr := "testStr"
	actual := returnStr(testStr)
	if actual != testStr {
		t.Error("エラー")
	}
}
