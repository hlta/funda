package utils

import (
	"encoding/json"
	"funda/internal/model"
	"os"
	"path/filepath"
)

func LoadDefaultAccounts(orgID uint) ([]model.Account, error) {
	var accounts []model.Account
	configPath := filepath.Join("configs", "default_accounts.json")
	data, err := os.ReadFile(configPath)
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
