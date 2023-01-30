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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetContextTagOctetString extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.OCTET_STRING;
  }

  // Properties.
  protected final BACnetTagPayloadOctetString payload;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetContextTagOctetString(
      BACnetTagHeader header, BACnetTagPayloadOctetString payload, Short tagNumberArgument) {
    super(header, tagNumberArgument);
    this.payload = payload;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagPayloadOctetString getPayload() {
    return payload;
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetContextTagOctetString");

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetContextTagOctetString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagOctetString _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetContextTagBuilder staticParseBACnetContextTagBuilder(
      ReadBuffer readBuffer,
      BACnetTagHeader header,
      Short tagNumberArgument,
      BACnetDataType dataType)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagOctetString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagPayloadOctetString payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetTagPayloadOctetString.staticParse(
                        readBuffer, (long) (header.getActualLength())),
                readBuffer));

    readBuffer.closeContext("BACnetContextTagOctetString");
    // Create the instance
    return new BACnetContextTagOctetStringBuilderImpl(payload, tagNumberArgument);
  }

  public static class BACnetContextTagOctetStringBuilderImpl
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final BACnetTagPayloadOctetString payload;
    private final Short tagNumberArgument;

    public BACnetContextTagOctetStringBuilderImpl(
        BACnetTagPayloadOctetString payload, Short tagNumberArgument) {
      this.payload = payload;
      this.tagNumberArgument = tagNumberArgument;
    }

    public BACnetContextTagOctetString build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagOctetString bACnetContextTagOctetString =
          new BACnetContextTagOctetString(header, payload, tagNumberArgument);
      return bACnetContextTagOctetString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagOctetString)) {
      return false;
    }
    BACnetContextTagOctetString that = (BACnetContextTagOctetString) o;
    return (getPayload() == that.getPayload()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPayload());
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
