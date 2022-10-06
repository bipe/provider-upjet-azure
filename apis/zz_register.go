/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/provider-azure/apis/apimanagement/v1beta1"
	v1beta1authorization "github.com/upbound/provider-azure/apis/authorization/v1beta1"
	v1beta1azure "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1cache "github.com/upbound/provider-azure/apis/cache/v1beta1"
	v1beta1compute "github.com/upbound/provider-azure/apis/compute/v1beta1"
	v1beta1containerregistry "github.com/upbound/provider-azure/apis/containerregistry/v1beta1"
	v1beta1containerservice "github.com/upbound/provider-azure/apis/containerservice/v1beta1"
	v1beta1cosmosdb "github.com/upbound/provider-azure/apis/cosmosdb/v1beta1"
	v1beta1dataprotection "github.com/upbound/provider-azure/apis/dataprotection/v1beta1"
	v1beta1datashare "github.com/upbound/provider-azure/apis/datashare/v1beta1"
	v1beta1dbformariadb "github.com/upbound/provider-azure/apis/dbformariadb/v1beta1"
	v1beta1dbformysql "github.com/upbound/provider-azure/apis/dbformysql/v1beta1"
	v1beta1dbforpostgresql "github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1"
	v1beta1devices "github.com/upbound/provider-azure/apis/devices/v1beta1"
	v1beta1eventhub "github.com/upbound/provider-azure/apis/eventhub/v1beta1"
	v1beta1insights "github.com/upbound/provider-azure/apis/insights/v1beta1"
	v1beta1keyvault "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta1kusto "github.com/upbound/provider-azure/apis/kusto/v1beta1"
	v1beta1logic "github.com/upbound/provider-azure/apis/logic/v1beta1"
	v1beta1management "github.com/upbound/provider-azure/apis/management/v1beta1"
	v1beta1marketplaceordering "github.com/upbound/provider-azure/apis/marketplaceordering/v1beta1"
	v1beta1media "github.com/upbound/provider-azure/apis/media/v1beta1"
	v1beta1mixedreality "github.com/upbound/provider-azure/apis/mixedreality/v1beta1"
	v1beta1netapp "github.com/upbound/provider-azure/apis/netapp/v1beta1"
	v1beta1network "github.com/upbound/provider-azure/apis/network/v1beta1"
	v1beta1notificationhubs "github.com/upbound/provider-azure/apis/notificationhubs/v1beta1"
	v1beta1operationalinsights "github.com/upbound/provider-azure/apis/operationalinsights/v1beta1"
	v1beta1resources "github.com/upbound/provider-azure/apis/resources/v1beta1"
	v1beta1security "github.com/upbound/provider-azure/apis/security/v1beta1"
	v1beta1sql "github.com/upbound/provider-azure/apis/sql/v1beta1"
	v1beta1storage "github.com/upbound/provider-azure/apis/storage/v1beta1"
	v1beta1storagecache "github.com/upbound/provider-azure/apis/storagecache/v1beta1"
	v1beta1storagesync "github.com/upbound/provider-azure/apis/storagesync/v1beta1"
	v1beta1streamanalytics "github.com/upbound/provider-azure/apis/streamanalytics/v1beta1"
	v1alpha1 "github.com/upbound/provider-azure/apis/v1alpha1"
	v1beta1apis "github.com/upbound/provider-azure/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1authorization.SchemeBuilder.AddToScheme,
		v1beta1azure.SchemeBuilder.AddToScheme,
		v1beta1cache.SchemeBuilder.AddToScheme,
		v1beta1compute.SchemeBuilder.AddToScheme,
		v1beta1containerregistry.SchemeBuilder.AddToScheme,
		v1beta1containerservice.SchemeBuilder.AddToScheme,
		v1beta1cosmosdb.SchemeBuilder.AddToScheme,
		v1beta1dataprotection.SchemeBuilder.AddToScheme,
		v1beta1datashare.SchemeBuilder.AddToScheme,
		v1beta1dbformariadb.SchemeBuilder.AddToScheme,
		v1beta1dbformysql.SchemeBuilder.AddToScheme,
		v1beta1dbforpostgresql.SchemeBuilder.AddToScheme,
		v1beta1devices.SchemeBuilder.AddToScheme,
		v1beta1eventhub.SchemeBuilder.AddToScheme,
		v1beta1insights.SchemeBuilder.AddToScheme,
		v1beta1keyvault.SchemeBuilder.AddToScheme,
		v1beta1kusto.SchemeBuilder.AddToScheme,
		v1beta1logic.SchemeBuilder.AddToScheme,
		v1beta1management.SchemeBuilder.AddToScheme,
		v1beta1marketplaceordering.SchemeBuilder.AddToScheme,
		v1beta1media.SchemeBuilder.AddToScheme,
		v1beta1mixedreality.SchemeBuilder.AddToScheme,
		v1beta1netapp.SchemeBuilder.AddToScheme,
		v1beta1network.SchemeBuilder.AddToScheme,
		v1beta1notificationhubs.SchemeBuilder.AddToScheme,
		v1beta1operationalinsights.SchemeBuilder.AddToScheme,
		v1beta1resources.SchemeBuilder.AddToScheme,
		v1beta1security.SchemeBuilder.AddToScheme,
		v1beta1sql.SchemeBuilder.AddToScheme,
		v1beta1storage.SchemeBuilder.AddToScheme,
		v1beta1storagecache.SchemeBuilder.AddToScheme,
		v1beta1storagesync.SchemeBuilder.AddToScheme,
		v1beta1streamanalytics.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
