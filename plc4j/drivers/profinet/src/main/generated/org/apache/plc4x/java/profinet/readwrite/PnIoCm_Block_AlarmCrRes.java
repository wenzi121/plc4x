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

public class PnIoCm_Block_AlarmCrRes extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.ALARM_CR_BLOCK_RES;
  }

  // Properties.
  protected final PnIoCm_AlarmCrType alarmType;
  protected final int localAlarmReference;
  protected final int maxAlarmDataLength;

  public PnIoCm_Block_AlarmCrRes(
      short blockVersionHigh,
      short blockVersionLow,
      PnIoCm_AlarmCrType alarmType,
      int localAlarmReference,
      int maxAlarmDataLength) {
    super(blockVersionHigh, blockVersionLow);
    this.alarmType = alarmType;
    this.localAlarmReference = localAlarmReference;
    this.maxAlarmDataLength = maxAlarmDataLength;
  }

  public PnIoCm_AlarmCrType getAlarmType() {
    return alarmType;
  }

  public int getLocalAlarmReference() {
    return localAlarmReference;
  }

  public int getMaxAlarmDataLength() {
    return maxAlarmDataLength;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Block_AlarmCrRes");

    // Simple Field (alarmType)
    writeSimpleEnumField(
        "alarmType",
        "PnIoCm_AlarmCrType",
        alarmType,
        new DataWriterEnumDefault<>(
            PnIoCm_AlarmCrType::getValue,
            PnIoCm_AlarmCrType::name,
            writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (localAlarmReference)
    writeSimpleField(
        "localAlarmReference",
        localAlarmReference,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (maxAlarmDataLength)
    writeSimpleField(
        "maxAlarmDataLength",
        maxAlarmDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Block_AlarmCrRes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_AlarmCrRes _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (alarmType)
    lengthInBits += 16;

    // Simple field (localAlarmReference)
    lengthInBits += 16;

    // Simple field (maxAlarmDataLength)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_AlarmCrRes");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PnIoCm_AlarmCrType alarmType =
        readEnumField(
            "alarmType",
            "PnIoCm_AlarmCrType",
            new DataReaderEnumDefault<>(
                PnIoCm_AlarmCrType::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int localAlarmReference =
        readSimpleField(
            "localAlarmReference",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int maxAlarmDataLength =
        readSimpleField(
            "maxAlarmDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_AlarmCrRes");
    // Create the instance
    return new PnIoCm_Block_AlarmCrResBuilderImpl(
        alarmType, localAlarmReference, maxAlarmDataLength);
  }

  public static class PnIoCm_Block_AlarmCrResBuilderImpl
      implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final PnIoCm_AlarmCrType alarmType;
    private final int localAlarmReference;
    private final int maxAlarmDataLength;

    public PnIoCm_Block_AlarmCrResBuilderImpl(
        PnIoCm_AlarmCrType alarmType, int localAlarmReference, int maxAlarmDataLength) {
      this.alarmType = alarmType;
      this.localAlarmReference = localAlarmReference;
      this.maxAlarmDataLength = maxAlarmDataLength;
    }

    public PnIoCm_Block_AlarmCrRes build(short blockVersionHigh, short blockVersionLow) {
      PnIoCm_Block_AlarmCrRes pnIoCm_Block_AlarmCrRes =
          new PnIoCm_Block_AlarmCrRes(
              blockVersionHigh,
              blockVersionLow,
              alarmType,
              localAlarmReference,
              maxAlarmDataLength);
      return pnIoCm_Block_AlarmCrRes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_AlarmCrRes)) {
      return false;
    }
    PnIoCm_Block_AlarmCrRes that = (PnIoCm_Block_AlarmCrRes) o;
    return (getAlarmType() == that.getAlarmType())
        && (getLocalAlarmReference() == that.getLocalAlarmReference())
        && (getMaxAlarmDataLength() == that.getMaxAlarmDataLength())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getAlarmType(), getLocalAlarmReference(), getMaxAlarmDataLength());
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
