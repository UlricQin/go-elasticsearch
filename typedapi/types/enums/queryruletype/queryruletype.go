// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe

// Package queryruletype
package queryruletype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe/specification/query_ruleset/_types/QueryRuleset.ts#L44-L46
type QueryRuleType struct {
	Name string
}

var (
	Pinned = QueryRuleType{"pinned"}
)

func (q QueryRuleType) MarshalText() (text []byte, err error) {
	return []byte(q.String()), nil
}

func (q *QueryRuleType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "pinned":
		*q = Pinned
	default:
		*q = QueryRuleType{string(text)}
	}

	return nil
}

func (q QueryRuleType) String() string {
	return q.Name
}