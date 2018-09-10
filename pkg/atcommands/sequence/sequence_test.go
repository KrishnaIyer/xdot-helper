// Copyright © 2018 Krishna Iyer Easwaran.  All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 	http:#www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sequence

import (
	"testing"

	"github.com/KrishnaIyer/xdot-helper/pkg/pbapi"
)

func TestSequenceOK(t *testing.T) {
	s := New("/dev/tty.usbmodem1451111", 1, false)
	defer s.CloseHandler()
	// Correct sequence
	cmd := pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	s.AddCommand(cmd)
	cmd = pbapi.Command{
		Request:         "AT+DI",
		WaitPeriod:      1,
		LinesInResponse: 1,
	}
	s.AddCommand(cmd)

	t.Log(s.ListCommands())

	r, err := s.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	// Error sequence, continue

	// Error sequence, exit

}

func TestSequenceErrorExit(t *testing.T) {
	s := New("/dev/tty.usbmodem146111", 1, true)
	defer s.CloseHandler()
	// Correct sequence
	cmd := pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	s.AddCommand(cmd)
	cmd = pbapi.Command{
		Request:         "ATTTT",
		WaitPeriod:      1,
		LinesInResponse: 1,
	}
	s.AddCommand(cmd)
	cmd = pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	s.AddCommand(cmd)

	t.Log(s.ListCommands())

	r, err := s.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

}

func TestSequenceErrorContinue(t *testing.T) {
	s := New("/dev/tty.usbmodem146111", 1, false)
	defer s.CloseHandler()
	// Correct sequence
	cmd := pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	s.AddCommand(cmd)
	cmd = pbapi.Command{
		Request:         "ATTTT",
		WaitPeriod:      1,
		LinesInResponse: 1,
	}
	s.AddCommand(cmd)
	cmd = pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	s.AddCommand(cmd)

	t.Log(s.ListCommands())

	r, err := s.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

}
