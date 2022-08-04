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

// BACnetServiceAckCreateObject is the corresponding interface of BACnetServiceAckCreateObject
type BACnetServiceAckCreateObject interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
}

// BACnetServiceAckCreateObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckCreateObject.
// This is useful for switch cases.
type BACnetServiceAckCreateObjectExactly interface {
	BACnetServiceAckCreateObject
	isBACnetServiceAckCreateObject() bool
}

// _BACnetServiceAckCreateObject is the data-structure of this message
type _BACnetServiceAckCreateObject struct {
	*_BACnetServiceAck
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckCreateObject) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckCreateObject) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckCreateObject factory function for _BACnetServiceAckCreateObject
func NewBACnetServiceAckCreateObject(objectIdentifier BACnetApplicationTagObjectIdentifier, serviceAckLength uint16) *_BACnetServiceAckCreateObject {
	_result := &_BACnetServiceAckCreateObject{
		ObjectIdentifier:  objectIdentifier,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckCreateObject(structType interface{}) BACnetServiceAckCreateObject {
	if casted, ok := structType.(BACnetServiceAckCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckCreateObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckCreateObject) GetTypeName() string {
	return "BACnetServiceAckCreateObject"
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckCreateObjectParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (BACnetServiceAckCreateObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckCreateObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckCreateObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetServiceAckCreateObject")
	}
	objectIdentifier := _objectIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckCreateObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckCreateObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckCreateObject{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier: objectIdentifier,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckCreateObject) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckCreateObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckCreateObject")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckCreateObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckCreateObject")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckCreateObject) isBACnetServiceAckCreateObject() bool {
	return true
}

func (m *_BACnetServiceAckCreateObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
