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

// CALDataRecall is the corresponding interface of CALDataRecall
type CALDataRecall interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCount returns Count (property field)
	GetCount() uint8
}

// CALDataRecallExactly can be used when we want exactly this type and not a type which fulfills CALDataRecall.
// This is useful for switch cases.
type CALDataRecallExactly interface {
	CALDataRecall
	isCALDataRecall() bool
}

// _CALDataRecall is the data-structure of this message
type _CALDataRecall struct {
	*_CALData
	ParamNo Parameter
	Count   uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataRecall) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataRecall) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataRecall) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataRecall) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataRecall factory function for _CALDataRecall
func NewCALDataRecall(paramNo Parameter, count uint8, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataRecall {
	_result := &_CALDataRecall{
		ParamNo:  paramNo,
		Count:    count,
		_CALData: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataRecall(structType interface{}) CALDataRecall {
	if casted, ok := structType.(CALDataRecall); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataRecall); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataRecall) GetTypeName() string {
	return "CALDataRecall"
}

func (m *_CALDataRecall) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataRecall) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataRecall) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRecallParse(readBuffer utils.ReadBuffer, requestContext RequestContext) (CALDataRecall, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRecall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataRecall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	if pullErr := readBuffer.PullContext("paramNo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for paramNo")
	}
	_paramNo, _paramNoErr := ParameterParse(readBuffer)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataRecall")
	}
	paramNo := _paramNo
	if closeErr := readBuffer.CloseContext("paramNo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for paramNo")
	}

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of CALDataRecall")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("CALDataRecall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataRecall")
	}

	// Create a partially initialized instance
	_child := &_CALDataRecall{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		ParamNo: paramNo,
		Count:   count,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataRecall) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRecall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataRecall")
		}

		// Simple Field (paramNo)
		if pushErr := writeBuffer.PushContext("paramNo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for paramNo")
		}
		_paramNoErr := writeBuffer.WriteSerializable(m.GetParamNo())
		if popErr := writeBuffer.PopContext("paramNo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for paramNo")
		}
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := writeBuffer.WriteUint8("count", 8, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRecall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataRecall")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataRecall) isCALDataRecall() bool {
	return true
}

func (m *_CALDataRecall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
