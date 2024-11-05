package lib

import (
	"encoding/json"
	"os"
	"path"
)

type AuthentificationData struct {
	UserId           string `json:"userId"`
	Seed             string `json:"seed"`
	Counter          int    `json:"counter"`
	NSSE             string `json:"nsse"`
	UserPublicId     string `json:"userPublicId"`
	QrCodePrivateKey string `json:"privateKey"`
}

const persistenceFileName = "_auth_.json"

func PersistAuthData(data *AuthentificationData) error {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	configPath, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(configPath, persistenceFileName), b, 0644)
}

func RetrieveAuthData() (*AuthentificationData, error) {
	configPath, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(path.Join(configPath, persistenceFileName))
	if err != nil {
		return nil, err
	}

	var data AuthentificationData

	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func IsAuthDataPersisted() bool {
	configPath, err := os.UserCacheDir()
	if err != nil {
		return false
	}

	_, err = os.Stat(path.Join(configPath, persistenceFileName))
	return err == nil
}

func DeletePersistedAuthData() error {
	configPath, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	return os.Remove(path.Join(configPath, persistenceFileName))
}
