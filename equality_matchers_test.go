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

type testCase struct {
	candidate interface{}
	expectedResult MatchResult
	expectedError string
}

func checkTestCases(t *testing.T, matcher Matcher, cases []testCase) {
	for i, c := range cases {
		result, err := matcher.Matches(c.candidate)

		if result != c.expectedResult {
			t.Errorf("Case %d: expected %v, got %v", i, c.expectedResult, result)
		}

		if err != c.expectedError {
			t.Errorf("Case %d: expected error %v, got %v", i, c.expectedError, err)
		}
	}
}

func TestNegativeInt8 (t *testing.T) {
	matcher := Equals(int8(-17))
	desc := matcher.Description()
	expectedDesc := "-17"

	if (desc != expectedDesc) {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase {
		// Various types of -17.
		testCase{ -17, MATCH_TRUE, "" },
		testCase{ -17.0, MATCH_TRUE, "" },
		testCase{ int(-17), MATCH_TRUE, "" },
		testCase{ int8(-17), MATCH_TRUE, "" },
		testCase{ int16(-17), MATCH_TRUE, "" },
		testCase{ int32(-17), MATCH_TRUE, "" },
		testCase{ int64(-17), MATCH_TRUE, "" },
		testCase{ float32(-17), MATCH_TRUE, "" },
		testCase{ float64(-17), MATCH_TRUE, "" },
		testCase{ complex64(-17), MATCH_TRUE, "" },
		testCase{ complex128(-17), MATCH_TRUE, "" },
		testCase{ interface{}(int(-17)), MATCH_TRUE, "" },

		// Values that would be -17 in two's complement.
		testCase{ uint((1 << 32) - 17), MATCH_FALSE, "" },
		testCase{ uint8((1 << 8) - 17), MATCH_FALSE, "" },
		testCase{ uint16((1 << 16) - 17), MATCH_FALSE, "" },
		testCase{ uint32((1 << 32) - 17), MATCH_FALSE, "" },
		testCase{ uint64((1 << 64) - 17), MATCH_FALSE, "" },

		// Other types.
		testCase{ uintptr((1 << 32) - 17), MATCH_UNDEFINED, "which is not numeric" },
		testCase{ true, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ [...]int { -17 }, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ make(chan int), MATCH_UNDEFINED, "which is not numeric" },
		testCase{ func () {}, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ map[int]int{}, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ &someInt, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ []int { -17 }, MATCH_UNDEFINED, "which is not numeric" },
		testCase{ "-17", MATCH_UNDEFINED, "which is not numeric" },
		testCase{ testCase{}, MATCH_UNDEFINED, "which is not numeric" },
	}

	checkTestCases(t, matcher, cases)
}
