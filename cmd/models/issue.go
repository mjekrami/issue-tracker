package models

type (
	IssueCreate struct {
		Name        string `json:"name" required:"true"`
		AssignedTo  string `json:"assignedto" required:"true"`
		Description string `json:"description"`
	}
	IssueUpdate struct {
		ID string `json:"id" required:"true"`
	}
	IssueResponse struct {
		ID        string `json:"id" required:"true"`
		CreatedAt string `json:"created_at" required:"true"`
		CreatedBy string `json:"created_by" required:"true"`
		IssueCreate
	}
)
