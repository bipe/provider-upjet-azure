/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AgentPool.
func (mg *AgentPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerRegistryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerRegistryNameRef,
		Selector:     mg.Spec.ForProvider.ContainerRegistryNameSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerRegistryName")
	}
	mg.Spec.ForProvider.ContainerRegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerRegistryNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkSubnetID")
	}
	mg.Spec.ForProvider.VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContainerConnectedRegistry.
func (mg *ContainerConnectedRegistry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerRegistryID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ContainerRegistryIDRef,
		Selector:     mg.Spec.ForProvider.ContainerRegistryIDSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerRegistryID")
	}
	mg.Spec.ForProvider.ContainerRegistryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerRegistryIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SyncTokenID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SyncTokenIDRef,
		Selector:     mg.Spec.ForProvider.SyncTokenIDSelector,
		To: reference.To{
			List:    &TokenList{},
			Managed: &Token{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SyncTokenID")
	}
	mg.Spec.ForProvider.SyncTokenID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SyncTokenIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Registry.
func (mg *Registry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkRuleSet); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetID")
			}
			mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkRuleSet[i3].VirtualNetwork[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScopeMap.
func (mg *ScopeMap) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerRegistryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerRegistryNameRef,
		Selector:     mg.Spec.ForProvider.ContainerRegistryNameSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerRegistryName")
	}
	mg.Spec.ForProvider.ContainerRegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerRegistryNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Token.
func (mg *Token) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerRegistryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerRegistryNameRef,
		Selector:     mg.Spec.ForProvider.ContainerRegistryNameSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerRegistryName")
	}
	mg.Spec.ForProvider.ContainerRegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerRegistryNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScopeMapID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ScopeMapIDRef,
		Selector:     mg.Spec.ForProvider.ScopeMapIDSelector,
		To: reference.To{
			List:    &ScopeMapList{},
			Managed: &ScopeMap{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScopeMapID")
	}
	mg.Spec.ForProvider.ScopeMapID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeMapIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Webhook.
func (mg *Webhook) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RegistryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RegistryNameRef,
		Selector:     mg.Spec.ForProvider.RegistryNameSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RegistryName")
	}
	mg.Spec.ForProvider.RegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegistryNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
