/*
 * Todo API
 *
 * The Todo API manages a list of todo items as described by the TodoMVC backend project: http://todobackend.com 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package todov1
// Error struct for Error
type Error struct {
	Type string `json:"type"`
	Title string `json:"title,omitempty"`
	Status int32 `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}
