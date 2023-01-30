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

public class CBusMessageToClient extends CBusMessage implements Message {

  // Accessors for discriminator values.
  public Boolean getIsResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final ReplyOrConfirmation reply;

  // Arguments.
  protected final RequestContext requestContext;
  protected final CBusOptions cBusOptions;

  public CBusMessageToClient(
      ReplyOrConfirmation reply, RequestContext requestContext, CBusOptions cBusOptions) {
    super(requestContext, cBusOptions);
    this.reply = reply;
    this.requestContext = requestContext;
    this.cBusOptions = cBusOptions;
  }

  public ReplyOrConfirmation getReply() {
    return reply;
  }

  @Override
  protected void serializeCBusMessageChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusMessageToClient");

    // Simple Field (reply)
    writeSimpleField("reply", reply, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CBusMessageToClient");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusMessageToClient _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (reply)
    lengthInBits += reply.getLengthInBits();

    return lengthInBits;
  }

  public static CBusMessageBuilder staticParseCBusMessageBuilder(
      ReadBuffer readBuffer,
      Boolean isResponse,
      RequestContext requestContext,
      CBusOptions cBusOptions)
      throws ParseException {
    readBuffer.pullContext("CBusMessageToClient");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReplyOrConfirmation reply =
        readSimpleField(
            "reply",
            new DataReaderComplexDefault<>(
                () ->
                    ReplyOrConfirmation.staticParse(
                        readBuffer, (CBusOptions) (cBusOptions), (RequestContext) (requestContext)),
                readBuffer));

    readBuffer.closeContext("CBusMessageToClient");
    // Create the instance
    return new CBusMessageToClientBuilderImpl(reply, requestContext, cBusOptions);
  }

  public static class CBusMessageToClientBuilderImpl implements CBusMessage.CBusMessageBuilder {
    private final ReplyOrConfirmation reply;
    private final RequestContext requestContext;
    private final CBusOptions cBusOptions;

    public CBusMessageToClientBuilderImpl(
        ReplyOrConfirmation reply, RequestContext requestContext, CBusOptions cBusOptions) {
      this.reply = reply;
      this.requestContext = requestContext;
      this.cBusOptions = cBusOptions;
    }

    public CBusMessageToClient build(RequestContext requestContext, CBusOptions cBusOptions) {

      CBusMessageToClient cBusMessageToClient =
          new CBusMessageToClient(reply, requestContext, cBusOptions);
      return cBusMessageToClient;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusMessageToClient)) {
      return false;
    }
    CBusMessageToClient that = (CBusMessageToClient) o;
    return (getReply() == that.getReply()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReply());
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
