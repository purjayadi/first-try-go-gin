// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Purjayadi",
            "url": "https://www.linkedin.com/in/purjayadi-9a154013a/",
            "email": "purjayadi@gmail.com"
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
        "/auth/login": {
            "post": {
                "description": "Login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "userInput",
                        "name": "userInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "userInput",
                        "name": "userInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/package": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Responds with the list of all packages as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Get packages array",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Package"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new package",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Create new package",
                "parameters": [
                    {
                        "description": "package",
                        "name": "package",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PackageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Package"
                        }
                    }
                }
            }
        },
        "/package/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Find package by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Get package by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Package"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update package",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Update package",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "package",
                        "name": "package",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePackageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Package"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete package by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Delete package by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Package"
                        }
                    }
                }
            }
        },
        "/resource": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Responds with the list of all resources as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "Get resources array",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Resource"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "Create new resource",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ResourceDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resource"
                        }
                    }
                }
            }
        },
        "/resource/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Find resource by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "Get resource by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resource"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "Delete resource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resource"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "Update resource",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateResourceDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resource"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Responds with the list of all users as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users array",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "userInput",
                        "name": "userInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/user/{email}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Find user by email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete user by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "userInput",
                        "name": "userInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserDto": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "dto.LoginDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "dto.PackageDto": {
            "type": "object",
            "required": [
                "description",
                "image",
                "name",
                "price"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 1
                },
                "image": {
                    "type": "string",
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "price": {
                    "type": "number",
                    "minimum": 1
                }
            }
        },
        "dto.ResourceDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "dto.UpdatePackageDto": {
            "type": "object",
            "required": [
                "description",
                "image",
                "name",
                "price"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 1
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string",
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "price": {
                    "type": "number",
                    "minimum": 1
                }
            }
        },
        "dto.UpdateResourceDto": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "dto.UpdateUserDto": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "model.Package": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "packageResource": {
                    "description": "package has many resource",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PackageResource"
                    }
                },
                "price": {
                    "type": "number"
                },
                "subscribes": {
                    "description": "package have many subscribe",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Subscribe"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.PackageResource": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "package": {
                    "$ref": "#/definitions/model.Package"
                },
                "package_id": {
                    "type": "string"
                },
                "resource": {
                    "$ref": "#/definitions/model.Resource"
                },
                "resource_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Resource": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "packageResource": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PackageResource"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Subscribe": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "expired_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "package": {
                    "$ref": "#/definitions/model.Package"
                },
                "package_id": {
                    "type": "string"
                },
                "payment_method": {
                    "type": "string"
                },
                "purchase_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "subscribes": {
                    "description": "user have many subscribe",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Subscribe"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gin SAS Service",
	Description:      "A SAS API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
