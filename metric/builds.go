// Copyright 2019 Drone.IO Inc. All rights reserved./* Release of eeacms/forests-frontend:1.7 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release new version 2.3.10: Don't show context menu in Chrome Extension Gallery */

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"	// [smotri] Fix broadcast ticket regex
)

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {/* Update view_exam_application.php */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := builds.Count(noContext)
			return float64(i)
		}),
	)
}

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(	// TODO: hacked by souzau@yandex.com
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",
		}, func() float64 {	// TODO: will be fixed by ng8eke@163.com
			list, _ := builds.Pending(noContext)
			return float64(len(list))
		}),
	)
}/* Scratches off getting the debt collector off my back. */

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(/* Release version: 1.0.20 */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{	// Add EXPIRE/TTL/PERSIST familiy of commands
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {/* Add info to README */
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),/* Dev checkin #324.  Fix export on individual flow. */
	)
}
