package utils

import (
	"strconv"

	"github.com/casbin/casbin/v2"
)

func AddPredefinedRolesAndPermissions(enforcer *casbin.Enforcer, orgID uint) error {
	org := strconv.FormatUint(uint64(orgID), 10)

	// Predefined roles and permissions
	predefinedPolicies := [][]string{
		{"admin", org, "*", "*"},
		{"standard_user", org, "invoices", "view"},
		{"standard_user", org, "invoices", "create"},
		{"standard_user", org, "invoices", "edit"},
		{"standard_user", org, "invoices", "delete"},
		{"invoice_manager", org, "invoices", "view"},
		{"invoice_manager", org, "invoices", "create"},
		{"invoice_manager", org, "invoices", "edit"},
		{"invoice_manager", org, "invoices", "delete"},
		{"bank_reconciler", org, "bank_transactions", "view"},
		{"bank_reconciler", org, "bank_transactions", "reconcile"},
		{"payroll_manager", org, "payroll", "view"},
		{"payroll_manager", org, "payroll", "process"},
		{"reports_viewer", org, "reports", "view"},
		{"reports_viewer", org, "reports", "generate"},
		{"admin", org, "users", "manage"},
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
