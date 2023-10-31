package request

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string  `json:"positionSide"`
	Side          string  `json:"side"`
	EntryQuantity float64 `json:"entryQuantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int     `json:"leverageLevel"`
}
