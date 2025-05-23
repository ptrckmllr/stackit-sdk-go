/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the DatabaseOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseOptions{}

/*
	types and functions for collationName
*/

// isNotNullableString
type DatabaseOptionsGetCollationNameAttributeType = *string

func getDatabaseOptionsGetCollationNameAttributeTypeOk(arg DatabaseOptionsGetCollationNameAttributeType) (ret DatabaseOptionsGetCollationNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDatabaseOptionsGetCollationNameAttributeType(arg *DatabaseOptionsGetCollationNameAttributeType, val DatabaseOptionsGetCollationNameRetType) {
	*arg = &val
}

type DatabaseOptionsGetCollationNameArgType = string
type DatabaseOptionsGetCollationNameRetType = string

/*
	types and functions for compatibilityLevel
*/

// isLong
type DatabaseOptionsGetCompatibilityLevelAttributeType = *int64
type DatabaseOptionsGetCompatibilityLevelArgType = int64
type DatabaseOptionsGetCompatibilityLevelRetType = int64

func getDatabaseOptionsGetCompatibilityLevelAttributeTypeOk(arg DatabaseOptionsGetCompatibilityLevelAttributeType) (ret DatabaseOptionsGetCompatibilityLevelRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDatabaseOptionsGetCompatibilityLevelAttributeType(arg *DatabaseOptionsGetCompatibilityLevelAttributeType, val DatabaseOptionsGetCompatibilityLevelRetType) {
	*arg = &val
}

/*
	types and functions for owner
*/

// isNotNullableString
type DatabaseOptionsGetOwnerAttributeType = *string

func getDatabaseOptionsGetOwnerAttributeTypeOk(arg DatabaseOptionsGetOwnerAttributeType) (ret DatabaseOptionsGetOwnerRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDatabaseOptionsGetOwnerAttributeType(arg *DatabaseOptionsGetOwnerAttributeType, val DatabaseOptionsGetOwnerRetType) {
	*arg = &val
}

type DatabaseOptionsGetOwnerArgType = string
type DatabaseOptionsGetOwnerRetType = string

// DatabaseOptions struct for DatabaseOptions
type DatabaseOptions struct {
	// Name of the collation of the database
	CollationName DatabaseOptionsGetCollationNameAttributeType `json:"collationName,omitempty"`
	// CompatibilityLevel of the Database.
	CompatibilityLevel DatabaseOptionsGetCompatibilityLevelAttributeType `json:"compatibilityLevel,omitempty"`
	// Name of the owner of the database.
	Owner DatabaseOptionsGetOwnerAttributeType `json:"owner,omitempty"`
}

// NewDatabaseOptions instantiates a new DatabaseOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseOptions() *DatabaseOptions {
	this := DatabaseOptions{}
	return &this
}

// NewDatabaseOptionsWithDefaults instantiates a new DatabaseOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseOptionsWithDefaults() *DatabaseOptions {
	this := DatabaseOptions{}
	return &this
}

// GetCollationName returns the CollationName field value if set, zero value otherwise.
func (o *DatabaseOptions) GetCollationName() (res DatabaseOptionsGetCollationNameRetType) {
	res, _ = o.GetCollationNameOk()
	return
}

// GetCollationNameOk returns a tuple with the CollationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptions) GetCollationNameOk() (ret DatabaseOptionsGetCollationNameRetType, ok bool) {
	return getDatabaseOptionsGetCollationNameAttributeTypeOk(o.CollationName)
}

// HasCollationName returns a boolean if a field has been set.
func (o *DatabaseOptions) HasCollationName() bool {
	_, ok := o.GetCollationNameOk()
	return ok
}

// SetCollationName gets a reference to the given string and assigns it to the CollationName field.
func (o *DatabaseOptions) SetCollationName(v DatabaseOptionsGetCollationNameRetType) {
	setDatabaseOptionsGetCollationNameAttributeType(&o.CollationName, v)
}

// GetCompatibilityLevel returns the CompatibilityLevel field value if set, zero value otherwise.
func (o *DatabaseOptions) GetCompatibilityLevel() (res DatabaseOptionsGetCompatibilityLevelRetType) {
	res, _ = o.GetCompatibilityLevelOk()
	return
}

// GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptions) GetCompatibilityLevelOk() (ret DatabaseOptionsGetCompatibilityLevelRetType, ok bool) {
	return getDatabaseOptionsGetCompatibilityLevelAttributeTypeOk(o.CompatibilityLevel)
}

// HasCompatibilityLevel returns a boolean if a field has been set.
func (o *DatabaseOptions) HasCompatibilityLevel() bool {
	_, ok := o.GetCompatibilityLevelOk()
	return ok
}

// SetCompatibilityLevel gets a reference to the given int64 and assigns it to the CompatibilityLevel field.
func (o *DatabaseOptions) SetCompatibilityLevel(v DatabaseOptionsGetCompatibilityLevelRetType) {
	setDatabaseOptionsGetCompatibilityLevelAttributeType(&o.CompatibilityLevel, v)
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DatabaseOptions) GetOwner() (res DatabaseOptionsGetOwnerRetType) {
	res, _ = o.GetOwnerOk()
	return
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptions) GetOwnerOk() (ret DatabaseOptionsGetOwnerRetType, ok bool) {
	return getDatabaseOptionsGetOwnerAttributeTypeOk(o.Owner)
}

// HasOwner returns a boolean if a field has been set.
func (o *DatabaseOptions) HasOwner() bool {
	_, ok := o.GetOwnerOk()
	return ok
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DatabaseOptions) SetOwner(v DatabaseOptionsGetOwnerRetType) {
	setDatabaseOptionsGetOwnerAttributeType(&o.Owner, v)
}

func (o DatabaseOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getDatabaseOptionsGetCollationNameAttributeTypeOk(o.CollationName); ok {
		toSerialize["CollationName"] = val
	}
	if val, ok := getDatabaseOptionsGetCompatibilityLevelAttributeTypeOk(o.CompatibilityLevel); ok {
		toSerialize["CompatibilityLevel"] = val
	}
	if val, ok := getDatabaseOptionsGetOwnerAttributeTypeOk(o.Owner); ok {
		toSerialize["Owner"] = val
	}
	return toSerialize, nil
}

type NullableDatabaseOptions struct {
	value *DatabaseOptions
	isSet bool
}

func (v NullableDatabaseOptions) Get() *DatabaseOptions {
	return v.value
}

func (v *NullableDatabaseOptions) Set(val *DatabaseOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseOptions(val *DatabaseOptions) *NullableDatabaseOptions {
	return &NullableDatabaseOptions{value: val, isSet: true}
}

func (v NullableDatabaseOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
