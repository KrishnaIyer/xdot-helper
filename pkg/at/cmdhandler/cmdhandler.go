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

package cmdhandler

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
	constErrorResponse     = "ERROR"
	constCorrectResponse   = "OK"
)

// CMDHandler handles AT commands
type CMDHandler struct {
	logger *zap.Logger
	us     *usbserial.USBSerial
}

// New creates a new AT commands Handler.
func New(device string) *CMDHandler {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	us, err := usbserial.New(device, true, 5)
	if err != nil {
		logger.Fatal("Failed to connect to Serial Device")
	}
	return &CMDHandler{
		us:     us,
		logger: logger,
	}
}

// Execute sends the AT Command to the usb device and extracts the response.
// The return values are:
//  - The response code; `OK`, `ERROR`
//  - The response data.
//  - Error
func (c *CMDHandler) Execute(cmd pbapi.Command) (string, []byte, error) {
	var data string
	if cmd.Arguments != "" {
		data = cmd.Request + constCmdSeparator + cmd.Arguments
	} else {
		data = cmd.Request
	}
	rawRes, err := c.us.SendData([]byte(data), int(cmd.WaitPeriod))
	if err != nil {
		return "", nil, errors.New("Failed to execute command: " + err.Error())
	}
	r, err := cleanResponse(rawRes)
	if err != nil {
		return "", nil, errors.New("Failed to decode response: " + err.Error())
	}
	lenR := len(r)
	if lenR == 0 {
		return "", nil, errors.New("Failed to decode response: " + err.Error())
	}
	if r[0] != cmd.Request || (len(r)-2) != int(cmd.LinesInResponse) {
		return "", nil, errors.New("Invalid Response received from Device")
	}
	var res bytes.Buffer
	for i := 1; i < (len(r) - 1); i++ {
		res.Write([]byte(r[i]))
	}
	return r[(lenR - 1)], res.Bytes(), err
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
