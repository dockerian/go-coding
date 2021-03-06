/*
 * Go-coding API
 *
 * Go-coding API Gateway.
 *
 * OpenAPI spec version: 0.0.1
 * Contact: jason.zhuy@gmail.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package client

// APISchema provides detailed info for current API spec.
type ApiSchema struct {

	// Go-coding API schema version.
	Version string `json:"version,omitempty"`

	// Go-coding API description.
	Description string `json:"description,omitempty"`

	// Go-coding API endpoint URL.
	EndpointURL string `json:"endpointURL,omitempty"`

	// Swagger yaml file.
	SwaggerYaml string `json:"swaggerYaml,omitempty"`

	DbInfo DbSchema `json:"dbInfo,omitempty"`
}
