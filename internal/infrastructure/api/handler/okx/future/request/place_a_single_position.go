package request

import entity "tradetoolv2/internal/app/domain/entity/okx"

type PlaceASinglePositionHandlerRequest struct {
	InstId  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	Side    string `json:"side"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	Sz      string `json:"sz"`
}

func (p *PlaceASinglePositionHandlerRequest) ToPlaceSingleOrderEntity() *entity.PlaceSingleOrderEntity {
	e := entity.PlaceSingleOrderEntity{
		InstId:  p.InstId,
		TdMode:  p.TdMode,
		Side:    p.Side,
		PosSide: p.PosSide,
		OrdType: p.OrdType,
		Sz:      p.Sz,
	}

	return &e
}
