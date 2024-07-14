package api

import (
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
)

// Handlers holds all the handlers
type Handlers struct {
	AuthHandler         *AuthHandler
	OrganizationHandler *OrganizationHandler
}

// NewHandlers creates a new Handlers struct
func NewHandlers(authService *service.AuthService, orgService *service.OrganizationService, enforcer *casbin.Enforcer) *Handlers {
	return &Handlers{
		AuthHandler:         NewAuthHandler(authService, enforcer),
		OrganizationHandler: NewOrganizationHandler(orgService, enforcer),
	}
}
