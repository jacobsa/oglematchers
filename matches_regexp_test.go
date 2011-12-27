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

type MatchesRegexpTest struct {
}

func init() { RegisterTestSuite(&MatchesRegexpTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *MatchesRegexpTest) Description() {
	m := MatchesRegexp("foo.*bar")
	ExpectEq("matches regexp \"foo.*bar\"", m.Description())
}

func (t *MatchesRegexpTest) InvalidRegexp() {
	ExpectThat(
		func() { MatchesRegexp("(foo") },
		Panics(HasSubstr("missing closing )")))
}

func (t *MatchesRegexpTest) CandidateIsNil() {
	m := MatchesRegexp("")
	res, err := m.Matches(nil)

	ExpectEq(false, res)
	ExpectThat(err, Error(Equals("which is not a string or []byte")))
	ExpectTrue(isFatal(err))
}

func (t *MatchesRegexpTest) CandidateIsInteger() {
	m := MatchesRegexp("")
	res, err := m.Matches(17)

	ExpectEq(false, res)
	ExpectThat(err, Error(Equals("which is not a string or []byte")))
	ExpectTrue(isFatal(err))
}

func (t *MatchesRegexpTest) NonMatchingCandidates() {
	m := MatchesRegexp("fo[op]\\s+x")
	var res bool
	var err error

	res, err = m.Matches("fon x")
	ExpectFalse(res)
	ExpectThat(err, Error(Equals("")))
	ExpectFalse(isFatal(err))

	res, err = m.Matches("fopx")
	ExpectFalse(res)
	ExpectThat(err, Error(Equals("")))
	ExpectFalse(isFatal(err))

	res, err = m.Matches("fop   ")
	ExpectFalse(res)
	ExpectThat(err, Error(Equals("")))
	ExpectFalse(isFatal(err))
}

func (t *MatchesRegexpTest) MatchingCandidates() {
	m := MatchesRegexp("fo[op]\\s+x")
	var res bool
	var err error

	res, err = m.Matches("foo x")
	ExpectTrue(res)
	ExpectEq(nil, err)

	res, err = m.Matches("fop     x")
	ExpectTrue(res)
	ExpectEq(nil, err)

	res, err = m.Matches("blah blah foo x blah blah")
	ExpectTrue(res)
	ExpectEq(nil, err)
}
