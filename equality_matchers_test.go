// Copyright 2011 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ogletest

import (
	"testing"
)

var someInt int = -17

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type testCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func checkTestCases(t *testing.T, matcher Matcher, cases []testCase) {
	for i, c := range cases {
		result, err := matcher.Matches(c.candidate)

		if result != c.expectedResult {
			t.Errorf(
				"Case %d (candidate %v): expected %v, got %v",
				i,
				c.candidate,
				c.expectedResult,
				result)
		}

		if err != c.expectedError {
			t.Errorf("Case %d: expected error %v, got %v", i, c.expectedError, err)
		}
	}
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func TestNegativeIntegerLiteral(t *testing.T) {
	// -2^30
	matcher := Equals(-1073741824)
	desc := matcher.Description()
	expectedDesc := "-1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegerLiteral(t *testing.T) {
	// 2^30
	matcher := Equals(1073741824)
	desc := matcher.Description()
	expectedDesc := "1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// Floating point literals
////////////////////////////////////////////////////////////

func TestNegativeIntegralFloatingPointLiteral(t *testing.T) {
	// -2^30
	matcher := Equals(-1073741824.0)
	desc := matcher.Description()
	expectedDesc := "-1073741824.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(float64(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloatingPointLiteral(t *testing.T) {
	// 2^30
	matcher := Equals(1073741824.0)
	desc := matcher.Description()
	expectedDesc := "1073741824.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(float64(1073741824)), MATCH_TRUE, ""},

		// Values that would be 1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{uint(1073741823), MATCH_FALSE, ""},
		testCase{uint32(1073741823), MATCH_FALSE, ""},
		testCase{uint64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonIntegralFloatingPointLiteral(t *testing.T) {
	matcher := Equals(17.5)
	desc := matcher.Description()
	expectedDesc := "17.5"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 17.1.
		testCase{17.1, MATCH_TRUE, ""},
		testCase{17.1, MATCH_TRUE, ""},
		testCase{17.1 + 0i, MATCH_TRUE, ""},
		testCase{float32(17.1), MATCH_TRUE, ""},
		testCase{float64(17.1), MATCH_TRUE, ""},
		testCase{complex64(17.1), MATCH_TRUE, ""},
		testCase{complex128(17.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{17, MATCH_FALSE, ""},
		testCase{17.2, MATCH_FALSE, ""},
		testCase{18, MATCH_FALSE, ""},
		testCase{int(17), MATCH_FALSE, ""},
		testCase{int(18), MATCH_FALSE, ""},
		testCase{int32(17), MATCH_FALSE, ""},
		testCase{int64(17), MATCH_FALSE, ""},
		testCase{uint(17), MATCH_FALSE, ""},
		testCase{uint32(17), MATCH_FALSE, ""},
		testCase{uint64(17), MATCH_FALSE, ""},
		testCase{complex128(17.1 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int8
////////////////////////////////////////////////////////////

func TestNegativeInt8(t *testing.T) {
	matcher := Equals(int8(-17))
	desc := matcher.Description()
	expectedDesc := "-17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -17.
		testCase{-17, MATCH_TRUE, ""},
		testCase{-17.0, MATCH_TRUE, ""},
		testCase{-17 + 0i, MATCH_TRUE, ""},
		testCase{int(-17), MATCH_TRUE, ""},
		testCase{int8(-17), MATCH_TRUE, ""},
		testCase{int16(-17), MATCH_TRUE, ""},
		testCase{int32(-17), MATCH_TRUE, ""},
		testCase{int64(-17), MATCH_TRUE, ""},
		testCase{float32(-17), MATCH_TRUE, ""},
		testCase{float64(-17), MATCH_TRUE, ""},
		testCase{complex64(-17), MATCH_TRUE, ""},
		testCase{complex128(-17), MATCH_TRUE, ""},
		testCase{interface{}(int(-17)), MATCH_TRUE, ""},

		// Values that would be -17 in two's complement.
		testCase{uint((1 << 32) - 17), MATCH_FALSE, ""},
		testCase{uint8((1 << 8) - 17), MATCH_FALSE, ""},
		testCase{uint16((1 << 16) - 17), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 17), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 17), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-16), MATCH_FALSE, ""},
		testCase{int8(-16), MATCH_FALSE, ""},
		testCase{int16(-16), MATCH_FALSE, ""},
		testCase{int32(-16), MATCH_FALSE, ""},
		testCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float32(-17.1), MATCH_FALSE, ""},
		testCase{float32(-16.9), MATCH_FALSE, ""},
		testCase{complex64(-16), MATCH_FALSE, ""},
		testCase{complex64(-17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr((1 << 32) - 17), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"-17", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroInt8(t *testing.T) {
	matcher := Equals(int8(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 0.
		testCase{0, MATCH_TRUE, ""},
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(int(0)), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1), MATCH_FALSE, ""},
		testCase{int8(1), MATCH_FALSE, ""},
		testCase{int16(1), MATCH_FALSE, ""},
		testCase{int32(1), MATCH_FALSE, ""},
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{float32(-0.1), MATCH_FALSE, ""},
		testCase{float32(0.1), MATCH_FALSE, ""},
		testCase{complex64(1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt8(t *testing.T) {
	matcher := Equals(int8(17))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 17.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(17), MATCH_TRUE, ""},
		testCase{int8(17), MATCH_TRUE, ""},
		testCase{int16(17), MATCH_TRUE, ""},
		testCase{int32(17), MATCH_TRUE, ""},
		testCase{int64(17), MATCH_TRUE, ""},
		testCase{float32(17), MATCH_TRUE, ""},
		testCase{float64(17), MATCH_TRUE, ""},
		testCase{complex64(17), MATCH_TRUE, ""},
		testCase{complex128(17), MATCH_TRUE, ""},
		testCase{interface{}(int(17)), MATCH_TRUE, ""},
		testCase{uint(17), MATCH_TRUE, ""},
		testCase{uint8(17), MATCH_TRUE, ""},
		testCase{uint16(17), MATCH_TRUE, ""},
		testCase{uint32(17), MATCH_TRUE, ""},
		testCase{uint64(17), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(16), MATCH_FALSE, ""},
		testCase{int8(16), MATCH_FALSE, ""},
		testCase{int16(16), MATCH_FALSE, ""},
		testCase{int32(16), MATCH_FALSE, ""},
		testCase{int64(16), MATCH_FALSE, ""},
		testCase{float32(16.9), MATCH_FALSE, ""},
		testCase{float32(17.1), MATCH_FALSE, ""},
		testCase{complex64(16), MATCH_FALSE, ""},
		testCase{complex64(17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(17), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"17", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int16
////////////////////////////////////////////////////////////

func TestNegativeInt16(t *testing.T) {
	matcher := Equals(int16(-32766))
	desc := matcher.Description()
	expectedDesc := "-32766"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -32766.
		testCase{-32766, MATCH_TRUE, ""},
		testCase{-32766.0, MATCH_TRUE, ""},
		testCase{-32766 + 0i, MATCH_TRUE, ""},
		testCase{int(-32766), MATCH_TRUE, ""},
		testCase{int16(-32766), MATCH_TRUE, ""},
		testCase{int32(-32766), MATCH_TRUE, ""},
		testCase{int64(-32766), MATCH_TRUE, ""},
		testCase{float32(-32766), MATCH_TRUE, ""},
		testCase{float64(-32766), MATCH_TRUE, ""},
		testCase{complex64(-32766), MATCH_TRUE, ""},
		testCase{complex128(-32766), MATCH_TRUE, ""},
		testCase{interface{}(int(-32766)), MATCH_TRUE, ""},

		// Values that would be -32766 in two's complement.
		testCase{uint((1 << 32) - 32766), MATCH_FALSE, ""},
		testCase{uint16((1 << 16) - 32766), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 32766), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 32766), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-16), MATCH_FALSE, ""},
		testCase{int8(-16), MATCH_FALSE, ""},
		testCase{int16(-16), MATCH_FALSE, ""},
		testCase{int32(-16), MATCH_FALSE, ""},
		testCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float32(-32766.1), MATCH_FALSE, ""},
		testCase{float32(-32765.9), MATCH_FALSE, ""},
		testCase{complex64(-32766.1), MATCH_FALSE, ""},
		testCase{complex64(-32766 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr((1 << 32) - 32766), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"-32766", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroInt16(t *testing.T) {
	matcher := Equals(int16(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 0.
		testCase{0, MATCH_TRUE, ""},
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(int(0)), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1), MATCH_FALSE, ""},
		testCase{int8(1), MATCH_FALSE, ""},
		testCase{int16(1), MATCH_FALSE, ""},
		testCase{int32(1), MATCH_FALSE, ""},
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{float32(-0.1), MATCH_FALSE, ""},
		testCase{float32(0.1), MATCH_FALSE, ""},
		testCase{complex64(1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt16(t *testing.T) {
	matcher := Equals(int16(32765))
	desc := matcher.Description()
	expectedDesc := "32765"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 32765.
		testCase{32765, MATCH_TRUE, ""},
		testCase{32765.0, MATCH_TRUE, ""},
		testCase{32765 + 0i, MATCH_TRUE, ""},
		testCase{int(32765), MATCH_TRUE, ""},
		testCase{int16(32765), MATCH_TRUE, ""},
		testCase{int32(32765), MATCH_TRUE, ""},
		testCase{int64(32765), MATCH_TRUE, ""},
		testCase{float32(32765), MATCH_TRUE, ""},
		testCase{float64(32765), MATCH_TRUE, ""},
		testCase{complex64(32765), MATCH_TRUE, ""},
		testCase{complex128(32765), MATCH_TRUE, ""},
		testCase{interface{}(int(32765)), MATCH_TRUE, ""},
		testCase{uint(32765), MATCH_TRUE, ""},
		testCase{uint16(32765), MATCH_TRUE, ""},
		testCase{uint32(32765), MATCH_TRUE, ""},
		testCase{uint64(32765), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(32764), MATCH_FALSE, ""},
		testCase{int16(32764), MATCH_FALSE, ""},
		testCase{int32(32764), MATCH_FALSE, ""},
		testCase{int64(32764), MATCH_FALSE, ""},
		testCase{float32(32764.9), MATCH_FALSE, ""},
		testCase{float32(32765.1), MATCH_FALSE, ""},
		testCase{complex64(32765.9), MATCH_FALSE, ""},
		testCase{complex64(32765 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(32765), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"32765", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int32
////////////////////////////////////////////////////////////

func TestNegativeInt32(t *testing.T) {
	// -2^30
	matcher := Equals(int32(-1073741824))
	desc := matcher.Description()
	expectedDesc := "-1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt32(t *testing.T) {
	// 2^30
	matcher := Equals(int32(1073741824))
	desc := matcher.Description()
	expectedDesc := "1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int64
////////////////////////////////////////////////////////////

func TestNegativeInt64(t *testing.T) {
	// -2^40
	matcher := Equals(int64(-1099511627776))
	desc := matcher.Description()
	expectedDesc := "-1099511627776"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -1099511627776.
		testCase{-1099511627776.0, MATCH_TRUE, ""},
		testCase{-1099511627776 + 0i, MATCH_TRUE, ""},
		testCase{int64(-1099511627776), MATCH_TRUE, ""},
		testCase{float32(-1099511627776), MATCH_TRUE, ""},
		testCase{float64(-1099511627776), MATCH_TRUE, ""},
		testCase{complex64(-1099511627776), MATCH_TRUE, ""},
		testCase{complex128(-1099511627776), MATCH_TRUE, ""},
		testCase{interface{}(int64(-1099511627776)), MATCH_TRUE, ""},

		// Values that would be -1099511627776 in two's complement.
		testCase{uint64((1 << 64) - 1099511627776), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int64(-1099511627775), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1099511627776.1), MATCH_FALSE, ""},
		testCase{float64(-1099511627775.9), MATCH_FALSE, ""},
		testCase{complex128(-1099511627775), MATCH_FALSE, ""},
		testCase{complex128(-1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt64(t *testing.T) {
	// 2^30
	matcher := Equals(int64(1099511627776))
	desc := matcher.Description()
	expectedDesc := "1099511627776"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 1099511627776.
		testCase{1099511627776.0, MATCH_TRUE, ""},
		testCase{1099511627776 + 0i, MATCH_TRUE, ""},
		testCase{int64(1099511627776), MATCH_TRUE, ""},
		testCase{uint64(1099511627776), MATCH_TRUE, ""},
		testCase{float32(1099511627776), MATCH_TRUE, ""},
		testCase{float64(1099511627776), MATCH_TRUE, ""},
		testCase{complex64(1099511627776), MATCH_TRUE, ""},
		testCase{complex128(1099511627776), MATCH_TRUE, ""},
		testCase{interface{}(int64(1099511627776)), MATCH_TRUE, ""},
		testCase{interface{}(uint64(1099511627776)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1099511627775), MATCH_FALSE, ""},
		testCase{uint64(1099511627775), MATCH_FALSE, ""},
		testCase{float64(1099511627776.1), MATCH_FALSE, ""},
		testCase{float64(1099511627775.9), MATCH_FALSE, ""},
		testCase{complex128(1099511627775), MATCH_FALSE, ""},
		testCase{complex128(1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}
