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

public class PnDcp_Block_DevicePropertiesDeviceInstance extends PnDcp_Block implements Message {

  // Accessors for discriminator values.
  public PnDcp_BlockOptions getOption() {
    return PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION;
  }

  public Short getSuboption() {
    return (short) 7;
  }

  // Properties.
  protected final short deviceInstanceHigh;
  protected final short deviceInstanceLow;

  public PnDcp_Block_DevicePropertiesDeviceInstance(
      short deviceInstanceHigh, short deviceInstanceLow) {
    super();
    this.deviceInstanceHigh = deviceInstanceHigh;
    this.deviceInstanceLow = deviceInstanceLow;
  }

  public short getDeviceInstanceHigh() {
    return deviceInstanceHigh;
  }

  public short getDeviceInstanceLow() {
    return deviceInstanceLow;
  }

  @Override
  protected void serializePnDcp_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnDcp_Block_DevicePropertiesDeviceInstance");

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x0000, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (deviceInstanceHigh)
    writeSimpleField("deviceInstanceHigh", deviceInstanceHigh, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (deviceInstanceLow)
    writeSimpleField("deviceInstanceLow", deviceInstanceLow, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("PnDcp_Block_DevicePropertiesDeviceInstance");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnDcp_Block_DevicePropertiesDeviceInstance _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Simple field (deviceInstanceHigh)
    lengthInBits += 8;

    // Simple field (deviceInstanceLow)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PnDcp_BlockBuilder staticParsePnDcp_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_Block_DevicePropertiesDeviceInstance");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    short deviceInstanceHigh =
        readSimpleField("deviceInstanceHigh", readUnsignedShort(readBuffer, 8));

    short deviceInstanceLow =
        readSimpleField("deviceInstanceLow", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("PnDcp_Block_DevicePropertiesDeviceInstance");
    // Create the instance
    return new PnDcp_Block_DevicePropertiesDeviceInstanceBuilderImpl(
        deviceInstanceHigh, deviceInstanceLow);
  }

  public static class PnDcp_Block_DevicePropertiesDeviceInstanceBuilderImpl
      implements PnDcp_Block.PnDcp_BlockBuilder {
    private final short deviceInstanceHigh;
    private final short deviceInstanceLow;

    public PnDcp_Block_DevicePropertiesDeviceInstanceBuilderImpl(
        short deviceInstanceHigh, short deviceInstanceLow) {
      this.deviceInstanceHigh = deviceInstanceHigh;
      this.deviceInstanceLow = deviceInstanceLow;
    }

    public PnDcp_Block_DevicePropertiesDeviceInstance build() {
      PnDcp_Block_DevicePropertiesDeviceInstance pnDcp_Block_DevicePropertiesDeviceInstance =
          new PnDcp_Block_DevicePropertiesDeviceInstance(deviceInstanceHigh, deviceInstanceLow);
      return pnDcp_Block_DevicePropertiesDeviceInstance;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block_DevicePropertiesDeviceInstance)) {
      return false;
    }
    PnDcp_Block_DevicePropertiesDeviceInstance that =
        (PnDcp_Block_DevicePropertiesDeviceInstance) o;
    return (getDeviceInstanceHigh() == that.getDeviceInstanceHigh())
        && (getDeviceInstanceLow() == that.getDeviceInstanceLow())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDeviceInstanceHigh(), getDeviceInstanceLow());
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
