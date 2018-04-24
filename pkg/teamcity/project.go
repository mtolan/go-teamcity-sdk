package teamcity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"errors"
	"fmt"

	"github.com/dghubble/sling"

	"github.com/go-openapi/swag"
)

// Represents a Project
type Project struct {

	// archived
	Archived *bool `json:"archived,omitempty" xml:"archived"`

	// build types
	// BuildTypes *BuildTypes `json:"buildTypes,omitempty"`

	// // default template
	// DefaultTemplate *BuildType `json:"defaultTemplate,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// internal Id
	InternalID string `json:"internalId,omitempty" xml:"internalId"`

	// links
	// Links *Links `json:"links,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// parameters
	// Parameters *Properties `json:"parameters,omitempty"`

	// parent project
	ParentProject *Project `json:"parentProject,omitempty"`

	// parent project Id
	ParentProjectID string `json:"parentProjectId,omitempty" xml:"parentProjectId"`

	// parent project internal Id
	ParentProjectInternalID string `json:"parentProjectInternalId,omitempty" xml:"parentProjectInternalId"`

	// parent project name
	ParentProjectName string `json:"parentProjectName,omitempty" xml:"parentProjectName"`

	// project features
	// ProjectFeatures *ProjectFeatures `json:"projectFeatures,omitempty"`

	// projects
	// Projects *Projects `json:"projects,omitempty"`

	// // read only UI
	// ReadOnlyUI *StateField `json:"readOnlyUI,omitempty"`

	// // templates
	// Templates *BuildTypes `json:"templates,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty" xml:"uuid"`

	// vcs roots
	// VcsRoots *VcsRoots `json:"vcsRoots,omitempty"`

	// web Url
	WebURL string `json:"webUrl,omitempty" xml:"webUrl"`
}

type ProjectReference struct {
	// id
	ID string `json:"id,omitempty" xml:"id"`
	// name
	Name string `json:"name,omitempty" xml:"name"`
}

type ProjectService struct {
	sling *sling.Sling
}

func newProjectService(base *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: base.Path("projects/"),
	}
}

// CreateProject Creates a new project at root project level
func (s *ProjectService) CreateProject(project *Project) (*ProjectReference, error) {
	var created ProjectReference
	if err := project.Validate(); err != nil {
		return nil, err
	}

	response, err := s.sling.New().BodyJSON(project).Post("").ReceiveSuccess(&created)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 400 {
		return nil, fmt.Errorf("A project with name '%s' already exists", project.Name)
	}

	return &created, nil
}

// GetProject Retrieves a project resource by ID
func (c *Client) GetProject(id string) (*Project, error) {
	var out Project

	err := c.doJSONRequest("GET", fmt.Sprintf("projects/%s", id), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, err
}

//DeleteProject - Deletes a project
func (c *Client) DeleteProject(id string) error {
	return c.doJSONRequest("DELETE", fmt.Sprintf("projects/%s", id), nil, nil)
}

// Validate validates this project
func (m *Project) Validate() error {
	//var res []error

	if len(m.Name) <= 0 {
		return errors.New("Project must have a name")
	}

	// if err := m.validateBuildTypes(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateDefaultTemplate(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateLinks(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateParameters(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateParentProject(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateProjectFeatures(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateProjects(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateReadOnlyUI(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateTemplates(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if err := m.validateVcsRoots(formats); err != nil {
	// 	// prop
	// 	res = append(res, err)
	// }

	// if len(res) > 0 {
	// 	return errors.CompositeValidationError(res...)
	// }
	return nil
}

// func (m *Project) validateBuildTypes(formats strfmt.Registry) error {

// 	if swag.IsZero(m.BuildTypes) { // not required
// 		return nil
// 	}

// 	if m.BuildTypes != nil {

// 		if err := m.BuildTypes.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("buildTypes")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateDefaultTemplate(formats strfmt.Registry) error {

// 	if swag.IsZero(m.DefaultTemplate) { // not required
// 		return nil
// 	}

// 	if m.DefaultTemplate != nil {

// 		if err := m.DefaultTemplate.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("defaultTemplate")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateLinks(formats strfmt.Registry) error {

// 	if swag.IsZero(m.Links) { // not required
// 		return nil
// 	}

// 	if m.Links != nil {

// 		if err := m.Links.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("links")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateParameters(formats strfmt.Registry) error {

// 	if swag.IsZero(m.Parameters) { // not required
// 		return nil
// 	}

// 	if m.Parameters != nil {

// 		if err := m.Parameters.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("parameters")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateParentProject(formats strfmt.Registry) error {

// 	if swag.IsZero(m.ParentProject) { // not required
// 		return nil
// 	}

// 	if m.ParentProject != nil {

// 		if err := m.ParentProject.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("parentProject")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateProjectFeatures(formats strfmt.Registry) error {

// 	if swag.IsZero(m.ProjectFeatures) { // not required
// 		return nil
// 	}

// 	if m.ProjectFeatures != nil {

// 		if err := m.ProjectFeatures.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("projectFeatures")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateProjects(formats strfmt.Registry) error {

// 	if swag.IsZero(m.Projects) { // not required
// 		return nil
// 	}

// 	if m.Projects != nil {

// 		if err := m.Projects.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("projects")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateReadOnlyUI(formats strfmt.Registry) error {

// 	if swag.IsZero(m.ReadOnlyUI) { // not required
// 		return nil
// 	}

// 	if m.ReadOnlyUI != nil {

// 		if err := m.ReadOnlyUI.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("readOnlyUI")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateTemplates(formats strfmt.Registry) error {

// 	if swag.IsZero(m.Templates) { // not required
// 		return nil
// 	}

// 	if m.Templates != nil {

// 		if err := m.Templates.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("templates")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Project) validateVcsRoots(formats strfmt.Registry) error {

// 	if swag.IsZero(m.VcsRoots) { // not required
// 		return nil
// 	}

// 	if m.VcsRoots != nil {

// 		if err := m.VcsRoots.Validate(formats); err != nil {
// 			if ve, ok := err.(*errors.Validation); ok {
// 				return ve.ValidateName("vcsRoots")
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}