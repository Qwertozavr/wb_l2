package main

import "testing"

// given test 1
// - "a4bc2d5e" => "aaaabccddddde"
func Test1(t *testing.T) {
	str, err := Unpack("a4bc2d5e")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "aaaabccddddde" {
		t.Errorf("Test 1: a4bc2d5e\nExpected %s, got %s", "aaaabccddddde", str)
	}
}

// given test 2
// - "abcd" => "abcd"
func Test2(t *testing.T) {
	str, err := Unpack("abcd")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "abcd" {
		t.Errorf("Test 2: abcd\nExpected %s, got %s", "abcd", str)
	}
}

// given test 3
// - "45" => ""
func Test3(t *testing.T) {
	str, err := Unpack("45")
	if (err != nil) != true {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "" {
		t.Errorf("Test 3: 45\nExpected %s (empty string), got %s", "", str)
	}
}

// given test 4
// - "" => ""
func Test4(t *testing.T) {
	str, err := Unpack("")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "" {
		t.Errorf("Test 4: (empty string)\nExpected %s (empty string), got %s", "", str)
	}
}

// given test 5
// - "qwe\4\5" => "qwe45"
func Test5(t *testing.T) {
	str, err := Unpack("qwe\\4\\5")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "qwe45" {
		t.Errorf("Test 5: qwe\\4\\5\nExpected %s, got %s", "qwe45", str)
	}
}

// given test 6
// - "qwe\45" => "qwe44444"
func Test6(t *testing.T) {
	str, err := Unpack("qwe\\45")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "qwe44444" {
		t.Errorf("Test 6: qwe\\45\nExpected %s, got %s", "qwe44444", str)
	}
}

// given test 7
// - "qwe\5" => "qwe\\\\"
func Test7(t *testing.T) {
	str, err := Unpack("qwe\\\\5")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "qwe\\\\\\\\\\" {
		t.Errorf("Test 7: qwe\\\\5\nExpected %s, got %s", "qwe\\\\\\\\\\", str)
	}
}

// custom test to ensure that escaped 'a'-'z' are considered as a wrong expression
// - "a\\b" => ""
func Test8(t *testing.T) {
	str, err := Unpack("a\\b")
	if (err != nil) != false {
		t.Errorf("Unpack() error = %v, wantErr = %v", err, false)
		return
	}
	if str != "ab" {
		t.Errorf("Test 8: a\\b\nExpected %s (empty string), got %s", "", str)
	}
}
