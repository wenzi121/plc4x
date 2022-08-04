/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestSmartConnectShortcut_PIPE byte = 0x7C

// RequestSmartConnectShortcut is the corresponding interface of RequestSmartConnectShortcut
type RequestSmartConnectShortcut interface {
	utils.LengthAware
	utils.Serializable
	Request
	// GetPipePeek returns PipePeek (property field)
	GetPipePeek() RequestType
	// GetSecondPipe returns SecondPipe (property field)
	GetSecondPipe() *byte
}

// RequestSmartConnectShortcutExactly can be used when we want exactly this type and not a type which fulfills RequestSmartConnectShortcut.
// This is useful for switch cases.
type RequestSmartConnectShortcutExactly interface {
	RequestSmartConnectShortcut
	isRequestSmartConnectShortcut() bool
}

// _RequestSmartConnectShortcut is the data-structure of this message
type _RequestSmartConnectShortcut struct {
	*_Request
	PipePeek   RequestType
	SecondPipe *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestSmartConnectShortcut) InitializeParent(parent Request, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestSmartConnectShortcut) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestSmartConnectShortcut) GetPipePeek() RequestType {
	return m.PipePeek
}

func (m *_RequestSmartConnectShortcut) GetSecondPipe() *byte {
	return m.SecondPipe
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestSmartConnectShortcut) GetPipe() byte {
	return RequestSmartConnectShortcut_PIPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestSmartConnectShortcut factory function for _RequestSmartConnectShortcut
func NewRequestSmartConnectShortcut(pipePeek RequestType, secondPipe *byte, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestSmartConnectShortcut {
	_result := &_RequestSmartConnectShortcut{
		PipePeek:   pipePeek,
		SecondPipe: secondPipe,
		_Request:   NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestSmartConnectShortcut(structType interface{}) RequestSmartConnectShortcut {
	if casted, ok := structType.(RequestSmartConnectShortcut); ok {
		return casted
	}
	if casted, ok := structType.(*RequestSmartConnectShortcut); ok {
		return *casted
	}
	return nil
}

func (m *_RequestSmartConnectShortcut) GetTypeName() string {
	return "RequestSmartConnectShortcut"
}

func (m *_RequestSmartConnectShortcut) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestSmartConnectShortcut) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (pipe)
	lengthInBits += 8

	// Optional Field (secondPipe)
	if m.SecondPipe != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_RequestSmartConnectShortcut) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestSmartConnectShortcutParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (RequestSmartConnectShortcut, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestSmartConnectShortcut"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestSmartConnectShortcut")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (pipe)
	pipe, _pipeErr := readBuffer.ReadByte("pipe")
	if _pipeErr != nil {
		return nil, errors.Wrap(_pipeErr, "Error parsing 'pipe' field of RequestSmartConnectShortcut")
	}
	if pipe != RequestSmartConnectShortcut_PIPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestSmartConnectShortcut_PIPE) + " but got " + fmt.Sprintf("%d", pipe))
	}

	// Peek Field (pipePeek)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("pipePeek"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pipePeek")
	}
	pipePeek, _err := RequestTypeParse(readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'pipePeek' field of RequestSmartConnectShortcut")
	}
	if closeErr := readBuffer.CloseContext("pipePeek"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pipePeek")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (secondPipe) (Can be skipped, if a given expression evaluates to false)
	var secondPipe *byte = nil
	if bool((pipePeek) == (RequestType_SMART_CONNECT_SHORTCUT)) {
		_val, _err := readBuffer.ReadByte("secondPipe")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'secondPipe' field of RequestSmartConnectShortcut")
		}
		secondPipe = &_val
	}

	if closeErr := readBuffer.CloseContext("RequestSmartConnectShortcut"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestSmartConnectShortcut")
	}

	// Create a partially initialized instance
	_child := &_RequestSmartConnectShortcut{
		_Request: &_Request{
			CBusOptions: cBusOptions,
		},
		PipePeek:   pipePeek,
		SecondPipe: secondPipe,
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestSmartConnectShortcut) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestSmartConnectShortcut"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestSmartConnectShortcut")
		}

		// Const Field (pipe)
		_pipeErr := writeBuffer.WriteByte("pipe", 0x7C)
		if _pipeErr != nil {
			return errors.Wrap(_pipeErr, "Error serializing 'pipe' field")
		}

		// Optional Field (secondPipe) (Can be skipped, if the value is null)
		var secondPipe *byte = nil
		if m.GetSecondPipe() != nil {
			secondPipe = m.GetSecondPipe()
			_secondPipeErr := writeBuffer.WriteByte("secondPipe", *(secondPipe))
			if _secondPipeErr != nil {
				return errors.Wrap(_secondPipeErr, "Error serializing 'secondPipe' field")
			}
		}

		if popErr := writeBuffer.PopContext("RequestSmartConnectShortcut"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestSmartConnectShortcut")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestSmartConnectShortcut) isRequestSmartConnectShortcut() bool {
	return true
}

func (m *_RequestSmartConnectShortcut) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
