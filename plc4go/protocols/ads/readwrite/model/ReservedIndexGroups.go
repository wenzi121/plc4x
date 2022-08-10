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
)

// Code generated by code-generation. DO NOT EDIT.

// ReservedIndexGroups is an enum
type ReservedIndexGroups uint32

type IReservedIndexGroups interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const(
	ReservedIndexGroups_ADSIGRP_SYMTAB ReservedIndexGroups = 0x0000F000
	ReservedIndexGroups_ADSIGRP_SYMNAME ReservedIndexGroups = 0x0000F001
	ReservedIndexGroups_ADSIGRP_SYMVAL ReservedIndexGroups = 0x0000F002
	ReservedIndexGroups_ADSIGRP_SYM_HNDBYNAME ReservedIndexGroups = 0x0000F003
	ReservedIndexGroups_ADSIGRP_SYM_VALBYNAME ReservedIndexGroups = 0x0000F004
	ReservedIndexGroups_ADSIGRP_SYM_VALBYHND ReservedIndexGroups = 0x0000F005
	ReservedIndexGroups_ADSIGRP_SYM_RELEASEHND ReservedIndexGroups = 0x0000F006
	ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAME ReservedIndexGroups = 0x0000F007
	ReservedIndexGroups_ADSIGRP_SYM_VERSION ReservedIndexGroups = 0x0000F008
	ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAMEEX ReservedIndexGroups = 0x0000F009
	ReservedIndexGroups_ADSIGRP_SYM_DOWNLOAD ReservedIndexGroups = 0x0000F00A
	ReservedIndexGroups_ADSIGRP_SYM_UPLOAD ReservedIndexGroups = 0x0000F00B
	ReservedIndexGroups_ADSIGRP_SYM_UPLOADINFO ReservedIndexGroups = 0x0000F00C
	ReservedIndexGroups_ADSIGRP_SYMNOTE ReservedIndexGroups = 0x0000F010
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIB ReservedIndexGroups = 0x0000F020
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIX ReservedIndexGroups = 0x0000F021
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RISIZE ReservedIndexGroups = 0x0000F025
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOB ReservedIndexGroups = 0x0000F030
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOX ReservedIndexGroups = 0x0000F031
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOSIZE ReservedIndexGroups = 0x0000F035
	ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARI ReservedIndexGroups = 0x0000F040
	ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARO ReservedIndexGroups = 0x0000F050
	ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIOB ReservedIndexGroups = 0x0000F060
	ReservedIndexGroups_ADSIGRP_MULTIPLE_READ ReservedIndexGroups = 0x0000F080
	ReservedIndexGroups_ADSIGRP_MULTIPLE_WRITE ReservedIndexGroups = 0x0000F081
	ReservedIndexGroups_ADSIGRP_MULTIPLE_READ_WRITE ReservedIndexGroups = 0x0000F082
	ReservedIndexGroups_ADSIGRP_MULTIPLE_RELEASE_HANDLE ReservedIndexGroups = 0x0000F083
	ReservedIndexGroups_ADSIGRP_DEVICE_DATA ReservedIndexGroups = 0x0000F100
	ReservedIndexGroups_ADSIOFFS_DEVDATA_ADSSTATE ReservedIndexGroups = 0x00000000
	ReservedIndexGroups_ADSIOFFS_DEVDATA_DEVSTATE ReservedIndexGroups = 0x00000002
)

var ReservedIndexGroupsValues []ReservedIndexGroups

func init() {
	_ = errors.New
	ReservedIndexGroupsValues = []ReservedIndexGroups {
		ReservedIndexGroups_ADSIGRP_SYMTAB,
		ReservedIndexGroups_ADSIGRP_SYMNAME,
		ReservedIndexGroups_ADSIGRP_SYMVAL,
		ReservedIndexGroups_ADSIGRP_SYM_HNDBYNAME,
		ReservedIndexGroups_ADSIGRP_SYM_VALBYNAME,
		ReservedIndexGroups_ADSIGRP_SYM_VALBYHND,
		ReservedIndexGroups_ADSIGRP_SYM_RELEASEHND,
		ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAME,
		ReservedIndexGroups_ADSIGRP_SYM_VERSION,
		ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAMEEX,
		ReservedIndexGroups_ADSIGRP_SYM_DOWNLOAD,
		ReservedIndexGroups_ADSIGRP_SYM_UPLOAD,
		ReservedIndexGroups_ADSIGRP_SYM_UPLOADINFO,
		ReservedIndexGroups_ADSIGRP_SYMNOTE,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIB,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIX,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RISIZE,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOB,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOX,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOSIZE,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARI,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARO,
		ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIOB,
		ReservedIndexGroups_ADSIGRP_MULTIPLE_READ,
		ReservedIndexGroups_ADSIGRP_MULTIPLE_WRITE,
		ReservedIndexGroups_ADSIGRP_MULTIPLE_READ_WRITE,
		ReservedIndexGroups_ADSIGRP_MULTIPLE_RELEASE_HANDLE,
		ReservedIndexGroups_ADSIGRP_DEVICE_DATA,
		ReservedIndexGroups_ADSIOFFS_DEVDATA_ADSSTATE,
		ReservedIndexGroups_ADSIOFFS_DEVDATA_DEVSTATE,
	}
}

func ReservedIndexGroupsByValue(value uint32) (enum ReservedIndexGroups, ok bool) {
	switch value {
		case 0x00000000:
			return ReservedIndexGroups_ADSIOFFS_DEVDATA_ADSSTATE, true
		case 0x00000002:
			return ReservedIndexGroups_ADSIOFFS_DEVDATA_DEVSTATE, true
		case 0x0000F000:
			return ReservedIndexGroups_ADSIGRP_SYMTAB, true
		case 0x0000F001:
			return ReservedIndexGroups_ADSIGRP_SYMNAME, true
		case 0x0000F002:
			return ReservedIndexGroups_ADSIGRP_SYMVAL, true
		case 0x0000F003:
			return ReservedIndexGroups_ADSIGRP_SYM_HNDBYNAME, true
		case 0x0000F004:
			return ReservedIndexGroups_ADSIGRP_SYM_VALBYNAME, true
		case 0x0000F005:
			return ReservedIndexGroups_ADSIGRP_SYM_VALBYHND, true
		case 0x0000F006:
			return ReservedIndexGroups_ADSIGRP_SYM_RELEASEHND, true
		case 0x0000F007:
			return ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAME, true
		case 0x0000F008:
			return ReservedIndexGroups_ADSIGRP_SYM_VERSION, true
		case 0x0000F009:
			return ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAMEEX, true
		case 0x0000F00A:
			return ReservedIndexGroups_ADSIGRP_SYM_DOWNLOAD, true
		case 0x0000F00B:
			return ReservedIndexGroups_ADSIGRP_SYM_UPLOAD, true
		case 0x0000F00C:
			return ReservedIndexGroups_ADSIGRP_SYM_UPLOADINFO, true
		case 0x0000F010:
			return ReservedIndexGroups_ADSIGRP_SYMNOTE, true
		case 0x0000F020:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIB, true
		case 0x0000F021:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIX, true
		case 0x0000F025:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RISIZE, true
		case 0x0000F030:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOB, true
		case 0x0000F031:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOX, true
		case 0x0000F035:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOSIZE, true
		case 0x0000F040:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARI, true
		case 0x0000F050:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARO, true
		case 0x0000F060:
			return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIOB, true
		case 0x0000F080:
			return ReservedIndexGroups_ADSIGRP_MULTIPLE_READ, true
		case 0x0000F081:
			return ReservedIndexGroups_ADSIGRP_MULTIPLE_WRITE, true
		case 0x0000F082:
			return ReservedIndexGroups_ADSIGRP_MULTIPLE_READ_WRITE, true
		case 0x0000F083:
			return ReservedIndexGroups_ADSIGRP_MULTIPLE_RELEASE_HANDLE, true
		case 0x0000F100:
			return ReservedIndexGroups_ADSIGRP_DEVICE_DATA, true
	}
	return 0, false
}

func ReservedIndexGroupsByName(value string) (enum ReservedIndexGroups, ok bool) {
	switch value {
	case "ADSIOFFS_DEVDATA_ADSSTATE":
		return ReservedIndexGroups_ADSIOFFS_DEVDATA_ADSSTATE, true
	case "ADSIOFFS_DEVDATA_DEVSTATE":
		return ReservedIndexGroups_ADSIOFFS_DEVDATA_DEVSTATE, true
	case "ADSIGRP_SYMTAB":
		return ReservedIndexGroups_ADSIGRP_SYMTAB, true
	case "ADSIGRP_SYMNAME":
		return ReservedIndexGroups_ADSIGRP_SYMNAME, true
	case "ADSIGRP_SYMVAL":
		return ReservedIndexGroups_ADSIGRP_SYMVAL, true
	case "ADSIGRP_SYM_HNDBYNAME":
		return ReservedIndexGroups_ADSIGRP_SYM_HNDBYNAME, true
	case "ADSIGRP_SYM_VALBYNAME":
		return ReservedIndexGroups_ADSIGRP_SYM_VALBYNAME, true
	case "ADSIGRP_SYM_VALBYHND":
		return ReservedIndexGroups_ADSIGRP_SYM_VALBYHND, true
	case "ADSIGRP_SYM_RELEASEHND":
		return ReservedIndexGroups_ADSIGRP_SYM_RELEASEHND, true
	case "ADSIGRP_SYM_INFOBYNAME":
		return ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAME, true
	case "ADSIGRP_SYM_VERSION":
		return ReservedIndexGroups_ADSIGRP_SYM_VERSION, true
	case "ADSIGRP_SYM_INFOBYNAMEEX":
		return ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAMEEX, true
	case "ADSIGRP_SYM_DOWNLOAD":
		return ReservedIndexGroups_ADSIGRP_SYM_DOWNLOAD, true
	case "ADSIGRP_SYM_UPLOAD":
		return ReservedIndexGroups_ADSIGRP_SYM_UPLOAD, true
	case "ADSIGRP_SYM_UPLOADINFO":
		return ReservedIndexGroups_ADSIGRP_SYM_UPLOADINFO, true
	case "ADSIGRP_SYMNOTE":
		return ReservedIndexGroups_ADSIGRP_SYMNOTE, true
	case "ADSIGRP_IOIMAGE_RWIB":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIB, true
	case "ADSIGRP_IOIMAGE_RWIX":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIX, true
	case "ADSIGRP_IOIMAGE_RISIZE":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RISIZE, true
	case "ADSIGRP_IOIMAGE_RWOB":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOB, true
	case "ADSIGRP_IOIMAGE_RWOX":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOX, true
	case "ADSIGRP_IOIMAGE_RWOSIZE":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOSIZE, true
	case "ADSIGRP_IOIMAGE_CLEARI":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARI, true
	case "ADSIGRP_IOIMAGE_CLEARO":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARO, true
	case "ADSIGRP_IOIMAGE_RWIOB":
		return ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIOB, true
	case "ADSIGRP_MULTIPLE_READ":
		return ReservedIndexGroups_ADSIGRP_MULTIPLE_READ, true
	case "ADSIGRP_MULTIPLE_WRITE":
		return ReservedIndexGroups_ADSIGRP_MULTIPLE_WRITE, true
	case "ADSIGRP_MULTIPLE_READ_WRITE":
		return ReservedIndexGroups_ADSIGRP_MULTIPLE_READ_WRITE, true
	case "ADSIGRP_MULTIPLE_RELEASE_HANDLE":
		return ReservedIndexGroups_ADSIGRP_MULTIPLE_RELEASE_HANDLE, true
	case "ADSIGRP_DEVICE_DATA":
		return ReservedIndexGroups_ADSIGRP_DEVICE_DATA, true
	}
	return 0, false
}

func ReservedIndexGroupsKnows(value uint32)  bool {
	for _, typeValue := range ReservedIndexGroupsValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastReservedIndexGroups(structType interface{}) ReservedIndexGroups {
	castFunc := func(typ interface{}) ReservedIndexGroups {
		if sReservedIndexGroups, ok := typ.(ReservedIndexGroups); ok {
			return sReservedIndexGroups
		}
		return 0
	}
	return castFunc(structType)
}

func (m ReservedIndexGroups) GetLengthInBits() uint16 {
	return 32
}

func (m ReservedIndexGroups) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReservedIndexGroupsParse(readBuffer utils.ReadBuffer) (ReservedIndexGroups, error) {
	val, err := readBuffer.ReadUint32("ReservedIndexGroups", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ReservedIndexGroups")
	}
	if enum, ok := ReservedIndexGroupsByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return ReservedIndexGroups(val), nil
	} else {
		return enum, nil
	}
}

func (e ReservedIndexGroups) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("ReservedIndexGroups", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ReservedIndexGroups) PLC4XEnumName() string {
	switch e {
	case ReservedIndexGroups_ADSIOFFS_DEVDATA_ADSSTATE:
		return "ADSIOFFS_DEVDATA_ADSSTATE"
	case ReservedIndexGroups_ADSIOFFS_DEVDATA_DEVSTATE:
		return "ADSIOFFS_DEVDATA_DEVSTATE"
	case ReservedIndexGroups_ADSIGRP_SYMTAB:
		return "ADSIGRP_SYMTAB"
	case ReservedIndexGroups_ADSIGRP_SYMNAME:
		return "ADSIGRP_SYMNAME"
	case ReservedIndexGroups_ADSIGRP_SYMVAL:
		return "ADSIGRP_SYMVAL"
	case ReservedIndexGroups_ADSIGRP_SYM_HNDBYNAME:
		return "ADSIGRP_SYM_HNDBYNAME"
	case ReservedIndexGroups_ADSIGRP_SYM_VALBYNAME:
		return "ADSIGRP_SYM_VALBYNAME"
	case ReservedIndexGroups_ADSIGRP_SYM_VALBYHND:
		return "ADSIGRP_SYM_VALBYHND"
	case ReservedIndexGroups_ADSIGRP_SYM_RELEASEHND:
		return "ADSIGRP_SYM_RELEASEHND"
	case ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAME:
		return "ADSIGRP_SYM_INFOBYNAME"
	case ReservedIndexGroups_ADSIGRP_SYM_VERSION:
		return "ADSIGRP_SYM_VERSION"
	case ReservedIndexGroups_ADSIGRP_SYM_INFOBYNAMEEX:
		return "ADSIGRP_SYM_INFOBYNAMEEX"
	case ReservedIndexGroups_ADSIGRP_SYM_DOWNLOAD:
		return "ADSIGRP_SYM_DOWNLOAD"
	case ReservedIndexGroups_ADSIGRP_SYM_UPLOAD:
		return "ADSIGRP_SYM_UPLOAD"
	case ReservedIndexGroups_ADSIGRP_SYM_UPLOADINFO:
		return "ADSIGRP_SYM_UPLOADINFO"
	case ReservedIndexGroups_ADSIGRP_SYMNOTE:
		return "ADSIGRP_SYMNOTE"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIB:
		return "ADSIGRP_IOIMAGE_RWIB"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIX:
		return "ADSIGRP_IOIMAGE_RWIX"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RISIZE:
		return "ADSIGRP_IOIMAGE_RISIZE"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOB:
		return "ADSIGRP_IOIMAGE_RWOB"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOX:
		return "ADSIGRP_IOIMAGE_RWOX"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWOSIZE:
		return "ADSIGRP_IOIMAGE_RWOSIZE"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARI:
		return "ADSIGRP_IOIMAGE_CLEARI"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_CLEARO:
		return "ADSIGRP_IOIMAGE_CLEARO"
	case ReservedIndexGroups_ADSIGRP_IOIMAGE_RWIOB:
		return "ADSIGRP_IOIMAGE_RWIOB"
	case ReservedIndexGroups_ADSIGRP_MULTIPLE_READ:
		return "ADSIGRP_MULTIPLE_READ"
	case ReservedIndexGroups_ADSIGRP_MULTIPLE_WRITE:
		return "ADSIGRP_MULTIPLE_WRITE"
	case ReservedIndexGroups_ADSIGRP_MULTIPLE_READ_WRITE:
		return "ADSIGRP_MULTIPLE_READ_WRITE"
	case ReservedIndexGroups_ADSIGRP_MULTIPLE_RELEASE_HANDLE:
		return "ADSIGRP_MULTIPLE_RELEASE_HANDLE"
	case ReservedIndexGroups_ADSIGRP_DEVICE_DATA:
		return "ADSIGRP_DEVICE_DATA"
	}
	return ""
}

func (e ReservedIndexGroups) String() string {
	return e.PLC4XEnumName()
}

