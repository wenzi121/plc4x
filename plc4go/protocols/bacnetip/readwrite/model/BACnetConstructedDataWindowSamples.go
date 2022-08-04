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

// BACnetConstructedDataWindowSamples is the corresponding interface of BACnetConstructedDataWindowSamples
type BACnetConstructedDataWindowSamples interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetWindowSamples returns WindowSamples (property field)
	GetWindowSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataWindowSamplesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataWindowSamples.
// This is useful for switch cases.
type BACnetConstructedDataWindowSamplesExactly interface {
	BACnetConstructedDataWindowSamples
	isBACnetConstructedDataWindowSamples() bool
}

// _BACnetConstructedDataWindowSamples is the data-structure of this message
type _BACnetConstructedDataWindowSamples struct {
	*_BACnetConstructedData
	WindowSamples BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWindowSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WINDOW_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWindowSamples) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataWindowSamples) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetWindowSamples() BACnetApplicationTagUnsignedInteger {
	return m.WindowSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetWindowSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWindowSamples factory function for _BACnetConstructedDataWindowSamples
func NewBACnetConstructedDataWindowSamples(windowSamples BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWindowSamples {
	_result := &_BACnetConstructedDataWindowSamples{
		WindowSamples:          windowSamples,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataWindowSamples(structType interface{}) BACnetConstructedDataWindowSamples {
	if casted, ok := structType.(BACnetConstructedDataWindowSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWindowSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWindowSamples) GetTypeName() string {
	return "BACnetConstructedDataWindowSamples"
}

func (m *_BACnetConstructedDataWindowSamples) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataWindowSamples) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (windowSamples)
	lengthInBits += m.WindowSamples.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataWindowSamples) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataWindowSamplesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataWindowSamples, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWindowSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWindowSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (windowSamples)
	if pullErr := readBuffer.PullContext("windowSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for windowSamples")
	}
	_windowSamples, _windowSamplesErr := BACnetApplicationTagParse(readBuffer)
	if _windowSamplesErr != nil {
		return nil, errors.Wrap(_windowSamplesErr, "Error parsing 'windowSamples' field of BACnetConstructedDataWindowSamples")
	}
	windowSamples := _windowSamples.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("windowSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for windowSamples")
	}

	// Virtual field
	_actualValue := windowSamples
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWindowSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWindowSamples")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataWindowSamples{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		WindowSamples: windowSamples,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataWindowSamples) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWindowSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWindowSamples")
		}

		// Simple Field (windowSamples)
		if pushErr := writeBuffer.PushContext("windowSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for windowSamples")
		}
		_windowSamplesErr := writeBuffer.WriteSerializable(m.GetWindowSamples())
		if popErr := writeBuffer.PopContext("windowSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for windowSamples")
		}
		if _windowSamplesErr != nil {
			return errors.Wrap(_windowSamplesErr, "Error serializing 'windowSamples' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWindowSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWindowSamples")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWindowSamples) isBACnetConstructedDataWindowSamples() bool {
	return true
}

func (m *_BACnetConstructedDataWindowSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
