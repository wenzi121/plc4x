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
package org.apache.plc4x.java.ads.readwrite;

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

public class AdsReadRequest extends AmsPacket implements Message {

  // Accessors for discriminator values.
  public CommandId getCommandId() {
    return CommandId.ADS_READ;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final long indexGroup;
  protected final long indexOffset;
  protected final long length;

  public AdsReadRequest(
      AmsNetId targetAmsNetId,
      int targetAmsPort,
      AmsNetId sourceAmsNetId,
      int sourceAmsPort,
      long errorCode,
      long invokeId,
      long indexGroup,
      long indexOffset,
      long length) {
    super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId);
    this.indexGroup = indexGroup;
    this.indexOffset = indexOffset;
    this.length = length;
  }

  public long getIndexGroup() {
    return indexGroup;
  }

  public long getIndexOffset() {
    return indexOffset;
  }

  public long getLength() {
    return length;
  }

  @Override
  protected void serializeAmsPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AdsReadRequest");

    // Simple Field (indexGroup)
    writeSimpleField("indexGroup", indexGroup, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (indexOffset)
    writeSimpleField("indexOffset", indexOffset, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (length)
    writeSimpleField("length", length, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("AdsReadRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsReadRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (indexGroup)
    lengthInBits += 32;

    // Simple field (indexOffset)
    lengthInBits += 32;

    // Simple field (length)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static AmsPacketBuilder staticParseAmsPacketBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsReadRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long indexGroup = readSimpleField("indexGroup", readUnsignedLong(readBuffer, 32));

    long indexOffset = readSimpleField("indexOffset", readUnsignedLong(readBuffer, 32));

    long length = readSimpleField("length", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("AdsReadRequest");
    // Create the instance
    return new AdsReadRequestBuilderImpl(indexGroup, indexOffset, length);
  }

  public static class AdsReadRequestBuilderImpl implements AmsPacket.AmsPacketBuilder {
    private final long indexGroup;
    private final long indexOffset;
    private final long length;

    public AdsReadRequestBuilderImpl(long indexGroup, long indexOffset, long length) {
      this.indexGroup = indexGroup;
      this.indexOffset = indexOffset;
      this.length = length;
    }

    public AdsReadRequest build(
        AmsNetId targetAmsNetId,
        int targetAmsPort,
        AmsNetId sourceAmsNetId,
        int sourceAmsPort,
        long errorCode,
        long invokeId) {
      AdsReadRequest adsReadRequest =
          new AdsReadRequest(
              targetAmsNetId,
              targetAmsPort,
              sourceAmsNetId,
              sourceAmsPort,
              errorCode,
              invokeId,
              indexGroup,
              indexOffset,
              length);
      return adsReadRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsReadRequest)) {
      return false;
    }
    AdsReadRequest that = (AdsReadRequest) o;
    return (getIndexGroup() == that.getIndexGroup())
        && (getIndexOffset() == that.getIndexOffset())
        && (getLength() == that.getLength())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIndexGroup(), getIndexOffset(), getLength());
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
