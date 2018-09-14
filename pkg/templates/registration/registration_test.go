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

package registration

import (
	"testing"

	sequence "github.com/KrishnaIyer/xdot-helper/pkg/atcommands/sequence"
	"github.com/KrishnaIyer/xdot-helper/pkg/pbapi"
)

func TestNewRegistationSequence(t *testing.T) {
	devEUISeq := sequence.New("/dev/tty.usbmodem1451111", 1, true)
	defer devEUISeq.CloseHandler()
	devEUISeq.AddCommand(pbapi.Command{
		Request:         "AT+DI",
		LinesInResponse: 1,
	})
	t.Log(devEUISeq.ListCommands())
	res, err := devEUISeq.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	s, err := NewRegistationSequence("/dev/tty.usbmodem1451111", "1E3C71683FC5B214C46E058C5C6D0B60", "70B3D57ED0012348")
}
