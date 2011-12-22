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
	"errors"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type allOfFakeMatcher struct {
	desc string
	res  MatchResult
	err  error
}

func (m *allOfFakeMatcher) Matches(c interface{}) (MatchResult, error) {
	return m.res, m.err
}

func (m *allOfFakeMatcher) Description() string {
	return m.desc
}

type AllOfTest struct {
}

func init() { RegisterTestSuite(&AllOfTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *AllOfTest) DescriptionWithEmptySet() {
	m := AllOf()
	ExpectEq("is anything", m.Description())
}

func (t *AllOfTest) DescriptionWithOneMatcher() {
	m := AllOf(&allOfFakeMatcher{"taco", MATCH_FALSE, nil})
	ExpectEq("taco", m.Description())
}

func (t *AllOfTest) DescriptionWithMultipleMatchers() {
	m := AllOf(
		&allOfFakeMatcher{"taco", MATCH_FALSE, nil},
		&allOfFakeMatcher{"burrito", MATCH_FALSE, nil},
		&allOfFakeMatcher{"enchilada", MATCH_FALSE, nil})

	ExpectEq("taco, and burrito, and enchilada", m.Description())
}

func (t *AllOfTest) EmptySet() {
	m := AllOf()
	res, err := m.Matches(17)

	ExpectEq(MATCH_TRUE, res)
	ExpectEq(nil, err)
}

func (t *AllOfTest) OneMatcherSaysUndefinedAndSomeSayFalse() {
	m := AllOf(
		&allOfFakeMatcher{"", MATCH_FALSE, errors.New("")},
		&allOfFakeMatcher{"", MATCH_UNDEFINED, errors.New("taco")},
		&allOfFakeMatcher{"", MATCH_FALSE, errors.New("")},
		&allOfFakeMatcher{"", MATCH_TRUE, nil})

	res, err := m.Matches(17)

	ExpectEq(MATCH_UNDEFINED, res)
	ExpectThat(err, Error(Equals("taco")))
}

func (t *AllOfTest) OneMatcherSaysFalseAndOthersSayTrue() {
	m := AllOf(
		&allOfFakeMatcher{"", MATCH_TRUE, nil},
		&allOfFakeMatcher{"", MATCH_FALSE, errors.New("taco")},
		&allOfFakeMatcher{"", MATCH_TRUE, nil})

	res, err := m.Matches(17)

	ExpectEq(MATCH_FALSE, res)
	ExpectThat(err, Error(Equals("taco")))
}

func (t *AllOfTest) AllMatchersSayTrue() {
	m := AllOf(
		&allOfFakeMatcher{"", MATCH_TRUE, nil},
		&allOfFakeMatcher{"", MATCH_TRUE, nil},
		&allOfFakeMatcher{"", MATCH_TRUE, nil})

	res, err := m.Matches(17)

	ExpectEq(MATCH_TRUE, res)
	ExpectEq(nil, err)
}
