package api

import (
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
)

// Handlers holds all the handlers
type Handlers struct {
	AuthHandler         *AuthHandler
	OrganizationHandler *OrganizationHandler
	AccountHandler      *AccountHandler
}

// NewHandlers creates a new Handlers struct
func NewHandlers(authService *service.AuthService, orgService *service.OrganizationService, accountService *service.AccountService, enforcer *casbin.Enforcer) *Handlers {
	return &Handlers{
		AuthHandler:         NewAuthHandler(authService, enforcer),
		OrganizationHandler: NewOrganizationHandler(orgService, enforcer),
		AccountHandler:      NewAccountHandler(*accountService, enforcer),
	}
}
