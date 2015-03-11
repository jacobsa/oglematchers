// Copyright 2015 Aaron Jacobs. All Rights Reserved.
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
	"testing"

	. "github.com/jacobsa/ogletest"
)

func TestHasSameTypeAs(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type HasSameTypeAsTest struct {
}

func init() { RegisterTestSuite(&HasSameTypeAsTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *HasSameTypeAsTest) Description() {
	AssertTrue(false, "TODO")
}

func (t *HasSameTypeAsTest) CandidateIsLiteralNil() {
	AssertTrue(false, "TODO")
}

func (t *HasSameTypeAsTest) CandidateIsNilMap() {
	AssertTrue(false, "TODO")
}

func (t *HasSameTypeAsTest) CandidateIsNilInterface() {
	AssertTrue(false, "TODO")
}

func (t *HasSameTypeAsTest) CandidateIsString() {
	AssertTrue(false, "TODO")
}

func (t *HasSameTypeAsTest) CandidateIsStringAlias() {
	AssertTrue(false, "TODO")
}
