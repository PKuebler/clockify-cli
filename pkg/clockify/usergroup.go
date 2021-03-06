package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

// UserGroupService serve the clockify user group api
type UserGroupService struct {
	client *APIClient
}

// UserGroup resource from clockify
type UserGroup struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	WorkspaceID string   `json:"workspaceId"`
	UserIDs     []string `json:"userIds"`
}

// UserGroupListFilter for the clockify list endpoint
type UserGroupListFilter struct {
	Name      *string
	ProjectID *string
	Page      int
	PageSize  int
}

// ToQuery formats the filters for the Get Query
func (u *UserGroupListFilter) ToQuery() string {
	v := url.Values{}

	if u.Name != nil {
		v.Add("name", *u.Name)
	}
	if u.ProjectID != nil {
		v.Add("projectId", *u.ProjectID)
	}

	if u.Page > 0 {
		v.Add("page", strconv.Itoa(u.Page))
	}
	if u.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(u.PageSize))
	}

	return v.Encode()
}

// List all user groups in a workspace
func (u *UserGroupService) List(workspaceID string, filter *UserGroupListFilter) (*[]UserGroup, error) {
	path := fmt.Sprintf("workspaces/%s/user-groups", workspaceID)
	req, err := u.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var groups []UserGroup
	_, err = u.client.do(req, &groups)
	if err != nil {
		return nil, err
	}

	return &groups, nil
}
