package requests

type Header struct {
	ProductionOrder            int     `json:"ProductionOrder"`
	IsCancelled          	   *bool   `json:"IsCancelled"`
}
