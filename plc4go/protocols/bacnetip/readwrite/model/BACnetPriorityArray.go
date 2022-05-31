/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetPriorityArray is the data-structure of this message
type BACnetPriorityArray struct {
	Data []*BACnetPriorityValue

	// Arguments.
	ObjectType BACnetObjectType
	TagNumber  uint8
}

// IBACnetPriorityArray is the corresponding interface of BACnetPriorityArray
type IBACnetPriorityArray interface {
	// GetData returns Data (property field)
	GetData() []*BACnetPriorityValue
	// GetPriorityValue01 returns PriorityValue01 (virtual field)
	GetPriorityValue01() *BACnetPriorityValue
	// GetPriorityValue02 returns PriorityValue02 (virtual field)
	GetPriorityValue02() *BACnetPriorityValue
	// GetPriorityValue03 returns PriorityValue03 (virtual field)
	GetPriorityValue03() *BACnetPriorityValue
	// GetPriorityValue04 returns PriorityValue04 (virtual field)
	GetPriorityValue04() *BACnetPriorityValue
	// GetPriorityValue05 returns PriorityValue05 (virtual field)
	GetPriorityValue05() *BACnetPriorityValue
	// GetPriorityValue06 returns PriorityValue06 (virtual field)
	GetPriorityValue06() *BACnetPriorityValue
	// GetPriorityValue07 returns PriorityValue07 (virtual field)
	GetPriorityValue07() *BACnetPriorityValue
	// GetPriorityValue08 returns PriorityValue08 (virtual field)
	GetPriorityValue08() *BACnetPriorityValue
	// GetPriorityValue09 returns PriorityValue09 (virtual field)
	GetPriorityValue09() *BACnetPriorityValue
	// GetPriorityValue10 returns PriorityValue10 (virtual field)
	GetPriorityValue10() *BACnetPriorityValue
	// GetPriorityValue11 returns PriorityValue11 (virtual field)
	GetPriorityValue11() *BACnetPriorityValue
	// GetPriorityValue12 returns PriorityValue12 (virtual field)
	GetPriorityValue12() *BACnetPriorityValue
	// GetPriorityValue13 returns PriorityValue13 (virtual field)
	GetPriorityValue13() *BACnetPriorityValue
	// GetPriorityValue14 returns PriorityValue14 (virtual field)
	GetPriorityValue14() *BACnetPriorityValue
	// GetPriorityValue15 returns PriorityValue15 (virtual field)
	GetPriorityValue15() *BACnetPriorityValue
	// GetPriorityValue16 returns PriorityValue16 (virtual field)
	GetPriorityValue16() *BACnetPriorityValue
	// GetIsValidPriorityArray returns IsValidPriorityArray (virtual field)
	GetIsValidPriorityArray() bool
	// GetIsIndexedAccess returns IsIndexedAccess (virtual field)
	GetIsIndexedAccess() bool
	// GetIndexEntry returns IndexEntry (virtual field)
	GetIndexEntry() *BACnetPriorityValue
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityArray) GetData() []*BACnetPriorityValue {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetPriorityArray) GetPriorityValue01() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (0)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[0]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue02() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (1)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[1]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue03() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (2)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[2]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue04() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (3)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[3]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue05() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (4)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[4]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue06() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (5)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[5]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue07() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (6)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[6]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue08() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (7)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[7]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue09() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (8)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[8]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue10() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (9)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[9]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue11() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (10)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[10]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue12() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (11)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[11]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue13() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (12)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[12]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue14() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (13)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[13]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue15() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (14)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[14]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetPriorityValue16() *BACnetPriorityValue {
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (15)), func() interface{} { return CastBACnetPriorityValue(m.GetData()[15]) }, func() interface{} { return CastBACnetPriorityValue(nil) })))
}

func (m *BACnetPriorityArray) GetIsValidPriorityArray() bool {
	return bool(bool((len(m.GetData())) == (16)))
}

func (m *BACnetPriorityArray) GetIsIndexedAccess() bool {
	return bool(bool((len(m.GetData())) == (1)))
}

func (m *BACnetPriorityArray) GetIndexEntry() *BACnetPriorityValue {
	return CastBACnetPriorityValue(m.GetPriorityValue01())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityArray factory function for BACnetPriorityArray
func NewBACnetPriorityArray(data []*BACnetPriorityValue, objectType BACnetObjectType, tagNumber uint8) *BACnetPriorityArray {
	return &BACnetPriorityArray{Data: data, ObjectType: objectType, TagNumber: tagNumber}
}

func CastBACnetPriorityArray(structType interface{}) *BACnetPriorityArray {
	if casted, ok := structType.(BACnetPriorityArray); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityArray); ok {
		return casted
	}
	return nil
}

func (m *BACnetPriorityArray) GetTypeName() string {
	return "BACnetPriorityArray"
}

func (m *BACnetPriorityArray) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityArray) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetPriorityArray) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityArrayParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType, tagNumber uint8) (*BACnetPriorityArray, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityArray"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	data := make([]*BACnetPriorityValue, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetPriorityValueParse(readBuffer, objectType)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, CastBACnetPriorityValue(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_priorityValue01 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (0)), func() interface{} { return CastBACnetPriorityValue(data[0]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue01 := CastBACnetPriorityValue(_priorityValue01)
	_ = priorityValue01

	// Virtual field
	_priorityValue02 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (1)), func() interface{} { return CastBACnetPriorityValue(data[1]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue02 := CastBACnetPriorityValue(_priorityValue02)
	_ = priorityValue02

	// Virtual field
	_priorityValue03 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (2)), func() interface{} { return CastBACnetPriorityValue(data[2]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue03 := CastBACnetPriorityValue(_priorityValue03)
	_ = priorityValue03

	// Virtual field
	_priorityValue04 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (3)), func() interface{} { return CastBACnetPriorityValue(data[3]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue04 := CastBACnetPriorityValue(_priorityValue04)
	_ = priorityValue04

	// Virtual field
	_priorityValue05 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (4)), func() interface{} { return CastBACnetPriorityValue(data[4]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue05 := CastBACnetPriorityValue(_priorityValue05)
	_ = priorityValue05

	// Virtual field
	_priorityValue06 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (5)), func() interface{} { return CastBACnetPriorityValue(data[5]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue06 := CastBACnetPriorityValue(_priorityValue06)
	_ = priorityValue06

	// Virtual field
	_priorityValue07 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (6)), func() interface{} { return CastBACnetPriorityValue(data[6]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue07 := CastBACnetPriorityValue(_priorityValue07)
	_ = priorityValue07

	// Virtual field
	_priorityValue08 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (7)), func() interface{} { return CastBACnetPriorityValue(data[7]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue08 := CastBACnetPriorityValue(_priorityValue08)
	_ = priorityValue08

	// Virtual field
	_priorityValue09 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (8)), func() interface{} { return CastBACnetPriorityValue(data[8]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue09 := CastBACnetPriorityValue(_priorityValue09)
	_ = priorityValue09

	// Virtual field
	_priorityValue10 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (9)), func() interface{} { return CastBACnetPriorityValue(data[9]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue10 := CastBACnetPriorityValue(_priorityValue10)
	_ = priorityValue10

	// Virtual field
	_priorityValue11 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (10)), func() interface{} { return CastBACnetPriorityValue(data[10]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue11 := CastBACnetPriorityValue(_priorityValue11)
	_ = priorityValue11

	// Virtual field
	_priorityValue12 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (11)), func() interface{} { return CastBACnetPriorityValue(data[11]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue12 := CastBACnetPriorityValue(_priorityValue12)
	_ = priorityValue12

	// Virtual field
	_priorityValue13 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (12)), func() interface{} { return CastBACnetPriorityValue(data[12]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue13 := CastBACnetPriorityValue(_priorityValue13)
	_ = priorityValue13

	// Virtual field
	_priorityValue14 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (13)), func() interface{} { return CastBACnetPriorityValue(data[13]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue14 := CastBACnetPriorityValue(_priorityValue14)
	_ = priorityValue14

	// Virtual field
	_priorityValue15 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (14)), func() interface{} { return CastBACnetPriorityValue(data[14]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue15 := CastBACnetPriorityValue(_priorityValue15)
	_ = priorityValue15

	// Virtual field
	_priorityValue16 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (15)), func() interface{} { return CastBACnetPriorityValue(data[15]) }, func() interface{} { return CastBACnetPriorityValue(nil) }))
	priorityValue16 := CastBACnetPriorityValue(_priorityValue16)
	_ = priorityValue16

	// Virtual field
	_isValidPriorityArray := bool((len(data)) == (16))
	isValidPriorityArray := bool(_isValidPriorityArray)
	_ = isValidPriorityArray

	// Virtual field
	_isIndexedAccess := bool((len(data)) == (1))
	isIndexedAccess := bool(_isIndexedAccess)
	_ = isIndexedAccess

	// Virtual field
	_indexEntry := priorityValue01
	indexEntry := CastBACnetPriorityValue(_indexEntry)
	_ = indexEntry

	if closeErr := readBuffer.CloseContext("BACnetPriorityArray"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetPriorityArray(data, objectType, tagNumber), nil
}

func (m *BACnetPriorityArray) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPriorityArray"); pushErr != nil {
		return pushErr
	}

	// Array Field (data)
	if m.Data != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Data {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}
	// Virtual field
	if _priorityValue01Err := writeBuffer.WriteVirtual("priorityValue01", m.GetPriorityValue01()); _priorityValue01Err != nil {
		return errors.Wrap(_priorityValue01Err, "Error serializing 'priorityValue01' field")
	}
	// Virtual field
	if _priorityValue02Err := writeBuffer.WriteVirtual("priorityValue02", m.GetPriorityValue02()); _priorityValue02Err != nil {
		return errors.Wrap(_priorityValue02Err, "Error serializing 'priorityValue02' field")
	}
	// Virtual field
	if _priorityValue03Err := writeBuffer.WriteVirtual("priorityValue03", m.GetPriorityValue03()); _priorityValue03Err != nil {
		return errors.Wrap(_priorityValue03Err, "Error serializing 'priorityValue03' field")
	}
	// Virtual field
	if _priorityValue04Err := writeBuffer.WriteVirtual("priorityValue04", m.GetPriorityValue04()); _priorityValue04Err != nil {
		return errors.Wrap(_priorityValue04Err, "Error serializing 'priorityValue04' field")
	}
	// Virtual field
	if _priorityValue05Err := writeBuffer.WriteVirtual("priorityValue05", m.GetPriorityValue05()); _priorityValue05Err != nil {
		return errors.Wrap(_priorityValue05Err, "Error serializing 'priorityValue05' field")
	}
	// Virtual field
	if _priorityValue06Err := writeBuffer.WriteVirtual("priorityValue06", m.GetPriorityValue06()); _priorityValue06Err != nil {
		return errors.Wrap(_priorityValue06Err, "Error serializing 'priorityValue06' field")
	}
	// Virtual field
	if _priorityValue07Err := writeBuffer.WriteVirtual("priorityValue07", m.GetPriorityValue07()); _priorityValue07Err != nil {
		return errors.Wrap(_priorityValue07Err, "Error serializing 'priorityValue07' field")
	}
	// Virtual field
	if _priorityValue08Err := writeBuffer.WriteVirtual("priorityValue08", m.GetPriorityValue08()); _priorityValue08Err != nil {
		return errors.Wrap(_priorityValue08Err, "Error serializing 'priorityValue08' field")
	}
	// Virtual field
	if _priorityValue09Err := writeBuffer.WriteVirtual("priorityValue09", m.GetPriorityValue09()); _priorityValue09Err != nil {
		return errors.Wrap(_priorityValue09Err, "Error serializing 'priorityValue09' field")
	}
	// Virtual field
	if _priorityValue10Err := writeBuffer.WriteVirtual("priorityValue10", m.GetPriorityValue10()); _priorityValue10Err != nil {
		return errors.Wrap(_priorityValue10Err, "Error serializing 'priorityValue10' field")
	}
	// Virtual field
	if _priorityValue11Err := writeBuffer.WriteVirtual("priorityValue11", m.GetPriorityValue11()); _priorityValue11Err != nil {
		return errors.Wrap(_priorityValue11Err, "Error serializing 'priorityValue11' field")
	}
	// Virtual field
	if _priorityValue12Err := writeBuffer.WriteVirtual("priorityValue12", m.GetPriorityValue12()); _priorityValue12Err != nil {
		return errors.Wrap(_priorityValue12Err, "Error serializing 'priorityValue12' field")
	}
	// Virtual field
	if _priorityValue13Err := writeBuffer.WriteVirtual("priorityValue13", m.GetPriorityValue13()); _priorityValue13Err != nil {
		return errors.Wrap(_priorityValue13Err, "Error serializing 'priorityValue13' field")
	}
	// Virtual field
	if _priorityValue14Err := writeBuffer.WriteVirtual("priorityValue14", m.GetPriorityValue14()); _priorityValue14Err != nil {
		return errors.Wrap(_priorityValue14Err, "Error serializing 'priorityValue14' field")
	}
	// Virtual field
	if _priorityValue15Err := writeBuffer.WriteVirtual("priorityValue15", m.GetPriorityValue15()); _priorityValue15Err != nil {
		return errors.Wrap(_priorityValue15Err, "Error serializing 'priorityValue15' field")
	}
	// Virtual field
	if _priorityValue16Err := writeBuffer.WriteVirtual("priorityValue16", m.GetPriorityValue16()); _priorityValue16Err != nil {
		return errors.Wrap(_priorityValue16Err, "Error serializing 'priorityValue16' field")
	}
	// Virtual field
	if _isValidPriorityArrayErr := writeBuffer.WriteVirtual("isValidPriorityArray", m.GetIsValidPriorityArray()); _isValidPriorityArrayErr != nil {
		return errors.Wrap(_isValidPriorityArrayErr, "Error serializing 'isValidPriorityArray' field")
	}
	// Virtual field
	if _isIndexedAccessErr := writeBuffer.WriteVirtual("isIndexedAccess", m.GetIsIndexedAccess()); _isIndexedAccessErr != nil {
		return errors.Wrap(_isIndexedAccessErr, "Error serializing 'isIndexedAccess' field")
	}
	// Virtual field
	if _indexEntryErr := writeBuffer.WriteVirtual("indexEntry", m.GetIndexEntry()); _indexEntryErr != nil {
		return errors.Wrap(_indexEntryErr, "Error serializing 'indexEntry' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityArray"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPriorityArray) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
