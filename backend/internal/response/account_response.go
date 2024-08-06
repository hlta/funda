package response

type AccountResponse struct {
	ID      uint    `json:"id"`
	Code    int     `json:"code"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	TaxRate string  `json:"tax_rate"`
	Balance float64 `json:"balance"`
	OrgID   uint    `json:"org_id"`
	OrgName string  `json:"org_name"`
	YTD     float64 `json:"ytd"` // Year to Date balance
}
