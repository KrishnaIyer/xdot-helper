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
	sequence "github.com/KrishnaIyer/xdot-helper/pkg/atcommands/sequence"
)

// Defines the commands that are to be executed during the Join Sequence.
const (
	XDOTLoraATPing                 = "AT"
	XDOTLoraATFactoryReset         = "AT&F"
	XDOTLoraATEnablePrivateNetwork = "AT+PN=1"
	XDOTLoraATGetDevEUI            = "AT+DI"
	XDOTLoraATRX2Freq              = "AT+RXF=869525000"
	XDOTLoraATJoinModeOTAA         = "AT+NJM=1"
	XDOTLoraATJoinRX2DataRate      = "AT+MAC=0503d2ad84"
	XDOTLoraATNoJoinRetries        = "AT+JR=0"
	XDOTLoraATSaveConfig           = "AT&W"
	XDOTLoraATJoin                 = "AT+JOIN"
	XDOTLoraATWriteNwkKey          = "AT+NK=0,"
	XDOTLoraATWriteAppEUI          = "AT+NI=0,"
	XDOTLoraATJoinStatus           = "AT+NJS"
	XDOTLoraATSetTxDataRate        = "AT+TXDR=7"
)

// New returns a new sequence.Sequence that contains the steps for registering a device to a Network.
func New(device, nwkKey, appEUI string) (*sequence.Sequence, error) {
	s := sequence.New(device, 1, true)
	//Create the sequence.
	reqs := []string{XDOTLoraATFactoryReset, XDOTLoraATEnablePrivateNetwork, XDOTLoraATWriteAppEUI + appEUI, XDOTLoraATWriteNwkKey + nwkKey, XDOTLoraATNoJoinRetries, XDOTLoraATSaveConfig}
	s.MakeSequenceFromReqList(reqs)
	return s, nil
}
