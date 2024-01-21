package request

import entity "tradetoolv2/internal/app/domain/entity/okx"

type SetNewLeverageHandlerRequest struct {
	InstId  string `json:"instId"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide"`
}

func (s *SetNewLeverageHandlerRequest) ToSetLeverageEntity() *entity.SetLeverageFuture {
	e := entity.SetLeverageFuture{
		InstId:  s.InstId,
		Lever:   s.Lever,
		MgnMode: s.MgnMode,
		PosSide: s.PosSide,
	}
	return &e
}
