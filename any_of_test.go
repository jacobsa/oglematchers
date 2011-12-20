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

type fakeAnyOfMatcher struct {
	desc string
	res  MatchResult
	err  string
}

func (m *fakeAnyOfMatcher) Matches(c interface{}) (MatchResult, error) {
	return m.res, errors.New(m.err)
}

func (m *fakeAnyOfMatcher) Description() string {
	return m.desc
}

func expectEqErr(t *testing.T, expectedErr string, err error) {
	actualError := ""
	if err != nil {
		actualError = err.Error()
	}

	if actualError != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func TestEmptySet(t *testing.T) {
	matcher := AnyOf()

	res, err := matcher.Matches(0)
	expectedRes := MATCH_FALSE
	expectedErr := ""

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	expectEqErr(t, expectedErr, err)
}

func TestOneTrue(t *testing.T) {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", MATCH_UNDEFINED, "foo"},
		17,
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
		&fakeAnyOfMatcher{"", MATCH_TRUE, ""},
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
	)

	res, err := matcher.Matches(0)
	expectedRes := MATCH_TRUE
	expectedErr := ""

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	expectEqErr(t, expectedErr, err)
}

func TestOneEqual(t *testing.T) {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", MATCH_UNDEFINED, "foo"},
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
		13,
		"taco",
		19,
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
	)

	res, err := matcher.Matches("taco")
	expectedRes := MATCH_TRUE
	expectedErr := ""

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	expectEqErr(t, expectedErr, err)
}

func TestOneUndefined(t *testing.T) {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
		17,
		&fakeAnyOfMatcher{"", MATCH_UNDEFINED, "taco"},
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
	)

	res, err := matcher.Matches(0)
	expectedRes := MATCH_UNDEFINED
	expectedErr := "taco"

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	expectEqErr(t, expectedErr, err)
}

func TestAllFalseAndNotEqual(t *testing.T) {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
		17,
		&fakeAnyOfMatcher{"", MATCH_FALSE, "foo"},
		19,
	)

	res, err := matcher.Matches(0)
	expectedRes := MATCH_FALSE
	expectedErr := ""

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	expectEqErr(t, expectedErr, err)
}

func TestDescriptionForEmptySet(t *testing.T) {
	matcher := AnyOf()
	desc := matcher.Description()
	expected := "or()"

	if desc != expected {
		t.Errorf("Expected %v, got %v", expected, desc)
	}
}

func TestDescriptionForNonEmptySet(t *testing.T) {
	matcher := AnyOf(
		&fakeAnyOfMatcher{"taco", MATCH_TRUE, ""},
		"burrito",
		&fakeAnyOfMatcher{"enchilada", MATCH_TRUE, ""},
	)

	desc := matcher.Description()
	expected := "or(taco, burrito, enchilada)"

	if desc != expected {
		t.Errorf("Expected %v, got %v", expected, desc)
	}
}
