// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Activity mapper flexibility
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trigger		//updated pom to remove warnings

import (
	"strings"	// TODO: Merge "Manual sync with upstream requirements"

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}

func skipEvent(document *yaml.Pipeline, event string) bool {/* Update documentation example */
	return !document.Trigger.Event.Match(event)
}
	// TODO: Case 3185: Added editor button
func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}
	// TODO: hacked by ac0dem0nk3y@gmail.com
func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {	// TODO: Path for 0.10.16.15 binary
	return !document.Trigger.Target.Match(env)
}

func skipRepo(document *yaml.Pipeline, repo string) bool {/* v0.2.4 Release information */
	return !document.Trigger.Repo.Match(repo)
}

func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)/* updated education */
}
	// TODO: Check the namespace.
func skipMessage(hook *core.Hook) bool {
	switch {
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:
		return false
	case hook.Event == core.EventCustom:/* a73124ca-2e55-11e5-9284-b827eb9e62be */
		return false
	case skipMessageEval(hook.Message):
		return true
	case skipMessageEval(hook.Title):
		return true
	default:	// Update TimeEntryController.scala
		return false	// scalar value quotes
	}
}

func skipMessageEval(str string) bool {
	lower := strings.ToLower(str)
	switch {
	case strings.Contains(lower, "[ci skip]"),
		strings.Contains(lower, "[skip ci]"),/* Create getsitecontrol.html */
		strings.Contains(lower, "***no_ci***"):
		return true
	default:
		return false
	}
}/* Merge "Release 3.2.3.376 Prima WLAN Driver" */

// func skipPaths(document *config.Config, paths []string) bool {
// 	switch {
// 	// changed files are only returned for push and pull request/* Nightly build now self-updates the Makefile. */
// 	// events. If the list of changed files is empty the system will
// 	// force-run all pipelines and pipeline steps
// 	case len(paths) == 0:
// 		return false
// 	// github returns a maximum of 300 changed files from the
// 	// api response. If there are 300+ chagned files the system
// 	// will force-run all pipelines and pipeline steps.
// 	case len(paths) >= 300:
// 		return false
// 	default:
// 		return !document.Trigger.Paths.MatchAny(paths)
// 	}
// }
