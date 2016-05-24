package common

// job 请求

const (
	EXECUTING = "EXECUTING"
	STOP ="STOP"
	TEST = "TEST"
)
type JobRequest struct {


	JobSnapshot int `json:"jobSnapshot"`
	Params string	`json:"params"`
	Status string	`json:"status"`



}