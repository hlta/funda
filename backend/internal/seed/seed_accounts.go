package seed

import (
	"funda/internal/logger"
	"funda/internal/store"
	"funda/internal/utils"

	"gorm.io/gorm"
)

// SeedAccountsForOrg seeds default accounts for the specified organization ID.
func SeedAccountsForOrg(db *gorm.DB, orgID uint, log logger.Logger) error {
	accountRepo := store.NewAccountRepository(db)
	defaultAccounts, err := utils.LoadDefaultAccounts(orgID)
	if err != nil {
		log.WithField("orgID", orgID).Error("Error loading default accounts: ", err)
		return err
	}

	for _, account := range defaultAccounts {
		_, err := accountRepo.FindByCodeAndOrg(account.Code, orgID)
		if err != nil && err == gorm.ErrRecordNotFound {
			if err := accountRepo.Create(&account); err != nil {
				log.WithField("account", account.Name).Error("Error seeding account: ", err)
				return err
			}
		}
	}
	return nil
}
