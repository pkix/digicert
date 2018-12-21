package digicert

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// ViewOrderResponse exports view order detail
type ViewOrderResponse struct {
	ID          int `json:"id,omitempty"`
	Certificate struct {
		ID           int       `json:"id,omitempty"`
		Thumbprint   string    `json:"thumbprint,omitempty"`
		SerialNumber string    `json:"serial_number,omitempty"`
		CommonName   string    `json:"common_name,omitempty"`
		DNSNames     []string  `json:"dns_names,omitempty"`
		DateCreated  time.Time `json:"date_created,omitempty"`
		ValidFrom    string    `json:"valid_from,omitempty"`
		ValidTill    string    `json:"valid_till,omitempty"`
		Csr          string    `json:"csr,omitempty"`
		Organization struct {
			ID int `json:"id,omitempty"`
		} `json:"organization,omitempty"`
		OrganizationUnits []string `json:"organization_units,omitempty"`
		ServerPlatform    struct {
			ID         int    `json:"id,omitempty"`
			Name       string `json:"name,omitempty"`
			InstallURL string `json:"install_url,omitempty"`
			CsrURL     string `json:"csr_url,omitempty"`
		} `json:"server_platform,omitempty"`
		SignatureHash string `json:"signature_hash,omitempty"`
		KeySize       int    `json:"key_size,omitempty"`
		CaCert        struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"ca_cert,omitempty"`
	} `json:"certificate,omitempty"`
	Status         string    `json:"status,omitempty"`
	IsRenewal      bool      `json:"is_renewal,omitempty"`
	IsRenewed      bool      `json:"is_renewed,omitempty"`
	RenewedOrderID int       `json:"renewed_order_id,omitempty"`
	DateCreated    time.Time `json:"date_created,omitempty"`
	Organization   struct {
		Name        string `json:"name,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
		IsActive    bool   `json:"is_active,omitempty"`
		City        string `json:"city,omitempty"`
		State       string `json:"state,omitempty"`
		Country     string `json:"country,omitempty"`
	} `json:"organization,omitempty"`
	DisableRenewalNotifications bool `json:"disable_renewal_notifications,omitempty"`
	Container                   struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"container,omitempty"`
	Product struct {
		NameID                string `json:"name_id,omitempty"`
		Name                  string `json:"name,omitempty"`
		Type                  string `json:"type,omitempty"`
		ValidationType        string `json:"validation_type,omitempty"`
		ValidationName        string `json:"validation_name,omitempty"`
		ValidationDescription string `json:"validation_description,omitempty"`
	} `json:"product,omitempty"`
	OrganizationContact struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Telephone string `json:"telephone,omitempty"`
	} `json:"organization_contact,omitempty"`
	TechnicalContact struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Telephone string `json:"telephone,omitempty"`
	} `json:"technical_contact,omitempty"`
	User struct {
		ID        int    `json:"id,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
	} `json:"user,omitempty"`
	Requests []struct {
		ID       int       `json:"id,omitempty"`
		Date     time.Time `json:"date,omitempty"`
		Type     string    `json:"type,omitempty"`
		Status   string    `json:"status,omitempty"`
		Comments string    `json:"comments,omitempty"`
	} `json:"requests,omitempty"`
	ReceiptID            int    `json:"receipt_id,omitempty"`
	CsProvisioningMethod string `json:"cs_provisioning_method,omitempty"`
	PublicID             string `json:"public_id,omitempty"`
	AllowDuplicates      bool   `json:"allow_duplicates,omitempty"`
	UserAssignments      []struct {
		ID        int    `json:"id,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
	} `json:"user_assignments,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"`
	DisableCt     bool   `json:"disable_ct,omitempty"`

	SchemeValidationErrors
}

// ListOrders export listing all of orders
type ListOrders struct {
	Orders []struct {
		ID          int `json:"id,omitempty"`
		Certificate struct {
			ID            int      `json:"id,omitempty"`
			CommonName    string   `json:"common_name,omitempty"`
			DNSNames      []string `json:"dns_names,omitempty"`
			ValidTill     string   `json:"valid_till,omitempty"`
			SignatureHash string   `json:"signature_hash,omitempty"`
		} `json:"certificate"`
		Status       string    `json:"status,omitempty"`
		DateCreated  time.Time `json:"date_created,omitempty"`
		Organization struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"organization"`
		ValidityYears int `json:"validity_years,omitempty"`
		Container     struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"container"`
		Product struct {
			NameID string `json:"name_id,omitempty"`
			Name   string `json:"name,omitempty"`
			Type   string `json:"type,omitempty"`
		} `json:"product"`
		Price int `json:"price,omitempty,omitempty"`
	} `json:"orders"`
	Page struct {
		Total  int `json:"total,omitempty"`
		Limit  int `json:"limit,omitempty"`
		Offset int `json:"offset,omitempty"`
	} `json:"page"`

	SchemeValidationErrors
}

// ViewOrder exports Use this endpoint to retrieve a certificate order. Note that the Technical Contact information is not currently used. Technical Contact information will be copied from the Organization Contact at this time.
func (c *Client) ViewOrder(orderID string) (*ViewOrderResponse, error) {
	c.result = new(ViewOrderResponse)
	data, err := c.apiconnect("GET", "/order/certificate/"+orderID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewOrderResponse), err
}

// ListOrders exports Use this endpoint to retrieve a list of all certificate orders.
func (c *Client) ListOrders(limit, offset int) (*ListOrders, error) {
	c.result = new(ListOrders)
	_limit := strconv.Itoa(limit)
	_offset := strconv.Itoa(offset)
	data, err := c.apiconnect("GET", "/order/certificate/?limit="+_limit+"&offset="+_offset, nil)
	if err != nil {
		return nil, err
	}

	jsonByte := []byte(strings.Replace(string(data), "[]", "{}", -1)) // DigiCert API returns array [] as empty of Organization
	if err := json.Unmarshal(jsonByte, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListOrders), err
}

// submitting OV/EV/DV, Client certifiates orders to digicert

// UnknownSSLRequest presents Order SSL Using Product Determinato
type UnknownSSLRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	RenewedThumbprint string `json:"renewed_thumbprint"`
	Organization      struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	Product                     struct {
		TypeHint string `json:"type_hint"`
	} `json:"product"`
	RenewalOfOrderID int  `json:"renewal_of_order_id"`
	DisableCt        bool `json:"disable_ct"`
}

// UnknownSSLResponse presents response of determinator SSL orders.
type UnknownSSLResponse struct {
	ID       int `json:"id"`
	Requests []struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	} `json:"requests"`
}

// OrderDVRequest presents Geotrust and RapidSSL requests
type OrderDVRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
	} `json:"certificate"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	TechnicalContact            struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		JobTitle  string `json:"job_title"`
		Telephone string `json:"telephone"`
	} `json:"technical_contact"`
	DisableCt          bool   `json:"disable_ct"`
	DcvMethod          string `json:"dcv_method"`
	Locale             string `json:"locale"`
	AlternativeOrderID string `json:"alternative_order_id"`
	DcvEmails          []struct {
		DNSName string `json:"dns_name"`
		Email   string `json:"email"`
	} `json:"dcv_emails"`
}

// OrderDVResponse presents response of Geotrust and RapidSSL DV
type OrderDVResponse struct {
	ID             int    `json:"id"`
	CertificateID  int    `json:"certificate_id"`
	DcvRandomValue string `json:"dcv_random_value"`
}

// OrderSSLByDeterminator exports To determine the appropriate SSL product being requested based on the data passed.
// If an product cannot be determined, it will return a 400 error with the message code: "ambiguous_product".
func (c *Client) OrderSSLByDeterminator(request *UnknownSSLRequest) (*UnknownSSLResponse, error) {
	c.result = new(UnknownSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*UnknownSSLResponse), err
}

// OrderDVSSL exports To order DVSSL certificates.
// DVSSL Brands supports: geotrust or rapidssl
func (c *Client) OrderDVSSL(dvBrand string, request *OrderDVRequest) (*OrderDVResponse, error) {
	c.result = new(OrderDVResponse)
	c.request = request
	switch dvBrand {
	case "geotrust":
		dvBrand = "ssl_dv_geotrust"
	case "rapidssl":
		dvBrand = "ssl_dv_rapidssl"
	default:
		return nil, errors.New("The DVSSL brands are not accepted")
	}
	data, err := c.apiconnect("POST", "/order/certificate/"+dvBrand, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderDVResponse), err
}

// OrderStandardSSLRequest exports request of order a digicert standard ssl
type OrderStandardSSLRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		ProfileOption string `json:"profile_option"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	PaymentMethod               string `json:"payment_method"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderOVEVSSLResponse presents a response of order digicert standard ssl
type OrderOVEVSSLResponse struct {
	ID       int `json:"id"`
	Requests []struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	} `json:"requests"`
}

// OrderStandardSSL exports To order DigiCert standard SSL certificate.
func (c *Client) OrderStandardSSL(request *OrderStandardSSLRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_plus", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderSSLMultiDomainRequest exports a request of multi-domain SSL
type OrderSSLMultiDomainRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		ProfileOption string `json:"profile_option"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderSSLMultiDomain exports To order DigiCert multi-domain SSL certificate.
func (c *Client) OrderSSLMultiDomain(request *OrderSSLMultiDomainRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_multi_domain", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderWildcardSSLRequest exports a request of wildcard certificate order
type OrderWildcardSSLRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		ProfileOption string `json:"profile_option"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderWildcardSSL exports To order DigiCert wildcard SSL certificate.
func (c *Client) OrderWildcardSSL(request *OrderWildcardSSLRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_wildcard", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderEVSSLRequest exports a request of Digicert EV SSL
type OrderEVSSLRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderEVPlusSSL exports To order DigiCert EV Plus SSL certificate.
func (c *Client) OrderEVPlusSSL(request *OrderEVSSLRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_ev_plus", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderEVMultiDomainRequest presents a request of Digicert EV Multi-domain SSL
type OrderEVMultiDomainRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		ProfileOption string `json:"profile_option"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderEVMultiDomainSSL exports To order DigiCert EV Plus Multi-Domain SSL certificate.
func (c *Client) OrderEVMultiDomainSSL(request *OrderEVMultiDomainRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_multi_domain", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderCloudSSLReqeust presents a request of Cloud wildcard combination SSL
type OrderCloudSSLReqeust struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
	DisableCt                   bool   `json:"disable_ct"`
}

// OrderCloudSSL exports To order DigiCert CloudSSL certificate, which enable multi-domain wildcards in one OVSSL certificate
func (c *Client) OrderCloudSSL(request *OrderEVMultiDomainRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/ssl_cloud_wildcard", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderClientPremiumRequest presents a requst of client premium order
type OrderClientPremiumRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Emails            []string `json:"emails"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		SignatureHash     string   `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears    int `json:"validity_years"`
	AutoRenew        int `json:"auto_renew"`
	RenewalOfOrderID int `json:"renewal_of_order_id"`
}

// OrderClientResponse presents a response of client certificate order
type OrderClientResponse struct {
	ID int `json:"id"`
}

// OrderClientPremium exports Use this endpoint to create a client certificate that can be used for email encryption and signing, client authentication, and document signing.
// NOTE: client certificates do not require additonal approval after the order is created.
func (c *Client) OrderClientPremium(request *OrderClientPremiumRequest) (*OrderClientResponse, error) {
	c.result = new(OrderClientResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/client_premium_sha2", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderClientResponse), err
}

// OrderClientEmailSecurityPlusRequest presents a request of client email security plus certificate
type OrderClientEmailSecurityPlusRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Emails            []string `json:"emails"`
		OrganizationUnits []string `json:"organization_units"`
		SignatureHash     string   `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears    int `json:"validity_years"`
	AutoRenew        int `json:"auto_renew"`
	RenewalOfOrderID int `json:"renewal_of_order_id"`
}

// OrderClientDigitalSignaturePlusRequest presents a request of Order Client Digital Signature Plus
type OrderClientDigitalSignaturePlusRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Emails            []string `json:"emails"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		SignatureHash     string   `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears    int `json:"validity_years"`
	AutoRenew        int `json:"auto_renew"`
	RenewalOfOrderID int `json:"renewal_of_order_id"`
}

// OrderClientEmailSecurityPlus exports Use this endpoint to create a client certificate that can be used for email encryption.
// NOTE: client certificates do not require additonal approval after the order is created.
func (c *Client) OrderClientEmailSecurityPlus(request *OrderClientEmailSecurityPlusRequest) (*OrderClientResponse, error) {
	c.result = new(OrderClientResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/client_email_security_plus", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderClientResponse), err
}

// OrderClientDigitalSignaturePlus exports Use this endpoint to create a client certificate that can be used for email signing, document signing, and client authentication.
// NOTE: client certificates do not require additonal approval after the order is created.
func (c *Client) OrderClientDigitalSignaturePlus(request *OrderClientDigitalSignaturePlusRequest) (*OrderClientResponse, error) {
	c.result = new(OrderClientResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/client_digital_signature_plus", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderClientResponse), err
}

// OrderPrivateSSLPlusRequest presents a request of Order Private SSL Plus Request
type OrderPrivateSSLPlusRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		CaCertID      string `json:"ca_cert_id"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
}

// OrderPrivateSSLPlus exports To order DigiCert private SSL Plus certificate.
func (c *Client) OrderPrivateSSLPlus(request *OrderPrivateSSLPlusRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/private_ssl_plus", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderPrivateSSLWildcardRequest presents a request of Order Private SSL Wildcard
type OrderPrivateSSLWildcardRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		CaCertID      string `json:"ca_cert_id"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
}

// OrderPrivateSSLWildcard exports To order DigiCert private wildcard SSL Plus certificate.
func (c *Client) OrderPrivateSSLWildcard(request *OrderPrivateSSLWildcardRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/private_ssl_wildcard", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderPrivateSSLMultiDomainRequest presents a request of Order Private SSL Multi Domain
type OrderPrivateSSLMultiDomainRequest struct {
	Certificate struct {
		CommonName        string   `json:"common_name"`
		DNSNames          []string `json:"dns_names"`
		Csr               string   `json:"csr"`
		OrganizationUnits []string `json:"organization_units"`
		ServerPlatform    struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
		CaCertID      string `json:"ca_cert_id"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears               int    `json:"validity_years"`
	CustomExpirationDate        string `json:"custom_expiration_date"`
	Comments                    string `json:"comments"`
	DisableRenewalNotifications bool   `json:"disable_renewal_notifications"`
	RenewalOfOrderID            int    `json:"renewal_of_order_id"`
}

// OrderPrivateSSLMultiDomain exports To order DigiCert private multi-domain SSL Plus certificate.
func (c *Client) OrderPrivateSSLMultiDomain(request *OrderPrivateSSLMultiDomainRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/private_ssl_multi_domain", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderCodeSigningRequest presents a request of Order a Code Signing
type OrderCodeSigningRequest struct {
	Certificate struct {
		Csr            string `json:"csr"`
		ServerPlatform struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears    int    `json:"validity_years"`
	Comments         string `json:"comments"`
	RenewalOfOrderID int    `json:"renewal_of_order_id"`
}

// OrderCodeSigning exports To order DigiCert standard code signing certificate.
func (c *Client) OrderCodeSigning(request *OrderCodeSigningRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/code_signing", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderEVCodeSigningRequest presents a request of Order an EV Code Signing
type OrderEVCodeSigningRequest struct {
	Certificate struct {
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears        int    `json:"validity_years"`
	Comments             string `json:"comments"`
	RenewalOfOrderID     int    `json:"renewal_of_order_id"`
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
}

// OrderEVCodeSigning exports To order DigiCert EV code signing certificate.
// signature_hash	Required	sha256, sha384, sha512 (sha1 on private certs)
// The certificate's signing algorithm hash, for code signing only sha256 is supported.
// method	Required	STANDARD,EXPEDITED
// The method to ship by, EXPEDITED carries an additional cost
func (c *Client) OrderEVCodeSigning(request *OrderEVCodeSigningRequest) (*OrderOVEVSSLResponse, error) {
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/code_signing_ev", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}

// OrderDocumentSigningOrganizationRequest presents a request of Order a Document Signing Organization
type OrderDocumentSigningOrganizationRequest struct {
	Certificate struct {
		ServerPlatform struct {
			ID int `json:"id"`
		} `json:"server_platform"`
		SignatureHash string `json:"signature_hash"`
	} `json:"certificate"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
	ValidityYears        int    `json:"validity_years"`
	Comments             string `json:"comments"`
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
	Subject struct {
		Name     string `json:"name"`
		JobTitle string `json:"job_title"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	} `json:"subject"`
}

// OrderDocumentSigningOrganization exports to order a Document Signing Organization (2000) or (5000) Certificate
// signature_hash	Required	sha256, sha384, sha512 (sha1 on private certs)
// The certificate's signing algorithm hash, for code signing only sha256 is supported.
// method	Required	STANDARD,EXPEDITED
// The method to ship by, EXPEDITED carries an additional cost
func (c *Client) OrderDocumentSigningOrganization(amount int, request *OrderDocumentSigningOrganizationRequest) (*OrderOVEVSSLResponse, error) {
	var product string
	switch amount {
	case 2000:
		product = "document_signing_org_1"
	case 5000:
		product = "document_signing_org_2"
	default:
		return nil, errors.New("There's no a product available for this amount")
	}
	c.result = new(OrderOVEVSSLResponse)
	c.request = request
	data, err := c.apiconnect("POST", "/order/certificate/"+product, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*OrderOVEVSSLResponse), err
}
