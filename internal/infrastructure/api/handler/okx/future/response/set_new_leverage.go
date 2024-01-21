package response

type SetNewLeverageHandlerResponse struct {
	Level   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
