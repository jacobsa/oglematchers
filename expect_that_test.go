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
	internal.CurrentTest = &internal.TestState{}
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
}

func TestMatchFalseWithoutMessages(t *testing.T) {
}

func TestMatchUndefinedWithoutMessages(t *testing.T) {
}

func TestFailureWithMatcherMessage(t *testing.T) {
}

func TestFailureWithUserMessage(t *testing.T) {
}

func TestAdditionalFailure(t *testing.T) {
}
