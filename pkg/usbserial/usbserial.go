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

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

const (
	constDefaultBaudRate = 115200
	constDefaultDataBits = 8
	constDefaultStopBits = 1
	constLineFeedASCII   = 0x0A
	constReadBufferSize  = 2000 // Configured to be larger than the response of the largest command.
)

// USBSerial wraps the underlying serial library used.
type USBSerial struct {
	Name           string
	Port           io.ReadWriteCloser
	BlockUntilData bool
	ReadTimeout    uint8 //in seconds.
}

// New creates a USBSerial object.
// Always put `defer USBSerial.close()` immediately after calling this function to close the port on exit.
// Leaving serial ports open will hamper subsequent connections.
func New(portname string, blockUntilData bool, readtimeoutseconds uint8) (*USBSerial, error) {
	var minreadsize uint
	if blockUntilData == true {
		minreadsize = 1
	} else {
		minreadsize = 0
	}

	options := serial.OpenOptions{
		PortName:        portname,
		BaudRate:        constDefaultBaudRate,
		DataBits:        constDefaultDataBits,
		StopBits:        constDefaultStopBits,
		MinimumReadSize: minreadsize,
	}
	port, err := serial.Open(options)
	if err != nil {
		return nil, err
	}
	us := &USBSerial{
		Name:           portname,
		Port:           port,
		BlockUntilData: blockUntilData,
		ReadTimeout:    readtimeoutseconds,
	}
	return us, nil
}

// Close closes the underlying serial port.
// This function must always be called before exiting the program.
func (us *USBSerial) Close() {
	us.Port.Close()
}

// SendData sends the data on the serial buffer.
// In the case of devices that echo inputs, the buffer has to be read twice; once for the echo and once for the response.
// This is handled by continuously reading until a timeout.
// This function is not thread-safe and should not be used concurrently.
// The `waitperiod` parameter defines the wait time before reading the response buffer.
//   - For normal commands, use 1.
//   - For commands with larger response sizes, use 2.
func (us *USBSerial) SendData(data []byte, waitperiod string) error {
	var err error
	var n int
	//Write the command
	var sendBuffer bytes.Buffer
	sendBuffer.Write([]byte(data))
	sendBuffer.WriteByte(constLineFeedASCII)
	n, err = us.Port.Write(sendBuffer.Bytes())
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("No bytes written on serial port")
	}

	//Read response
	var waittime time.Duration
	t, err := time.ParseDuration(waitperiod)
	if err != nil {
		waittime = 1 * time.Second
	} else {
		waittime = t * time.Second
	}
	recBuf := make([]byte, constReadBufferSize) // The serial interface needs a buffer of fixed size and hence this intermediate buffer is used.
	ch := make(chan int)
	for i := 0; i < 1; i++ {
		go func() {
			time.Sleep(waittime) // allow some delay to read let the buffer get filled with the entire response.
			n, err = us.Port.Read(recBuf)
			fmt.Println(recBuf)
			fmt.Println(string(recBuf))
			ch <- 1
		}()

		<-ch //The function is blocked here until the go routine above is completed.
		if err != nil {
			return err
		}
	}
	return nil
}

// ScanPorts scans all serial ports and returns a list of available ports.
func (us *USBSerial) ScanPorts() []string {
	return nil
}
