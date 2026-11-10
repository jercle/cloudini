package lib

import "os"

const (
	// SecretFilePerm is used for files holding credentials, tokens or keys.
	SecretFilePerm os.FileMode = 0o600
	// SecretDirPerm is used for the config and cache directories.
	SecretDirPerm os.FileMode = 0o700
	// OutputFilePerm is used for non-sensitive generated output.
	OutputFilePerm os.FileMode = 0o644
	// OutputDirPerm is used for non-sensitive output directories.
	OutputDirPerm os.FileMode = 0o755
)

// WriteSecretFile writes data to path readable by the current user only. The
// mode is re-applied afterwards because os.WriteFile keeps the existing mode of
// a file that is already present.
func WriteSecretFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, SecretFilePerm); err != nil {
		return err
	}
	return os.Chmod(path, SecretFilePerm)
}

// ensureSecretDir creates dir if needed and removes group/other access from it.
func ensureSecretDir(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return os.MkdirAll(dir, SecretDirPerm)
	}
	if info.Mode().Perm()&0o077 != 0 {
		return os.Chmod(dir, SecretDirPerm)
	}
	return nil
}
