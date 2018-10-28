package lokalise

import (
	"context"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/lokalise/go-lokalise-api/model"
)

type ProjectsService struct {
	client *Client
}

const (
	pathProjects = "projects"
)

type ProjectsOptions struct {
	PageOptions
	TeamID int64 `json:"teamID,omitempty"`
}

func (options ProjectsOptions) Apply(req *resty.Request) {
	options.PageOptions.Apply(req)
	if options.TeamID != 0 {
		req.SetQueryParam("filter_team_id", fmt.Sprintf("%d", options.TeamID))
	}
}

func (c *ProjectsService) List(ctx context.Context, pageOptions ProjectsOptions) (model.ProjectsResponse, error) {
	var res model.ProjectsResponse
	resp, err := c.client.getList(ctx, pathProjects, &res, pageOptions)
	if err != nil {
		return model.ProjectsResponse{}, err
	}
	applyPaged(resp, &res.Paged)
	return res, apiError(resp)
}

func (c *ProjectsService) Create(ctx context.Context, name, description string) (model.Project, error) {
	var res model.Project
	resp, err := c.client.post(ctx, pathProjects, &res, map[string]interface{}{
		"name":        name,
		"description": description,
	})
	if err != nil {
		return model.Project{}, err
	}
	return res, apiError(resp)
}

func (c *ProjectsService) CreateForTeam(ctx context.Context, name, description string, teamID int64) (model.Project, error) {
	var res model.Project
	resp, err := c.client.post(ctx, pathProjects, &res, map[string]interface{}{
		"name":        name,
		"description": description,
		"team_id":     teamID,
	})
	if err != nil {
		return model.Project{}, err
	}
	return res, apiError(resp)
}

func (c *ProjectsService) Retrieve(ctx context.Context, projectID string) (model.Project, error) {
	var res model.Project
	resp, err := c.client.get(ctx, fmt.Sprintf("%s/%s", pathProjects, projectID), &res)
	if err != nil {
		return model.Project{}, err
	}
	return res, apiError(resp)
}

func (c *ProjectsService) Update(ctx context.Context, projectID, name, description string) (model.Project, error) {
	var res model.Project
	resp, err := c.client.put(ctx, fmt.Sprintf("%s/%s", pathProjects, projectID), &res, map[string]interface{}{
		"name":        name,
		"description": description,
	})
	if err != nil {
		return model.Project{}, err
	}
	return res, apiError(resp)
}
