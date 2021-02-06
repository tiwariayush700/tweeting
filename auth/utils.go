package auth

import "encoding/json"

func GetDataFromTokenClaims(claims map[string]interface{}, val interface{}) error {
	claimsBytes, _ := json.Marshal(claims)

	err := json.Unmarshal(claimsBytes, val)
	if err != nil {
		return err
	}

	return err
}
