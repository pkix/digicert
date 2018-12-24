package digicert

import (
	"encoding/json"
	"log"
)

// CheckUserNameResponse presents checking username available.
type CheckUserNameResponse struct {
	Available bool `json:"available"`

	SchemeValidationErrors
}

// CheckUserName exports Use this endpoint to check to see if the specified username is available.
func (c *Client) CheckUserName(username string) bool {
	var check CheckUserNameResponse
	data, err := c.makeRequest("GET", "/user/availability/"+username, nil)
	if err != nil {
		log.Println("err", err)
		return false
	}
	if err := json.Unmarshal(data, &check); err != nil {
		return false
	}
	if c.statusCode == 200 && check.Available == true {
		return true
	}
	return false
}

// ListRolesResponse exports available roles
type ListRolesResponse struct {
	AccessRoles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"access_roles"`

	SchemeValidationErrors
}

// NewUserRequest presents a new user request
type NewUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Telephone string `json:"telephone"`
	Container struct {
		ID int `json:"id"`
	} `json:"container"`
	AccessRoles []struct {
		ID int `json:"id"`
	} `json:"access_roles"`
	ContainerIDAssignments []int `json:"container_id_assignments"`
}

// NewUserResponse present a new user with ID
type NewUserResponse struct {
	ID int `json:"id"`

	SchemeValidationErrors
}

// UserResponse presents a user detail
type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	AccountID int    `json:"account_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Telephone string `json:"telephone"`
	Status    string `json:"status"`
	Container struct {
		ID         int    `json:"id"`
		PublicID   string `json:"public_id"`
		Name       string `json:"name"`
		ParentID   int    `json:"parent_id"`
		TemplateID int    `json:"template_id"`
		HasLogo    bool   `json:"has_logo"`
		IsActive   bool   `json:"is_active"`
	} `json:"container"`
	AccessRoles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"access_roles"`
	IsCertCentral           bool `json:"is_cert_central"`
	IsEnterprise            bool `json:"is_enterprise"`
	HasContainerAssignments bool `json:"has_container_assignments"`
	ContainerVisibility     []struct {
		ID         int    `json:"id"`
		PublicID   string `json:"public_id"`
		Name       string `json:"name"`
		ParentID   int    `json:"parent_id"`
		TemplateID int    `json:"template_id"`
		HasLogo    bool   `json:"has_logo"`
		IsActive   bool   `json:"is_active"`
	} `json:"container_visibility"`

	SchemeValidationErrors
}

// UpdateUserRequest presents update user request
type UpdateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Telephone string `json:"telephone"`
}

// UpdateUserRoleRequest presents update user role request
type UpdateUserRoleRequest struct {
	AccessRoles []struct {
		ID int `json:"id"`
	} `json:"access_roles"`
}

// ListUsersResponse presents all of users
type ListUsersResponse struct {
	Users []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		JobTitle  string `json:"job_title"`
		Status    string `json:"status"`
		Container struct {
			ID         int    `json:"id"`
			PublicID   string `json:"public_id"`
			Name       string `json:"name"`
			ParentID   int    `json:"parent_id"`
			TemplateID int    `json:"template_id"`
			HasLogo    bool   `json:"has_logo"`
			IsActive   bool   `json:"is_active"`
		} `json:"container"`
		AccessRoles []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"access_roles"`
		HasContainerAssignments bool `json:"has_container_assignments"`
	} `json:"users"`

	SchemeValidationErrors
}

// ListRoles exports Use this endpoint to retrieve a list of access roles that are available for the specified container. These roles can be used to create or update a user in the container.
func (c *Client) ListRoles(containerID string) (*ListRolesResponse, error) {
	c.result = new(ListRolesResponse)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/role", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListRolesResponse), err
}

// NewUser exports Use this endpoint to create a new user. By default the user will be created in the same container as the currently authenticated user.
func (c *Client) NewUser(request *NewUserRequest) (*NewUserResponse, error) {
	c.result = new(NewUserResponse)
	c.request = request
	data, err := c.makeRequest("POST", "/user", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*NewUserResponse), err
}

// ResendCreateUserEmail exports Use this endpoint to resend a create user email to a specific user.
func (c *Client) ResendCreateUserEmail(userID string) bool {
	_, err := c.makeRequest("GET", "/user/"+userID+"/resend-create-email/", nil)
	if err != nil {
		return false
	}
	if c.statusCode == 204 {
		return true
	}
	return false
}

// ViewUser exports Use this endpoint to view the specified user.
func (c *Client) ViewUser(userID string) (*UserResponse, error) {
	c.result = new(UserResponse)
	data, err := c.makeRequest("GET", "/user/"+userID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*UserResponse), err
}

// UpdateUser exports Use this endpoint to update the specified user.
// username	Required	[string]
// first_name	Required	[string]
// last_name	Required	[string]
// email	Required	[string]
// job_title	Optional	[string]		This is required for to allow the User to be an approver for Extended Validation certificates
// telephone	Optional	[string]		This is required for to allow the User to be an approver for Extended Validation certificates
func (c *Client) UpdateUser(userID string, request *UpdateUserRequest) bool {
	c.request = request
	_, err := c.makeRequest("PUT", "/user/"+userID, nil)
	if err != nil {
		return false
	}
	if c.statusCode == 204 {
		return true
	}
	return false
}

// UpdateUserRole exports Use this endpoint to update the access roles of the specified user.
func (c *Client) UpdateUserRole(userID string, request *UpdateUserRoleRequest) bool {
	c.request = request
	_, err := c.makeRequest("PUT", "/user/"+userID+"/role", nil)
	if err != nil {
		return false
	}
	if c.statusCode == 204 {
		return true
	}
	return false
}

// DeleteUser exports Use this endpoint with the DELETE method to delete the specified user.
func (c *Client) DeleteUser(userID string) bool {
	_, err := c.makeRequest("DELETE", "/user/"+userID, nil)
	if err != nil {
		return false
	}
	if c.statusCode == 204 {
		return true
	}
	return false
}

// ListUsers exports Use this endpoint to retrieve a list of users in the current container and all child containers or from the specified container.
func (c *Client) ListUsers(containerID string) (*ListUsersResponse, error) {
	c.result = new(ListUsersResponse)
	data, err := c.makeRequest("GET", "/user?container_id="+containerID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListUsersResponse), err
}
