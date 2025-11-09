package settings

import (
	// "log"
	"net/url"
	"os"
)

// Check if the vault path is valid
func IsVaultPathValid(vaultPath string) bool {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		// log.Fatalf("Vault path '%s' is invalid", vaultPath) // TODO: Esses logs vão ser registrados no arquivo
		return false
	}

	return true
}

// Check if the vault repository is valid
func IsVaultRepositoryValid(vaultRepository string) bool {
	parsed, err := url.ParseRequestURI(vaultRepository)
	if err != nil {
		// log.Fatalf("Vault repository '%s' is invalid", vaultRepository) // TODO: Esses logs vão ser registrados no arquivo
		return false
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		// log.Fatalf("Vault repository '%s' is invalid", vaultRepository) // TODO: Esses logs vão ser registrados no arquivo
		return false
	}

	return true
}
