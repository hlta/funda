package seed

import (
	"funda/internal/store"
	"funda/internal/utils"
	"log"

	"gorm.io/gorm"
)

func SeedAccountsForOrg(db *gorm.DB, orgID uint) error {
	accountRepo := store.NewAccountRepository(db)
	defaultAccounts, err := utils.LoadDefaultAccounts(orgID)
	if err != nil {
		log.Printf("Error loading default accounts: %v", err)
		return err
	}

	for _, account := range defaultAccounts {
		_, err := accountRepo.FindByCodeAndOrg(account.Code, orgID)
		if err != nil && err == gorm.ErrRecordNotFound {
			if err := accountRepo.Create(&account); err != nil {
				log.Printf("Error seeding account %s: %v", account.Name, err)
				return err
			}
		}
	}
	return nil
}
