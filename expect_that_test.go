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
	"github.com/jacobsa/ogletest/internal"
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

// Set up a new test state with empty fields.
func setUpCurrentTest() {
	internal.CurrentTest = internal.NewTestState()
}

type fakeExpectThatMatcher struct {
	desc string
	res  MatchResult
	err  string
}

func (m *fakeExpectThatMatcher) Matches(c interface{}) (MatchResult, string) {
	return m.res, m.err
}

func (m *fakeExpectThatMatcher) Description() string {
	return m.desc
}

func assertEqInt(t *testing.T, e, c int) {
	if e != c {
		t.Fatalf("Expected %d, got %d", e, c)
	}
}

func expectEqUint(t *testing.T, e, c uint) {
	if e != c {
		t.Errorf("Expected %u, got %u", e, c)
	}
}

func expectEqStr(t *testing.T, e, c string) {
	if e != c {
		t.Errorf("Expected %s, got %s", e, c)
	}
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func TestNoCurrentTest(t *testing.T) {
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

	internal.CurrentTest = nil
	ExpectThat(17, Equals(17))
}

func TestNoFailure(t *testing.T) {
	setUpCurrentTest()
	matcher := &fakeExpectThatMatcher{"", MATCH_TRUE, ""}
	ExpectThat(17, matcher)

	assertEqInt(t, 0, len(internal.CurrentTest.FailureRecords))
}

func TestMatchFalseWithoutMessages(t *testing.T) {
	setUpCurrentTest()
	matcher := &fakeExpectThatMatcher{"taco", MATCH_FALSE, ""}
	ExpectThat(17, matcher)

	assertEqInt(t, 1, len(internal.CurrentTest.FailureRecords))

	record := internal.CurrentTest.FailureRecords[0]
	expectEqStr(t, "expect_that_test.go", record.FileName)
	expectEqUint(t, 98, record.LineNumber)
	expectEqStr(t, "Expected: taco\nActual:   17", record.GeneratedError)
	expectEqStr(t, "", record.UserError)
}

func TestMatchUndefinedWithoutMessages(t *testing.T) {
	setUpCurrentTest()
	matcher := &fakeExpectThatMatcher{"taco", MATCH_UNDEFINED, ""}
	ExpectThat(17, matcher)

	assertEqInt(t, 1, len(internal.CurrentTest.FailureRecords))

	record := internal.CurrentTest.FailureRecords[0]
	expectEqStr(t, "expect_that_test.go", record.FileName)
	expectEqUint(t, 112, record.LineNumber)
	expectEqStr(t, "Expected: taco\nActual:   17", record.GeneratedError)
	expectEqStr(t, "", record.UserError)
}

func TestFailureWithMatcherMessage(t *testing.T) {
	setUpCurrentTest()
	matcher := &fakeExpectThatMatcher{"taco", MATCH_UNDEFINED, "which is foo"}
	ExpectThat(17, matcher)

	assertEqInt(t, 1, len(internal.CurrentTest.FailureRecords))

	record := internal.CurrentTest.FailureRecords[0]
	expectEqStr(t, "Expected: taco\nActual:   17, which is foo", record.GeneratedError)
}

func TestFailureWithUserMessage(t *testing.T) {
	setUpCurrentTest()
	matcher := &fakeExpectThatMatcher{"taco", MATCH_UNDEFINED, ""}
	ExpectThat(17, matcher, "Asd: %d %s", 19, "taco")

	assertEqInt(t, 1, len(internal.CurrentTest.FailureRecords))
	record := internal.CurrentTest.FailureRecords[0]

	expectEqStr(t, "Asd: 19 taco", record.UserError)
}

func TestAdditionalFailure(t *testing.T) {
}
