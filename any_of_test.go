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
	"errors"
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type fakeAnyOfMatcher struct {
	desc string
	res  bool
	err  error
}

func (m *fakeAnyOfMatcher) Matches(c interface{}) (bool, error) {
	return m.res, m.err
}

func (m *fakeAnyOfMatcher) Description() string {
	return m.desc
}

type AnyOfTest struct {
}

func init() { RegisterTestSuite(&AnyOfTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *AnyOfTest) EmptySet() {
	matcher := AnyOf()

	res, err := matcher.Matches(0)

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("")))
}

func (t *AnyOfTest) OneTrue() {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", false, NewFatalError("foo")},
		17,
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
		&fakeAnyOfMatcher{"", true, nil},
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
	)

	res, err := matcher.Matches(0)

	ExpectTrue(res)
	ExpectEq(nil, err)
}

func (t *AnyOfTest) OneEqual() {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", false, NewFatalError("foo")},
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
		13,
		"taco",
		19,
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
	)

	res, err := matcher.Matches("taco")

	ExpectTrue(res)
	ExpectEq(nil, err)
}

func (t *AnyOfTest) OneFatal() {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
		17,
		&fakeAnyOfMatcher{"", false, NewFatalError("taco")},
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
	)

	res, err := matcher.Matches(0)

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("taco")))
}

func (t *AnyOfTest) AllFalseAndNotEqual() {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
		17,
		&fakeAnyOfMatcher{"", false, errors.New("foo")},
		19,
	)

	res, err := matcher.Matches(0)

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("")))
}

func (t *AnyOfTest) DescriptionForEmptySet() {
	matcher := AnyOf()
	ExpectEq("or()", matcher.Description())
}

func (t *AnyOfTest) DescriptionForNonEmptySet() {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"taco", true, nil},
		"burrito",
		&fakeAnyOfMatcher{"enchilada", true, nil},
	)

	ExpectEq("or(taco, burrito, enchilada)", matcher.Description())
}
