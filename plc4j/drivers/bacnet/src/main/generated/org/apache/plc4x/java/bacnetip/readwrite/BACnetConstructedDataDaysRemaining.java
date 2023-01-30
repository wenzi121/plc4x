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

public class BACnetConstructedDataDaysRemaining extends BACnetConstructedData implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.DAYS_REMAINING;
  }

  // Properties.
  protected final BACnetApplicationTagSignedInteger daysRemaining;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataDaysRemaining(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagSignedInteger daysRemaining,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.daysRemaining = daysRemaining;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetApplicationTagSignedInteger getDaysRemaining() {
    return daysRemaining;
  }

  public BACnetApplicationTagSignedInteger getActualValue() {
    return (BACnetApplicationTagSignedInteger) (getDaysRemaining());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConstructedDataDaysRemaining");

    // Simple Field (daysRemaining)
    writeSimpleField("daysRemaining", daysRemaining, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetApplicationTagSignedInteger actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataDaysRemaining");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataDaysRemaining _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (daysRemaining)
    lengthInBits += daysRemaining.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataDaysRemaining");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagSignedInteger daysRemaining =
        readSimpleField(
            "daysRemaining",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagSignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));
    BACnetApplicationTagSignedInteger actualValue =
        readVirtualField("actualValue", BACnetApplicationTagSignedInteger.class, daysRemaining);

    readBuffer.closeContext("BACnetConstructedDataDaysRemaining");
    // Create the instance
    return new BACnetConstructedDataDaysRemainingBuilderImpl(
        daysRemaining, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataDaysRemainingBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetApplicationTagSignedInteger daysRemaining;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataDaysRemainingBuilderImpl(
        BACnetApplicationTagSignedInteger daysRemaining,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.daysRemaining = daysRemaining;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataDaysRemaining build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataDaysRemaining bACnetConstructedDataDaysRemaining =
          new BACnetConstructedDataDaysRemaining(
              openingTag,
              peekedTagHeader,
              closingTag,
              daysRemaining,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataDaysRemaining;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataDaysRemaining)) {
      return false;
    }
    BACnetConstructedDataDaysRemaining that = (BACnetConstructedDataDaysRemaining) o;
    return (getDaysRemaining() == that.getDaysRemaining()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDaysRemaining());
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
