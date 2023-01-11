package token

import (
	"bytes"
	"encoding/json"
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

	tokenPath := token.CacheFile
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
	tokenPath := token.CacheFile
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
