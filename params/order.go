package params

type RequestCreateOrder struct {
	CustomerName string        `json:"customer_name"`
	Items        []RequestItem `json:"items"`
}

type RequestItem struct {
	ItemID      uint   `json:"item_id,omitempty"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}
