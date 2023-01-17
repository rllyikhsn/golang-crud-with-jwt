package presenter

type (
	Status  string
	Message string
	Data    interface{}
)

var (
	BadRequestSeason     Status = "failed"
	SuccessRequestSeason Status = "success"
)
