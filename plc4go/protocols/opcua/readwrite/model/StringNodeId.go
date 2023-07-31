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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// StringNodeId is the corresponding interface of StringNodeId
type StringNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() PascalString
}

// StringNodeIdExactly can be used when we want exactly this type and not a type which fulfills StringNodeId.
// This is useful for switch cases.
type StringNodeIdExactly interface {
	StringNodeId
	isStringNodeId() bool
}

// _StringNodeId is the data-structure of this message
type _StringNodeId struct {
	NamespaceIndex uint16
	Identifier     PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StringNodeId) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_StringNodeId) GetIdentifier() PascalString {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStringNodeId factory function for _StringNodeId
func NewStringNodeId(namespaceIndex uint16, identifier PascalString) *_StringNodeId {
	return &_StringNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

// Deprecated: use the interface for direct cast
func CastStringNodeId(structType any) StringNodeId {
	if casted, ok := structType.(StringNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*StringNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_StringNodeId) GetTypeName() string {
	return "StringNodeId"
}

func (m *_StringNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (identifier)
	lengthInBits += m.Identifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StringNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StringNodeIdParse(ctx context.Context, theBytes []byte) (StringNodeId, error) {
	return StringNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StringNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StringNodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("StringNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StringNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of StringNodeId")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (identifier)
	if pullErr := readBuffer.PullContext("identifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for identifier")
	}
	_identifier, _identifierErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _identifierErr != nil {
		return nil, errors.Wrap(_identifierErr, "Error parsing 'identifier' field of StringNodeId")
	}
	identifier := _identifier.(PascalString)
	if closeErr := readBuffer.CloseContext("identifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for identifier")
	}

	if closeErr := readBuffer.CloseContext("StringNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StringNodeId")
	}

	// Create the instance
	return &_StringNodeId{
		NamespaceIndex: namespaceIndex,
		Identifier:     identifier,
	}, nil
}

func (m *_StringNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StringNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("StringNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StringNodeId")
	}

	// Simple Field (namespaceIndex)
	namespaceIndex := uint16(m.GetNamespaceIndex())
	_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, (namespaceIndex))
	if _namespaceIndexErr != nil {
		return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
	}

	// Simple Field (identifier)
	if pushErr := writeBuffer.PushContext("identifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for identifier")
	}
	_identifierErr := writeBuffer.WriteSerializable(ctx, m.GetIdentifier())
	if popErr := writeBuffer.PopContext("identifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for identifier")
	}
	if _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("StringNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StringNodeId")
	}
	return nil
}

func (m *_StringNodeId) isStringNodeId() bool {
	return true
}

func (m *_StringNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}