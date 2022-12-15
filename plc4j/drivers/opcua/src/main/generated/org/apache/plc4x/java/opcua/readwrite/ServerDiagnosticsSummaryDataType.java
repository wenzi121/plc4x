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

public class ServerDiagnosticsSummaryDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "861";
  }

  // Properties.
  protected final long serverViewCount;
  protected final long currentSessionCount;
  protected final long cumulatedSessionCount;
  protected final long securityRejectedSessionCount;
  protected final long rejectedSessionCount;
  protected final long sessionTimeoutCount;
  protected final long sessionAbortCount;
  protected final long currentSubscriptionCount;
  protected final long cumulatedSubscriptionCount;
  protected final long publishingIntervalCount;
  protected final long securityRejectedRequestsCount;
  protected final long rejectedRequestsCount;

  public ServerDiagnosticsSummaryDataType(
      long serverViewCount,
      long currentSessionCount,
      long cumulatedSessionCount,
      long securityRejectedSessionCount,
      long rejectedSessionCount,
      long sessionTimeoutCount,
      long sessionAbortCount,
      long currentSubscriptionCount,
      long cumulatedSubscriptionCount,
      long publishingIntervalCount,
      long securityRejectedRequestsCount,
      long rejectedRequestsCount) {
    super();
    this.serverViewCount = serverViewCount;
    this.currentSessionCount = currentSessionCount;
    this.cumulatedSessionCount = cumulatedSessionCount;
    this.securityRejectedSessionCount = securityRejectedSessionCount;
    this.rejectedSessionCount = rejectedSessionCount;
    this.sessionTimeoutCount = sessionTimeoutCount;
    this.sessionAbortCount = sessionAbortCount;
    this.currentSubscriptionCount = currentSubscriptionCount;
    this.cumulatedSubscriptionCount = cumulatedSubscriptionCount;
    this.publishingIntervalCount = publishingIntervalCount;
    this.securityRejectedRequestsCount = securityRejectedRequestsCount;
    this.rejectedRequestsCount = rejectedRequestsCount;
  }

  public long getServerViewCount() {
    return serverViewCount;
  }

  public long getCurrentSessionCount() {
    return currentSessionCount;
  }

  public long getCumulatedSessionCount() {
    return cumulatedSessionCount;
  }

  public long getSecurityRejectedSessionCount() {
    return securityRejectedSessionCount;
  }

  public long getRejectedSessionCount() {
    return rejectedSessionCount;
  }

  public long getSessionTimeoutCount() {
    return sessionTimeoutCount;
  }

  public long getSessionAbortCount() {
    return sessionAbortCount;
  }

  public long getCurrentSubscriptionCount() {
    return currentSubscriptionCount;
  }

  public long getCumulatedSubscriptionCount() {
    return cumulatedSubscriptionCount;
  }

  public long getPublishingIntervalCount() {
    return publishingIntervalCount;
  }

  public long getSecurityRejectedRequestsCount() {
    return securityRejectedRequestsCount;
  }

  public long getRejectedRequestsCount() {
    return rejectedRequestsCount;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ServerDiagnosticsSummaryDataType");

    // Simple Field (serverViewCount)
    writeSimpleField("serverViewCount", serverViewCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (currentSessionCount)
    writeSimpleField(
        "currentSessionCount", currentSessionCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (cumulatedSessionCount)
    writeSimpleField(
        "cumulatedSessionCount", cumulatedSessionCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (securityRejectedSessionCount)
    writeSimpleField(
        "securityRejectedSessionCount",
        securityRejectedSessionCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (rejectedSessionCount)
    writeSimpleField(
        "rejectedSessionCount", rejectedSessionCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (sessionTimeoutCount)
    writeSimpleField(
        "sessionTimeoutCount", sessionTimeoutCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (sessionAbortCount)
    writeSimpleField("sessionAbortCount", sessionAbortCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (currentSubscriptionCount)
    writeSimpleField(
        "currentSubscriptionCount", currentSubscriptionCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (cumulatedSubscriptionCount)
    writeSimpleField(
        "cumulatedSubscriptionCount",
        cumulatedSubscriptionCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (publishingIntervalCount)
    writeSimpleField(
        "publishingIntervalCount", publishingIntervalCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (securityRejectedRequestsCount)
    writeSimpleField(
        "securityRejectedRequestsCount",
        securityRejectedRequestsCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (rejectedRequestsCount)
    writeSimpleField(
        "rejectedRequestsCount", rejectedRequestsCount, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("ServerDiagnosticsSummaryDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ServerDiagnosticsSummaryDataType _value = this;

    // Simple field (serverViewCount)
    lengthInBits += 32;

    // Simple field (currentSessionCount)
    lengthInBits += 32;

    // Simple field (cumulatedSessionCount)
    lengthInBits += 32;

    // Simple field (securityRejectedSessionCount)
    lengthInBits += 32;

    // Simple field (rejectedSessionCount)
    lengthInBits += 32;

    // Simple field (sessionTimeoutCount)
    lengthInBits += 32;

    // Simple field (sessionAbortCount)
    lengthInBits += 32;

    // Simple field (currentSubscriptionCount)
    lengthInBits += 32;

    // Simple field (cumulatedSubscriptionCount)
    lengthInBits += 32;

    // Simple field (publishingIntervalCount)
    lengthInBits += 32;

    // Simple field (securityRejectedRequestsCount)
    lengthInBits += 32;

    // Simple field (rejectedRequestsCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ServerDiagnosticsSummaryDataTypeBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ServerDiagnosticsSummaryDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    long serverViewCount = readSimpleField("serverViewCount", readUnsignedLong(readBuffer, 32));

    long currentSessionCount =
        readSimpleField("currentSessionCount", readUnsignedLong(readBuffer, 32));

    long cumulatedSessionCount =
        readSimpleField("cumulatedSessionCount", readUnsignedLong(readBuffer, 32));

    long securityRejectedSessionCount =
        readSimpleField("securityRejectedSessionCount", readUnsignedLong(readBuffer, 32));

    long rejectedSessionCount =
        readSimpleField("rejectedSessionCount", readUnsignedLong(readBuffer, 32));

    long sessionTimeoutCount =
        readSimpleField("sessionTimeoutCount", readUnsignedLong(readBuffer, 32));

    long sessionAbortCount = readSimpleField("sessionAbortCount", readUnsignedLong(readBuffer, 32));

    long currentSubscriptionCount =
        readSimpleField("currentSubscriptionCount", readUnsignedLong(readBuffer, 32));

    long cumulatedSubscriptionCount =
        readSimpleField("cumulatedSubscriptionCount", readUnsignedLong(readBuffer, 32));

    long publishingIntervalCount =
        readSimpleField("publishingIntervalCount", readUnsignedLong(readBuffer, 32));

    long securityRejectedRequestsCount =
        readSimpleField("securityRejectedRequestsCount", readUnsignedLong(readBuffer, 32));

    long rejectedRequestsCount =
        readSimpleField("rejectedRequestsCount", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("ServerDiagnosticsSummaryDataType");
    // Create the instance
    return new ServerDiagnosticsSummaryDataTypeBuilder(
        serverViewCount,
        currentSessionCount,
        cumulatedSessionCount,
        securityRejectedSessionCount,
        rejectedSessionCount,
        sessionTimeoutCount,
        sessionAbortCount,
        currentSubscriptionCount,
        cumulatedSubscriptionCount,
        publishingIntervalCount,
        securityRejectedRequestsCount,
        rejectedRequestsCount);
  }

  public static class ServerDiagnosticsSummaryDataTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long serverViewCount;
    private final long currentSessionCount;
    private final long cumulatedSessionCount;
    private final long securityRejectedSessionCount;
    private final long rejectedSessionCount;
    private final long sessionTimeoutCount;
    private final long sessionAbortCount;
    private final long currentSubscriptionCount;
    private final long cumulatedSubscriptionCount;
    private final long publishingIntervalCount;
    private final long securityRejectedRequestsCount;
    private final long rejectedRequestsCount;

    public ServerDiagnosticsSummaryDataTypeBuilder(
        long serverViewCount,
        long currentSessionCount,
        long cumulatedSessionCount,
        long securityRejectedSessionCount,
        long rejectedSessionCount,
        long sessionTimeoutCount,
        long sessionAbortCount,
        long currentSubscriptionCount,
        long cumulatedSubscriptionCount,
        long publishingIntervalCount,
        long securityRejectedRequestsCount,
        long rejectedRequestsCount) {

      this.serverViewCount = serverViewCount;
      this.currentSessionCount = currentSessionCount;
      this.cumulatedSessionCount = cumulatedSessionCount;
      this.securityRejectedSessionCount = securityRejectedSessionCount;
      this.rejectedSessionCount = rejectedSessionCount;
      this.sessionTimeoutCount = sessionTimeoutCount;
      this.sessionAbortCount = sessionAbortCount;
      this.currentSubscriptionCount = currentSubscriptionCount;
      this.cumulatedSubscriptionCount = cumulatedSubscriptionCount;
      this.publishingIntervalCount = publishingIntervalCount;
      this.securityRejectedRequestsCount = securityRejectedRequestsCount;
      this.rejectedRequestsCount = rejectedRequestsCount;
    }

    public ServerDiagnosticsSummaryDataType build() {
      ServerDiagnosticsSummaryDataType serverDiagnosticsSummaryDataType =
          new ServerDiagnosticsSummaryDataType(
              serverViewCount,
              currentSessionCount,
              cumulatedSessionCount,
              securityRejectedSessionCount,
              rejectedSessionCount,
              sessionTimeoutCount,
              sessionAbortCount,
              currentSubscriptionCount,
              cumulatedSubscriptionCount,
              publishingIntervalCount,
              securityRejectedRequestsCount,
              rejectedRequestsCount);
      return serverDiagnosticsSummaryDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ServerDiagnosticsSummaryDataType)) {
      return false;
    }
    ServerDiagnosticsSummaryDataType that = (ServerDiagnosticsSummaryDataType) o;
    return (getServerViewCount() == that.getServerViewCount())
        && (getCurrentSessionCount() == that.getCurrentSessionCount())
        && (getCumulatedSessionCount() == that.getCumulatedSessionCount())
        && (getSecurityRejectedSessionCount() == that.getSecurityRejectedSessionCount())
        && (getRejectedSessionCount() == that.getRejectedSessionCount())
        && (getSessionTimeoutCount() == that.getSessionTimeoutCount())
        && (getSessionAbortCount() == that.getSessionAbortCount())
        && (getCurrentSubscriptionCount() == that.getCurrentSubscriptionCount())
        && (getCumulatedSubscriptionCount() == that.getCumulatedSubscriptionCount())
        && (getPublishingIntervalCount() == that.getPublishingIntervalCount())
        && (getSecurityRejectedRequestsCount() == that.getSecurityRejectedRequestsCount())
        && (getRejectedRequestsCount() == that.getRejectedRequestsCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getServerViewCount(),
        getCurrentSessionCount(),
        getCumulatedSessionCount(),
        getSecurityRejectedSessionCount(),
        getRejectedSessionCount(),
        getSessionTimeoutCount(),
        getSessionAbortCount(),
        getCurrentSubscriptionCount(),
        getCumulatedSubscriptionCount(),
        getPublishingIntervalCount(),
        getSecurityRejectedRequestsCount(),
        getRejectedRequestsCount());
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