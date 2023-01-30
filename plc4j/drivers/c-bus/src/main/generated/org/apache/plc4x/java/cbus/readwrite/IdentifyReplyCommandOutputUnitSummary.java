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
package org.apache.plc4x.java.cbus.readwrite;

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

public class IdentifyReplyCommandOutputUnitSummary extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.OutputUnitSummary;
  }

  // Properties.
  protected final IdentifyReplyCommandUnitSummary unitFlags;
  protected final Byte gavStoreEnabledByte1;
  protected final Byte gavStoreEnabledByte2;
  protected final short timeFromLastRecoverOfMainsInSeconds;

  // Arguments.
  protected final Short numBytes;

  public IdentifyReplyCommandOutputUnitSummary(
      IdentifyReplyCommandUnitSummary unitFlags,
      Byte gavStoreEnabledByte1,
      Byte gavStoreEnabledByte2,
      short timeFromLastRecoverOfMainsInSeconds,
      Short numBytes) {
    super(numBytes);
    this.unitFlags = unitFlags;
    this.gavStoreEnabledByte1 = gavStoreEnabledByte1;
    this.gavStoreEnabledByte2 = gavStoreEnabledByte2;
    this.timeFromLastRecoverOfMainsInSeconds = timeFromLastRecoverOfMainsInSeconds;
    this.numBytes = numBytes;
  }

  public IdentifyReplyCommandUnitSummary getUnitFlags() {
    return unitFlags;
  }

  public Byte getGavStoreEnabledByte1() {
    return gavStoreEnabledByte1;
  }

  public Byte getGavStoreEnabledByte2() {
    return gavStoreEnabledByte2;
  }

  public short getTimeFromLastRecoverOfMainsInSeconds() {
    return timeFromLastRecoverOfMainsInSeconds;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("IdentifyReplyCommandOutputUnitSummary");

    // Simple Field (unitFlags)
    writeSimpleField("unitFlags", unitFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (gavStoreEnabledByte1) (Can be skipped, if the value is null)
    writeOptionalField(
        "gavStoreEnabledByte1", gavStoreEnabledByte1, writeByte(writeBuffer, 8), (numBytes) > (1));

    // Optional Field (gavStoreEnabledByte2) (Can be skipped, if the value is null)
    writeOptionalField(
        "gavStoreEnabledByte2", gavStoreEnabledByte2, writeByte(writeBuffer, 8), (numBytes) > (2));

    // Simple Field (timeFromLastRecoverOfMainsInSeconds)
    writeSimpleField(
        "timeFromLastRecoverOfMainsInSeconds",
        timeFromLastRecoverOfMainsInSeconds,
        writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("IdentifyReplyCommandOutputUnitSummary");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandOutputUnitSummary _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (unitFlags)
    lengthInBits += unitFlags.getLengthInBits();

    // Optional Field (gavStoreEnabledByte1)
    if (gavStoreEnabledByte1 != null) {
      lengthInBits += 8;
    }

    // Optional Field (gavStoreEnabledByte2)
    if (gavStoreEnabledByte2 != null) {
      lengthInBits += 8;
    }

    // Simple field (timeFromLastRecoverOfMainsInSeconds)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static IdentifyReplyCommandBuilder staticParseIdentifyReplyCommandBuilder(
      ReadBuffer readBuffer, Attribute attribute, Short numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandOutputUnitSummary");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    IdentifyReplyCommandUnitSummary unitFlags =
        readSimpleField(
            "unitFlags",
            new DataReaderComplexDefault<>(
                () -> IdentifyReplyCommandUnitSummary.staticParse(readBuffer), readBuffer));

    Byte gavStoreEnabledByte1 =
        readOptionalField("gavStoreEnabledByte1", readByte(readBuffer, 8), (numBytes) > (1));

    Byte gavStoreEnabledByte2 =
        readOptionalField("gavStoreEnabledByte2", readByte(readBuffer, 8), (numBytes) > (2));

    short timeFromLastRecoverOfMainsInSeconds =
        readSimpleField("timeFromLastRecoverOfMainsInSeconds", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("IdentifyReplyCommandOutputUnitSummary");
    // Create the instance
    return new IdentifyReplyCommandOutputUnitSummaryBuilderImpl(
        unitFlags,
        gavStoreEnabledByte1,
        gavStoreEnabledByte2,
        timeFromLastRecoverOfMainsInSeconds,
        numBytes);
  }

  public static class IdentifyReplyCommandOutputUnitSummaryBuilderImpl
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final IdentifyReplyCommandUnitSummary unitFlags;
    private final Byte gavStoreEnabledByte1;
    private final Byte gavStoreEnabledByte2;
    private final short timeFromLastRecoverOfMainsInSeconds;
    private final Short numBytes;

    public IdentifyReplyCommandOutputUnitSummaryBuilderImpl(
        IdentifyReplyCommandUnitSummary unitFlags,
        Byte gavStoreEnabledByte1,
        Byte gavStoreEnabledByte2,
        short timeFromLastRecoverOfMainsInSeconds,
        Short numBytes) {
      this.unitFlags = unitFlags;
      this.gavStoreEnabledByte1 = gavStoreEnabledByte1;
      this.gavStoreEnabledByte2 = gavStoreEnabledByte2;
      this.timeFromLastRecoverOfMainsInSeconds = timeFromLastRecoverOfMainsInSeconds;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandOutputUnitSummary build(Short numBytes) {

      IdentifyReplyCommandOutputUnitSummary identifyReplyCommandOutputUnitSummary =
          new IdentifyReplyCommandOutputUnitSummary(
              unitFlags,
              gavStoreEnabledByte1,
              gavStoreEnabledByte2,
              timeFromLastRecoverOfMainsInSeconds,
              numBytes);
      return identifyReplyCommandOutputUnitSummary;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandOutputUnitSummary)) {
      return false;
    }
    IdentifyReplyCommandOutputUnitSummary that = (IdentifyReplyCommandOutputUnitSummary) o;
    return (getUnitFlags() == that.getUnitFlags())
        && (getGavStoreEnabledByte1() == that.getGavStoreEnabledByte1())
        && (getGavStoreEnabledByte2() == that.getGavStoreEnabledByte2())
        && (getTimeFromLastRecoverOfMainsInSeconds()
            == that.getTimeFromLastRecoverOfMainsInSeconds())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getUnitFlags(),
        getGavStoreEnabledByte1(),
        getGavStoreEnabledByte2(),
        getTimeFromLastRecoverOfMainsInSeconds());
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
