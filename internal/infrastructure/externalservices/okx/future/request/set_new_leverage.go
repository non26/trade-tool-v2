package request

type SetNewLeverageOKXServiceRequest struct {
	InstId  string `json:"instId"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
}
