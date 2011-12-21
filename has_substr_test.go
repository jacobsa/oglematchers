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
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type HasSubstrTest struct {

}

func init() { RegisterTestSuite(&HasSubstrTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *HasSubstrTest) Description() {
	matcher := HasSubstr("taco")
	ExpectThat(matcher.Description(), Equals("has substring \"taco\""))
}

func (t *HasSubstrTest) CandidateIsNil() {
	matcher := HasSubstr("")
	res, err := matcher.Matches(nil)

	ExpectThat(res, Equals(MATCH_UNDEFINED))
	ExpectThat(err, Error(Equals("which is not a string")))
}

func (t *HasSubstrTest) CandidateIsInteger() {
	matcher := HasSubstr("")
	res, err := matcher.Matches(17)

	ExpectThat(res, Equals(MATCH_UNDEFINED))
	ExpectThat(err, Error(Equals("which is not a string")))
}

func (t *HasSubstrTest) CandidateIsByteSlice() {
	matcher := HasSubstr("")
	res, err := matcher.Matches([]byte{17})

	ExpectThat(res, Equals(MATCH_UNDEFINED))
	ExpectThat(err, Error(Equals("which is not a string")))
}

func (t *HasSubstrTest) CandidateDoesntHaveSubstring() {
	matcher := HasSubstr("taco")
	res, err := matcher.Matches("tac")

	ExpectThat(res, Equals(MATCH_FALSE))
	ExpectThat(err, Equals(nil))
}

func (t *HasSubstrTest) CandidateEqualsArg() {
	matcher := HasSubstr("taco")
	res, err := matcher.Matches("taco")

	ExpectThat(res, Equals(MATCH_TRUE))
	ExpectThat(err, Equals(nil))
}

func (t *HasSubstrTest) CandidateHasProperSubstring() {
	matcher := HasSubstr("taco")
	res, err := matcher.Matches("burritos and tacos")

	ExpectThat(res, Equals(MATCH_TRUE))
	ExpectThat(err, Equals(nil))
}

func (t *HasSubstrTest) EmptyStringIsAlwaysSubString() {
	matcher := HasSubstr("")
	res, err := matcher.Matches("asdf")

	ExpectThat(res, Equals(MATCH_TRUE))
	ExpectThat(err, Equals(nil))
}
