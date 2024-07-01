package constants

const (
	// Auth Routes
	SignupRoute               = "/signup"
	LoginRoute                = "/login"
	LogoutRoute               = "/logout"
	CheckAuthRoute            = "/auth/check"
	GetUserOrganizationsRoute = "/auth/orgs"
	SwitchOrgRoute            = "/auth/switch-org"

	// Organization Routes
	CreateOrganizationRoute = "/organizations"
	GetOrganizationRoute    = "/organizations/:id"
	UpdateOrganizationRoute = "/organizations/:id"
)
