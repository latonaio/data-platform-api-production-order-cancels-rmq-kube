package requests

type Item struct {
	ProductionOrder    		int     `json:"ProductionOrder"`
	ProductionOrderItem		int     `json:"ProductionOrderItem"`
	IsCancelled        		*bool   `json:"IsCancelled"`
}
