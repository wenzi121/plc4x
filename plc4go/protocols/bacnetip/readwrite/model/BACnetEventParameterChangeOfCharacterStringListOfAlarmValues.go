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

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the corresponding interface of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() []BACnetApplicationTagCharacterString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfCharacterStringListOfAlarmValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesExactly interface {
	BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	isBACnetEventParameterChangeOfCharacterStringListOfAlarmValues() bool
}

// _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfAlarmValues []BACnetApplicationTagCharacterString
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetListOfAlarmValues() []BACnetApplicationTagCharacterString {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues factory function for _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
func NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{OpeningTag: openingTag, ListOfAlarmValues: listOfAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(structType interface{}) BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfAlarmValues) > 0 {
		for _, element := range m.ListOfAlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfAlarmValues)
	if pullErr := readBuffer.PullContext("listOfAlarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfAlarmValues")
	}
	// Terminated array
	var listOfAlarmValues []BACnetApplicationTagCharacterString
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfAlarmValues' field of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
			}
			listOfAlarmValues = append(listOfAlarmValues, _item.(BACnetApplicationTagCharacterString))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfAlarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfAlarmValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{
		TagNumber:         tagNumber,
		OpeningTag:        openingTag,
		ListOfAlarmValues: listOfAlarmValues,
		ClosingTag:        closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfAlarmValues)
	if pushErr := writeBuffer.PushContext("listOfAlarmValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfAlarmValues")
	}
	for _, _element := range m.GetListOfAlarmValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfAlarmValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfAlarmValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfAlarmValues")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) isBACnetEventParameterChangeOfCharacterStringListOfAlarmValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
