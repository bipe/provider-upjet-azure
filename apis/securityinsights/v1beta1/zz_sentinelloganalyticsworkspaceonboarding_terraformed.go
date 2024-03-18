// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SentinelLogAnalyticsWorkspaceOnboarding
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetTerraformResourceType() string {
	return "azurerm_sentinel_log_analytics_workspace_onboarding"
}

// GetConnectionDetailsMapping for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this SentinelLogAnalyticsWorkspaceOnboarding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelLogAnalyticsWorkspaceOnboardingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetTerraformSchemaVersion() int {
	return 0
}
