package constants

const (
	// Auth Messages
	InvalidRequestDetails  = "Invalid request details"
	InvalidCredentials     = "Invalid credentials"
	SignupSuccessful       = "Signup successful"
	LoginSuccessful        = "Login successful"
	LogoutSuccessful       = "Logout successful"
	Authenticated          = "Authenticated"
	NotAuthenticated       = "Not authenticated"
	InvalidToken           = "Invalid token"
	FailedCreateUserAndOrg = "Failed to create user and organization"
	FailedGenerateToken    = "Failed to generate token"

	// Organization Messages
	InvalidOrganizationID           = "Invalid organization ID"
	InvalidOrganizationIDFormat     = "Invalid organization ID format"
	OrganizationNotFound            = "Organization not found"
	AccessDenied                    = "Access denied"
	FailedCreateOrganization        = "Failed to create organization"
	FailedUpdateOrganization        = "Failed to update organization"
	OrganizationsRetrieved          = "Organizations retrieved successfully"
	OrganizationCreatedSuccessfully = "Organization created successfully"
	OrganizationUpdatedSuccessfully = "Organization updated successfully"
	OrganizationSwitched            = "Organization switched successfully"
	FailedRetrieveOrganizations     = "Failed to retrieve organizations"

	// Middleware Messages
	AuthorizationHeaderRequired = "authorization header is required"
	InvalidAuthorizationHeader  = "authorization header format must be 'Bearer {token}'"
	InvalidOrExpiredToken       = "invalid or expired JWT token"
	UserClaimsNotAvailable      = "user claims are not available"
	TokenNotFound               = "token not found in context"

	FailedAssignRole          = "Failed to assign role"
	FailedRetrieveRoles       = "Failed to retrieve role"
	FailedRetrievePermissions = "Failed to retrieve permissions"
)
