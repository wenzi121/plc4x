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

// OpcuaNodeIdServicesVariableRedundant is an enum
type OpcuaNodeIdServicesVariableRedundant int32

type IOpcuaNodeIdServicesVariableRedundant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableRedundant_RedundantServerMode_EnumStrings OpcuaNodeIdServicesVariableRedundant = 32418
)

var OpcuaNodeIdServicesVariableRedundantValues []OpcuaNodeIdServicesVariableRedundant

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableRedundantValues = []OpcuaNodeIdServicesVariableRedundant{
		OpcuaNodeIdServicesVariableRedundant_RedundantServerMode_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableRedundantByValue(value int32) (enum OpcuaNodeIdServicesVariableRedundant, ok bool) {
	switch value {
	case 32418:
		return OpcuaNodeIdServicesVariableRedundant_RedundantServerMode_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRedundantByName(value string) (enum OpcuaNodeIdServicesVariableRedundant, ok bool) {
	switch value {
	case "RedundantServerMode_EnumStrings":
		return OpcuaNodeIdServicesVariableRedundant_RedundantServerMode_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRedundantKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableRedundantValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableRedundant(structType any) OpcuaNodeIdServicesVariableRedundant {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRedundant {
		if sOpcuaNodeIdServicesVariableRedundant, ok := typ.(OpcuaNodeIdServicesVariableRedundant); ok {
			return sOpcuaNodeIdServicesVariableRedundant
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRedundant) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRedundant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableRedundantParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRedundant, error) {
	return OpcuaNodeIdServicesVariableRedundantParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableRedundantParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRedundant, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRedundant", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRedundant")
	}
	if enum, ok := OpcuaNodeIdServicesVariableRedundantByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRedundant")
		return OpcuaNodeIdServicesVariableRedundant(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRedundant) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRedundant) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRedundant", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRedundant) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRedundant_RedundantServerMode_EnumStrings:
		return "RedundantServerMode_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRedundant) String() string {
	return e.PLC4XEnumName()
}