/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the CreateServerNetworking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerNetworking{}

/*
	types and functions for networkId
*/

// isNotNullableString
type CreateServerNetworkingGetNetworkIdAttributeType = *string

func getCreateServerNetworkingGetNetworkIdAttributeTypeOk(arg CreateServerNetworkingGetNetworkIdAttributeType) (ret CreateServerNetworkingGetNetworkIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateServerNetworkingGetNetworkIdAttributeType(arg *CreateServerNetworkingGetNetworkIdAttributeType, val CreateServerNetworkingGetNetworkIdRetType) {
	*arg = &val
}

type CreateServerNetworkingGetNetworkIdArgType = string
type CreateServerNetworkingGetNetworkIdRetType = string

// CreateServerNetworking The initial networking setup for the server creation with a network.
type CreateServerNetworking struct {
	// Universally Unique Identifier (UUID).
	NetworkId CreateServerNetworkingGetNetworkIdAttributeType `json:"networkId,omitempty"`
}

// NewCreateServerNetworking instantiates a new CreateServerNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerNetworking() *CreateServerNetworking {
	this := CreateServerNetworking{}
	return &this
}

// NewCreateServerNetworkingWithDefaults instantiates a new CreateServerNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerNetworkingWithDefaults() *CreateServerNetworking {
	this := CreateServerNetworking{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *CreateServerNetworking) GetNetworkId() (res CreateServerNetworkingGetNetworkIdRetType) {
	res, _ = o.GetNetworkIdOk()
	return
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetworking) GetNetworkIdOk() (ret CreateServerNetworkingGetNetworkIdRetType, ok bool) {
	return getCreateServerNetworkingGetNetworkIdAttributeTypeOk(o.NetworkId)
}

// HasNetworkId returns a boolean if a field has been set.
func (o *CreateServerNetworking) HasNetworkId() bool {
	_, ok := o.GetNetworkIdOk()
	return ok
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *CreateServerNetworking) SetNetworkId(v CreateServerNetworkingGetNetworkIdRetType) {
	setCreateServerNetworkingGetNetworkIdAttributeType(&o.NetworkId, v)
}

func (o CreateServerNetworking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateServerNetworkingGetNetworkIdAttributeTypeOk(o.NetworkId); ok {
		toSerialize["NetworkId"] = val
	}
	return toSerialize, nil
}

type NullableCreateServerNetworking struct {
	value *CreateServerNetworking
	isSet bool
}

func (v NullableCreateServerNetworking) Get() *CreateServerNetworking {
	return v.value
}

func (v *NullableCreateServerNetworking) Set(val *CreateServerNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerNetworking(val *CreateServerNetworking) *NullableCreateServerNetworking {
	return &NullableCreateServerNetworking{value: val, isSet: true}
}

func (v NullableCreateServerNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
