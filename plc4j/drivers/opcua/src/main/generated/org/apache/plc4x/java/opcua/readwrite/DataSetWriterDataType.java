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

public class DataSetWriterDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "15599";
  }

  // Properties.
  protected final PascalString name;
  protected final boolean enabled;
  protected final int dataSetWriterId;
  protected final DataSetFieldContentMask dataSetFieldContentMask;
  protected final long keyFrameCount;
  protected final PascalString dataSetName;
  protected final int noOfDataSetWriterProperties;
  protected final List<ExtensionObjectDefinition> dataSetWriterProperties;
  protected final ExtensionObject transportSettings;
  protected final ExtensionObject messageSettings;

  public DataSetWriterDataType(
      PascalString name,
      boolean enabled,
      int dataSetWriterId,
      DataSetFieldContentMask dataSetFieldContentMask,
      long keyFrameCount,
      PascalString dataSetName,
      int noOfDataSetWriterProperties,
      List<ExtensionObjectDefinition> dataSetWriterProperties,
      ExtensionObject transportSettings,
      ExtensionObject messageSettings) {
    super();
    this.name = name;
    this.enabled = enabled;
    this.dataSetWriterId = dataSetWriterId;
    this.dataSetFieldContentMask = dataSetFieldContentMask;
    this.keyFrameCount = keyFrameCount;
    this.dataSetName = dataSetName;
    this.noOfDataSetWriterProperties = noOfDataSetWriterProperties;
    this.dataSetWriterProperties = dataSetWriterProperties;
    this.transportSettings = transportSettings;
    this.messageSettings = messageSettings;
  }

  public PascalString getName() {
    return name;
  }

  public boolean getEnabled() {
    return enabled;
  }

  public int getDataSetWriterId() {
    return dataSetWriterId;
  }

  public DataSetFieldContentMask getDataSetFieldContentMask() {
    return dataSetFieldContentMask;
  }

  public long getKeyFrameCount() {
    return keyFrameCount;
  }

  public PascalString getDataSetName() {
    return dataSetName;
  }

  public int getNoOfDataSetWriterProperties() {
    return noOfDataSetWriterProperties;
  }

  public List<ExtensionObjectDefinition> getDataSetWriterProperties() {
    return dataSetWriterProperties;
  }

  public ExtensionObject getTransportSettings() {
    return transportSettings;
  }

  public ExtensionObject getMessageSettings() {
    return messageSettings;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DataSetWriterDataType");

    // Simple Field (name)
    writeSimpleField("name", name, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 7));

    // Simple Field (enabled)
    writeSimpleField("enabled", enabled, writeBoolean(writeBuffer));

    // Simple Field (dataSetWriterId)
    writeSimpleField("dataSetWriterId", dataSetWriterId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (dataSetFieldContentMask)
    writeSimpleEnumField(
        "dataSetFieldContentMask",
        "DataSetFieldContentMask",
        dataSetFieldContentMask,
        new DataWriterEnumDefault<>(
            DataSetFieldContentMask::getValue,
            DataSetFieldContentMask::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (keyFrameCount)
    writeSimpleField("keyFrameCount", keyFrameCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (dataSetName)
    writeSimpleField("dataSetName", dataSetName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfDataSetWriterProperties)
    writeSimpleField(
        "noOfDataSetWriterProperties",
        noOfDataSetWriterProperties,
        writeSignedInt(writeBuffer, 32));

    // Array Field (dataSetWriterProperties)
    writeComplexTypeArrayField("dataSetWriterProperties", dataSetWriterProperties, writeBuffer);

    // Simple Field (transportSettings)
    writeSimpleField(
        "transportSettings", transportSettings, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (messageSettings)
    writeSimpleField(
        "messageSettings", messageSettings, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("DataSetWriterDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DataSetWriterDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (enabled)
    lengthInBits += 1;

    // Simple field (dataSetWriterId)
    lengthInBits += 16;

    // Simple field (dataSetFieldContentMask)
    lengthInBits += 32;

    // Simple field (keyFrameCount)
    lengthInBits += 32;

    // Simple field (dataSetName)
    lengthInBits += dataSetName.getLengthInBits();

    // Simple field (noOfDataSetWriterProperties)
    lengthInBits += 32;

    // Array field
    if (dataSetWriterProperties != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : dataSetWriterProperties) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= dataSetWriterProperties.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (transportSettings)
    lengthInBits += transportSettings.getLengthInBits();

    // Simple field (messageSettings)
    lengthInBits += messageSettings.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("DataSetWriterDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString name =
        readSimpleField(
            "name",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean enabled = readSimpleField("enabled", readBoolean(readBuffer));

    int dataSetWriterId = readSimpleField("dataSetWriterId", readUnsignedInt(readBuffer, 16));

    DataSetFieldContentMask dataSetFieldContentMask =
        readEnumField(
            "dataSetFieldContentMask",
            "DataSetFieldContentMask",
            new DataReaderEnumDefault<>(
                DataSetFieldContentMask::enumForValue, readUnsignedLong(readBuffer, 32)));

    long keyFrameCount = readSimpleField("keyFrameCount", readUnsignedLong(readBuffer, 32));

    PascalString dataSetName =
        readSimpleField(
            "dataSetName",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfDataSetWriterProperties =
        readSimpleField("noOfDataSetWriterProperties", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> dataSetWriterProperties =
        readCountArrayField(
            "dataSetWriterProperties",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("14535")),
                readBuffer),
            noOfDataSetWriterProperties);

    ExtensionObject transportSettings =
        readSimpleField(
            "transportSettings",
            new DataReaderComplexDefault<>(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    ExtensionObject messageSettings =
        readSimpleField(
            "messageSettings",
            new DataReaderComplexDefault<>(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("DataSetWriterDataType");
    // Create the instance
    return new DataSetWriterDataTypeBuilderImpl(
        name,
        enabled,
        dataSetWriterId,
        dataSetFieldContentMask,
        keyFrameCount,
        dataSetName,
        noOfDataSetWriterProperties,
        dataSetWriterProperties,
        transportSettings,
        messageSettings);
  }

  public static class DataSetWriterDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final boolean enabled;
    private final int dataSetWriterId;
    private final DataSetFieldContentMask dataSetFieldContentMask;
    private final long keyFrameCount;
    private final PascalString dataSetName;
    private final int noOfDataSetWriterProperties;
    private final List<ExtensionObjectDefinition> dataSetWriterProperties;
    private final ExtensionObject transportSettings;
    private final ExtensionObject messageSettings;

    public DataSetWriterDataTypeBuilderImpl(
        PascalString name,
        boolean enabled,
        int dataSetWriterId,
        DataSetFieldContentMask dataSetFieldContentMask,
        long keyFrameCount,
        PascalString dataSetName,
        int noOfDataSetWriterProperties,
        List<ExtensionObjectDefinition> dataSetWriterProperties,
        ExtensionObject transportSettings,
        ExtensionObject messageSettings) {
      this.name = name;
      this.enabled = enabled;
      this.dataSetWriterId = dataSetWriterId;
      this.dataSetFieldContentMask = dataSetFieldContentMask;
      this.keyFrameCount = keyFrameCount;
      this.dataSetName = dataSetName;
      this.noOfDataSetWriterProperties = noOfDataSetWriterProperties;
      this.dataSetWriterProperties = dataSetWriterProperties;
      this.transportSettings = transportSettings;
      this.messageSettings = messageSettings;
    }

    public DataSetWriterDataType build() {
      DataSetWriterDataType dataSetWriterDataType =
          new DataSetWriterDataType(
              name,
              enabled,
              dataSetWriterId,
              dataSetFieldContentMask,
              keyFrameCount,
              dataSetName,
              noOfDataSetWriterProperties,
              dataSetWriterProperties,
              transportSettings,
              messageSettings);
      return dataSetWriterDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DataSetWriterDataType)) {
      return false;
    }
    DataSetWriterDataType that = (DataSetWriterDataType) o;
    return (getName() == that.getName())
        && (getEnabled() == that.getEnabled())
        && (getDataSetWriterId() == that.getDataSetWriterId())
        && (getDataSetFieldContentMask() == that.getDataSetFieldContentMask())
        && (getKeyFrameCount() == that.getKeyFrameCount())
        && (getDataSetName() == that.getDataSetName())
        && (getNoOfDataSetWriterProperties() == that.getNoOfDataSetWriterProperties())
        && (getDataSetWriterProperties() == that.getDataSetWriterProperties())
        && (getTransportSettings() == that.getTransportSettings())
        && (getMessageSettings() == that.getMessageSettings())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getEnabled(),
        getDataSetWriterId(),
        getDataSetFieldContentMask(),
        getKeyFrameCount(),
        getDataSetName(),
        getNoOfDataSetWriterProperties(),
        getDataSetWriterProperties(),
        getTransportSettings(),
        getMessageSettings());
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
