/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountPrivilegeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccountPrivilegeParameters struct {

	// Account host, default is `%`.
	// +kubebuilder:validation:Optional
	AccountHost *string `json:"accountHost,omitempty" tf:"account_host,omitempty"`

	// Account name.
	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// List of specified database name.
	// +kubebuilder:validation:Required
	DatabaseNames []*string `json:"databaseNames" tf:"database_names,omitempty"`

	// Instance ID.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	MySQLID *string `json:"mysqlId,omitempty" tf:"mysql_id,omitempty"`

	// +kubebuilder:validation:Optional
	MySQLIDRef *v1.Reference `json:"mySqlidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	MySQLIDSelector *v1.Selector `json:"mySqlidSelector,omitempty" tf:"-"`

	// Database permissions. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT` and `TRIGGER``.
	// +kubebuilder:validation:Optional
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`
}

// AccountPrivilegeSpec defines the desired state of AccountPrivilege
type AccountPrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountPrivilegeParameters `json:"forProvider"`
}

// AccountPrivilegeStatus defines the observed state of AccountPrivilege.
type AccountPrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountPrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountPrivilege is the Schema for the AccountPrivileges API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type AccountPrivilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountPrivilegeSpec   `json:"spec"`
	Status            AccountPrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountPrivilegeList contains a list of AccountPrivileges
type AccountPrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountPrivilege `json:"items"`
}

// Repository type metadata.
var (
	AccountPrivilege_Kind             = "AccountPrivilege"
	AccountPrivilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountPrivilege_Kind}.String()
	AccountPrivilege_KindAPIVersion   = AccountPrivilege_Kind + "." + CRDGroupVersion.String()
	AccountPrivilege_GroupVersionKind = CRDGroupVersion.WithKind(AccountPrivilege_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountPrivilege{}, &AccountPrivilegeList{})
}
