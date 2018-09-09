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
	"bytes"
	"errors"

	handler "github.com/KrishnaIyer/xdot-helper/pkg/atcommands/handler"
	"github.com/KrishnaIyer/xdot-helper/pkg/pbapi"
)

// Sequence is chain of AT commands executed sequencially with a delay of Delay between each command.
type Sequence struct {
	CMDs        []pbapi.Command
	Delay       uint16
	Device      string
	ExitOnError bool
	Handler     *handler.Handler
}

const constCMDListSeparator = 0x3B // Ascii for ;

// New returns an empty sequence and creates the handler
func New(device string, delay uint16, exitOnError bool) *Sequence {
	return &Sequence{
		Device:      device,
		Delay:       delay,
		ExitOnError: exitOnError,
		Handler:     handler.New("/dev/tty.usbmodem146111"),
	}
}

// AddCommand appends a new command to the end of the sequence.
func (seq *Sequence) AddCommand(cmd pbapi.Command) error {
	if cmd.Request == "" {
		return errors.New("Invalid Command")
	}
	seq.CMDs = append(seq.CMDs, cmd)
	return nil
}

// ListCommands returns a semicolon separated string of AT Commands in this sequence in the order they are to be executed.
func (seq *Sequence) ListCommands() string {
	var b bytes.Buffer
	for _, cmd := range seq.CMDs {
		b.Write([]byte(cmd.Request))
		b.WriteByte(constCMDListSeparator)
	}
	return b.String()
}

// Execute executes each command in the sequence in the order in which they are added, with a time delay of `Sequence.Delay` between each command.
// The `ExitOnError` parameter can be used to stop execution in case of errors.
func (seq *Sequence) Execute() ([]pbapi.Result, error) {
	var seqres []pbapi.Result
	for _, cmd := range seq.CMDs {
		rescode, res, err := seq.Handler.Execute(cmd)
		if err != nil {
			return nil, err
		}
		r := pbapi.Result{
			Request:      cmd.Request,
			ResponseCode: rescode,
			Response:     string(res),
		}
		seqres = append(seqres, r)
		if rescode != pbapi.Result_ERROR && seq.ExitOnError == true {
			return seqres, nil
		}
	}
	return seqres, nil
}
