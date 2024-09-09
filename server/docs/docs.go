// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/category": {
            "get": {
                "description": "get categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "List categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Category"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create category",
                "parameters": [
                    {
                        "description": "Add category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "get": {
                "description": "get category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "update category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "description": "Update category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete category",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/category/{id}/mark": {
            "post": {
                "description": "mark category entries as read/unread up to a timestamp",
                "tags": [
                    "category"
                ],
                "summary": "Mark as read/unread",
                "parameters": [
                    {
                        "type": "string",
                        "format": "date-time",
                        "example": "\"2006-01-02 15:04:05\"",
                        "description": "Timestamp to mark/unread as read to",
                        "name": "before",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "read",
                            "unread"
                        ],
                        "type": "string",
                        "description": "New status",
                        "name": "as",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/category/{id}/refresh": {
            "post": {
                "description": "refresh all feeds in the category",
                "tags": [
                    "category"
                ],
                "summary": "Refresh feeds",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.AddCategory": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.UpdateCategory": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gopherss API",
	Description:      "RSS feed management app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
