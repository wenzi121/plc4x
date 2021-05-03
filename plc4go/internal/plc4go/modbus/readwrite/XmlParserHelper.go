//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package readwrite

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/modbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

type ModbusXmlParserHelper struct {
}

func (m ModbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	_ = strconv.Atoi
	switch typeName {
	case "ModbusPDUWriteFileRecordRequestItem":
		return model.ModbusPDUWriteFileRecordRequestItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusPDUReadFileRecordResponseItem":
		return model.ModbusPDUReadFileRecordResponseItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusConstants":
		return model.ModbusConstantsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusTcpADU":
		//var response bool
		response := parserArguments[0] == "true"
		return model.ModbusTcpADUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	case "ModbusPDUWriteFileRecordResponseItem":
		return model.ModbusPDUWriteFileRecordResponseItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusPDU":
		return nil, errors.New("ModbusPDU unmappable")
	case "ModbusPDUReadFileRecordRequestItem":
		return model.ModbusPDUReadFileRecordRequestItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusSerialADU":
		//var response bool
		response := parserArguments[0] == "true"
		return model.ModbusSerialADUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}

// Deprecated: will be removed in favor of Parse soon
func (m ModbusXmlParserHelper) ParseOld(typeName string, xmlString string) (interface{}, error) {
	switch typeName {
	case "ModbusPDUWriteFileRecordRequestItem":
		var obj *model.ModbusPDUWriteFileRecordRequestItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusPDUReadFileRecordResponseItem":
		var obj *model.ModbusPDUReadFileRecordResponseItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusConstants":
		var obj *model.ModbusConstants
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusTcpADU":
		var obj *model.ModbusTcpADU
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusPDUWriteFileRecordResponseItem":
		var obj *model.ModbusPDUWriteFileRecordResponseItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusPDU":
		var obj *model.ModbusPDU
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusPDUReadFileRecordRequestItem":
		var obj *model.ModbusPDUReadFileRecordRequestItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "ModbusSerialADU":
		var obj *model.ModbusSerialADU
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
