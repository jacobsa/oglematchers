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

package internal

import (
)

// TestState represents the state of a currently running or previously running
// test.
type TestState struct {
	// The name of the test suite, for example "UserInfoTest".
	SuiteName []string

	// The name of the test function within the test suite, for example
	// "ReturnsCorrectPhoneNumber".
	TestName []string

	// A set of failure messages that the test has produced.
	FailureMessages []string
}

// CurrentTest is the state for the currently running test, if any.
var CurrentTest *TestState
