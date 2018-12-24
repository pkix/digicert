package digicert

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

// DownloadCertificateResponse exports download certificate
type DownloadCertificateResponse struct {
	Raw string `json:"raw"`
	// Validations []struct {
	// 	DateCreated    string `json:"date_created,omitempty"`
	// 	Description    string `json:"description,omitempty"`
	// 	Name           string `json:"name,omitempty"`
	// 	Status         string `json:"status,omitempty"`
	// 	Type           string `json:"type,omitempty"`
	// 	ValidatedUntil string `json:"validated_until,omitempty"`
	// } `json:"validations,omitempty"`
	SchemeValidationErrors
}

// RevokeCertificateRequest exports the request comment
type RevokeCertificateRequest struct {
	Comment string
}

// RevokeCertificateResponse exports revoke response
type RevokeCertificateResponse struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Requester struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"requester"`
	Comments string `json:"comments"`
	SchemeValidationErrors
}

// CancelRequest presents a certificate cancel request
type CancelRequest struct {
	Status     string `json:"status"`
	Note       string `json:"note"`
	SendEmails bool   `json:"send_emails"`
}

// ReissueRequest presents a reissue certificate
type ReissueRequest struct {
	Certificate struct {
		CommonName     string   `json:"common_name"`
		DNSNames       []string `json:"dns_names"`
		Csr            string   `json:"csr"`
		ServerPlatform struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
}

// ReissueResponse presents a reissue response
type ReissueResponse struct {
	ID       int `json:"id"`
	Requests []struct {
		ID int `json:"id"`
	} `json:"requests"`
	SchemeValidationErrors
}

// DuplicateRequest presents a duplicate certificate request
type DuplicateRequest struct {
	Certificate struct {
		CommonName     string   `json:"common_name"`
		DNSNames       []string `json:"dns_names"`
		Csr            string   `json:"csr"`
		ServerPlatform struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
}

// DuplicateResponse presents a duplicate certificate response
type DuplicateResponse struct {
	ID       int `json:"id"`
	Requests []struct {
		ID int `json:"id"`
	} `json:"requests"`

	SchemeValidationErrors
}

// ListDuplicateResponse presents all duplicate certificates.
type ListDuplicateResponse struct {
	Certificates []struct {
		ID             int       `json:"id"`
		Thumbprint     string    `json:"thumbprint"`
		SerialNumber   string    `json:"serial_number"`
		CommonName     string    `json:"common_name"`
		DNSNames       []string  `json:"dns_names"`
		Status         string    `json:"status"`
		DateCreated    time.Time `json:"date_created"`
		ValidFrom      string    `json:"valid_from"`
		ValidTill      string    `json:"valid_till"`
		Csr            string    `json:"csr"`
		ServerPlatform struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			InstallURL string `json:"install_url"`
			CsrURL     string `json:"csr_url"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		KeySize       int    `json:"key_size"`
		CaCertID      string `json:"ca_cert_id"`
		SubID         string `json:"sub_id"`
		PublicID      string `json:"public_id"`
	} `json:"certificates"`

	SchemeValidationErrors
}

// OrderStatusResponse presents a orderstatus within minutes.
type OrderStatusResponse struct {
	Orders []struct {
		OrderID       int    `json:"order_id"`
		CertificateID int    `json:"certificate_id"`
		Status        string `json:"status"`
	} `json:"orders"`

	SchemeValidationErrors
}

// DVChangeDCVMethodRequest presents changing DCV method.
type DVChangeDCVMethodRequest struct {
	DcvMethod string `json:"dcv_method"`
}

// DVRandomValue presents changing DCV method response.
type DVRandomValue struct {
	DcvRandomValue string `json:"dcv_random_value"`
}

// DVCheckDCVResponse presents checking dcv response.
type DVCheckDCVResponse struct {
	OrderStatus   string `json:"order_status"`
	CertificateID int    `json:"certificate_id"`
	DcvStatus     string `json:"dcv_status"`

	SchemeValidationErrors
}

// AddCSRRequest presents csr raw.
type AddCSRRequest struct {
	CSR string `json:"csr"`
}

// DownloadCertificate exports method download the certificate
func (c *Client) DownloadCertificate(certificateID string) (string, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/x-pem-file")
	data, err := c.makeRequest("GET", "/certificate/"+certificateID+"/download/platform", headers)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DownloadPKCS7Certificate exports method download a client certificate
func (c *Client) DownloadPKCS7Certificate(certificateID string) (string, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/x-pkcs7-certificates")
	data, err := c.makeRequest("GET", "/certificate/"+certificateID+"/download/format/p7b", headers)
	if err != nil {
		return "", err
	}
	// log.Println("data", string(data))
	return string(data), nil
}

// Revoke exports method revoke the certificate
func (c *Client) Revoke(certificateID, comment string) (*RevokeCertificateResponse, error) {
	c.request = &RevokeCertificateRequest{
		Comment: comment,
	}
	c.result = new(RevokeCertificateResponse)
	data, err := c.makeRequest("PUT", "/certificate/"+certificateID+"/revoke", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*RevokeCertificateResponse), err
}

// Cancel to update the status of an order. Currently this endpoint only allows updating the status to 'CANCELED'
func (c *Client) Cancel(orderID, comment string) (bool, error) {
	c.request = &CancelRequest{
		Status:     "CANCELED",
		Note:       comment,
		SendEmails: true,
	}
	_, err := c.makeRequest("PUT", "/order/certificate/"+orderID+"/status", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// Reissue to reissue a certificate order. A reissue replaces the existing certificate with a new one that has different information such as common name, CSR, etc.
func (c *Client) Reissue(orderID string, request *ReissueRequest) (*ReissueResponse, error) {
	c.request = request
	c.result = new(ReissueResponse)
	data, err := c.makeRequest("POST", "/order/certificate/"+orderID+"/reissue", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil || c.statusCode != 204 {
		return nil, err
	}
	return c.result.(*ReissueResponse), err
}

// Duplicate exports use this endpoint to request a duplicate certificate for an order. A duplicate shares the expiration date as the existing certificate and is identical with the exception of the CSR and a possible change in the server platform and signature hash. The common name and sans need to be the same as the original order. Multi-Domain SSL Certs allow a san to be moved to the common name. Wildcard certs allow for additional sans (as long as they are subdomains of the wildcard).
func (c *Client) Duplicate(orderID string, request *DuplicateRequest) (*DuplicateResponse, error) {
	c.request = request
	c.result = new(DuplicateResponse)
	data, err := c.makeRequest("POST", "/order/certificate/"+orderID+"/duplicate", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil || c.statusCode != 201 {
		return nil, err
	}
	return c.result.(*DuplicateResponse), err
}

// ListDuplicateCertificates exports view all duplicate certificates for an order.
func (c *Client) ListDuplicateCertificates(orderID string) (*ListDuplicateResponse, error) {
	c.result = new(ListDuplicateResponse)
	data, err := c.makeRequest("GET", "/order/certificate/"+orderID+"/duplicate", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListDuplicateResponse), err
}

// ListOrganizationsRequest presents retrive a list of organizations
type ListOrganizationsRequest struct {
	Organizations []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"organizations"`
}

// ListOrganizations exports retrieve a list of organizations.
func (c *Client) ListOrganizations(containerID string) (*ListOrganizationsRequest, error) {
	c.result = new(ListOrganizationsRequest)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/order/organization", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListOrganizationsRequest), err
}

// ListEmailValidationsResponse presents list of email validations.
type ListEmailValidationsResponse struct {
	DeliveryOptions []string `json:"delivery_options"`
	Emails          []struct {
		Email       string `json:"email"`
		Status      string `json:"status"`
		DateEmailed string `json:"date_emailed"`
	} `json:"emails"`
}

// ListEmailValidations exprots Use this endpoint to view the status of all emails that require validation on a client certificate order.
func (c *Client) ListEmailValidations(orderID string) (*ListEmailValidationsResponse, error) {
	c.result = new(ListEmailValidationsResponse)
	data, err := c.makeRequest("GET", "/order/certificate/"+orderID+"/email-validation", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListEmailValidationsResponse), err
}

// OrderStatus exports Use this endpoint to check on order status changes within a supplied time range up to a week (10080 minutes).
func (c *Client) OrderStatus(minutes int) (*OrderStatusResponse, error) {
	c.result = new(OrderStatusResponse)
	data, err := c.makeRequest("GET", "/order/certificate/status-changes?minutes="+strconv.Itoa(minutes), nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderStatusResponse), err
}

// DVChangeDCVMethod exports Use this endpoint on pending DV SSL orders to change the DCV method to use to prove control over the domain on the order. Method: email, dns-txt-token, http-token
func (c *Client) DVChangeDCVMethod(orderID, method string) (*DVRandomValue, error) {
	switch method {
	case "email":
		method = "email"
	case "dns-txt-token":
		method = "dns-txt-token"
	case "http-token":
		method = "http-token"
	default:
		return nil, errors.New("The wrong method")
	}
	c.request = &DVChangeDCVMethodRequest{
		DcvMethod: method,
	}
	c.result = new(DVRandomValue)
	data, err := c.makeRequest("PUT", "/order/certificate/"+orderID+"/dcv-method", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*DVRandomValue), err
}

// DVResendDCVEmail exprots Use this endpoint on pending DV SSL orders to resend DCV emails for a certificate order.
func (c *Client) DVResendDCVEmail(orderID, comment string) (bool, error) {
	_, err := c.makeRequest("PUT", "/order/certificate/"+orderID+"/resend-emails", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// DVDCVRandomValue exports Use this endpoint on pending DV SSL orders to generate a random value for dns-txt-token and http-token DCV methods.
func (c *Client) DVDCVRandomValue(orderID string) (*DVRandomValue, error) {
	c.result = new(DVRandomValue)
	data, err := c.makeRequest("PUT", "/order/certificate/"+orderID+"/dcv-random-value", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*DVRandomValue), err
}

// DVCheckDCV exports Use this endpoint on pending DV SSL orders to perform Domain Control Validation (DCV) over a domain one a random value for dns-txt-token or http-token is in place.
func (c *Client) DVCheckDCV(orderID string) (*DVCheckDCVResponse, error) {
	c.result = new(DVCheckDCVResponse)
	data, err := c.makeRequest("PUT", "/order/certificate/"+orderID+"/check-dcv", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*DVCheckDCVResponse), err
}

// AddCSR exports Use this endpoint to add or update a CSR on a pending certificate order.
// For Code Signing, CSR is only required for Java platform (server_platform.id = 55).
// For EV Code Signing, CSR is only required for email provisioning.
// For client certificates, the CSR is optional.
func (c *Client) AddCSR(orderID, csr string) (bool, error) {
	c.request = &AddCSRRequest{
		CSR: csr,
	}
	_, err := c.makeRequest("POST", "/order/certificate/"+orderID+"/csr", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}
