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

package oglematchers

import (
	"errors"
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type fakeMatcher struct {
	matchFunc   func(interface{}) (MatchResult, string)
	description string
}

func (m *fakeMatcher) Matches(c interface{}) (MatchResult, error) {
	res, err := m.matchFunc(c)
	return res, errors.New(err)
}

func (m *fakeMatcher) Description() string {
	return m.description
}

type NotTest struct {

}

func init()                     { RegisterTestSuite(&NotTest{}) }
func TestOgletest(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *NotTest) CallsWrapped() {
	var suppliedCandidate interface{}
	matchFunc := func(c interface{}) (MatchResult, string) {
		suppliedCandidate = c
		return MATCH_TRUE, ""
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	matcher.Matches(17)
	ExpectThat(suppliedCandidate, Equals(17))
}

func (t *NotTest) WrappedReturnsMatchTrue() {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_TRUE, ""
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, _ := matcher.Matches(0)
	ExpectThat(res, Equals(MATCH_FALSE))
}

func (t *NotTest) WrappedReturnsMatchFalse() {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_FALSE, "taco"
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, err := matcher.Matches(0)
	ExpectThat(res, Equals(MATCH_TRUE))
	ExpectThat(err, Equals(nil))
}

func (t *NotTest) WrappedReturnsMatchUndefined() {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_UNDEFINED, "taco"
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, err := matcher.Matches(0)
	ExpectThat(res, Equals(MATCH_UNDEFINED))
	ExpectThat(err, Equals("taco"))
}

func (t *NotTest) Description() {
	wrapped := &fakeMatcher{nil, "taco"}
	matcher := Not(wrapped)

	ExpectThat(matcher.Description(), Equals("not(taco)"))
}
