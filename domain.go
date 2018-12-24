package digicert

import (
	"encoding/json"
	"time"
)

// NewDomainRequest represents a request of creating a new domain
type NewDomainRequest struct {
	Name         string `json:"name"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	Validations []struct {
		Type string `json:"type"`
		User struct {
			ID int `json:"id"`
		} `json:"user,omitempty"`
	} `json:"validations"`
	Dcv struct {
		Method string `json:"method"`
	} `json:"dcv"`
}

// NewDomainResponse presents a response of creating a new domain
type NewDomainResponse struct {
	ID int `json:"id"`

	SchemeValidationErrors
}

// ViewADomainResponse presents a domain detail, with dcv and validation
type ViewADomainResponse struct {
	ID           int       `json:"id"`
	IsActive     bool      `json:"is_active"`
	Status       string    `json:"status"`
	Name         string    `json:"name"`
	DateCreated  time.Time `json:"date_created"`
	Organization struct {
		ID          int    `json:"id"`
		Status      string `json:"status"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		IsActive    string `json:"is_active"`
	} `json:"organization"`
	Validations []struct {
		Type           string    `json:"type"`
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		DateCreated    time.Time `json:"date_created"`
		ValidatedUntil time.Time `json:"validated_until"`
		Status         string    `json:"status"`
		DcvStatus      string    `json:"dcv_status"`
		OrgStatus      string    `json:"org_status"`
		VerifiedUsers  []struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			JobTitle  string `json:"job_title"`
			Telephone string `json:"telephone"`
		} `json:"verified_users,omitempty"`
	} `json:"validations"`
	Container struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"container"`

	SchemeValidationErrors
}

// ListDoaminsResponse presents all domains
type ListDoaminsResponse struct {
	Domains []struct {
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		DateCreated  time.Time `json:"date_created"`
		Organization struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			AssumedName string `json:"assumed_name"`
			DisplayName string `json:"display_name"`
		} `json:"organization"`
		Validations []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
		} `json:"validations,omitempty"`
		Container struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"container"`
	} `json:"domains"`

	SchemeValidationErrors
}

// ValidationTypesResponse presents domain validation types
type ValidationTypesResponse struct {
	ValidationTypes []struct {
		Type         string `json:"type"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		RequiresUser bool   `json:"requires_user"`
		RequiresDcv  bool   `json:"requires_dcv"`
	} `json:"validation_types"`

	SchemeValidationErrors
}

// ValidationRequest presents OV/EV validations
type ValidationRequest struct {
	Validations []struct {
		Type string `json:"type"`
		User struct {
			ID int `json:"id"`
		} `json:"user,omitempty"`
	} `json:"validations"`
}

// ViewValidationResponse presents a domain's validation detail
type ViewValidationResponse struct {
	Validations []struct {
		Type           string    `json:"type"`
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		DateCreated    time.Time `json:"date_created,omitempty"`
		ValidatedUntil time.Time `json:"validated_until,omitempty"`
		Status         string    `json:"status"`
		DcvStatus      string    `json:"dcv_status"`
		VerifiedUsers  []struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"verified_users,omitempty"`
	} `json:"validations"`

	SchemeValidationErrors
}

// ListDomainControlMethodsResponse presents a list of domain control methods
type ListDomainControlMethodsResponse struct {
	Methods []struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"methods"`

	SchemeValidationErrors
}

// DomainControlMethodRequest presents a request of domain control
type DomainControlMethodRequest struct {
	Method string `json:"method"`
}

// DomainControlEmailsResponse presents a domain control emails.
type DomainControlEmailsResponse struct {
	NameScope   string   `json:"name_scope"`
	BaseEmails  []string `json:"base_emails"`
	WhoisEmails []string `json:"whois_emails"`

	SchemeValidationErrors
}

// ResendDCVEmailReqeust presents DCV email of a domain
type ResendDCVEmailReqeust struct {
	NameScope string `json:"name_scope"`
}

// EmailApprove presents submit email approve request.
type EmailApprove struct {
	Method         string `json:"method"`
	NameScope      string `json:"name_scope"`
	DcvInvitations []struct {
		InvitationID int `json:"invitation_id"`
	} `json:"dcv_invitations"`
}

// DNSApprove presents submit dns approve request
type DNSApprove struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

// ApproveStatuesResponse presents a status of approval process
type ApproveStatuesResponse struct {
	Status string `json:"status"`

	SchemeValidationErrors
}

// NewDomain exports to add a domain for an organization in a container. You also must specify at least one validation type for the domain.
func (c *Client) NewDomain(request *NewDomainRequest) (*NewDomainResponse, error) {
	c.result = new(NewDomainResponse)
	c.request = request
	data, err := c.makeRequest("POST", "/domain", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*NewDomainResponse), err
}

// ActiveDomain exports to activate a domain that was previously deactivated.
func (c *Client) ActiveDomain(id string) (bool, error) {
	_, err := c.makeRequest("PUT", "/domain/"+id+"/activate", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// DeactiveDomain exports to deactivate a domain.
func (c *Client) DeactiveDomain(id string) (bool, error) {
	_, err := c.makeRequest("PUT", "/domain/"+id+"/deactivate", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// ViewADomain exports to view a domain detail.
func (c *Client) ViewADomain(id string) (*ViewADomainResponse, error) {
	c.result = new(ViewADomainResponse)
	data, err := c.makeRequest("GET", "/domain/"+id+"?include_dcv=true&include_validation=true", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewADomainResponse), err
}

// ListDomains exports to retrieve a list of domains.
func (c *Client) ListDomains(containerID string) (*ListDoaminsResponse, error) {
	c.result = new(ListDoaminsResponse)
	data, err := c.makeRequest("GET", "/domain?container_id="+containerID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListDoaminsResponse), err
}

// ListValidationTypes exports to retrieve a list of validation types available for domains.
func (c *Client) ListValidationTypes() (*ValidationTypesResponse, error) {
	c.result = new(ValidationTypesResponse)
	data, err := c.makeRequest("GET", "/domain/validation-type", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ValidationTypesResponse), err
}

// SubmitValidation exports to submit an existing domain for validation for one or more validation types, or to resubmit a domain for validation that has expired.
func (c *Client) SubmitValidation(domainID string, request *ValidationRequest) (bool, error) {
	c.request = request
	_, err := c.makeRequest("POST", "/domain/"+domainID+"/validation", nil)
	if err != nil {
		return false, err
	}
	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// ViewValidaton exports Use this domain to get a list of the validation types for which a domain has been submitted.
func (c *Client) ViewValidaton(domainID string) (*ViewValidationResponse, error) {
	c.result = new(ViewValidationResponse)
	data, err := c.makeRequest("GET", "/domain/"+domainID+"validation", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewValidationResponse), err
}

// ListDomainControlMethods exports Use this endpoint to retrieve a list of domain control validation (DCV) methods types available for domains.
func (c *Client) ListDomainControlMethods() (*ListDomainControlMethodsResponse, error) {
	c.result = new(ListDomainControlMethodsResponse)
	data, err := c.makeRequest("GET", "/domain/dcv/method", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListDomainControlMethodsResponse), err
}

// ChangeDomainControlMethod exports Use this endpoint to set the Domain Control Validation (DCV) method for the domain.
func (c *Client) ChangeDomainControlMethod(domainID string, request *DomainControlMethodRequest) (bool, error) {
	c.request = request
	_, err := c.makeRequest("POST", "/domain/"+domainID+"/dcv/method", nil)
	if err != nil {
		return false, err
	}
	if c.statusCode == 200 {
		return true, err
	}
	return false, err
}

// GetDomainControlEmails exports Use this endpoint to retrieve domain email addresses for Domain Control Validation (DCV).
func (c *Client) GetDomainControlEmails(domainID string) (*DomainControlEmailsResponse, error) {
	c.result = new(DomainControlEmailsResponse)
	data, err := c.makeRequest("GET", "/domain/"+domainID+"/dcv/emails", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*DomainControlEmailsResponse), err
}

// ResendDCVEmail exports Use this endpoint to resend emails for Domain Control Validation (DCV).
func (c *Client) ResendDCVEmail(domainID string, request *ResendDCVEmailReqeust) (bool, error) {
	c.request = request
	_, err := c.makeRequest("POST", "/domain/"+domainID+"/dcv/emails", nil)
	if err != nil {
		return false, err
	}
	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// ApproveEmail exports Use this endpoint to submit the Domain Control Validation (DCV) approval
func (c *Client) ApproveEmail(domainID string, request *EmailApprove) (*ApproveStatuesResponse, error) {
	c.result = new(ApproveStatuesResponse)
	c.request = request
	data, err := c.makeRequest("POST", "/domain/"+domainID+"/dcv", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ApproveStatuesResponse), err
}

// ApproveDNS exports Use this endpoint to submit the Domain Control Validation (DCV) approval
func (c *Client) ApproveDNS(domainID string, request *DNSApprove) (*ApproveStatuesResponse, error) {
	c.result = new(ApproveStatuesResponse)
	c.request = request
	data, err := c.makeRequest("POST", "/domain/"+domainID+"/dcv/cname", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ApproveStatuesResponse), err
}

// ValidateToken exports Use this endpoint to submit an email token for Domain Control Validation (DCV).
func (c *Client) ValidateToken(token string) (*ApproveStatuesResponse, error) {
	c.result = new(ApproveStatuesResponse)
	data, err := c.makeRequest("PUT", "/domain/dcv/email/token/"+token, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ApproveStatuesResponse), err
}
