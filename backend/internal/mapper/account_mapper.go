package mapper

import (
	"funda/internal/model"
	"funda/internal/response"
)

// ToAccountResponse maps a model.Account and YTD to a response.AccountResponse
func ToAccountResponse(account model.Account, ytd float64) response.AccountResponse {
	return response.AccountResponse{
		ID:      account.ID,
		Code:    account.Code,
		Name:    account.Name,
		Type:    account.Type,
		TaxRate: account.TaxRate,
		Balance: account.Balance,
		OrgID:   account.OrgID,
		YTD:     ytd,
	}
}
