// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bulky-vote": {
            "post": {
                "description": "you can vote with id or name (every porduct that matches that conditions will consider your point)",
                "tags": [
                    "bulkyVote"
                ],
                "summary": "This is for voting",
                "operationId": "bulkyVote",
                "parameters": [
                    {
                        "description": "The body to create a vote",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.product"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/filter": {
            "post": {
                "description": "you can get product with id or name",
                "tags": [
                    "filter"
                ],
                "summary": "This is for filtering",
                "operationId": "filter",
                "parameters": [
                    {
                        "description": "filter products",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.product"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/get-all": {
            "get": {
                "description": "you can vote with id or name (every porduct that matches that conditions will consider your point)",
                "tags": [
                    "get-all"
                ],
                "summary": "This is for voting",
                "operationId": "get-all",
                "responses": {}
            }
        },
        "/vote": {
            "post": {
                "description": "you can vote with id or name (every porduct that matches that conditions will consider your point)",
                "tags": [
                    "vote"
                ],
                "summary": "This is for voting",
                "operationId": "vote",
                "parameters": [
                    {
                        "description": "The body to create a vote",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.product"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "api.product": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "example: 1\nrequired: false",
                    "type": "integer"
                },
                "name": {
                    "description": "The name of product\nexample: TV\nrequired: false",
                    "type": "string"
                },
                "point": {
                    "description": "The value for a vote point\nexample: 6.5\nrequired: true",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
