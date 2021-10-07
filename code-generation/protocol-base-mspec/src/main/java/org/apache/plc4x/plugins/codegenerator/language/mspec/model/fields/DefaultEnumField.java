/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields;

import org.apache.plc4x.plugins.codegenerator.types.fields.EnumField;
import org.apache.plc4x.plugins.codegenerator.types.references.TypeReference;
import org.apache.plc4x.plugins.codegenerator.types.terms.Term;

import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;

public class DefaultEnumField extends DefaultField implements EnumField {

    private final TypeReference type;
    private final String name;
    private final String fieldName;
    private final List<Term> params;

    public DefaultEnumField(Map<String, Term> attributes, TypeReference type, String name, String fieldName, List<Term> params) {
        super(attributes);
        this.type = Objects.requireNonNull(type);
        this.name = Objects.requireNonNull(name);
        this.fieldName = fieldName;
        this.params = params;
    }

    public TypeReference getType() {
        return type;
    }

    public String getName() {
        return name;
    }

    public Optional<String> getFieldName() {
        return Optional.ofNullable(fieldName);
    }

    @Override
    public Optional<List<Term>> getParams() {
        return Optional.ofNullable(params);
    }

}
