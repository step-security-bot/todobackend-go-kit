/*
 * Todo API
 *
 * The Todo API manages a list of todo items as described by the TodoMVC backend project: http://todobackend.com 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package todov1

import (
	"encoding/json"
)

// UpdateTodoItemRequest struct for UpdateTodoItemRequest
type UpdateTodoItemRequest struct {
	Title NullableString `json:"title,omitempty"`
	Completed NullableBool `json:"completed,omitempty"`
	Order NullableInt32 `json:"order,omitempty"`
}

// NewUpdateTodoItemRequest instantiates a new UpdateTodoItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTodoItemRequest() *UpdateTodoItemRequest {
	this := UpdateTodoItemRequest{}
	return &this
}

// NewUpdateTodoItemRequestWithDefaults instantiates a new UpdateTodoItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTodoItemRequestWithDefaults() *UpdateTodoItemRequest {
	this := UpdateTodoItemRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTodoItemRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTodoItemRequest) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateTodoItemRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UpdateTodoItemRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UpdateTodoItemRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UpdateTodoItemRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTodoItemRequest) GetCompleted() bool {
	if o == nil || o.Completed.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTodoItemRequest) GetCompletedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *UpdateTodoItemRequest) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableBool and assigns it to the Completed field.
func (o *UpdateTodoItemRequest) SetCompleted(v bool) {
	o.Completed.Set(&v)
}
// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *UpdateTodoItemRequest) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *UpdateTodoItemRequest) UnsetCompleted() {
	o.Completed.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTodoItemRequest) GetOrder() int32 {
	if o == nil || o.Order.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTodoItemRequest) GetOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *UpdateTodoItemRequest) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *UpdateTodoItemRequest) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *UpdateTodoItemRequest) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *UpdateTodoItemRequest) UnsetOrder() {
	o.Order.Unset()
}

func (o UpdateTodoItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTodoItemRequest struct {
	value *UpdateTodoItemRequest
	isSet bool
}

func (v NullableUpdateTodoItemRequest) Get() *UpdateTodoItemRequest {
	return v.value
}

func (v *NullableUpdateTodoItemRequest) Set(val *UpdateTodoItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTodoItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTodoItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTodoItemRequest(val *UpdateTodoItemRequest) *NullableUpdateTodoItemRequest {
	return &NullableUpdateTodoItemRequest{value: val, isSet: true}
}

func (v NullableUpdateTodoItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTodoItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


