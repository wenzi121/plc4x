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
package org.apache.plc4x.java.opcua.readwrite;

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

public class UnsignedRationalNumber extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "24109";
  }

  // Properties.
  protected final long numerator;
  protected final long denominator;

  public UnsignedRationalNumber(long numerator, long denominator) {
    super();
    this.numerator = numerator;
    this.denominator = denominator;
  }

  public long getNumerator() {
    return numerator;
  }

  public long getDenominator() {
    return denominator;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("UnsignedRationalNumber");

    // Simple Field (numerator)
    writeSimpleField("numerator", numerator, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (denominator)
    writeSimpleField("denominator", denominator, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("UnsignedRationalNumber");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UnsignedRationalNumber _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (numerator)
    lengthInBits += 32;

    // Simple field (denominator)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("UnsignedRationalNumber");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long numerator = readSimpleField("numerator", readUnsignedLong(readBuffer, 32));

    long denominator = readSimpleField("denominator", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("UnsignedRationalNumber");
    // Create the instance
    return new UnsignedRationalNumberBuilderImpl(numerator, denominator);
  }

  public static class UnsignedRationalNumberBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long numerator;
    private final long denominator;

    public UnsignedRationalNumberBuilderImpl(long numerator, long denominator) {
      this.numerator = numerator;
      this.denominator = denominator;
    }

    public UnsignedRationalNumber build() {
      UnsignedRationalNumber unsignedRationalNumber =
          new UnsignedRationalNumber(numerator, denominator);
      return unsignedRationalNumber;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UnsignedRationalNumber)) {
      return false;
    }
    UnsignedRationalNumber that = (UnsignedRationalNumber) o;
    return (getNumerator() == that.getNumerator())
        && (getDenominator() == that.getDenominator())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNumerator(), getDenominator());
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
