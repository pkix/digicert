package digicert

import (
	"encoding/json"
	"time"
)

// NewContainerRequest represents a new request of container detail
type NewContainerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TemplateID  int    `json:"template_id"`
	User        struct {
		Username    string `json:"username"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		AccessRoles []struct {
			ID int `json:"id"`
		} `json:"access_roles"`
	} `json:"user"`
}

// NewContainerResponse represents a new container ID
type NewContainerResponse struct {
	ID int `json:"id"`

	SchemeValidationErrors
}

// UpdateAContainerRequest represents updating container reqeust
type UpdateAContainerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Ekey        string `json:"ekey"`
}

//ViewContainerDetails represents a container details
type ViewContainerDetails struct {
	AccessRoles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"access_roles"`
	DateCreated string `json:"date_created"`
	ID          int    `json:"id"`
	Name        string `json:"name"`

	SchemeValidationErrors
}

// ListContainerTempaltesResponse represents the container templates
type ListContainerTempaltesResponse struct {
	ContainerTemplates []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		DateCreated time.Time `json:"date_created"`
	} `json:"container_templates"`

	SchemeValidationErrors
}

// ViewAContainerTempl presents a container template
type ViewAContainerTempl struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"date_created"`
	AccessRoles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"access_roles"`
}

// ListChilContainersResponse presents a list children containers of a container
type ListChilContainersResponse struct {
	Containers []struct {
		ID                 int      `json:"id"`
		PublicID           string   `json:"public_id"`
		Name               string   `json:"name"`
		Description        string   `json:"description"`
		ParentID           int      `json:"parent_id"`
		TemplateID         int      `json:"template_id"`
		HasLogo            bool     `json:"has_logo"`
		IsActive           bool     `json:"is_active"`
		AllowedDomainNames []string `json:"allowed_domain_names"`
	} `json:"containers"`

	SchemeValidationErrors
}

// ViewAContainerOfParentResponse presents a detail of a container's parent
type ViewAContainerOfParentResponse struct {
	ID                 int      `json:"id"`
	PublicID           string   `json:"public_id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	ParentID           int      `json:"parent_id"`
	TemplateID         int      `json:"template_id"`
	HasLogo            bool     `json:"has_logo"`
	IsActive           bool     `json:"is_active"`
	AllowedDomainNames []string `json:"allowed_domain_names"`

	SchemeValidationErrors
}

// NewContainer exports A Container is an Operational Division used to model your organizational structure. The features of the container you create are determined by its Container Template. When you create a new container, you may also specify the name, email address, and access role of a user for the container. An email will be sent to the new user to set up an account password.
func (c *Client) NewContainer(containerID string, request *NewContainerRequest) (*NewContainerResponse, error) {
	c.result = new(NewContainerResponse)
	c.request = request
	data, err := c.makeRequest("POST", "/container/"+containerID+"/children", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*NewContainerResponse), err
}

// UpdateAContainer exports to update a container to change its name or description.
func (c *Client) UpdateAContainer(containerID string, request *UpdateAContainerRequest) (bool, error) {
	c.request = request
	_, err := c.makeRequest("PUT", "/container/"+containerID, nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// DeactiveContainer exports deactivates the given container and all its children.
func (c *Client) DeactiveContainer(containerID string) (bool, error) {
	_, err := c.makeRequest("PUT", "/container/"+containerID+"/deactivate", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// ActiveContainer exports activates the given container and all its children.
func (c *Client) ActiveContainer(containerID string) (bool, error) {
	_, err := c.makeRequest("PUT", "/container/"+containerID+"/active", nil)
	if err != nil {
		return false, err
	}

	if c.statusCode == 204 {
		return true, err
	}
	return false, err
}

// ViewContainer exports information about a specific container can be retrieved through this endpoint, including its name, description, template, and parent container id.
func (c *Client) ViewContainer(containerID string) (*ViewContainerDetails, error) {
	c.result = new(ViewContainerDetails)
	data, err := c.makeRequest("GET", "/container/"+containerID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewContainerDetails), err
}

// ListContainerTempaltes exports container Templates define a set of features that are available to a container. Use this endpoint to retrieve a list of the templates that are available to use to create child containers.
func (c *Client) ListContainerTempaltes(containerID string) (*ListContainerTempaltesResponse, error) {
	c.result = new(ListContainerTempaltesResponse)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/template", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListContainerTempaltesResponse), err
}

// ViewAContainerTempl exports Use this endpoint to retrieve information about a specific container template, including which user access roles are available under this template.
func (c *Client) ViewAContainerTempl(containerID, templID string) (*ViewAContainerTempl, error) {
	c.result = new(ViewAContainerTempl)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/template/"+templID, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewAContainerTempl), err
}

// ListChilContainers exports retrieves a list of child containers for the specified container. The list only includes the immediate children of the container.
func (c *Client) ListChilContainers(containerID string) (*ListChilContainersResponse, error) {
	c.result = new(ListChilContainersResponse)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/children", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ListChilContainersResponse), err
}

// ViewAContainerOfParent exports Retieves information about the parent container of the specified container.
func (c *Client) ViewAContainerOfParent(containerID string) (*ViewAContainerOfParentResponse, error) {
	c.result = new(ViewAContainerOfParentResponse)
	data, err := c.makeRequest("GET", "/container/"+containerID+"/parent", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.result); err != nil {
		return nil, err
	}
	return c.result.(*ViewAContainerOfParentResponse), err
}
