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

// CALDataNormalRequestReset is the corresponding interface of CALDataNormalRequestReset
type CALDataNormalRequestReset interface {
	utils.LengthAware
	utils.Serializable
	CALDataNormal
}

// CALDataNormalRequestResetExactly can be used when we want exactly this type and not a type which fulfills CALDataNormalRequestReset.
// This is useful for switch cases.
type CALDataNormalRequestResetExactly interface {
	CALDataNormalRequestReset
	isCALDataNormalRequestReset() bool
}

// _CALDataNormalRequestReset is the data-structure of this message
type _CALDataNormalRequestReset struct {
	*_CALDataNormal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataNormalRequestReset) InitializeParent(parent CALDataNormal, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataNormalRequestReset) GetParent() CALDataNormal {
	return m._CALDataNormal
}

// NewCALDataNormalRequestReset factory function for _CALDataNormalRequestReset
func NewCALDataNormalRequestReset(commandTypeContainer CALCommandTypeContainer) *_CALDataNormalRequestReset {
	_result := &_CALDataNormalRequestReset{
		_CALDataNormal: NewCALDataNormal(commandTypeContainer),
	}
	_result._CALDataNormal._CALDataNormalChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataNormalRequestReset(structType interface{}) CALDataNormalRequestReset {
	if casted, ok := structType.(CALDataNormalRequestReset); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataNormalRequestReset); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataNormalRequestReset) GetTypeName() string {
	return "CALDataNormalRequestReset"
}

func (m *_CALDataNormalRequestReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataNormalRequestReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_CALDataNormalRequestReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataNormalRequestResetParse(readBuffer utils.ReadBuffer) (CALDataNormalRequestReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataNormalRequestReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataNormalRequestReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALDataNormalRequestReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataNormalRequestReset")
	}

	// Create a partially initialized instance
	_child := &_CALDataNormalRequestReset{
		_CALDataNormal: &_CALDataNormal{},
	}
	_child._CALDataNormal._CALDataNormalChildRequirements = _child
	return _child, nil
}

func (m *_CALDataNormalRequestReset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataNormalRequestReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataNormalRequestReset")
		}

		if popErr := writeBuffer.PopContext("CALDataNormalRequestReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataNormalRequestReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataNormalRequestReset) isCALDataNormalRequestReset() bool {
	return true
}

func (m *_CALDataNormalRequestReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
