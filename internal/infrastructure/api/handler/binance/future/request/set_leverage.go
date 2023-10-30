package request

type SetLeverageHandlerRequest struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}
