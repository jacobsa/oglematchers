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

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type ltTestCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func checkLtTestCases(t *testing.T, matcher Matcher, cases []ltTestCase) {
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

func TestLtIntegerBadTypes(t *testing.T) {
	matcher := LessThan(int(-150))

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	checkLtTestCases(t, matcher, cases)
}

func TestLtFloatBadTypes(t *testing.T) {
	matcher := LessThan(float32(-150))

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	checkLtTestCases(t, matcher, cases)
}

func TestLtStringBadTypes(t *testing.T) {
	matcher := LessThan("17")

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int8(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int16(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint8(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint16(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{float32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{float64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	checkLtTestCases(t, matcher, cases)
}

func TestLtBadArg(t *testing.T) {
	panicked := false

	defer func() {
		if !panicked {
			t.Errorf("Expected panic; got none.")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	LessThan(complex128(0))
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func TestLtNegativeIntegerLiteral(t *testing.T) {
	matcher := LessThan(-150)
	desc := matcher.Description()
	expectedDesc := "less than -150"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-(1 << 30), MATCH_TRUE, ""},
		ltTestCase{-151, MATCH_TRUE, ""},
		ltTestCase{-150, MATCH_FALSE, ""},
		ltTestCase{0, MATCH_FALSE, ""},
		ltTestCase{17, MATCH_FALSE, ""},

		ltTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int(-151), MATCH_TRUE, ""},
		ltTestCase{int(-150), MATCH_FALSE, ""},
		ltTestCase{int(0), MATCH_FALSE, ""},
		ltTestCase{int(17), MATCH_FALSE, ""},

		ltTestCase{int8(-127), MATCH_FALSE, ""},
		ltTestCase{int8(0), MATCH_FALSE, ""},
		ltTestCase{int8(17), MATCH_FALSE, ""},

		ltTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		ltTestCase{int16(-151), MATCH_TRUE, ""},
		ltTestCase{int16(-150), MATCH_FALSE, ""},
		ltTestCase{int16(0), MATCH_FALSE, ""},
		ltTestCase{int16(17), MATCH_FALSE, ""},

		ltTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int32(-151), MATCH_TRUE, ""},
		ltTestCase{int32(-150), MATCH_FALSE, ""},
		ltTestCase{int32(0), MATCH_FALSE, ""},
		ltTestCase{int32(17), MATCH_FALSE, ""},

		ltTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int64(-151), MATCH_TRUE, ""},
		ltTestCase{int64(-150), MATCH_FALSE, ""},
		ltTestCase{int64(0), MATCH_FALSE, ""},
		ltTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint(0), MATCH_FALSE, ""},
		ltTestCase{uint(17), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_FALSE, ""},
		ltTestCase{uint8(17), MATCH_FALSE, ""},
		ltTestCase{uint8(253), MATCH_FALSE, ""},

		ltTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		ltTestCase{uint16(0), MATCH_FALSE, ""},
		ltTestCase{uint16(17), MATCH_FALSE, ""},

		ltTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint32(0), MATCH_FALSE, ""},
		ltTestCase{uint32(17), MATCH_FALSE, ""},

		ltTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		ltTestCase{uint64(0), MATCH_FALSE, ""},
		ltTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float32(-151), MATCH_TRUE, ""},
		ltTestCase{float32(-150.1), MATCH_TRUE, ""},
		ltTestCase{float32(-150), MATCH_FALSE, ""},
		ltTestCase{float32(-149.9), MATCH_FALSE, ""},
		ltTestCase{float32(0), MATCH_FALSE, ""},
		ltTestCase{float32(17), MATCH_FALSE, ""},
		ltTestCase{float32(160), MATCH_FALSE, ""},

		ltTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float64(-151), MATCH_TRUE, ""},
		ltTestCase{float64(-150.1), MATCH_TRUE, ""},
		ltTestCase{float64(-150), MATCH_FALSE, ""},
		ltTestCase{float64(-149.9), MATCH_FALSE, ""},
		ltTestCase{float64(0), MATCH_FALSE, ""},
		ltTestCase{float64(17), MATCH_FALSE, ""},
		ltTestCase{float64(160), MATCH_FALSE, ""},
	}

	checkLtTestCases(t, matcher, cases)
}

func TestLtZeroIntegerLiteral(t *testing.T) {
	matcher := LessThan(0)
	desc := matcher.Description()
	expectedDesc := "less than 0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-(1 << 30), MATCH_TRUE, ""},
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{0, MATCH_FALSE, ""},
		ltTestCase{1, MATCH_FALSE, ""},
		ltTestCase{17, MATCH_FALSE, ""},
		ltTestCase{(1 << 30), MATCH_FALSE, ""},

		ltTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(0), MATCH_FALSE, ""},
		ltTestCase{int(1), MATCH_FALSE, ""},
		ltTestCase{int(17), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(0), MATCH_FALSE, ""},
		ltTestCase{int8(1), MATCH_FALSE, ""},

		ltTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_FALSE, ""},
		ltTestCase{int16(1), MATCH_FALSE, ""},
		ltTestCase{int16(17), MATCH_FALSE, ""},

		ltTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(0), MATCH_FALSE, ""},
		ltTestCase{int32(1), MATCH_FALSE, ""},
		ltTestCase{int32(17), MATCH_FALSE, ""},

		ltTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(0), MATCH_FALSE, ""},
		ltTestCase{int64(1), MATCH_FALSE, ""},
		ltTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint((1 << 32) - 1), MATCH_FALSE, ""},
		ltTestCase{uint(0), MATCH_FALSE, ""},
		ltTestCase{uint(17), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_FALSE, ""},
		ltTestCase{uint8(17), MATCH_FALSE, ""},
		ltTestCase{uint8(253), MATCH_FALSE, ""},

		ltTestCase{uint16((1 << 16) - 1), MATCH_FALSE, ""},
		ltTestCase{uint16(0), MATCH_FALSE, ""},
		ltTestCase{uint16(17), MATCH_FALSE, ""},

		ltTestCase{uint32((1 << 32) - 1), MATCH_FALSE, ""},
		ltTestCase{uint32(0), MATCH_FALSE, ""},
		ltTestCase{uint32(17), MATCH_FALSE, ""},

		ltTestCase{uint64((1 << 64) - 1), MATCH_FALSE, ""},
		ltTestCase{uint64(0), MATCH_FALSE, ""},
		ltTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(-0.1), MATCH_TRUE, ""},
		ltTestCase{float32(-0.0), MATCH_FALSE, ""},
		ltTestCase{float32(0), MATCH_FALSE, ""},
		ltTestCase{float32(0.1), MATCH_FALSE, ""},
		ltTestCase{float32(17), MATCH_FALSE, ""},
		ltTestCase{float32(160), MATCH_FALSE, ""},

		ltTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(-0.1), MATCH_TRUE, ""},
		ltTestCase{float64(-0), MATCH_FALSE, ""},
		ltTestCase{float64(0), MATCH_FALSE, ""},
		ltTestCase{float64(17), MATCH_FALSE, ""},
		ltTestCase{float64(160), MATCH_FALSE, ""},
	}

	checkLtTestCases(t, matcher, cases)
}

func TestLtPositiveIntegerLiteral(t *testing.T) {
	matcher := LessThan(150)
	desc := matcher.Description()
	expectedDesc := "less than 150"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{149, MATCH_TRUE, ""},
		ltTestCase{150, MATCH_FALSE, ""},
		ltTestCase{151, MATCH_FALSE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(149), MATCH_TRUE, ""},
		ltTestCase{int(150), MATCH_FALSE, ""},
		ltTestCase{int(151), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(0), MATCH_TRUE, ""},
		ltTestCase{int8(17), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(149), MATCH_TRUE, ""},
		ltTestCase{int16(150), MATCH_FALSE, ""},
		ltTestCase{int16(151), MATCH_FALSE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(149), MATCH_TRUE, ""},
		ltTestCase{int32(150), MATCH_FALSE, ""},
		ltTestCase{int32(151), MATCH_FALSE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(149), MATCH_TRUE, ""},
		ltTestCase{int64(150), MATCH_FALSE, ""},
		ltTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(149), MATCH_TRUE, ""},
		ltTestCase{uint(150), MATCH_FALSE, ""},
		ltTestCase{uint(151), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(127), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(149), MATCH_TRUE, ""},
		ltTestCase{uint16(150), MATCH_FALSE, ""},
		ltTestCase{uint16(151), MATCH_FALSE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(149), MATCH_TRUE, ""},
		ltTestCase{uint32(150), MATCH_FALSE, ""},
		ltTestCase{uint32(151), MATCH_FALSE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(149), MATCH_TRUE, ""},
		ltTestCase{uint64(150), MATCH_FALSE, ""},
		ltTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(149), MATCH_TRUE, ""},
		ltTestCase{float32(149.9), MATCH_TRUE, ""},
		ltTestCase{float32(150), MATCH_FALSE, ""},
		ltTestCase{float32(150.1), MATCH_FALSE, ""},
		ltTestCase{float32(151), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(149), MATCH_TRUE, ""},
		ltTestCase{float64(149.9), MATCH_TRUE, ""},
		ltTestCase{float64(150), MATCH_FALSE, ""},
		ltTestCase{float64(150.1), MATCH_FALSE, ""},
		ltTestCase{float64(151), MATCH_FALSE, ""},
	}

	checkLtTestCases(t, matcher, cases)
}
