package settings

import (
	"log"
	"os"
)

// Check if the vault path is valid
func IsVaultPathValid(vaultPath string) bool {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		log.Fatalf("Vault path '%s' is invalid.", vaultPath)
	}

	return true
}
