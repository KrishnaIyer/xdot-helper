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

package athandler

import (
	"testing"

	"github.com/KrishnaIyer/xdot-helper/pkg/pbapi"
)

func TestCleanResponse(t *testing.T) {
	data := []byte{65, 84, 43, 68, 73, 13, 13, 10, 48, 48, 45, 56, 48, 45, 48, 48, 45, 48, 48, 45, 48, 52, 45, 48, 48, 45, 48, 51, 45, 53, 102, 13, 10, 13, 10, 79, 75, 13, 10}
	res, err := cleanResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(res))
	for _, str := range res {
		t.Log([]byte(str))
	}
}

func TestExecute(t *testing.T) {
	ath := New("/dev/tty.usbmodem146111")
	cmd := pbapi.Command{
		Request:         "AT",
		WaitPeriod:      1,
		LinesInResponse: 0,
	}
	res, err := ath.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))

	cmd = pbapi.Command{
		Request:         "AT+DI",
		WaitPeriod:      1,
		LinesInResponse: 1,
	}
	res, err = ath.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}
