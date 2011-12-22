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

package oglematchers_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"math"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type GreaterThanTest struct {
}

func init() { RegisterTestSuite(&GreaterThanTest{}) }

type gtTestCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func (t *GreaterThanTest) checkTestCases(matcher Matcher, cases []gtTestCase) {
	for i, c := range cases {
		result, err := matcher.Matches(c.candidate)

		ExpectThat(
			result,
			Equals(c.expectedResult),
			"Case %d (candidate %v)",
			i,
			c.candidate)

		errorMatcher := Error(Equals(c.expectedError))
		if c.expectedError == "" {
			errorMatcher = Equals(nil)
		}

		ExpectThat(
			err,
			errorMatcher,
			"Case %d (candidate %v)",
			i,
			c.candidate)
	}
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func (t *GreaterThanTest) IntegerCandidateBadTypes() {
	matcher := GreaterThan(int(-150))

	cases := []gtTestCase{
		gtTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{&gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) FloatCandidateBadTypes() {
	matcher := GreaterThan(float32(-150))

	cases := []gtTestCase{
		gtTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{&gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) StringCandidateBadTypes() {
	matcher := GreaterThan("17")

	cases := []gtTestCase{
		gtTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{int(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{int8(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{int16(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{int32(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{int64(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uint(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uint8(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uint16(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uint32(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uint64(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{float32(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{float64(0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{&gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		gtTestCase{gtTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) BadArgument() {
	panicked := false

	defer func() {
		ExpectThat(panicked, Equals(true))
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	GreaterThan(complex128(0))
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func (t *GreaterThanTest) NegativeIntegerLiteral() {
	matcher := GreaterThan(-150)
	desc := matcher.Description()
	expectedDesc := "greater than -150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-(1 << 30), MATCH_TRUE, ""},
		gtTestCase{-151, MATCH_TRUE, ""},
		gtTestCase{-150, MATCH_TRUE, ""},
		gtTestCase{-149, MATCH_FALSE, ""},
		gtTestCase{0, MATCH_FALSE, ""},
		gtTestCase{17, MATCH_FALSE, ""},

		gtTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int(-151), MATCH_TRUE, ""},
		gtTestCase{int(-150), MATCH_TRUE, ""},
		gtTestCase{int(-149), MATCH_FALSE, ""},
		gtTestCase{int(0), MATCH_FALSE, ""},
		gtTestCase{int(17), MATCH_FALSE, ""},

		gtTestCase{int8(-127), MATCH_FALSE, ""},
		gtTestCase{int8(0), MATCH_FALSE, ""},
		gtTestCase{int8(17), MATCH_FALSE, ""},

		gtTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		gtTestCase{int16(-151), MATCH_TRUE, ""},
		gtTestCase{int16(-150), MATCH_TRUE, ""},
		gtTestCase{int16(-149), MATCH_FALSE, ""},
		gtTestCase{int16(0), MATCH_FALSE, ""},
		gtTestCase{int16(17), MATCH_FALSE, ""},

		gtTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int32(-151), MATCH_TRUE, ""},
		gtTestCase{int32(-150), MATCH_TRUE, ""},
		gtTestCase{int32(-149), MATCH_FALSE, ""},
		gtTestCase{int32(0), MATCH_FALSE, ""},
		gtTestCase{int32(17), MATCH_FALSE, ""},

		gtTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int64(-151), MATCH_TRUE, ""},
		gtTestCase{int64(-150), MATCH_TRUE, ""},
		gtTestCase{int64(-149), MATCH_FALSE, ""},
		gtTestCase{int64(0), MATCH_FALSE, ""},
		gtTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		gtTestCase{uint(0), MATCH_FALSE, ""},
		gtTestCase{uint(17), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_FALSE, ""},
		gtTestCase{uint8(17), MATCH_FALSE, ""},
		gtTestCase{uint8(253), MATCH_FALSE, ""},

		gtTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		gtTestCase{uint16(0), MATCH_FALSE, ""},
		gtTestCase{uint16(17), MATCH_FALSE, ""},

		gtTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		gtTestCase{uint32(0), MATCH_FALSE, ""},
		gtTestCase{uint32(17), MATCH_FALSE, ""},

		gtTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		gtTestCase{uint64(0), MATCH_FALSE, ""},
		gtTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float32(-151), MATCH_TRUE, ""},
		gtTestCase{float32(-150.1), MATCH_TRUE, ""},
		gtTestCase{float32(-150), MATCH_TRUE, ""},
		gtTestCase{float32(-149.9), MATCH_FALSE, ""},
		gtTestCase{float32(0), MATCH_FALSE, ""},
		gtTestCase{float32(17), MATCH_FALSE, ""},
		gtTestCase{float32(160), MATCH_FALSE, ""},

		gtTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float64(-151), MATCH_TRUE, ""},
		gtTestCase{float64(-150.1), MATCH_TRUE, ""},
		gtTestCase{float64(-150), MATCH_TRUE, ""},
		gtTestCase{float64(-149.9), MATCH_FALSE, ""},
		gtTestCase{float64(0), MATCH_FALSE, ""},
		gtTestCase{float64(17), MATCH_FALSE, ""},
		gtTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) ZeroIntegerLiteral() {
	matcher := GreaterThan(0)
	desc := matcher.Description()
	expectedDesc := "greater than 0"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-(1 << 30), MATCH_TRUE, ""},
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{0, MATCH_TRUE, ""},
		gtTestCase{1, MATCH_FALSE, ""},
		gtTestCase{17, MATCH_FALSE, ""},
		gtTestCase{(1 << 30), MATCH_FALSE, ""},

		gtTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(0), MATCH_TRUE, ""},
		gtTestCase{int(1), MATCH_FALSE, ""},
		gtTestCase{int(17), MATCH_FALSE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(0), MATCH_TRUE, ""},
		gtTestCase{int8(1), MATCH_FALSE, ""},

		gtTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(0), MATCH_TRUE, ""},
		gtTestCase{int16(1), MATCH_FALSE, ""},
		gtTestCase{int16(17), MATCH_FALSE, ""},

		gtTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(0), MATCH_TRUE, ""},
		gtTestCase{int32(1), MATCH_FALSE, ""},
		gtTestCase{int32(17), MATCH_FALSE, ""},

		gtTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(0), MATCH_TRUE, ""},
		gtTestCase{int64(1), MATCH_FALSE, ""},
		gtTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint((1 << 32) - 1), MATCH_FALSE, ""},
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(1), MATCH_FALSE, ""},
		gtTestCase{uint(17), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(1), MATCH_FALSE, ""},
		gtTestCase{uint8(17), MATCH_FALSE, ""},
		gtTestCase{uint8(253), MATCH_FALSE, ""},

		gtTestCase{uint16((1 << 16) - 1), MATCH_FALSE, ""},
		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(1), MATCH_FALSE, ""},
		gtTestCase{uint16(17), MATCH_FALSE, ""},

		gtTestCase{uint32((1 << 32) - 1), MATCH_FALSE, ""},
		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(1), MATCH_FALSE, ""},
		gtTestCase{uint32(17), MATCH_FALSE, ""},

		gtTestCase{uint64((1 << 64) - 1), MATCH_FALSE, ""},
		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(1), MATCH_FALSE, ""},
		gtTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(-0.1), MATCH_TRUE, ""},
		gtTestCase{float32(-0.0), MATCH_TRUE, ""},
		gtTestCase{float32(0), MATCH_TRUE, ""},
		gtTestCase{float32(0.1), MATCH_FALSE, ""},
		gtTestCase{float32(17), MATCH_FALSE, ""},
		gtTestCase{float32(160), MATCH_FALSE, ""},

		gtTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(-0.1), MATCH_TRUE, ""},
		gtTestCase{float64(-0), MATCH_TRUE, ""},
		gtTestCase{float64(0), MATCH_TRUE, ""},
		gtTestCase{float64(0.1), MATCH_FALSE, ""},
		gtTestCase{float64(17), MATCH_FALSE, ""},
		gtTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) PositiveIntegerLiteral() {
	matcher := GreaterThan(150)
	desc := matcher.Description()
	expectedDesc := "greater than 150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{149, MATCH_TRUE, ""},
		gtTestCase{150, MATCH_TRUE, ""},
		gtTestCase{151, MATCH_FALSE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(149), MATCH_TRUE, ""},
		gtTestCase{int(150), MATCH_TRUE, ""},
		gtTestCase{int(151), MATCH_FALSE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(0), MATCH_TRUE, ""},
		gtTestCase{int8(17), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(149), MATCH_TRUE, ""},
		gtTestCase{int16(150), MATCH_TRUE, ""},
		gtTestCase{int16(151), MATCH_FALSE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(149), MATCH_TRUE, ""},
		gtTestCase{int32(150), MATCH_TRUE, ""},
		gtTestCase{int32(151), MATCH_FALSE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(149), MATCH_TRUE, ""},
		gtTestCase{int64(150), MATCH_TRUE, ""},
		gtTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(149), MATCH_TRUE, ""},
		gtTestCase{uint(150), MATCH_TRUE, ""},
		gtTestCase{uint(151), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(127), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(149), MATCH_TRUE, ""},
		gtTestCase{uint16(150), MATCH_TRUE, ""},
		gtTestCase{uint16(151), MATCH_FALSE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(149), MATCH_TRUE, ""},
		gtTestCase{uint32(150), MATCH_TRUE, ""},
		gtTestCase{uint32(151), MATCH_FALSE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(149), MATCH_TRUE, ""},
		gtTestCase{uint64(150), MATCH_TRUE, ""},
		gtTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(149), MATCH_TRUE, ""},
		gtTestCase{float32(149.9), MATCH_TRUE, ""},
		gtTestCase{float32(150), MATCH_TRUE, ""},
		gtTestCase{float32(150.1), MATCH_FALSE, ""},
		gtTestCase{float32(151), MATCH_FALSE, ""},

		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(149), MATCH_TRUE, ""},
		gtTestCase{float64(149.9), MATCH_TRUE, ""},
		gtTestCase{float64(150), MATCH_TRUE, ""},
		gtTestCase{float64(150.1), MATCH_FALSE, ""},
		gtTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Float literals
////////////////////////////////////////////////////////////

func (t *GreaterThanTest) NegativeFloatLiteral() {
	matcher := GreaterThan(-150.1)
	desc := matcher.Description()
	expectedDesc := "greater than -150.1"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-(1 << 30), MATCH_TRUE, ""},
		gtTestCase{-151, MATCH_TRUE, ""},
		gtTestCase{-150.1, MATCH_TRUE, ""},
		gtTestCase{-150, MATCH_FALSE, ""},
		gtTestCase{-149, MATCH_FALSE, ""},
		gtTestCase{0, MATCH_FALSE, ""},
		gtTestCase{17, MATCH_FALSE, ""},

		gtTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int(-151), MATCH_TRUE, ""},
		gtTestCase{int(-150), MATCH_FALSE, ""},
		gtTestCase{int(-149), MATCH_FALSE, ""},
		gtTestCase{int(0), MATCH_FALSE, ""},
		gtTestCase{int(17), MATCH_FALSE, ""},

		gtTestCase{int8(-127), MATCH_FALSE, ""},
		gtTestCase{int8(0), MATCH_FALSE, ""},
		gtTestCase{int8(17), MATCH_FALSE, ""},

		gtTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		gtTestCase{int16(-151), MATCH_TRUE, ""},
		gtTestCase{int16(-150), MATCH_FALSE, ""},
		gtTestCase{int16(-149), MATCH_FALSE, ""},
		gtTestCase{int16(0), MATCH_FALSE, ""},
		gtTestCase{int16(17), MATCH_FALSE, ""},

		gtTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int32(-151), MATCH_TRUE, ""},
		gtTestCase{int32(-150), MATCH_FALSE, ""},
		gtTestCase{int32(-149), MATCH_FALSE, ""},
		gtTestCase{int32(0), MATCH_FALSE, ""},
		gtTestCase{int32(17), MATCH_FALSE, ""},

		gtTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{int64(-151), MATCH_TRUE, ""},
		gtTestCase{int64(-150), MATCH_FALSE, ""},
		gtTestCase{int64(-149), MATCH_FALSE, ""},
		gtTestCase{int64(0), MATCH_FALSE, ""},
		gtTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		gtTestCase{uint(0), MATCH_FALSE, ""},
		gtTestCase{uint(17), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_FALSE, ""},
		gtTestCase{uint8(17), MATCH_FALSE, ""},
		gtTestCase{uint8(253), MATCH_FALSE, ""},

		gtTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		gtTestCase{uint16(0), MATCH_FALSE, ""},
		gtTestCase{uint16(17), MATCH_FALSE, ""},

		gtTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		gtTestCase{uint32(0), MATCH_FALSE, ""},
		gtTestCase{uint32(17), MATCH_FALSE, ""},

		gtTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		gtTestCase{uint64(0), MATCH_FALSE, ""},
		gtTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float32(-151), MATCH_TRUE, ""},
		gtTestCase{float32(-150.2), MATCH_TRUE, ""},
		gtTestCase{float32(-150.1), MATCH_TRUE, ""},
		gtTestCase{float32(-150), MATCH_FALSE, ""},
		gtTestCase{float32(0), MATCH_FALSE, ""},
		gtTestCase{float32(17), MATCH_FALSE, ""},
		gtTestCase{float32(160), MATCH_FALSE, ""},

		gtTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		gtTestCase{float64(-151), MATCH_TRUE, ""},
		gtTestCase{float64(-150.2), MATCH_TRUE, ""},
		gtTestCase{float64(-150.1), MATCH_TRUE, ""},
		gtTestCase{float64(-150), MATCH_FALSE, ""},
		gtTestCase{float64(0), MATCH_FALSE, ""},
		gtTestCase{float64(17), MATCH_FALSE, ""},
		gtTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) PositiveFloatLiteral() {
	matcher := GreaterThan(149.9)
	desc := matcher.Description()
	expectedDesc := "greater than 149.9"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{149, MATCH_TRUE, ""},
		gtTestCase{149.9, MATCH_TRUE, ""},
		gtTestCase{150, MATCH_FALSE, ""},
		gtTestCase{151, MATCH_FALSE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(149), MATCH_TRUE, ""},
		gtTestCase{int(150), MATCH_FALSE, ""},
		gtTestCase{int(151), MATCH_FALSE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(0), MATCH_TRUE, ""},
		gtTestCase{int8(17), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(149), MATCH_TRUE, ""},
		gtTestCase{int16(150), MATCH_FALSE, ""},
		gtTestCase{int16(151), MATCH_FALSE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(149), MATCH_TRUE, ""},
		gtTestCase{int32(150), MATCH_FALSE, ""},
		gtTestCase{int32(151), MATCH_FALSE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(149), MATCH_TRUE, ""},
		gtTestCase{int64(150), MATCH_FALSE, ""},
		gtTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(149), MATCH_TRUE, ""},
		gtTestCase{uint(150), MATCH_FALSE, ""},
		gtTestCase{uint(151), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(127), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(149), MATCH_TRUE, ""},
		gtTestCase{uint16(150), MATCH_FALSE, ""},
		gtTestCase{uint16(151), MATCH_FALSE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(149), MATCH_TRUE, ""},
		gtTestCase{uint32(150), MATCH_FALSE, ""},
		gtTestCase{uint32(151), MATCH_FALSE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(149), MATCH_TRUE, ""},
		gtTestCase{uint64(150), MATCH_FALSE, ""},
		gtTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(149), MATCH_TRUE, ""},
		gtTestCase{float32(149.8), MATCH_TRUE, ""},
		gtTestCase{float32(149.9), MATCH_TRUE, ""},
		gtTestCase{float32(150), MATCH_FALSE, ""},
		gtTestCase{float32(151), MATCH_FALSE, ""},

		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(149), MATCH_TRUE, ""},
		gtTestCase{float64(149.8), MATCH_TRUE, ""},
		gtTestCase{float64(149.9), MATCH_TRUE, ""},
		gtTestCase{float64(150), MATCH_FALSE, ""},
		gtTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Subtle cases
////////////////////////////////////////////////////////////

func (t *GreaterThanTest) Int64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterThan(int64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 1, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(0), MATCH_TRUE, ""},
		gtTestCase{int16(32767), MATCH_TRUE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(255), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(65535), MATCH_TRUE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		gtTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) Int64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterThan(int64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{1 << 30, MATCH_TRUE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(0), MATCH_TRUE, ""},
		gtTestCase{int16(32767), MATCH_TRUE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(255), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(65535), MATCH_TRUE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) Uint64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterThan(uint64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 1, MATCH_TRUE, ""},
		gtTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(0), MATCH_TRUE, ""},
		gtTestCase{int16(32767), MATCH_TRUE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(255), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(65535), MATCH_TRUE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		gtTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) Uint64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterThan(uint64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{-1, MATCH_TRUE, ""},
		gtTestCase{1 << 30, MATCH_TRUE, ""},

		gtTestCase{int(-1), MATCH_TRUE, ""},
		gtTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		gtTestCase{int8(-1), MATCH_TRUE, ""},
		gtTestCase{int8(127), MATCH_TRUE, ""},

		gtTestCase{int16(-1), MATCH_TRUE, ""},
		gtTestCase{int16(0), MATCH_TRUE, ""},
		gtTestCase{int16(32767), MATCH_TRUE, ""},

		gtTestCase{int32(-1), MATCH_TRUE, ""},
		gtTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint(0), MATCH_TRUE, ""},
		gtTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		gtTestCase{uint8(0), MATCH_TRUE, ""},
		gtTestCase{uint8(255), MATCH_TRUE, ""},

		gtTestCase{uint16(0), MATCH_TRUE, ""},
		gtTestCase{uint16(65535), MATCH_TRUE, ""},

		gtTestCase{uint32(0), MATCH_TRUE, ""},
		gtTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) Float32AboveExactIntegerRange() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterThan(float32(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 3.3554432e+07"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float32(-1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) Float64AboveExactIntegerRange() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterThan(float64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than 1.8014398509481984e+16"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		// Signed integers.
		gtTestCase{int64(-1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 - 2), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 2), MATCH_TRUE, ""},
		gtTestCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		gtTestCase{uint64(0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 - 2), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 2), MATCH_TRUE, ""},
		gtTestCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Floating point.
		gtTestCase{float64(-1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		gtTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// String literals
////////////////////////////////////////////////////////////

func (t *GreaterThanTest) EmptyString() {
	matcher := GreaterThan("")
	desc := matcher.Description()
	expectedDesc := "greater than \"\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		gtTestCase{"", MATCH_TRUE, ""},
		gtTestCase{"\x00", MATCH_FALSE, ""},
		gtTestCase{"a", MATCH_FALSE, ""},
		gtTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) SingleNullByte() {
	matcher := GreaterThan("\x00")
	desc := matcher.Description()
	expectedDesc := "greater than \"\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		gtTestCase{"", MATCH_TRUE, ""},
		gtTestCase{"\x00", MATCH_TRUE, ""},
		gtTestCase{"\x00\x00", MATCH_FALSE, ""},
		gtTestCase{"a", MATCH_FALSE, ""},
		gtTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterThanTest) LongerString() {
	matcher := GreaterThan("foo\x00")
	desc := matcher.Description()
	expectedDesc := "greater than \"foo\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []gtTestCase{
		gtTestCase{"", MATCH_TRUE, ""},
		gtTestCase{"\x00", MATCH_TRUE, ""},
		gtTestCase{"bar", MATCH_TRUE, ""},
		gtTestCase{"foo", MATCH_TRUE, ""},
		gtTestCase{"foo\x00", MATCH_TRUE, ""},
		gtTestCase{"foo\x00\x00", MATCH_FALSE, ""},
		gtTestCase{"fooa", MATCH_FALSE, ""},
		gtTestCase{"qux", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}
