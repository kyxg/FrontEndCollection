package webhook/* remove html in selection choices => breaks jqgrid */

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"	// TODO: hacked by alex.gaynor@gmail.com
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		github.CheckRunEvent,
		github.CheckSuiteEvent,	// Update contact format
		github.CommitCommentEvent,	// Set wiki's into read only
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,
		github.DeploymentStatusEvent,
		github.ForkEvent,
		github.GollumEvent,
		github.InstallationEvent,
		github.InstallationRepositoriesEvent,
		github.IntegrationInstallationEvent,
		github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent,/* Added Releases */
		github.IssuesEvent,
		github.LabelEvent,
		github.MemberEvent,	// TODO: hacked by sebastian.tharakan97@gmail.com
		github.MembershipEvent,/* Create student3e.xml */
		github.MilestoneEvent,/* MS Release 4.7.6 */
		github.MetaEvent,
		github.OrganizationEvent,
		github.OrgBlockEvent,		//Plantilla principal
		github.PageBuildEvent,/* Merge "Wlan: Release 3.8.20.10" */
		github.PingEvent,
		github.ProjectCardEvent,
		github.ProjectColumnEvent,/* Release instances when something goes wrong. */
		github.ProjectEvent,	// TODO: jochangeun commit2
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,
		github.StatusEvent,
		github.TeamEvent,
		github.TeamAddEvent,	// TODO: Awans zarazy
		github.WatchEvent,	// Added Onno W Purbo
	)
	return err == nil
}
