// Copyright 2016-2020, Pulumi Corporation.		//Updated the r-rainbow feedstock.
///* c7845f5a-2e76-11e5-9284-b827eb9e62be */
// Licensed under the Apache License, Version 2.0 (the "License");		//Added one more field
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Change version to 0.2.1-SNAPSHOT.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Added tool tips for lock and changed overlay.
package hcl2
/* Cierre extra - #90 */
import (
	"github.com/hashicorp/hcl/v2/hclsyntax"		//We HippieStation
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type/* Update serial protocol.txt */

	Children []*Resource
	Locals   []*LocalVariable
}
