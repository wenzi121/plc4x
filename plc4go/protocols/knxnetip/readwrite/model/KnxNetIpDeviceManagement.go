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

// KnxNetIpDeviceManagement is the corresponding interface of KnxNetIpDeviceManagement
type KnxNetIpDeviceManagement interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpDeviceManagementExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpDeviceManagement.
// This is useful for switch cases.
type KnxNetIpDeviceManagementExactly interface {
	KnxNetIpDeviceManagement
	isKnxNetIpDeviceManagement() bool
}

// _KnxNetIpDeviceManagement is the data-structure of this message
type _KnxNetIpDeviceManagement struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetServiceType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpDeviceManagement) InitializeParent(parent ServiceId) {}

func (m *_KnxNetIpDeviceManagement) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpDeviceManagement factory function for _KnxNetIpDeviceManagement
func NewKnxNetIpDeviceManagement(version uint8) *_KnxNetIpDeviceManagement {
	_result := &_KnxNetIpDeviceManagement{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpDeviceManagement(structType interface{}) KnxNetIpDeviceManagement {
	if casted, ok := structType.(KnxNetIpDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpDeviceManagement) GetTypeName() string {
	return "KnxNetIpDeviceManagement"
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpDeviceManagementParse(readBuffer utils.ReadBuffer) (KnxNetIpDeviceManagement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of KnxNetIpDeviceManagement")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpDeviceManagement")
	}

	// Create a partially initialized instance
	_child := &_KnxNetIpDeviceManagement{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetIpDeviceManagement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpDeviceManagement")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpDeviceManagement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetIpDeviceManagement) isKnxNetIpDeviceManagement() bool {
	return true
}

func (m *_KnxNetIpDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
