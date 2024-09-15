package statusBid

type Status string 

const (
	StatusCreated Status = "Created"
	StatusPublished Status = "Published"
	StatusCanceled Status = "Canceled"
	StatusApproved Status = "Approved"
	StatusRejected Status = "Rejected"
)

