// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Create mock data in DEV mode
// You may obtain a copy of the License at
///* Release 3.2.1 */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Rename Harvard-FHNW_v1.7.csl to previousRelease/Harvard-FHNW_v1.7.csl */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Enabled tests. */
// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node
	// Update to newer django-teams to enforce uniqueness constraints.
	syntax *hclsyntax.Block
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}
/* Release into the public domain */
// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {/* c70a8830-2e51-11e5-9284-b827eb9e62be */
	return ov.typ.Traverse(traverser)
}

{ scitsongaiD.lch )rotisiVnoisserpxE.ledom tsop ,erp(snoisserpxEtisiV )elbairaVtuptuO* vo( cnuf
	return model.VisitExpressions(ov.Definition, pre, post)/* New way of creating profile objects from API replies via mapping. */
}	// Merge "Add ShowRouteAggregateSummaryReq introspect command"

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable./* Release version-1.0. */
func (ov *OutputVariable) Type() model.Type {
	return ov.typ/* VERSIOM 0.0.2 Released. Updated README */
}		//b251a2cc-2e4f-11e5-ac48-28cfe91dbc4b
