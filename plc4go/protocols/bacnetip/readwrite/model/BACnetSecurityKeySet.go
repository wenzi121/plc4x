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

// BACnetSecurityKeySet is the corresponding interface of BACnetSecurityKeySet
type BACnetSecurityKeySet interface {
	utils.LengthAware
	utils.Serializable
	// GetKeyRevision returns KeyRevision (property field)
	GetKeyRevision() BACnetContextTagUnsignedInteger
	// GetActivationTime returns ActivationTime (property field)
	GetActivationTime() BACnetDateTimeEnclosed
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() BACnetDateTimeEnclosed
	// GetKeyIds returns KeyIds (property field)
	GetKeyIds() BACnetSecurityKeySetKeyIds
}

// BACnetSecurityKeySetExactly can be used when we want exactly this type and not a type which fulfills BACnetSecurityKeySet.
// This is useful for switch cases.
type BACnetSecurityKeySetExactly interface {
	BACnetSecurityKeySet
	isBACnetSecurityKeySet() bool
}

// _BACnetSecurityKeySet is the data-structure of this message
type _BACnetSecurityKeySet struct {
	KeyRevision    BACnetContextTagUnsignedInteger
	ActivationTime BACnetDateTimeEnclosed
	ExpirationTime BACnetDateTimeEnclosed
	KeyIds         BACnetSecurityKeySetKeyIds
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityKeySet) GetKeyRevision() BACnetContextTagUnsignedInteger {
	return m.KeyRevision
}

func (m *_BACnetSecurityKeySet) GetActivationTime() BACnetDateTimeEnclosed {
	return m.ActivationTime
}

func (m *_BACnetSecurityKeySet) GetExpirationTime() BACnetDateTimeEnclosed {
	return m.ExpirationTime
}

func (m *_BACnetSecurityKeySet) GetKeyIds() BACnetSecurityKeySetKeyIds {
	return m.KeyIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSecurityKeySet factory function for _BACnetSecurityKeySet
func NewBACnetSecurityKeySet(keyRevision BACnetContextTagUnsignedInteger, activationTime BACnetDateTimeEnclosed, expirationTime BACnetDateTimeEnclosed, keyIds BACnetSecurityKeySetKeyIds) *_BACnetSecurityKeySet {
	return &_BACnetSecurityKeySet{KeyRevision: keyRevision, ActivationTime: activationTime, ExpirationTime: expirationTime, KeyIds: keyIds}
}

// Deprecated: use the interface for direct cast
func CastBACnetSecurityKeySet(structType interface{}) BACnetSecurityKeySet {
	if casted, ok := structType.(BACnetSecurityKeySet); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityKeySet); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityKeySet) GetTypeName() string {
	return "BACnetSecurityKeySet"
}

func (m *_BACnetSecurityKeySet) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetSecurityKeySet) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (keyRevision)
	lengthInBits += m.KeyRevision.GetLengthInBits()

	// Simple field (activationTime)
	lengthInBits += m.ActivationTime.GetLengthInBits()

	// Simple field (expirationTime)
	lengthInBits += m.ExpirationTime.GetLengthInBits()

	// Simple field (keyIds)
	lengthInBits += m.KeyIds.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetSecurityKeySet) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSecurityKeySetParse(readBuffer utils.ReadBuffer) (BACnetSecurityKeySet, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSecurityKeySet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityKeySet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (keyRevision)
	if pullErr := readBuffer.PullContext("keyRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for keyRevision")
	}
	_keyRevision, _keyRevisionErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _keyRevisionErr != nil {
		return nil, errors.Wrap(_keyRevisionErr, "Error parsing 'keyRevision' field of BACnetSecurityKeySet")
	}
	keyRevision := _keyRevision.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("keyRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for keyRevision")
	}

	// Simple Field (activationTime)
	if pullErr := readBuffer.PullContext("activationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activationTime")
	}
	_activationTime, _activationTimeErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(1)))
	if _activationTimeErr != nil {
		return nil, errors.Wrap(_activationTimeErr, "Error parsing 'activationTime' field of BACnetSecurityKeySet")
	}
	activationTime := _activationTime.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("activationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activationTime")
	}

	// Simple Field (expirationTime)
	if pullErr := readBuffer.PullContext("expirationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for expirationTime")
	}
	_expirationTime, _expirationTimeErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(2)))
	if _expirationTimeErr != nil {
		return nil, errors.Wrap(_expirationTimeErr, "Error parsing 'expirationTime' field of BACnetSecurityKeySet")
	}
	expirationTime := _expirationTime.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("expirationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for expirationTime")
	}

	// Simple Field (keyIds)
	if pullErr := readBuffer.PullContext("keyIds"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for keyIds")
	}
	_keyIds, _keyIdsErr := BACnetSecurityKeySetKeyIdsParse(readBuffer, uint8(uint8(3)))
	if _keyIdsErr != nil {
		return nil, errors.Wrap(_keyIdsErr, "Error parsing 'keyIds' field of BACnetSecurityKeySet")
	}
	keyIds := _keyIds.(BACnetSecurityKeySetKeyIds)
	if closeErr := readBuffer.CloseContext("keyIds"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for keyIds")
	}

	if closeErr := readBuffer.CloseContext("BACnetSecurityKeySet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityKeySet")
	}

	// Create the instance
	return &_BACnetSecurityKeySet{
		KeyRevision:    keyRevision,
		ActivationTime: activationTime,
		ExpirationTime: expirationTime,
		KeyIds:         keyIds,
	}, nil
}

func (m *_BACnetSecurityKeySet) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetSecurityKeySet"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityKeySet")
	}

	// Simple Field (keyRevision)
	if pushErr := writeBuffer.PushContext("keyRevision"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for keyRevision")
	}
	_keyRevisionErr := writeBuffer.WriteSerializable(m.GetKeyRevision())
	if popErr := writeBuffer.PopContext("keyRevision"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for keyRevision")
	}
	if _keyRevisionErr != nil {
		return errors.Wrap(_keyRevisionErr, "Error serializing 'keyRevision' field")
	}

	// Simple Field (activationTime)
	if pushErr := writeBuffer.PushContext("activationTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for activationTime")
	}
	_activationTimeErr := writeBuffer.WriteSerializable(m.GetActivationTime())
	if popErr := writeBuffer.PopContext("activationTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for activationTime")
	}
	if _activationTimeErr != nil {
		return errors.Wrap(_activationTimeErr, "Error serializing 'activationTime' field")
	}

	// Simple Field (expirationTime)
	if pushErr := writeBuffer.PushContext("expirationTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for expirationTime")
	}
	_expirationTimeErr := writeBuffer.WriteSerializable(m.GetExpirationTime())
	if popErr := writeBuffer.PopContext("expirationTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for expirationTime")
	}
	if _expirationTimeErr != nil {
		return errors.Wrap(_expirationTimeErr, "Error serializing 'expirationTime' field")
	}

	// Simple Field (keyIds)
	if pushErr := writeBuffer.PushContext("keyIds"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for keyIds")
	}
	_keyIdsErr := writeBuffer.WriteSerializable(m.GetKeyIds())
	if popErr := writeBuffer.PopContext("keyIds"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for keyIds")
	}
	if _keyIdsErr != nil {
		return errors.Wrap(_keyIdsErr, "Error serializing 'keyIds' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSecurityKeySet"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityKeySet")
	}
	return nil
}

func (m *_BACnetSecurityKeySet) isBACnetSecurityKeySet() bool {
	return true
}

func (m *_BACnetSecurityKeySet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
