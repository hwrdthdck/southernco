// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

// import (
// 	conventions "go.opentelemetry.io/collector/semconv/v1.27.0"
// )

// model.go contains specific attributes from the 1.28 and 1.29 releases of
// SemConv. They are manually added due to issue
// https://github.com/open-telemetry/weaver/issues/227 which will migrate code
// gen to weaver. Once that is done, these attributes will be migrated to the
// semantic conventions package.

const (
	// vcs.change.state with enum values of open, closed, or merged.
	AttributeVCSChangeState       = "vcs.change.state"
	AttributeVCSChangeStateOpen   = "open"
	AttributeVCSChangeStateClosed = "closed"
	AttributeVCSChangeStateMerged = "merged"

	// vcs.change.title
	AttributeVCSChangeTitle = "vcs.change.title"

	// vcs.change.id
	AttributeVCSChangeID = "vcs.change.id"

	// vcs.revision_delta.direction with enum values of behind or ahead.
	AttributeVCSRevisionDeltaDirection       = "vcs.revision_delta.direction"
	AttributeVCSRevisionDeltaDirectionBehind = "behind"
	AttributeVCSRevisionDeltaDirectionAhead  = "ahead"

	// vcs.line_change.type with enum values of added or removed.
	AttributeVCSLineChangeType        = "vcs.line_change.type"
	AttributeVCSLineChangeTypeAdded   = "added"
	AttributeVCSLineChangeTypeRemoved = "removed"

	// vcs.ref.type with enum values of branch or tag.
	AttributeVCSRefType       = "vcs.ref.type"
	AttributeVCSRefTypeBranch = "branch"
	AttributeVCSRefTypeTag    = "tag"

	// vcs.repository.name
	AttributeVCSRepositoryName = "vcs.repository.name"

	// vcs.ref.base.name
	AttributeVCSRefBase = "vcs.ref.base"

	// vcs.ref.base.revision
	AttributeVCSRefBaseRevision = "vcs.ref.base.revision"

	//vcs.ref.base.type with enum values of branch or tag.
	AttributeVCSRefBaseType       = "vcs.ref.base.type"
	AttributeVCSRefBaseTypeBranch = "branch"
	AttributeVCSRefBaseTypeTag    = "tag"

	// vcs.ref.head.name
	AttributeVCSRefHead = "vcs.ref.head"

	// vcs.ref.head.revision
	AttributeVCSRefHeadRevision = "vcs.ref.head.revision"

	// vcs.ref.head.type with enum values of branch or tag.
	AttributeVCSRefHeadType       = "vcs.ref.head.type"
	AttributeVCSRefHeadTypeBranch = "branch"
	AttributeVCSRefHeadTypeTag    = "tag"

	// The following prototype attributes that do not exist yet in semconv.
	// They are highly experimental and subject to change.

	AttributeCICDPipelineRunURLFull            = "cicd.pipeline.run.url.full" // equivalent to GitHub's `html_url`

    // These are being added in https://github.com/open-telemetry/semantic-conventions/pull/1681
	AttributeCICDPipelineRunStatus             = "cicd.pipeline.run.status"   // equivalent to GitHub's `conclusion`
	AttributeCICDPipelineRunStatusSuccess      = "success"
	AttributeCICDPipelineRunStatusFailure      = "failure"
	AttributeCICDPipelineRunStatusCancellation = "cancellation"
	AttributeCICDPipelineRunStatusError        = "error"
	AttributeCICDPipelineRunStatusSkip         = "skip"

	AttributeCICDPipelineTaskRunStatus             = "cicd.pipeline.run.task.status"   // equivalent to GitHub's `conclusion`
	AttributeCICDPipelineTaskRunStatusSuccess      = "success"
	AttributeCICDPipelineTaskRunStatusFailure      = "failure"
	AttributeCICDPipelineTaskRunStatusCancellation = "cancellation"
	AttributeCICDPipelineTaskRunStatusError        = "error"
	AttributeCICDPipelineTaskRunStatusSkip         = "skip"

    // TODO: Evaluate these
    AttributeCICDPipelineRunSenderLogin = "cicd.pipeline.run.sender.login" // GitHub's Run Sender Login
    AttributeCICDPipelineTaskRunSenderLogin = "cicd.pipeline.task.run.sender.login" // GitHub's Task Sender Login
    AttributeVCSVendorName = "vcs.vendor.name" // GitHub
    AttributeVCSRepositoryOwner = "vcs.repository.owner" // GitHub's Owner Login
    
    AttributeCICDPipelineFilePath = "cicd.pipeline.file.path" // GitHub's Path in workflow_run
    // path
    // previous attempt
    // referenced workflows
    // author email name & committer, message
    // pr url
    // status vs conclusion
    // associated changes (prs)
    // run attempt
    // installation (for the github app)

    


)
