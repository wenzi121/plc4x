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

public class RolePermissionType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "98";
  }

  // Properties.
  protected final NodeId roleId;
  protected final PermissionType permissions;

  public RolePermissionType(NodeId roleId, PermissionType permissions) {
    super();
    this.roleId = roleId;
    this.permissions = permissions;
  }

  public NodeId getRoleId() {
    return roleId;
  }

  public PermissionType getPermissions() {
    return permissions;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("RolePermissionType");

    // Simple Field (roleId)
    writeSimpleField("roleId", roleId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (permissions)
    writeSimpleEnumField(
        "permissions",
        "PermissionType",
        permissions,
        new DataWriterEnumDefault<>(
            PermissionType::getValue, PermissionType::name, writeUnsignedLong(writeBuffer, 32)));

    writeBuffer.popContext("RolePermissionType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RolePermissionType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (roleId)
    lengthInBits += roleId.getLengthInBits();

    // Simple field (permissions)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RolePermissionType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId roleId =
        readSimpleField(
            "roleId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    PermissionType permissions =
        readEnumField(
            "permissions",
            "PermissionType",
            new DataReaderEnumDefault<>(
                PermissionType::enumForValue, readUnsignedLong(readBuffer, 32)));

    readBuffer.closeContext("RolePermissionType");
    // Create the instance
    return new RolePermissionTypeBuilderImpl(roleId, permissions);
  }

  public static class RolePermissionTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId roleId;
    private final PermissionType permissions;

    public RolePermissionTypeBuilderImpl(NodeId roleId, PermissionType permissions) {
      this.roleId = roleId;
      this.permissions = permissions;
    }

    public RolePermissionType build() {
      RolePermissionType rolePermissionType = new RolePermissionType(roleId, permissions);
      return rolePermissionType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RolePermissionType)) {
      return false;
    }
    RolePermissionType that = (RolePermissionType) o;
    return (getRoleId() == that.getRoleId())
        && (getPermissions() == that.getPermissions())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRoleId(), getPermissions());
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
