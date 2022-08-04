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

// BACnetConstructedDataAccessEventTag is the corresponding interface of BACnetConstructedDataAccessEventTag
type BACnetConstructedDataAccessEventTag interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEventTag returns AccessEventTag (property field)
	GetAccessEventTag() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccessEventTagExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessEventTag.
// This is useful for switch cases.
type BACnetConstructedDataAccessEventTagExactly interface {
	BACnetConstructedDataAccessEventTag
	isBACnetConstructedDataAccessEventTag() bool
}

// _BACnetConstructedDataAccessEventTag is the data-structure of this message
type _BACnetConstructedDataAccessEventTag struct {
	*_BACnetConstructedData
	AccessEventTag BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTag) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventTag) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_TAG
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventTag) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessEventTag) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTag) GetAccessEventTag() BACnetApplicationTagUnsignedInteger {
	return m.AccessEventTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTag) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetAccessEventTag())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventTag factory function for _BACnetConstructedDataAccessEventTag
func NewBACnetConstructedDataAccessEventTag(accessEventTag BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventTag {
	_result := &_BACnetConstructedDataAccessEventTag{
		AccessEventTag:         accessEventTag,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventTag(structType interface{}) BACnetConstructedDataAccessEventTag {
	if casted, ok := structType.(BACnetConstructedDataAccessEventTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventTag) GetTypeName() string {
	return "BACnetConstructedDataAccessEventTag"
}

func (m *_BACnetConstructedDataAccessEventTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccessEventTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessEventTag)
	lengthInBits += m.AccessEventTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessEventTagParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEventTag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEventTag)
	if pullErr := readBuffer.PullContext("accessEventTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventTag")
	}
	_accessEventTag, _accessEventTagErr := BACnetApplicationTagParse(readBuffer)
	if _accessEventTagErr != nil {
		return nil, errors.Wrap(_accessEventTagErr, "Error parsing 'accessEventTag' field of BACnetConstructedDataAccessEventTag")
	}
	accessEventTag := _accessEventTag.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("accessEventTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventTag")
	}

	// Virtual field
	_actualValue := accessEventTag
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventTag")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessEventTag{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AccessEventTag: accessEventTag,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessEventTag) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventTag")
		}

		// Simple Field (accessEventTag)
		if pushErr := writeBuffer.PushContext("accessEventTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventTag")
		}
		_accessEventTagErr := writeBuffer.WriteSerializable(m.GetAccessEventTag())
		if popErr := writeBuffer.PopContext("accessEventTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventTag")
		}
		if _accessEventTagErr != nil {
			return errors.Wrap(_accessEventTagErr, "Error serializing 'accessEventTag' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventTag")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventTag) isBACnetConstructedDataAccessEventTag() bool {
	return true
}

func (m *_BACnetConstructedDataAccessEventTag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
