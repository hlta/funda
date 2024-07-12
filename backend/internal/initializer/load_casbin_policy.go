package initializer

import (
	"bufio"
	"fmt"
	"funda/internal/logger"
	"os"
	"strings"

	"github.com/casbin/casbin/v2"
)

func LoadPoliciesFromCSV(enforcer *casbin.Enforcer, filepath string, logger logger.Logger) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open policy file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comments
		}

		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		if len(parts) < 4 {
			logger.Warn("invalid policy line: ", line)
			continue
		}

		// Convert []string to []interface{}
		policy := make([]interface{}, len(parts))
		for i, v := range parts {
			policy[i] = v
		}

		// Check if the policy already exists before adding it
		exists, _ := enforcer.HasPolicy(policy...)
		if !exists {
			if _, err := enforcer.AddPolicy(policy...); err != nil {
				return fmt.Errorf("failed to add policy: %w", err)
			}
		} else {
			logger.Info("policy already exists: ", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading policy file: %w", err)
	}

	return nil
}
