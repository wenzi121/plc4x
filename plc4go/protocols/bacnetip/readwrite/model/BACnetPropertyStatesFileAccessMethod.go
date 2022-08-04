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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesFileAccessMethod is the corresponding interface of BACnetPropertyStatesFileAccessMethod
type BACnetPropertyStatesFileAccessMethod interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetFileAccessMethod returns FileAccessMethod (property field)
	GetFileAccessMethod() BACnetFileAccessMethodTagged
}

// BACnetPropertyStatesFileAccessMethodExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesFileAccessMethod.
// This is useful for switch cases.
type BACnetPropertyStatesFileAccessMethodExactly interface {
	BACnetPropertyStatesFileAccessMethod
	isBACnetPropertyStatesFileAccessMethod() bool
}

// _BACnetPropertyStatesFileAccessMethod is the data-structure of this message
type _BACnetPropertyStatesFileAccessMethod struct {
	*_BACnetPropertyStates
	FileAccessMethod BACnetFileAccessMethodTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesFileAccessMethod) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesFileAccessMethod) GetFileAccessMethod() BACnetFileAccessMethodTagged {
	return m.FileAccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesFileAccessMethod factory function for _BACnetPropertyStatesFileAccessMethod
func NewBACnetPropertyStatesFileAccessMethod(fileAccessMethod BACnetFileAccessMethodTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesFileAccessMethod {
	_result := &_BACnetPropertyStatesFileAccessMethod{
		FileAccessMethod:      fileAccessMethod,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesFileAccessMethod(structType interface{}) BACnetPropertyStatesFileAccessMethod {
	if casted, ok := structType.(BACnetPropertyStatesFileAccessMethod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesFileAccessMethod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetTypeName() string {
	return "BACnetPropertyStatesFileAccessMethod"
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileAccessMethod)
	lengthInBits += m.FileAccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesFileAccessMethodParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesFileAccessMethod, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesFileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesFileAccessMethod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileAccessMethod)
	if pullErr := readBuffer.PullContext("fileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileAccessMethod")
	}
	_fileAccessMethod, _fileAccessMethodErr := BACnetFileAccessMethodTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _fileAccessMethodErr != nil {
		return nil, errors.Wrap(_fileAccessMethodErr, "Error parsing 'fileAccessMethod' field of BACnetPropertyStatesFileAccessMethod")
	}
	fileAccessMethod := _fileAccessMethod.(BACnetFileAccessMethodTagged)
	if closeErr := readBuffer.CloseContext("fileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileAccessMethod")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesFileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesFileAccessMethod")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesFileAccessMethod{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		FileAccessMethod:      fileAccessMethod,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesFileAccessMethod) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesFileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesFileAccessMethod")
		}

		// Simple Field (fileAccessMethod)
		if pushErr := writeBuffer.PushContext("fileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileAccessMethod")
		}
		_fileAccessMethodErr := writeBuffer.WriteSerializable(m.GetFileAccessMethod())
		if popErr := writeBuffer.PopContext("fileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileAccessMethod")
		}
		if _fileAccessMethodErr != nil {
			return errors.Wrap(_fileAccessMethodErr, "Error serializing 'fileAccessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesFileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesFileAccessMethod")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesFileAccessMethod) isBACnetPropertyStatesFileAccessMethod() bool {
	return true
}

func (m *_BACnetPropertyStatesFileAccessMethod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
