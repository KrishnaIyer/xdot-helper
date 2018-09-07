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

package usbserial

import "testing"

func TestSendData(t *testing.T) {

	usb, err := New("/dev/tty.usbmodem144111", true, 5)
	if err != nil {
		t.Fatal(err)
	}
	defer usb.Close()

	err = usb.SendData([]byte("AT&V"), 1)
	if err != nil {
		t.Fatal(err)
	}
}
