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

public class Ethernet_FramePayload_IPv4 extends Ethernet_FramePayload implements Message {

  // Accessors for discriminator values.
  public Integer getPacketType() {
    return (int) 0x0800;
  }

  // Constant values.
  public static final Byte VERSION = 0x4;
  public static final Byte HEADERLENGTH = 0x5;
  public static final Short DIFFERENTIATEDSERVICESCODEPOINT = 0x00;
  public static final Byte EXPLICITCONGESTIONNOTIFICATION = 0x0;
  public static final Byte FLAGS = 0x00;
  public static final Integer FRAGMENTOFFSET = 0x00;
  public static final Short PROTOCOL = 0x11;

  // Properties.
  protected final int identification;
  protected final short timeToLive;
  protected final IpAddress sourceAddress;
  protected final IpAddress destinationAddress;
  protected final int sourcePort;
  protected final int destinationPort;
  protected final DceRpc_Packet payload;

  public Ethernet_FramePayload_IPv4(
      int identification,
      short timeToLive,
      IpAddress sourceAddress,
      IpAddress destinationAddress,
      int sourcePort,
      int destinationPort,
      DceRpc_Packet payload) {
    super();
    this.identification = identification;
    this.timeToLive = timeToLive;
    this.sourceAddress = sourceAddress;
    this.destinationAddress = destinationAddress;
    this.sourcePort = sourcePort;
    this.destinationPort = destinationPort;
    this.payload = payload;
  }

  public int getIdentification() {
    return identification;
  }

  public short getTimeToLive() {
    return timeToLive;
  }

  public IpAddress getSourceAddress() {
    return sourceAddress;
  }

  public IpAddress getDestinationAddress() {
    return destinationAddress;
  }

  public int getSourcePort() {
    return sourcePort;
  }

  public int getDestinationPort() {
    return destinationPort;
  }

  public DceRpc_Packet getPayload() {
    return payload;
  }

  public byte getVersion() {
    return VERSION;
  }

  public byte getHeaderLength() {
    return HEADERLENGTH;
  }

  public short getDifferentiatedServicesCodepoint() {
    return DIFFERENTIATEDSERVICESCODEPOINT;
  }

  public byte getExplicitCongestionNotification() {
    return EXPLICITCONGESTIONNOTIFICATION;
  }

  public byte getFlags() {
    return FLAGS;
  }

  public int getFragmentOffset() {
    return FRAGMENTOFFSET;
  }

  public short getProtocol() {
    return PROTOCOL;
  }

  @Override
  protected void serializeEthernet_FramePayloadChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Ethernet_FramePayload_IPv4");

    // Const Field (version)
    writeConstField("version", VERSION, writeUnsignedByte(writeBuffer, 4));

    // Const Field (headerLength)
    writeConstField("headerLength", HEADERLENGTH, writeUnsignedByte(writeBuffer, 4));

    // Const Field (differentiatedServicesCodepoint)
    writeConstField(
        "differentiatedServicesCodepoint",
        DIFFERENTIATEDSERVICESCODEPOINT,
        writeUnsignedShort(writeBuffer, 6));

    // Const Field (explicitCongestionNotification)
    writeConstField(
        "explicitCongestionNotification",
        EXPLICITCONGESTIONNOTIFICATION,
        writeUnsignedByte(writeBuffer, 2));

    // Implicit Field (totalLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int totalLength = (int) ((28) + (getPayload().getLengthInBytes()));
    writeImplicitField("totalLength", totalLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (identification)
    writeSimpleField("identification", identification, writeUnsignedInt(writeBuffer, 16));

    // Const Field (flags)
    writeConstField("flags", FLAGS, writeUnsignedByte(writeBuffer, 3));

    // Const Field (fragmentOffset)
    writeConstField("fragmentOffset", FRAGMENTOFFSET, writeUnsignedInt(writeBuffer, 13));

    // Simple Field (timeToLive)
    writeSimpleField("timeToLive", timeToLive, writeUnsignedShort(writeBuffer, 8));

    // Const Field (protocol)
    writeConstField("protocol", PROTOCOL, writeUnsignedShort(writeBuffer, 8));

    // Implicit Field (headerChecksum) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int headerChecksum =
        (int)
            (org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.calculateIPv4Checksum(
                (28) + (getPayload().getLengthInBytes()),
                getIdentification(),
                getTimeToLive(),
                getSourceAddress(),
                getDestinationAddress()));
    writeImplicitField("headerChecksum", headerChecksum, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (sourceAddress)
    writeSimpleField("sourceAddress", sourceAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (destinationAddress)
    writeSimpleField(
        "destinationAddress", destinationAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (sourcePort)
    writeSimpleField("sourcePort", sourcePort, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (destinationPort)
    writeSimpleField("destinationPort", destinationPort, writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (packetLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int packetLength = (int) (getLengthInBytes());
    writeImplicitField("packetLength", packetLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("Ethernet_FramePayload_IPv4");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Ethernet_FramePayload_IPv4 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (version)
    lengthInBits += 4;

    // Const Field (headerLength)
    lengthInBits += 4;

    // Const Field (differentiatedServicesCodepoint)
    lengthInBits += 6;

    // Const Field (explicitCongestionNotification)
    lengthInBits += 2;

    // Implicit Field (totalLength)
    lengthInBits += 16;

    // Simple field (identification)
    lengthInBits += 16;

    // Const Field (flags)
    lengthInBits += 3;

    // Const Field (fragmentOffset)
    lengthInBits += 13;

    // Simple field (timeToLive)
    lengthInBits += 8;

    // Const Field (protocol)
    lengthInBits += 8;

    // Implicit Field (headerChecksum)
    lengthInBits += 16;

    // Simple field (sourceAddress)
    lengthInBits += sourceAddress.getLengthInBits();

    // Simple field (destinationAddress)
    lengthInBits += destinationAddress.getLengthInBits();

    // Simple field (sourcePort)
    lengthInBits += 16;

    // Simple field (destinationPort)
    lengthInBits += 16;

    // Implicit Field (packetLength)
    lengthInBits += 16;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static Ethernet_FramePayloadBuilder staticParseEthernet_FramePayloadBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Ethernet_FramePayload_IPv4");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte version =
        readConstField(
            "version", readUnsignedByte(readBuffer, 4), Ethernet_FramePayload_IPv4.VERSION);

    byte headerLength =
        readConstField(
            "headerLength",
            readUnsignedByte(readBuffer, 4),
            Ethernet_FramePayload_IPv4.HEADERLENGTH);

    short differentiatedServicesCodepoint =
        readConstField(
            "differentiatedServicesCodepoint",
            readUnsignedShort(readBuffer, 6),
            Ethernet_FramePayload_IPv4.DIFFERENTIATEDSERVICESCODEPOINT);

    byte explicitCongestionNotification =
        readConstField(
            "explicitCongestionNotification",
            readUnsignedByte(readBuffer, 2),
            Ethernet_FramePayload_IPv4.EXPLICITCONGESTIONNOTIFICATION);

    int totalLength = readImplicitField("totalLength", readUnsignedInt(readBuffer, 16));

    int identification = readSimpleField("identification", readUnsignedInt(readBuffer, 16));

    byte flags =
        readConstField("flags", readUnsignedByte(readBuffer, 3), Ethernet_FramePayload_IPv4.FLAGS);

    int fragmentOffset =
        readConstField(
            "fragmentOffset",
            readUnsignedInt(readBuffer, 13),
            Ethernet_FramePayload_IPv4.FRAGMENTOFFSET);

    short timeToLive = readSimpleField("timeToLive", readUnsignedShort(readBuffer, 8));

    short protocol =
        readConstField(
            "protocol", readUnsignedShort(readBuffer, 8), Ethernet_FramePayload_IPv4.PROTOCOL);

    int headerChecksum = readImplicitField("headerChecksum", readUnsignedInt(readBuffer, 16));

    IpAddress sourceAddress =
        readSimpleField(
            "sourceAddress",
            new DataReaderComplexDefault<>(() -> IpAddress.staticParse(readBuffer), readBuffer));

    IpAddress destinationAddress =
        readSimpleField(
            "destinationAddress",
            new DataReaderComplexDefault<>(() -> IpAddress.staticParse(readBuffer), readBuffer));

    int sourcePort = readSimpleField("sourcePort", readUnsignedInt(readBuffer, 16));

    int destinationPort = readSimpleField("destinationPort", readUnsignedInt(readBuffer, 16));

    int packetLength = readImplicitField("packetLength", readUnsignedInt(readBuffer, 16));

    DceRpc_Packet payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> DceRpc_Packet.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("Ethernet_FramePayload_IPv4");
    // Create the instance
    return new Ethernet_FramePayload_IPv4BuilderImpl(
        identification,
        timeToLive,
        sourceAddress,
        destinationAddress,
        sourcePort,
        destinationPort,
        payload);
  }

  public static class Ethernet_FramePayload_IPv4BuilderImpl
      implements Ethernet_FramePayload.Ethernet_FramePayloadBuilder {
    private final int identification;
    private final short timeToLive;
    private final IpAddress sourceAddress;
    private final IpAddress destinationAddress;
    private final int sourcePort;
    private final int destinationPort;
    private final DceRpc_Packet payload;

    public Ethernet_FramePayload_IPv4BuilderImpl(
        int identification,
        short timeToLive,
        IpAddress sourceAddress,
        IpAddress destinationAddress,
        int sourcePort,
        int destinationPort,
        DceRpc_Packet payload) {
      this.identification = identification;
      this.timeToLive = timeToLive;
      this.sourceAddress = sourceAddress;
      this.destinationAddress = destinationAddress;
      this.sourcePort = sourcePort;
      this.destinationPort = destinationPort;
      this.payload = payload;
    }

    public Ethernet_FramePayload_IPv4 build() {
      Ethernet_FramePayload_IPv4 ethernet_FramePayload_IPv4 =
          new Ethernet_FramePayload_IPv4(
              identification,
              timeToLive,
              sourceAddress,
              destinationAddress,
              sourcePort,
              destinationPort,
              payload);
      return ethernet_FramePayload_IPv4;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Ethernet_FramePayload_IPv4)) {
      return false;
    }
    Ethernet_FramePayload_IPv4 that = (Ethernet_FramePayload_IPv4) o;
    return (getIdentification() == that.getIdentification())
        && (getTimeToLive() == that.getTimeToLive())
        && (getSourceAddress() == that.getSourceAddress())
        && (getDestinationAddress() == that.getDestinationAddress())
        && (getSourcePort() == that.getSourcePort())
        && (getDestinationPort() == that.getDestinationPort())
        && (getPayload() == that.getPayload())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getIdentification(),
        getTimeToLive(),
        getSourceAddress(),
        getDestinationAddress(),
        getSourcePort(),
        getDestinationPort(),
        getPayload());
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
