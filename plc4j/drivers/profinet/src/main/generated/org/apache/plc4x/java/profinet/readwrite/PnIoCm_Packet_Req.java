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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_Packet_Req extends PnIoCm_Packet implements Message {

  // Accessors for discriminator values.
  public DceRpc_PacketType getPacketType() {
    return DceRpc_PacketType.REQUEST;
  }

  // Properties.
  protected final long argsMaximum;
  protected final long argsLength;
  protected final long arrayMaximumCount;
  protected final long arrayOffset;
  protected final long arrayActualCount;
  protected final List<PnIoCm_Block> blocks;

  public PnIoCm_Packet_Req(
      long argsMaximum,
      long argsLength,
      long arrayMaximumCount,
      long arrayOffset,
      long arrayActualCount,
      List<PnIoCm_Block> blocks) {
    super();
    this.argsMaximum = argsMaximum;
    this.argsLength = argsLength;
    this.arrayMaximumCount = arrayMaximumCount;
    this.arrayOffset = arrayOffset;
    this.arrayActualCount = arrayActualCount;
    this.blocks = blocks;
  }

  public long getArgsMaximum() {
    return argsMaximum;
  }

  public long getArgsLength() {
    return argsLength;
  }

  public long getArrayMaximumCount() {
    return arrayMaximumCount;
  }

  public long getArrayOffset() {
    return arrayOffset;
  }

  public long getArrayActualCount() {
    return arrayActualCount;
  }

  public List<PnIoCm_Block> getBlocks() {
    return blocks;
  }

  @Override
  protected void serializePnIoCm_PacketChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Packet_Req");

    // Simple Field (argsMaximum)
    writeSimpleField("argsMaximum", argsMaximum, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (argsLength)
    writeSimpleField("argsLength", argsLength, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (arrayMaximumCount)
    writeSimpleField("arrayMaximumCount", arrayMaximumCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (arrayOffset)
    writeSimpleField("arrayOffset", arrayOffset, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (arrayActualCount)
    writeSimpleField("arrayActualCount", arrayActualCount, writeUnsignedLong(writeBuffer, 32));

    // Array Field (blocks)
    writeComplexTypeArrayField("blocks", blocks, writeBuffer);

    writeBuffer.popContext("PnIoCm_Packet_Req");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Packet_Req _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (argsMaximum)
    lengthInBits += 32;

    // Simple field (argsLength)
    lengthInBits += 32;

    // Simple field (arrayMaximumCount)
    lengthInBits += 32;

    // Simple field (arrayOffset)
    lengthInBits += 32;

    // Simple field (arrayActualCount)
    lengthInBits += 32;

    // Array field
    if (blocks != null) {
      for (Message element : blocks) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static PnIoCm_PacketBuilder staticParsePnIoCm_PacketBuilder(
      ReadBuffer readBuffer, DceRpc_PacketType packetType) throws ParseException {
    readBuffer.pullContext("PnIoCm_Packet_Req");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long argsMaximum = readSimpleField("argsMaximum", readUnsignedLong(readBuffer, 32));

    long argsLength = readSimpleField("argsLength", readUnsignedLong(readBuffer, 32));

    long arrayMaximumCount = readSimpleField("arrayMaximumCount", readUnsignedLong(readBuffer, 32));

    long arrayOffset = readSimpleField("arrayOffset", readUnsignedLong(readBuffer, 32));

    long arrayActualCount = readSimpleField("arrayActualCount", readUnsignedLong(readBuffer, 32));

    List<PnIoCm_Block> blocks =
        readLengthArrayField(
            "blocks",
            new DataReaderComplexDefault<>(() -> PnIoCm_Block.staticParse(readBuffer), readBuffer),
            argsLength);

    readBuffer.closeContext("PnIoCm_Packet_Req");
    // Create the instance
    return new PnIoCm_Packet_ReqBuilderImpl(
        argsMaximum, argsLength, arrayMaximumCount, arrayOffset, arrayActualCount, blocks);
  }

  public static class PnIoCm_Packet_ReqBuilderImpl implements PnIoCm_Packet.PnIoCm_PacketBuilder {
    private final long argsMaximum;
    private final long argsLength;
    private final long arrayMaximumCount;
    private final long arrayOffset;
    private final long arrayActualCount;
    private final List<PnIoCm_Block> blocks;

    public PnIoCm_Packet_ReqBuilderImpl(
        long argsMaximum,
        long argsLength,
        long arrayMaximumCount,
        long arrayOffset,
        long arrayActualCount,
        List<PnIoCm_Block> blocks) {
      this.argsMaximum = argsMaximum;
      this.argsLength = argsLength;
      this.arrayMaximumCount = arrayMaximumCount;
      this.arrayOffset = arrayOffset;
      this.arrayActualCount = arrayActualCount;
      this.blocks = blocks;
    }

    public PnIoCm_Packet_Req build() {
      PnIoCm_Packet_Req pnIoCm_Packet_Req =
          new PnIoCm_Packet_Req(
              argsMaximum, argsLength, arrayMaximumCount, arrayOffset, arrayActualCount, blocks);
      return pnIoCm_Packet_Req;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Packet_Req)) {
      return false;
    }
    PnIoCm_Packet_Req that = (PnIoCm_Packet_Req) o;
    return (getArgsMaximum() == that.getArgsMaximum())
        && (getArgsLength() == that.getArgsLength())
        && (getArrayMaximumCount() == that.getArrayMaximumCount())
        && (getArrayOffset() == that.getArrayOffset())
        && (getArrayActualCount() == that.getArrayActualCount())
        && (getBlocks() == that.getBlocks())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getArgsMaximum(),
        getArgsLength(),
        getArrayMaximumCount(),
        getArrayOffset(),
        getArrayActualCount(),
        getBlocks());
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
