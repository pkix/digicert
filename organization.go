package digicert

import "encoding/json"

// NewOrganizationRequest represents that creating new organization in CertCentral.
type NewOrganizationRequest struct {
	Address     string `json:"address"`
	Address2    string `json:"address2"`
	AssumedName string `json:"assumed_name"`
	City        string `json:"city"`
	// Container   struct {
	// 	ID int `json:"id"`
	// } `json:"container"`
	Country             string  `json:"country"`
	Name                string  `json:"name"`
	OrganizationContact Contact `json:"organization_contact"`
	State               string  `json:"state"`
	Telephone           int     `json:"telephone"`
	Zip                 int     `json:"zip"`
}

// Contact exports contact stuct to new organization
type Contact struct {
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	JobTitle           string `json:"job_title"`
	LastName           string `json:"last_name"`
	Telephone          int    `json:"telephone"`
	TelephoneExtension int    `json:"telephone_extension"`
}

// ViewOrganizationValidationResponse exports organization valiation status
type ViewOrganizationValidationResponse struct {
	Validations []struct {
		DateCreated    string `json:"date_created,omitempty"`
		Description    string `json:"description,omitempty"`
		Name           string `json:"name,omitempty"`
		Status         string `json:"status,omitempty"`
		Type           string `json:"type,omitempty"`
		ValidatedUntil string `json:"validated_until,omitempty"`
	} `json:"validations,omitempty"`

	SchemeValidationErrors
}

// ViewOrganizationDetails represents the organization details
type ViewOrganizationDetails struct {
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	AssumedName string `json:"assumed_name,omitempty"`
	City        string `json:"city,omitempty"`
	Container   struct {
		ID       int    `json:"id,omitempty"`
		IsActive bool   `json:"is_active,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"container,omitempty"`
	Country     string `json:"country,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	EvApprovers []struct {
		FirstName string `json:"first_name,omitempty"`
		ID        int    `json:"id,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	} `json:"ev_approvers,omitempty"`
	ID                  int    `json:"id,omitempty"`
	IsActive            bool   `json:"is_active,omitempty"`
	Name                string `json:"name,omitempty"`
	OrganizationContact struct {
		Email              string `json:"email,omitempty"`
		FirstName          string `json:"first_name,omitempty"`
		ID                 int    `json:"id,omitempty"`
		LastName           string `json:"last_name,omitempty"`
		Telephone          string `json:"telephone,omitempty"`
		TelephoneExtension string `json:"telephone_extension,omitempty"`
	} `json:"organization_contact,omitempty"`
	State     string `json:"state,omitempty"`
	Status    string `json:"status,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Zip       string `json:"zip,omitempty"`

	SchemeValidationErrors
}

// AllOrganizationsResponse exports all organizations
type AllOrganizationsResponse struct {
	Organizations []struct {
		Address   string `json:"address,omitempty"`
		City      string `json:"city,omitempty"`
		Container struct {
			ID       int    `json:"id,omitempty"`
			IsActive bool   `json:"is_active,omitempty"`
			Name     string `json:"name,omitempty"`
			ParentID int    `json:"parent_id,omitempty"`
		} `json:"container,omitempty"`
		Country     string `json:"country,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
		ID          int    `json:"id,omitempty"`
		IsActive    bool   `json:"is_active,omitempty"`
		Name        string `json:"name,omitempty"`
		State       string `json:"state,omitempty"`
		Status      string `json:"status,omitempty"`
		Telephone   string `json:"telephone,omitempty"`
		Zip         string `json:"zip,omitempty"`
	} `json:"organizations,omitempty"`
	Page struct {
		Limit  int `json:"limit,omitempty"`
		Offset int `json:"offset,omitempty"`
		Total  int `json:"total,omitempty"`
	} `json:"page,omitempty"`

	SchemeValidationErrors
}

// ValidateOrganizationRequest exports that request validation for organization
type ValidateOrganizationRequest struct {
	Validations []struct {
		Type          string `json:"type"`
		VerifiedUsers []struct {
			ID int `json:"id"`
		} `json:"verified_users,omitempty"`
	} `json:"validations"`
}

// ViewOrganization exports Use this endpoint to view information about an organization.
func (c *Client) ViewOrganization(orgID string) (*ViewOrganizationDetails, error) {
	c.result = new(ViewOrganizationDetails)
	data, err := c.apiconnect("GET", "/organization/"+orgID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewOrganizationDetails), err
}

// ListAllOrganizations exports Use this endpoint to retrieve a list of organizations.
func (c *Client) ListAllOrganizations() (*AllOrganizationsResponse, error) {
	c.result = new(AllOrganizationsResponse)
	data, err := c.apiconnect("GET", "/organization/", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*AllOrganizationsResponse), err
}

// NewOrganization exports Use this endpoint to create a new organization. The organization information will be used by DigiCert for validation and may appear on certificates.
func (c *Client) NewOrganization(request *NewOrganizationRequest) (*ViewOrganizationDetails, error) {
	c.result = new(ViewOrganizationDetails)
	c.request = request
	data, err := c.apiconnect("POST", "/organization/", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewOrganizationDetails), err
}

// ViewOrganizationValidation exports Use this endpoint to obtain validation statuses for an organization.
func (c *Client) ViewOrganizationValidation(orgID string) (*ViewOrganizationValidationResponse, error) {
	c.result = new(ViewOrganizationValidationResponse)
	data, err := c.apiconnect("GET", "/organization/"+orgID+"/validation", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewOrganizationValidationResponse), err
}

// ValidateOrganization exports Use this endpoint to submit an organization to DigiCert for the specified validations.
func (c *Client) ValidateOrganization(orgID string, request *ValidateOrganizationRequest) bool {
	c.result = new(ViewOrganizationValidationResponse)
	c.request = request
	_, err := c.apiconnect("POST", "/organization/"+orgID+"/validation", nil)
	if err != nil {
		return false
	}

	if c.statusCode == 204 {
		return true
	}
	return false
}
