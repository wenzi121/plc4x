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

// ReplyEncodedReply is the corresponding interface of ReplyEncodedReply
type ReplyEncodedReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetEncodedReply returns EncodedReply (property field)
	GetEncodedReply() EncodedReply
	// GetChksum returns Chksum (property field)
	GetChksum() Checksum
}

// ReplyEncodedReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyEncodedReply.
// This is useful for switch cases.
type ReplyEncodedReplyExactly interface {
	ReplyEncodedReply
	isReplyEncodedReply() bool
}

// _ReplyEncodedReply is the data-structure of this message
type _ReplyEncodedReply struct {
	*_Reply
	EncodedReply EncodedReply
	Chksum       Checksum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyEncodedReply) InitializeParent(parent Reply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyEncodedReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyEncodedReply) GetEncodedReply() EncodedReply {
	return m.EncodedReply
}

func (m *_ReplyEncodedReply) GetChksum() Checksum {
	return m.Chksum
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyEncodedReply factory function for _ReplyEncodedReply
func NewReplyEncodedReply(encodedReply EncodedReply, chksum Checksum, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyEncodedReply {
	_result := &_ReplyEncodedReply{
		EncodedReply: encodedReply,
		Chksum:       chksum,
		_Reply:       NewReply(peekedByte, cBusOptions, requestContext),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyEncodedReply(structType interface{}) ReplyEncodedReply {
	if casted, ok := structType.(ReplyEncodedReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyEncodedReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyEncodedReply) GetTypeName() string {
	return "ReplyEncodedReply"
}

func (m *_ReplyEncodedReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ReplyEncodedReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Manual Field (encodedReply)
	lengthInBits += uint16(int32((int32(m.GetEncodedReply().GetLengthInBytes()) * int32(int32(2)))) * int32(int32(8)))

	// Manual Field (chksum)
	lengthInBits += uint16(utils.InlineIf((m.CBusOptions.GetSrchk()), func() interface{} { return int32((int32(16))) }, func() interface{} { return int32((int32(0))) }).(int32))

	return lengthInBits
}

func (m *_ReplyEncodedReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyEncodedReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ReplyEncodedReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyEncodedReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyEncodedReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (encodedReply)
	_encodedReply, _encodedReplyErr := ReadEncodedReply(readBuffer, cBusOptions, requestContext, cBusOptions.GetSrchk())
	if _encodedReplyErr != nil {
		return nil, errors.Wrap(_encodedReplyErr, "Error parsing 'encodedReply' field of ReplyEncodedReply")
	}
	var encodedReply EncodedReply
	if _encodedReply != nil {
		encodedReply = _encodedReply.(EncodedReply)
	}

	// Manual Field (chksum)
	_chksum, _chksumErr := ReadAndValidateChecksum(readBuffer, encodedReply, cBusOptions.GetSrchk())
	if _chksumErr != nil {
		return nil, errors.Wrap(_chksumErr, "Error parsing 'chksum' field of ReplyEncodedReply")
	}
	var chksum Checksum
	if _chksum != nil {
		chksum = _chksum.(Checksum)
	}

	if closeErr := readBuffer.CloseContext("ReplyEncodedReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyEncodedReply")
	}

	// Create a partially initialized instance
	_child := &_ReplyEncodedReply{
		_Reply: &_Reply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		EncodedReply: encodedReply,
		Chksum:       chksum,
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ReplyEncodedReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyEncodedReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyEncodedReply")
		}

		// Manual Field (encodedReply)
		_encodedReplyErr := WriteEncodedReply(writeBuffer, m.GetEncodedReply())
		if _encodedReplyErr != nil {
			return errors.Wrap(_encodedReplyErr, "Error serializing 'encodedReply' field")
		}

		// Manual Field (chksum)
		_chksumErr := CalculateChecksum(writeBuffer, m.GetEncodedReply(), m.CBusOptions.GetSrchk())
		if _chksumErr != nil {
			return errors.Wrap(_chksumErr, "Error serializing 'chksum' field")
		}

		if popErr := writeBuffer.PopContext("ReplyEncodedReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyEncodedReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ReplyEncodedReply) isReplyEncodedReply() bool {
	return true
}

func (m *_ReplyEncodedReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
