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

type fakeAnyOfMatcher struct {
	desc string
	res  MatchResult
	err  string
}

func (m *fakeAnyOfMatcher) Matches(c interface{}) (MatchResult, string) {
	return m.res, m.err
}

func (m *fakeAnyOfMatcher) Description() string {
	return m.desc
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func TestAnyOfCallsWrapped(t *testing.T) {
	t.Fail()
}

func TestEmptySet(t *testing.T) {
	matcher := AnyOf()

	res, err := matcher.Matches(0)
	expectedRes := MATCH_FALSE
	expectedErr := ""

	if res != expectedRes {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	if err != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
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

	if err != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
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

	if err != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
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

	if err != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
}

func TestAllFalse(t *testing.T) {
}

func TestDescriptionForEmptySet(t *testing.T) {
}

func TestDescriptionForNonEmptySet(t *testing.T) {
}
