package settings

import (
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

// Check if the vault path is valid
func IsVaultPathValid(vaultPath string) bool {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		log.WithField("level", "FATAL").Errorf("Vault path '%s' is invalid", vaultPath)
		return false
	}

	return true
}

// Check if the vault repository is valid
func IsVaultRepositoryValid(vaultRepository string) bool {
	parsed, err := url.ParseRequestURI(vaultRepository)
	if err != nil {
		log.WithField("level", "FATAL").Errorf("Vault repository '%s' is invalid", vaultRepository)
		return false
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		log.WithField("level", "FATAL").Errorf("Vault repository '%s' is invalid", vaultRepository)
		return false
	}

	return true
}
