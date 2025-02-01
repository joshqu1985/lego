package pagetoken

import (
	"encoding/base64"
	"encoding/json"
)

type PageToken struct {
	Mark int `json:"mark"`
	Size int `json:"size"`
}

func Encode(ptoken *PageToken) string {
	bytes, _ := json.Marshal(ptoken)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Decode(s string, defaultSize int) (PageToken, error) {
	if len(s) == 0 {
		return PageToken{Size: defaultSize}, nil
	}

	var token PageToken
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return token, err
	}
	return token, json.Unmarshal(bytes, &token)
}
