package request

import entity "tradetoolv2/internal/app/domain/entity/okx"

type PlaceASinglePositionOKXServiceRequest struct {
	InstId  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	Side    string `json:"side"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	Sz      string `json:"sz"`
	Px      string `json:"px"`
	PxUsd   string `json:"pxUsd"`
	PxVol   string `json:"pxVol"`
	TgtCcy  string `json:"tgtCcy"`
}

func (p *PlaceASinglePositionOKXServiceRequest) ToPlaceSingleOrderRequest(e *entity.PlaceSingleOrderEntity) {
	p.InstId = e.InstId
	p.TdMode = e.TdMode
	p.Side = e.Side
	p.PosSide = e.PosSide
	p.OrdType = e.OrdType
	p.Sz = e.Sz
	p.Px = e.Px
	p.PxUsd = e.PxUsd
	p.PxVol = e.PxVol
	p.TgtCcy = e.TgtCcy
}
