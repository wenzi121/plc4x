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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReadProperty is the corresponding interface of BACnetConfirmedServiceRequestReadProperty
type BACnetConfirmedServiceRequestReadProperty interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestReadPropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadProperty.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadPropertyExactly interface {
	BACnetConfirmedServiceRequestReadProperty
	isBACnetConfirmedServiceRequestReadProperty() bool
}

// _BACnetConfirmedServiceRequestReadProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestReadProperty struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadProperty factory function for _BACnetConfirmedServiceRequestReadProperty
func NewBACnetConfirmedServiceRequestReadProperty(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestReadProperty {
	_result := &_BACnetConfirmedServiceRequestReadProperty{
		ObjectIdentifier:               objectIdentifier,
		PropertyIdentifier:             propertyIdentifier,
		ArrayIndex:                     arrayIndex,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadProperty(structType interface{}) BACnetConfirmedServiceRequestReadProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadProperty"
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadPropertyParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestReadProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetConfirmedServiceRequestReadProperty")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetConfirmedServiceRequestReadProperty")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetConfirmedServiceRequestReadProperty")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadProperty{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadProperty")
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

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
		}
		_propertyIdentifierErr := writeBuffer.WriteSerializable(m.GetPropertyIdentifier())
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyIdentifier")
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (arrayIndex) (Can be skipped, if the value is null)
		var arrayIndex BACnetContextTagUnsignedInteger = nil
		if m.GetArrayIndex() != nil {
			if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for arrayIndex")
			}
			arrayIndex = m.GetArrayIndex()
			_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
			if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for arrayIndex")
			}
			if _arrayIndexErr != nil {
				return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadProperty) isBACnetConfirmedServiceRequestReadProperty() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
