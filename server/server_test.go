package main

import (
	"testing"
	"strconv"
	)

func TestPrimeCheck(t *testing.T) {
	test1 := prime_check("1")
	test1Result := false

	if test1 != test1Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test1), strconv.FormatBool(test1Result))
	}

	test2 := prime_check("3")
	test2Result := true

	if test2 != test2Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test2), strconv.FormatBool(test2Result))
	}

	test3 := prime_check("4")
	test3Result := false

	if test3 != test3Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test3), strconv.FormatBool(test3Result))
	}

	test4 := prime_check("3389")
	test4Result := true

	if test4 != test4Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test4), strconv.FormatBool(test4Result))
	}

	test5 := prime_check("12")
	test5Result := false

	if test5 != test5Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test5), strconv.FormatBool(test5Result))
	}

	test6 := prime_check("390077316841")
	test6Result := true

	if test6 != test6Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test6), strconv.FormatBool(test6Result))
	}

	test7 := prime_check("-2")
	test7Result := false

	if test7 != test7Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test6), strconv.FormatBool(test6Result))
	}

	test8 := prime_check("shfisudhf")
	test8Result := false

	if test8 != test8Result {
		t.Errorf("FAIL! Got '%s', result should've been '%s'", strconv.FormatBool(test6), strconv.FormatBool(test6Result))
	}
}