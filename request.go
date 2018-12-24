package digicert

import (
	"encoding/json"
	"errors"
	"time"
)

// ListRequestsResponse presents a list of request
type ListRequestsResponse struct {
	Requests []struct {
		ID        int       `json:"id"`
		Date      time.Time `json:"date"`
		Type      string    `json:"type"`
		Status    string    `json:"status"`
		Requester struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		} `json:"requester,omitempty"`
		Processor struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		} `json:"processor,omitempty"`
		Order struct {
			ID          int `json:"id"`
			Certificate struct {
				CommonName string `json:"common_name"`
			} `json:"certificate"`
			Organization struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"organization"`
			Container struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"container"`
			Product struct {
				NameID string `json:"name_id"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"product"`
		} `json:"order"`
	} `json:"requests"`

	SchemeValidationErrors
}

// ViewRequestResponse presents view a request detail.
type ViewRequestResponse struct {
	ID            int       `json:"id"`
	Date          time.Time `json:"date"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	DateProcessed time.Time `json:"date_processed"`
	Requester     struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"requester"`
	Processor struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"processor"`
	Order struct {
		ID          int `json:"id"`
		Certificate struct {
			CommonName   string    `json:"common_name"`
			DNSNames     []string  `json:"dns_names"`
			DateCreated  time.Time `json:"date_created"`
			Csr          string    `json:"csr"`
			Organization struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				City    string `json:"city"`
				State   string `json:"state"`
				Country string `json:"country"`
			} `json:"organization"`
			ServerPlatform struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				InstallURL string `json:"install_url"`
				CsrURL     string `json:"csr_url"`
			} `json:"server_platform"`
			SignatureHash string `json:"signature_hash"`
			KeySize       int    `json:"key_size"`
			CaCert        struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"ca_cert"`
		} `json:"certificate"`
		Status       string    `json:"status"`
		IsRenewal    bool      `json:"is_renewal"`
		DateCreated  time.Time `json:"date_created"`
		Organization struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
		} `json:"organization"`
		ValidityYears               int  `json:"validity_years"`
		DisableRenewalNotifications bool `json:"disable_renewal_notifications"`
		AutoRenew                   int  `json:"auto_renew"`
		Container                   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"container"`
		Product struct {
			NameID                string `json:"name_id"`
			Name                  string `json:"name"`
			Type                  string `json:"type"`
			ValidationType        string `json:"validation_type"`
			ValidationName        string `json:"validation_name"`
			ValidationDescription string `json:"validation_description"`
		} `json:"product"`
		OrganizationContact struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			JobTitle  string `json:"job_title"`
			Telephone string `json:"telephone"`
		} `json:"organization_contact"`
		TechnicalContact struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			JobTitle  string `json:"job_title"`
			Telephone string `json:"telephone"`
		} `json:"technical_contact"`
		User struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		} `json:"user"`
		Requests []struct {
			ID       int       `json:"id"`
			Date     time.Time `json:"date"`
			Type     string    `json:"type"`
			Status   string    `json:"status"`
			Comments string    `json:"comments"`
		} `json:"requests"`
		CsProvisioningMethod string `json:"cs_provisioning_method"`
		ShipInfo             struct {
			Name    string `json:"name"`
			Addr1   string `json:"addr1"`
			Addr2   string `json:"addr2"`
			City    string `json:"city"`
			State   string `json:"state"`
			Zip     int    `json:"zip"`
			Country string `json:"country"`
			Method  string `json:"method"`
		} `json:"ship_info"`
		DisableCt bool `json:"disable_ct"`
	} `json:"order"`
	Comments         string `json:"comments"`
	ProcessorComment string `json:"processor_comment"`

	SchemeValidationErrors
}

// UpdateRequestStatusRequest presents update "request" status
type UpdateRequestStatusRequest struct {
	Status           string `json:"status"`
	ProcessorComment string `json:"processor_comment"`
}

// ListRequests exports Use this endpoint to retrieve a list of certificate requests.
// Request statuses include: pending, approved, rejected and empty(returns all requests).
func (c *Client) ListRequests(status string) (*ListRequestsResponse, error) {
	switch status {
	case "pending":
		status = "pending"
	case "approved":
		status = "approved"
	case "rejected":
		status = "rejected"
	case "":
		status = ""
	default:
		return nil, errors.New("The status are not accepted")
	}
	c.result = new(ListRequestsResponse)
	data, err := c.makeRequest("GET", "/request?"+status, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListRequestsResponse), err
}

// ViewRequest exports Use this endpoint to retrieve a certificate request.
func (c *Client) ViewRequest(requestID string) (*ViewRequestResponse, error) {
	c.result = new(ViewRequestResponse)
	data, err := c.makeRequest("GET", "/request/"+requestID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewRequestResponse), err
}

// UpdateRequestStatus exports Use this endpoint to retrieve the status of a previously submitted certificate request.
// Statuses [REQUIRED]: submitted, pending, approved, rejected
func (c *Client) UpdateRequestStatus(requestID string) (bool, error) {
	_, err := c.makeRequest("PUT", "/request/"+requestID+"/status", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}
