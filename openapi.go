// Copyright (c) 2021, Microsoft Corporation, Sean Hinchee
// Licensed under the MIT License.

// Package openapi is a specific-use-case data structure for OpenAPI v3 JSON specification files.
// This package should not be considered an authoritative implementation of the OpenAPI v3 JSON specification structure.
package openapi

import (
	"bufio"
	"encoding/json"
	"io"
)

// API represents an OpenAPI specification instance.
// This is the top-level type.
type API struct {
	Version    string                       `json:"openapi"`    // OpenAPI semantic version
	Info       Info                         `json:"info"`       // Meta-information about the API
	Servers    []Server                     `json:"servers"`    // Servers the API may be accessible from
	Paths      map[string]map[string]Method `json:"paths"`      // Paths the API serves for callers
	Components map[string]map[string]Type   `json:"components"` // Types, etc. present within the API paths
}

// Type is a schema super type definition
type Type struct {
	Required []string `json:"required,omitempty"` // List of required, dependant, entries
	Is       string   `json:"type"`               // A value such as "object"

	// Properties has a structure similar to: `["SomeId"]{type, items}`
	Properties map[string]Property `json:"properties"`
	/* Structure:
	"properties" {
		architectures
		{
			type: array
			items {
				enum []string
				type string
			}
		}
		platformName
		{
			string
		}
		minVersion {
			string
		}
		maxVersion {
			string
		}
	*/
}

// Property is an entry in a map `["component"]{"properties"}` for a Type.Properties.
type Property struct {
	Type     string `json:"type,omitempty"`
	Ref      string `json:"$ref,omitempty"`
	Items    Schema `json:"items,omitempty"`
	Format   string `json:"format,omitempty"`
	Nullable bool   `json:"nullable,omitempty"`

	Enums []string `json:"enum,omitempty"`
}

// Schema represents the scheme for a given item or object.
type Schema struct {
	// Enums is the enumerated values possible in the scheme, if any.
	Enums []string `json:"enum,omitempty"`

	// Items, if empty, indicates the scheme is not that of an array.
	Items Item `json:"items,omitempty"` // Items expected in an array(?)

	// Type, if empty, is not an array.
	Type string `json:"type,omitempty"` // Type expected for input

	// Ref's value, if omitted, is probably in Property.Items["$ref"].
	Ref string `json:"$ref,omitempty"` // Reference path

	// Default is the default value of the scheme.
	Default string `json:"default,omitempty"`
}

// Item represents an item in a set.
type Item struct {
	// Enums is the enumerated values possible in the item, if any.
	Enums []string `json:"enum,omitempty"`

	// Type is the type of the item, if any.
	Type string `json:"type,omitempty"`

	// Ref is the reference identifier of the item, if any.
	Ref string `json:"$ref,omitempty"`
}

// Info stores meta-information about the API.
type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

// Server URL the API is called from.
type Server struct {
	URL string `json:"url"`
}

// Method describes the calling information for an API Path.
type Method struct {
	Tags        []string                       `json:"tags"`        // Tags (if any) for classifying the method
	Summary     string                         `json:"summary"`     // What does the method call provide/do?
	Description string                         `json:"description"` // ↑ ⊻ with Summary
	OperationID string                         `json:"operationId"` // Identifier for what is done
	Parameters  []Parameter                    `json:"parameters"`  // Parameters that the method may be called with
	Responses   map[string]Response            `json:"responses"`   // Expected responses for call in the form of `["HTTP code"]description`
	RequestBody `json:"requestBody,omitempty"` // Body of the Response, if any
}

// Content is the "content" structure within an HTTP request or response.
type Content map[string]map[string]Schema

// RequestBody represents the structure of a request body for HTTP methods such as POST.
type RequestBody struct {
	Description string           `json:"description"` // What does the body represent
	Content     `json:"content"` // Contents of body
	Required    bool             `json:"required"` // Is the body mandatory?
}

// Parameter describes how a given API parameter should be provided and valued.
type Parameter struct {
	Name        string          `json:"name"`        // Parameter name — ex. "accountId"
	In          string          `json:"in"`          // Where the parameter occurs in the HTTP call
	Description string          `json:"description"` // What does this parameter represent?
	Required    bool            `json:"required"`    // Is the parameter mandatory?
	Schema      `json:"schema"` // Describes the type and value scheme of a parameter
}

// Response holds information about an HTTP response.
type Response struct {
	Description string `json:"description"` // What the response provides

	// Content has the structure `[content-type]["schema"]Schema`.
	Content `json:"content"` // Contents of the response
}

// Parse takes a io.Reader which provides an OpenAPI v3 JSON specification and deserializes to an API.
func Parse(r io.Reader) (API, error) {
	br := bufio.NewReader(r)

	var api API

	dec := json.NewDecoder(br)
	err := dec.Decode(&api)

	return api, err
}
