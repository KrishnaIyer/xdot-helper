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

package handler

import (
	"bytes"
	"errors"
	"strings"

	"go.uber.org/zap"

	"github.com/KrishnaIyer/xdot-helper/pkg/pbapi"
	"github.com/KrishnaIyer/xdot-helper/pkg/usbserial"
)

const (
	constCmdSeparator      = "="
	constResDelimiter      = "\r\n"
	constResSplCharacterLF = "\n"
	constResSplCharacterCR = "\r"
	constERRORResponse     = "ERROR"
	constOKResponse        = "OK"
	constDefaultWaitPeriod = 1 // 1 sec
)

// Handler handles AT commands
type Handler struct {
	logger *zap.Logger
	us     *usbserial.USBSerial
}

// New creates a new AT commands Handler.
// Make sure to call `Handler.Close()`` upon exit.
func New(device string) *Handler {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	us, err := usbserial.New(device, true, 5)
	if err != nil {
		logger.Fatal("Failed to connect to Serial Device")
	}
	return &Handler{
		us:     us,
		logger: logger,
	}
}

// Execute sends the AT Command to the usb device and extracts the response.
// The return values are:
//  - The response code; `OK`, `ERROR`
//  - The response data.
//  - Error
func (h *Handler) Execute(cmd pbapi.Command) (pbapi.Result_ResCode, []byte, error) {
	if cmd.WaitPeriod == 0 {
		cmd.WaitPeriod = constDefaultWaitPeriod
	}
	rawRes, err := h.us.SendData([]byte(cmd.Request), int(cmd.WaitPeriod))
	if err != nil {
		return pbapi.Result_NONE, nil, errors.New("Failed to execute command: " + err.Error())
	}
	r, err := cleanResponse(rawRes)
	if err != nil {
		return pbapi.Result_NONE, nil, errors.New("Failed to decode response: " + err.Error())
	}
	lenR := len(r)
	if lenR == 0 {
		return pbapi.Result_NONE, nil, errors.New("Failed to decode response: " + err.Error())
	}
	if r[0] != cmd.Request || (len(r)-2) != int(cmd.LinesInResponse) {
		return pbapi.Result_NONE, nil, errors.New("Invalid Response received from Device")
	}
	var res bytes.Buffer
	for i := 1; i < (len(r) - 1); i++ {
		res.Write([]byte(r[i]))
	}
	var rescode pbapi.Result_ResCode
	switch r[(lenR - 1)] {
	case constOKResponse:
		rescode = pbapi.Result_OK
		break
	case constERRORResponse:
		rescode = pbapi.Result_ERROR
		break
	default:
		return pbapi.Result_NONE, nil, errors.New("Invalid Response received from Device")
	}
	return rescode, res.Bytes(), err
}

// Close closes the underlying serial connection.
// This function needs to be called upon exit.
func (h *Handler) Close() {
	h.us.Close()
}

// cleanResponse removes the special charaters  response byte stream.
func cleanResponse(rawRes []byte) ([]string, error) {
	s := strings.Split(string(rawRes), constResDelimiter)
	if len(s) == 0 {
		return nil, errors.New("Invalid raw byte stream")
	}
	var res []string
	var i, j int //golang defaults int values to 0
	for ; i < len(s)-1; i++ {
		if s[i] != "" {
			temp := strings.Trim(s[i], constResSplCharacterCR)
			temp = strings.Trim(temp, constResSplCharacterLF)
			res = append(res, temp) //Append is not thread-safe
			j++
		}
	}
	return res, nil
}
