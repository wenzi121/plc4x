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
package org.apache.plc4x.java.modbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ModbusPDUReadCoilsResponse extends ModbusPDU implements Message {

  // Accessors for discriminator values.
  public Boolean getErrorFlag() {
    return (boolean) false;
  }

  public Short getFunctionFlag() {
    return (short) 0x01;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final byte[] value;

  public ModbusPDUReadCoilsResponse(byte[] value) {
    super();
    this.value = value;
  }

  public byte[] getValue() {
    return value;
  }

  @Override
  protected void serializeModbusPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModbusPDUReadCoilsResponse");

    // Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short byteCount = (short) (COUNT(getValue()));
    writeImplicitField("byteCount", byteCount, writeUnsignedShort(writeBuffer, 8));

    // Array Field (value)
    writeByteArrayField("value", value, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("ModbusPDUReadCoilsResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusPDUReadCoilsResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (byteCount)
    lengthInBits += 8;

    // Array field
    if (value != null) {
      lengthInBits += 8 * value.length;
    }

    return lengthInBits;
  }

  public static ModbusPDUBuilder staticParseModbusPDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusPDUReadCoilsResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short byteCount = readImplicitField("byteCount", readUnsignedShort(readBuffer, 8));

    byte[] value = readBuffer.readByteArray("value", Math.toIntExact(byteCount));

    readBuffer.closeContext("ModbusPDUReadCoilsResponse");
    // Create the instance
    return new ModbusPDUReadCoilsResponseBuilderImpl(value);
  }

  public static class ModbusPDUReadCoilsResponseBuilderImpl implements ModbusPDU.ModbusPDUBuilder {
    private final byte[] value;

    public ModbusPDUReadCoilsResponseBuilderImpl(byte[] value) {
      this.value = value;
    }

    public ModbusPDUReadCoilsResponse build() {
      ModbusPDUReadCoilsResponse modbusPDUReadCoilsResponse = new ModbusPDUReadCoilsResponse(value);
      return modbusPDUReadCoilsResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUReadCoilsResponse)) {
      return false;
    }
    ModbusPDUReadCoilsResponse that = (ModbusPDUReadCoilsResponse) o;
    return (getValue() == that.getValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
