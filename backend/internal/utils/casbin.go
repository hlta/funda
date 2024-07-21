package utils

import (
	"github.com/casbin/casbin/v2"
)

func AddPredefinedRolesAndPermissions(enforcer *casbin.Enforcer, orgID uint) error {
	org := UintToString(orgID)

	// Predefined roles and permissions
	predefinedPolicies := [][]string{
		{"admin", org, "*", "*"},
		{"standard_user", org, "/invoices", "GET"},
		{"standard_user", org, "/invoices", "POST"},
		{"standard_user", org, "/invoices/:id", "PUT"},
		{"standard_user", org, "/invoices/:id", "DELETE"},
		{"invoice_manager", org, "/invoices", "GET"},
		{"invoice_manager", org, "/invoices", "POST"},
		{"invoice_manager", org, "/invoices/:id", "PUT"},
		{"invoice_manager", org, "/invoices/:id", "DELETE"},
		{"bank_reconciler", org, "/bank_transactions", "GET"},
		{"bank_reconciler", org, "/bank_transactions/reconcile", "POST"},
		{"payroll_manager", org, "/payroll", "GET"},
		{"payroll_manager", org, "/payroll/process", "POST"},
		{"reports_viewer", org, "/reports", "GET"},
		{"reports_viewer", org, "/reports/generate", "POST"},
	}

	// Add predefined policies to the enforcer
	for _, policy := range predefinedPolicies {
		// Convert []string to []interface{}
		policyInterface := make([]interface{}, len(policy))
		for i, v := range policy {
			policyInterface[i] = v
		}
		if _, err := enforcer.AddPolicy(policyInterface...); err != nil {
			return err
		}
	}

	return nil
}
