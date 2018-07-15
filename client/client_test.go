package main

import (
	"testing"
	"strconv"
	)

func TestClient(t *testing.T) {
	
	test1 := addr_check("127.0.0.1:8000")
	test1Result := true

	if test1 != test1Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test1), strconv.FormatBool(test1Result))
	}

	test2 := addr_check("127.0.0.1:80ss")
	test2Result := false

	if test2 != test2Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test2), strconv.FormatBool(test2Result))
	}

	test3 := addr_check("-23")
	test3Result := false

	if test3 != test3Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test3), strconv.FormatBool(test3Result))
	}

	test4 := addr_check("sadfasdf")
	test4Result := false

	if test4 != test4Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test4), strconv.FormatBool(test4Result))
	}

	test5 := input_check("sadfasdf")
	test5Result := false

	if test5 != test5Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test5), strconv.FormatBool(test5Result))
	}

	test6 := input_check("-23")
	test6Result := false

	if test6 != test6Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test6), strconv.FormatBool(test6Result))
	}

	test7 := input_check("3.14")
	test7Result := false

	if test7 != test7Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test7), strconv.FormatBool(test7Result))
	}

	test8 := input_check("34")
	test8Result := true

	if test8 != test8Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test8), strconv.FormatBool(test8Result))
	}

}