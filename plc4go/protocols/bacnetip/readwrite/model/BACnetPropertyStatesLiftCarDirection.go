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

// BACnetPropertyStatesLiftCarDirection is the corresponding interface of BACnetPropertyStatesLiftCarDirection
type BACnetPropertyStatesLiftCarDirection interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftCarDirection returns LiftCarDirection (property field)
	GetLiftCarDirection() BACnetLiftCarDirectionTagged
}

// BACnetPropertyStatesLiftCarDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLiftCarDirection.
// This is useful for switch cases.
type BACnetPropertyStatesLiftCarDirectionExactly interface {
	BACnetPropertyStatesLiftCarDirection
	isBACnetPropertyStatesLiftCarDirection() bool
}

// _BACnetPropertyStatesLiftCarDirection is the data-structure of this message
type _BACnetPropertyStatesLiftCarDirection struct {
	*_BACnetPropertyStates
	LiftCarDirection BACnetLiftCarDirectionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftCarDirection) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLiftCarDirection) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarDirection) GetLiftCarDirection() BACnetLiftCarDirectionTagged {
	return m.LiftCarDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftCarDirection factory function for _BACnetPropertyStatesLiftCarDirection
func NewBACnetPropertyStatesLiftCarDirection(liftCarDirection BACnetLiftCarDirectionTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLiftCarDirection {
	_result := &_BACnetPropertyStatesLiftCarDirection{
		LiftCarDirection:      liftCarDirection,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarDirection(structType interface{}) BACnetPropertyStatesLiftCarDirection {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarDirection) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarDirection"
}

func (m *_BACnetPropertyStatesLiftCarDirection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLiftCarDirection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftCarDirection)
	lengthInBits += m.LiftCarDirection.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftCarDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftCarDirectionParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLiftCarDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftCarDirection)
	if pullErr := readBuffer.PullContext("liftCarDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for liftCarDirection")
	}
	_liftCarDirection, _liftCarDirectionErr := BACnetLiftCarDirectionTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftCarDirectionErr != nil {
		return nil, errors.Wrap(_liftCarDirectionErr, "Error parsing 'liftCarDirection' field of BACnetPropertyStatesLiftCarDirection")
	}
	liftCarDirection := _liftCarDirection.(BACnetLiftCarDirectionTagged)
	if closeErr := readBuffer.CloseContext("liftCarDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for liftCarDirection")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarDirection")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLiftCarDirection{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LiftCarDirection:      liftCarDirection,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLiftCarDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarDirection")
		}

		// Simple Field (liftCarDirection)
		if pushErr := writeBuffer.PushContext("liftCarDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for liftCarDirection")
		}
		_liftCarDirectionErr := writeBuffer.WriteSerializable(m.GetLiftCarDirection())
		if popErr := writeBuffer.PopContext("liftCarDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for liftCarDirection")
		}
		if _liftCarDirectionErr != nil {
			return errors.Wrap(_liftCarDirectionErr, "Error serializing 'liftCarDirection' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarDirection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftCarDirection) isBACnetPropertyStatesLiftCarDirection() bool {
	return true
}

func (m *_BACnetPropertyStatesLiftCarDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
