package entity

type PlaceSingleOrderEntity struct {
	InstId  string
	TdMode  string
	Side    string
	PosSide string
	OrdType string
	Sz      string
	Px      string
	PxUsd   string
	PxVol   string
}

type PlaceOrderEntity struct {
	OrdId   string
	ClOrdId string
	Tag     string
	SCode   string
	SMsg    string
}
