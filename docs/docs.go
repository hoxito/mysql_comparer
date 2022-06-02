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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "jose aranciba",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/db/simple/diff": {
            "get": {
                "description": "differenciates databases",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getdiff"
                ],
                "summary": "Gets the difference between 2 databases",
                "parameters": [
                    {
                        "maxLength": 10,
                        "minLength": 5,
                        "type": "string",
                        "description": "string valid",
                        "name": "string",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "int valid",
                        "name": "int",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TableDiff"
                            }
                        }
                    }
                }
            }
        },
        "/v1/db/simple/diff/notables": {
            "get": {
                "description": "differenciates databases",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getdiff"
                ],
                "summary": "Gets the difference between 2 databases without showing every table in both dbs",
                "parameters": [
                    {
                        "maxLength": 10,
                        "minLength": 5,
                        "type": "string",
                        "description": "string valid",
                        "name": "string",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "int valid",
                        "name": "int",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TableDiff"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.TableDiff": {
            "type": "object",
            "properties": {
                "db1": {
                    "type": "string"
                },
                "db2": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "script1": {
                    "type": "string"
                },
                "script2": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3010",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Diff API",
	Description:      "This is a database differenciator API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
