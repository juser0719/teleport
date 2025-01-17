// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package databaseobjectimportrule

import (
	"testing"

	"github.com/gravitational/trace"
	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/defaults"
	dbobjectimportrulev1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/dbobjectimportrule/v1"
	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	"github.com/gravitational/teleport/api/types"
)

func TestNewDatabaseObjectImportRule(t *testing.T) {
	tests := []struct {
		name          string
		spec          *dbobjectimportrulev1.DatabaseObjectImportRuleSpec
		expectedError error
	}{
		{
			name: "valid rule",
			spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{
				DbLabels: types.Labels{"key": {"value"}}.ToProto(),
				Mappings: []*dbobjectimportrulev1.DatabaseObjectImportRuleMapping{{}},
			},
			expectedError: nil,
		},
		{
			name:          "nil spec",
			spec:          nil,
			expectedError: trace.BadParameter("missing spec"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDatabaseObjectImportRule(tt.name, tt.spec)
			require.ErrorIs(t, err, tt.expectedError)
		})
	}
}

func TestValidateDatabaseObjectImportRule(t *testing.T) {
	tests := []struct {
		name          string
		rule          *dbobjectimportrulev1.DatabaseObjectImportRule
		expectedError error
	}{
		{
			name: "valid rule",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    types.KindDatabaseObjectImportRule,
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "test",
					Namespace: defaults.Namespace,
				},
				Spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{
					DbLabels: types.Labels{"key": {"value"}}.ToProto(),
					Mappings: []*dbobjectimportrulev1.DatabaseObjectImportRuleMapping{{}},
				},
			},
			expectedError: nil,
		},
		{
			name:          "nil rule",
			rule:          nil,
			expectedError: trace.BadParameter("database object import rule must be non-nil"),
		},
		{
			name: "missing metadata",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:     types.KindDatabaseObjectImportRule,
				Version:  types.V1,
				Metadata: nil,
				Spec:     &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{},
			},
			expectedError: trace.BadParameter("metadata: must be non-nil"),
		},
		{
			name: "missing name",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    types.KindDatabaseObjectImportRule,
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "",
					Namespace: defaults.Namespace,
				},
				Spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{},
			},
			expectedError: trace.BadParameter("metadata.name: must be non-empty"),
		},
		{
			name: "invalid kind",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    "InvalidKind",
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "test",
					Namespace: defaults.Namespace,
				},
				Spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{},
			},
			expectedError: trace.BadParameter("invalid kind InvalidKind, expected db_object_import_rule"),
		},
		{
			name: "missing spec",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    types.KindDatabaseObjectImportRule,
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "test",
					Namespace: defaults.Namespace,
				},
				Spec: nil,
			},
			expectedError: trace.BadParameter("missing spec"),
		},
		{
			name: "missing db_labels",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    types.KindDatabaseObjectImportRule,
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "test",
					Namespace: defaults.Namespace,
				},
				Spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{
					Mappings: []*dbobjectimportrulev1.DatabaseObjectImportRuleMapping{{}},
				},
			},
			expectedError: trace.BadParameter("missing db_labels"),
		},
		{
			name: "missing mappings",
			rule: &dbobjectimportrulev1.DatabaseObjectImportRule{
				Kind:    types.KindDatabaseObjectImportRule,
				Version: types.V1,
				Metadata: &headerv1.Metadata{
					Name:      "test",
					Namespace: defaults.Namespace,
				},
				Spec: &dbobjectimportrulev1.DatabaseObjectImportRuleSpec{
					DbLabels: types.Labels{"key": {"value"}}.ToProto(),
				},
			},
			expectedError: trace.BadParameter("missing mappings"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDatabaseObjectImportRule(tt.rule)
			require.ErrorIs(t, err, tt.expectedError)
		})
	}
}
