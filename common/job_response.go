package common

// job response wapper
type JobResponse struct {

	Success bool `json:"success"`
	Message string `json:"message"`
	Content string `json:"content"`


}
