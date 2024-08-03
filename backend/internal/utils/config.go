package utils

import (
	"encoding/json"
	"funda/internal/model"
	"io/ioutil"
)

func LoadDefaultAccounts(orgID uint) ([]model.Account, error) {
	var accounts []model.Account
	data, err := ioutil.ReadFile("config/default_accounts.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, err
	}

	for i := range accounts {
		accounts[i].OrgID = orgID
	}

	return accounts, nil
}
