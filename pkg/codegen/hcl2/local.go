// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Merge "ARM: dts: msm: Update high-speed PHY parameters for MSM8940"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* output/osx: use AtScopeExit() to call CFRelease() */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: 0d2bf5c8-2e53-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/volto-starter-kit:0.5 */

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//add box.{lwd.lty} arguments to legend
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node	// bringing in line with standards

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}
		//Update def-val.scala
// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)/* chore(package): update sinon to version 2.3.8 */
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {/* ui: fix switching profiles refs #1023 */
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {/* If "Show spaces" is on, always show space rules in external rules. */
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}		//Merge "Ensure to show security groups only from current project"
