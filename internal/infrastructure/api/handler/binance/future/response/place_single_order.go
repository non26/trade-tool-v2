package response

type PlaceSignleOrderHandlerResponse struct {
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
}
