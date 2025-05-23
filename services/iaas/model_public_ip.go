/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the PublicIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIp{}

/*
	types and functions for id
*/

// isNotNullableString
type PublicIpGetIdAttributeType = *string

func getPublicIpGetIdAttributeTypeOk(arg PublicIpGetIdAttributeType) (ret PublicIpGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPublicIpGetIdAttributeType(arg *PublicIpGetIdAttributeType, val PublicIpGetIdRetType) {
	*arg = &val
}

type PublicIpGetIdArgType = string
type PublicIpGetIdRetType = string

/*
	types and functions for ip
*/

// isNotNullableString
type PublicIpGetIpAttributeType = *string

func getPublicIpGetIpAttributeTypeOk(arg PublicIpGetIpAttributeType) (ret PublicIpGetIpRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPublicIpGetIpAttributeType(arg *PublicIpGetIpAttributeType, val PublicIpGetIpRetType) {
	*arg = &val
}

type PublicIpGetIpArgType = string
type PublicIpGetIpRetType = string

/*
	types and functions for labels
*/

// isFreeform
type PublicIpGetLabelsAttributeType = *map[string]interface{}
type PublicIpGetLabelsArgType = map[string]interface{}
type PublicIpGetLabelsRetType = map[string]interface{}

func getPublicIpGetLabelsAttributeTypeOk(arg PublicIpGetLabelsAttributeType) (ret PublicIpGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPublicIpGetLabelsAttributeType(arg *PublicIpGetLabelsAttributeType, val PublicIpGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for networkInterface
*/

// isNullableString
type PublicIpGetNetworkInterfaceAttributeType = *NullableString

func getPublicIpGetNetworkInterfaceAttributeTypeOk(arg PublicIpGetNetworkInterfaceAttributeType) (ret PublicIpGetNetworkInterfaceRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg.Get(), true
}

func setPublicIpGetNetworkInterfaceAttributeType(arg *PublicIpGetNetworkInterfaceAttributeType, val PublicIpGetNetworkInterfaceRetType) {
	if IsNil(*arg) {
		*arg = NewNullableString(val)
	} else {
		(*arg).Set(val)
	}
}

type PublicIpGetNetworkInterfaceArgType = *string
type PublicIpGetNetworkInterfaceRetType = *string

// PublicIp Object that represents a public IP.
type PublicIp struct {
	// Universally Unique Identifier (UUID).
	Id PublicIpGetIdAttributeType `json:"id,omitempty"`
	// Object that represents an IP address.
	Ip PublicIpGetIpAttributeType `json:"ip,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels PublicIpGetLabelsAttributeType `json:"labels,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkInterface PublicIpGetNetworkInterfaceAttributeType `json:"networkInterface,omitempty"`
}

// NewPublicIp instantiates a new PublicIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIp() *PublicIp {
	this := PublicIp{}
	return &this
}

// NewPublicIpWithDefaults instantiates a new PublicIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpWithDefaults() *PublicIp {
	this := PublicIp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicIp) GetId() (res PublicIpGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetIdOk() (ret PublicIpGetIdRetType, ok bool) {
	return getPublicIpGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *PublicIp) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicIp) SetId(v PublicIpGetIdRetType) {
	setPublicIpGetIdAttributeType(&o.Id, v)
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PublicIp) GetIp() (res PublicIpGetIpRetType) {
	res, _ = o.GetIpOk()
	return
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetIpOk() (ret PublicIpGetIpRetType, ok bool) {
	return getPublicIpGetIpAttributeTypeOk(o.Ip)
}

// HasIp returns a boolean if a field has been set.
func (o *PublicIp) HasIp() bool {
	_, ok := o.GetIpOk()
	return ok
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PublicIp) SetIp(v PublicIpGetIpRetType) {
	setPublicIpGetIpAttributeType(&o.Ip, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PublicIp) GetLabels() (res PublicIpGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetLabelsOk() (ret PublicIpGetLabelsRetType, ok bool) {
	return getPublicIpGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *PublicIp) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *PublicIp) SetLabels(v PublicIpGetLabelsRetType) {
	setPublicIpGetLabelsAttributeType(&o.Labels, v)
}

// GetNetworkInterface returns the NetworkInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIp) GetNetworkInterface() (res PublicIpGetNetworkInterfaceRetType) {
	res, _ = o.GetNetworkInterfaceOk()
	return
}

// GetNetworkInterfaceOk returns a tuple with the NetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIp) GetNetworkInterfaceOk() (ret PublicIpGetNetworkInterfaceRetType, ok bool) {
	return getPublicIpGetNetworkInterfaceAttributeTypeOk(o.NetworkInterface)
}

// HasNetworkInterface returns a boolean if a field has been set.
func (o *PublicIp) HasNetworkInterface() bool {
	_, ok := o.GetNetworkInterfaceOk()
	return ok
}

// SetNetworkInterface gets a reference to the given string and assigns it to the NetworkInterface field.
func (o *PublicIp) SetNetworkInterface(v PublicIpGetNetworkInterfaceRetType) {
	setPublicIpGetNetworkInterfaceAttributeType(&o.NetworkInterface, v)
}

// SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil
func (o *PublicIp) SetNetworkInterfaceNil() {
	o.NetworkInterface = nil
}

// UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
func (o *PublicIp) UnsetNetworkInterface() {
	o.NetworkInterface = nil
}

func (o PublicIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPublicIpGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getPublicIpGetIpAttributeTypeOk(o.Ip); ok {
		toSerialize["Ip"] = val
	}
	if val, ok := getPublicIpGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getPublicIpGetNetworkInterfaceAttributeTypeOk(o.NetworkInterface); ok {
		toSerialize["NetworkInterface"] = val
	}
	return toSerialize, nil
}

type NullablePublicIp struct {
	value *PublicIp
	isSet bool
}

func (v NullablePublicIp) Get() *PublicIp {
	return v.value
}

func (v *NullablePublicIp) Set(val *PublicIp) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIp) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIp(val *PublicIp) *NullablePublicIp {
	return &NullablePublicIp{value: val, isSet: true}
}

func (v NullablePublicIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
