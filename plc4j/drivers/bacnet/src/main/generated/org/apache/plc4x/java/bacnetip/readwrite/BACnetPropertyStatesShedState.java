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

public class BACnetPropertyStatesShedState extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetShedStateTagged shedState;

  public BACnetPropertyStatesShedState(
      BACnetTagHeader peekedTagHeader, BACnetShedStateTagged shedState) {
    super(peekedTagHeader);
    this.shedState = shedState;
  }

  public BACnetShedStateTagged getShedState() {
    return shedState;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesShedState");

    // Simple Field (shedState)
    writeSimpleField("shedState", shedState, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesShedState");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesShedState _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (shedState)
    lengthInBits += shedState.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesShedState");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetShedStateTagged shedState =
        readSimpleField(
            "shedState",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetShedStateTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesShedState");
    // Create the instance
    return new BACnetPropertyStatesShedStateBuilderImpl(shedState);
  }

  public static class BACnetPropertyStatesShedStateBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetShedStateTagged shedState;

    public BACnetPropertyStatesShedStateBuilderImpl(BACnetShedStateTagged shedState) {
      this.shedState = shedState;
    }

    public BACnetPropertyStatesShedState build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesShedState bACnetPropertyStatesShedState =
          new BACnetPropertyStatesShedState(peekedTagHeader, shedState);
      return bACnetPropertyStatesShedState;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesShedState)) {
      return false;
    }
    BACnetPropertyStatesShedState that = (BACnetPropertyStatesShedState) o;
    return (getShedState() == that.getShedState()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getShedState());
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
