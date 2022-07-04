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

// ReplyNormalReply is the corresponding interface of ReplyNormalReply
type ReplyNormalReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetReply returns Reply (property field)
	GetReply() NormalReply
	// GetTermination returns Termination (property field)
	GetTermination() ResponseTermination
}

// ReplyNormalReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyNormalReply.
// This is useful for switch cases.
type ReplyNormalReplyExactly interface {
	ReplyNormalReply
	isReplyNormalReply() bool
}

// _ReplyNormalReply is the data-structure of this message
type _ReplyNormalReply struct {
	*_Reply
	Reply       NormalReply
	Termination ResponseTermination
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyNormalReply) InitializeParent(parent Reply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyNormalReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyNormalReply) GetReply() NormalReply {
	return m.Reply
}

func (m *_ReplyNormalReply) GetTermination() ResponseTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyNormalReply factory function for _ReplyNormalReply
func NewReplyNormalReply(reply NormalReply, termination ResponseTermination, peekedByte byte, messageLength uint16) *_ReplyNormalReply {
	_result := &_ReplyNormalReply{
		Reply:       reply,
		Termination: termination,
		_Reply:      NewReply(peekedByte, messageLength),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyNormalReply(structType interface{}) ReplyNormalReply {
	if casted, ok := structType.(ReplyNormalReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyNormalReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyNormalReply) GetTypeName() string {
	return "ReplyNormalReply"
}

func (m *_ReplyNormalReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ReplyNormalReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits()

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits()

	return lengthInBits
}

func (m *_ReplyNormalReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyNormalReplyParse(readBuffer utils.ReadBuffer, messageLength uint16) (ReplyNormalReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyNormalReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyNormalReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reply)
	if pullErr := readBuffer.PullContext("reply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reply")
	}
	_reply, _replyErr := NormalReplyParse(readBuffer, uint16(messageLength))
	if _replyErr != nil {
		return nil, errors.Wrap(_replyErr, "Error parsing 'reply' field of ReplyNormalReply")
	}
	reply := _reply.(NormalReply)
	if closeErr := readBuffer.CloseContext("reply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reply")
	}

	// Simple Field (termination)
	if pullErr := readBuffer.PullContext("termination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for termination")
	}
	_termination, _terminationErr := ResponseTerminationParse(readBuffer)
	if _terminationErr != nil {
		return nil, errors.Wrap(_terminationErr, "Error parsing 'termination' field of ReplyNormalReply")
	}
	termination := _termination.(ResponseTermination)
	if closeErr := readBuffer.CloseContext("termination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for termination")
	}

	if closeErr := readBuffer.CloseContext("ReplyNormalReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyNormalReply")
	}

	// Create a partially initialized instance
	_child := &_ReplyNormalReply{
		Reply:       reply,
		Termination: termination,
		_Reply: &_Reply{
			MessageLength: messageLength,
		},
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ReplyNormalReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyNormalReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyNormalReply")
		}

		// Simple Field (reply)
		if pushErr := writeBuffer.PushContext("reply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reply")
		}
		_replyErr := writeBuffer.WriteSerializable(m.GetReply())
		if popErr := writeBuffer.PopContext("reply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reply")
		}
		if _replyErr != nil {
			return errors.Wrap(_replyErr, "Error serializing 'reply' field")
		}

		// Simple Field (termination)
		if pushErr := writeBuffer.PushContext("termination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for termination")
		}
		_terminationErr := writeBuffer.WriteSerializable(m.GetTermination())
		if popErr := writeBuffer.PopContext("termination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for termination")
		}
		if _terminationErr != nil {
			return errors.Wrap(_terminationErr, "Error serializing 'termination' field")
		}

		if popErr := writeBuffer.PopContext("ReplyNormalReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyNormalReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ReplyNormalReply) isReplyNormalReply() bool {
	return true
}

func (m *_ReplyNormalReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
