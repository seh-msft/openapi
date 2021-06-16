# openapi

[![PkgGoDev](https://pkg.go.dev/badge/github.com/seh-msft/openapi)](https://pkg.go.dev/github.com/seh-msft/openapi)

[Go](https://golang.org) module for parsing OpenAPI v3 JSON files. 

## Build

	go build

## Install

	go install

## Test

None yet.

	go test

## Contributing

Please do. 

## Examples

Tools that use openapi:

- [inspector](https://github.com/seh-msft/inspector)
- [correlator](https://github.com/seh-msft/correlator)
- [generator](https://github.com/seh-msft/generator)

## What's missing

- Openapi has not been tested against a large number of Swagger JSON files
- Some things probably don't have to be `map` types, but are at the moment

## Documentation

Via `go doc --all`:

```
package openapi // import "github.com/seh-msft/openapi"

Package openapi is a specific-use-case data structure for OpenAPI v3 JSON
specification files. This package should not be considered an authoritative
implementation of the OpenAPI v3 JSON specification structure.

TYPES

type API struct {
	Version    string                       `json:"openapi"`    // OpenAPI semantic version
	Info       Info                         `json:"info"`       // Meta-information about the API
	Servers    []Server                     `json:"servers"`    // Servers the API may be accessible from
	Paths      map[string]map[string]Method `json:"paths"`      // Paths the API serves for callers
	Components map[string]map[string]Type   `json:"components"` // Types, etc. present within the API paths
}
    API represents an OpenAPI specification instance. This is the top-level
    type.

func Parse(r io.Reader) (API, error)
    Parse takes a io.Reader which provides an OpenAPI v3 JSON specification and
    deserializes to an API.

type Content map[string]map[string]Schema
    Content is the "content" structure within an HTTP request or response.

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}
    Info stores meta-information about the API.

type Item struct {
	// Enums is the enumerated values possible in the item, if any.
	Enums []string `json:"enum,omitempty"`

	// Type is the type of the item, if any.
	Type string `json:"type,omitempty"`

	// Ref is the reference identifier of the item, if any.
	Ref string `json:"$ref,omitempty"`
}
    Item represents an item in a set.

type Method struct {
	Tags        []string                       `json:"tags"`        // Tags (if any) for classifying the method
	Summary     string                         `json:"summary"`     // What does the method call provide/do?
	Description string                         `json:"description"` // ↑ ⊻ with Summary
	OperationID string                         `json:"operationId"` // Identifier for what is done
	Parameters  []Parameter                    `json:"parameters"`  // Parameters that the method may be called with
	Responses   map[string]Response            `json:"responses"`   // Expected responses for call in the form of `["HTTP code"]description`
	RequestBody `json:"requestBody,omitempty"` // Body of the Response, if any
}
    Method describes the calling information for an API Path.

type Parameter struct {
	Name        string          `json:"name"`        // Parameter name — ex. "accountId"
	In          string          `json:"in"`          // Where the parameter occurs in the HTTP call
	Description string          `json:"description"` // What does this parameter represent?
	Required    bool            `json:"required"`    // Is the parameter mandatory?
	Schema      `json:"schema"` // Describes the type and value scheme of a parameter
}
    Parameter describes how a given API parameter should be provided and valued.

type Property struct {
	Type     string `json:"type,omitempty"`
	Ref      string `json:"$ref,omitempty"`
	Items    Schema `json:"items,omitempty"`
	Format   string `json:"format,omitempty"`
	Nullable bool   `json:"nullable,omitempty"`

	Enums []string `json:"enum,omitempty"`
}
    Property is an entry in a map `["component"]{"properties"}` for a
    Type.Properties.

type RequestBody struct {
	Description string           `json:"description"` // What does the body represent
	Content     `json:"content"` // Contents of body
	Required    bool             `json:"required"` // Is the body mandatory?
}
    RequestBody represents the structure of a request body for HTTP methods such
    as POST.

type Response struct {
	Description string `json:"description"` // What the response provides

	// Content has the structure `[content-type]["schema"]Schema`.
	Content `json:"content"` // Contents of the response
}
    Response holds information about an HTTP response.

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
    Schema represents the scheme for a given item or object.

type Server struct {
	URL string `json:"url"`
}
    Server URL the API is called from.

type Type struct {
	Required []string `json:"required,omitempty"` // List of required, dependant, entries
	Is       string   `json:"type"`               // A value such as "object"

	// Properties has a structure similar to: `["SomeId"]{type, items}`
	Properties map[string]Property `json:"properties"`
}
    Type is a schema super type definition
```
