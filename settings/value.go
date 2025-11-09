package settings

import (
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
