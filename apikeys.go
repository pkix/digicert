package digicert

import (
	"encoding/json"
	"log"
	"time"
)

// APIKeyStatus present API Key status request
type APIKeyStatus struct {
	Status string `json:"status"`
}

//ListAPIKeysResponse represents an api keys details
type ListAPIKeysResponse struct {
	APIKeys []struct {
		CreateDate   string `json:"create_date"`
		ID           int    `json:"id"`
		LastUsedDate string `json:"last_used_date"`
		Name         string `json:"name"`
		Status       string `json:"status"`
		User         struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
		} `json:"user"`
	} `json:"api_keys"`
	SchemeValidationErrors
}

// NewAPIKeyRequest presents a new api key request
type NewAPIKeyRequest struct {
	Name string `json:"name"`
}

// NewAPIKeyResponse presents a new api key response
type NewAPIKeyResponse struct {
	ID     int    `json:"id"`
	APIKey string `json:"api_key"`
	SchemeValidationErrors
}

// ViewAPIKeyResponse presents a api key information
type ViewAPIKeyResponse struct {
	ID   int `json:"id"`
	User struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user"`
	Status       string    `json:"status"`
	CreateDate   time.Time `json:"create_date"`
	LastUsedDate time.Time `json:"last_used_date"`
	Name         string    `json:"name"`
	SchemeValidationErrors
}

// NewAPIKey create a new API Key for the specified user. The name parameter is a convenient identifier that you can use to help remember why the key was issued. The response will contain a unique id that can be used to manage the key. For security, the API key will only be shown this one time in the creation response and after that it will never be shown again. There is no way to retrieve it afterward.
func (c *Client) NewAPIKey(userID, keyName string) (*NewAPIKeyResponse, error) {
	c.request = &NewAPIKeyRequest{
		Name: keyName,
	}
	c.result = new(NewAPIKeyResponse)
	data, err := c.apiconnect("POST", "/key/user/"+userID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*NewAPIKeyResponse), err
}

// ListAPIKeys exports to retrieve a list of API Keys.
func (c *Client) ListAPIKeys() (*ListAPIKeysResponse, error) {
	c.result = new(ListAPIKeysResponse)
	data, err := c.apiconnect("GET", "/key/", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListAPIKeysResponse), err
}

// UpdateAPIKeyStatus update api key status, status only can be set to active or revoked.
func (c *Client) UpdateAPIKeyStatus(apiKeyID, status string) bool {
	c.request = &APIKeyStatus{
		Status: status,
	}
	_, err := c.apiconnect("PUT", "/key/"+apiKeyID+"/status", nil)
	if err != nil {
		log.Println("err", err)
		return false
	}

	// log.Println("status_code", c.statusCode)
	if c.statusCode == 204 {
		return true
	}
	return false
}

// ViewAPIKey exports to view information about the specified API Key. Note that the API Key itself will not be returned. For security, it is only ever returned one time during the initial key creation.
func (c *Client) ViewAPIKey(keyID string) (*ViewAPIKeyResponse, error) {
	c.result = new(ViewAPIKeyResponse)
	data, err := c.apiconnect("GET", "/key/"+keyID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewAPIKeyResponse), err
}
