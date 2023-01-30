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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ApduDataMemoryResponse extends ApduData implements Message {

  // Accessors for discriminator values.
  public Byte getApciType() {
    return (byte) 0x9;
  }

  // Properties.
  protected final int address;
  protected final byte[] data;

  public ApduDataMemoryResponse(int address, byte[] data) {
    super();
    this.address = address;
    this.data = data;
  }

  public int getAddress() {
    return address;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeApduDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduDataMemoryResponse");

    // Implicit Field (numBytes) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short numBytes = (short) (COUNT(getData()));
    writeImplicitField("numBytes", numBytes, writeUnsignedShort(writeBuffer, 6));

    // Simple Field (address)
    writeSimpleField("address", address, writeUnsignedInt(writeBuffer, 16));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("ApduDataMemoryResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataMemoryResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (numBytes)
    lengthInBits += 6;

    // Simple field (address)
    lengthInBits += 16;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static ApduDataBuilder staticParseApduDataBuilder(ReadBuffer readBuffer, Short dataLength)
      throws ParseException {
    readBuffer.pullContext("ApduDataMemoryResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short numBytes = readImplicitField("numBytes", readUnsignedShort(readBuffer, 6));

    int address = readSimpleField("address", readUnsignedInt(readBuffer, 16));

    byte[] data = readBuffer.readByteArray("data", Math.toIntExact(numBytes));

    readBuffer.closeContext("ApduDataMemoryResponse");
    // Create the instance
    return new ApduDataMemoryResponseBuilderImpl(address, data);
  }

  public static class ApduDataMemoryResponseBuilderImpl implements ApduData.ApduDataBuilder {
    private final int address;
    private final byte[] data;

    public ApduDataMemoryResponseBuilderImpl(int address, byte[] data) {
      this.address = address;
      this.data = data;
    }

    public ApduDataMemoryResponse build() {
      ApduDataMemoryResponse apduDataMemoryResponse = new ApduDataMemoryResponse(address, data);
      return apduDataMemoryResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataMemoryResponse)) {
      return false;
    }
    ApduDataMemoryResponse that = (ApduDataMemoryResponse) o;
    return (getAddress() == that.getAddress())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAddress(), getData());
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
