package constants

import "github.com/tiwariayush700/tweeting/models"

const File = "file"
const DefaultConfig = "local.json"
const FileUsage = "File to read all configuration"

const ContextKeyUserId = "userID"
const ContextKeyRole = "role"

const (
	Authorize = "Authorize"
	Reject    = "Reject"
)

const (
	ActionStatusPending  = models.ActionStatus("pending")
	ActionStatusApproved = models.ActionStatus("approved")
	ActionStatusRejected = models.ActionStatus("rejected")
)
