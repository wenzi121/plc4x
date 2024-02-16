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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ChunkType {
  CONTINUE((String) "C"),
  FINAL((String) "F"),
  ABORT((String) "A");
  private static final Map<String, ChunkType> map;

  static {
    map = new HashMap<>();
    for (ChunkType value : ChunkType.values()) {
      map.put((String) value.getValue(), value);
    }
  }

  private final String value;

  ChunkType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  public static ChunkType enumForValue(String value) {
    return map.get(value);
  }

  public static Boolean isDefined(String value) {
    return map.containsKey(value);
  }
}