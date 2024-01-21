package request

import entity "tradetoolv2/internal/app/domain/entity/okx"

type SetNewLeverageOKXServiceRequest struct {
	InstId  string `json:"instId"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
}

func (s *SetNewLeverageOKXServiceRequest) ToSetNewLeverageRequest(e *entity.SetLeverageFuture) {
	s.InstId = e.InstId
	s.Lever = e.Lever
	s.MgnMode = e.MgnMode
}
