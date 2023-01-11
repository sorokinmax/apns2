package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type FileToken struct {
	AuthKeyBase64 string
	KeyID         string
	TeamID        string
	IssuedAt      int64
	Bearer        string
}

func (token *Token) UpdateCacheFile() error {
	var fileToken FileToken
	fileToken.AuthKeyBase64 = token.AuthKeyBase64
	fileToken.IssuedAt = token.IssuedAt
	fileToken.KeyID = token.KeyID
	fileToken.TeamID = token.TeamID
	fileToken.Bearer = token.Bearer

	tokenPath := fmt.Sprintf("%s/%s.token", token.CacheDir, token.KeyID)
	file, err := os.Create(tokenPath)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(fileToken); err != nil {
		return err
	}
	return nil
}

func (token *Token) ReadCacheFile() error {
	tokenPath := fmt.Sprintf("%s/%s.token", token.CacheDir, token.KeyID)
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return err
	}
	fileToken := &FileToken{}
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&fileToken); err != nil {
		return err
	} else {
		if token.KeyID == fileToken.KeyID && token.TeamID == fileToken.TeamID && token.AuthKeyBase64 == fileToken.AuthKeyBase64 {
			authKey, _ := AuthKeyFromBytes([]byte(token.AuthKeyBase64))
			token.AuthKey = authKey
			token.Bearer = fileToken.Bearer
			token.IssuedAt = fileToken.IssuedAt
		}
		return nil
	}
}
