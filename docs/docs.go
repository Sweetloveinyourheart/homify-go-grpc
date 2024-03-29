// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
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
        "/assets": {
            "get": {
                "description": "Get assets based on the provided asset type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "Get assets by asset type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Asset type to filter by",
                        "name": "asset_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The assets list",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Creates a new asset based on the provided JSON payload.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "Create a new asset",
                "operationId": "addNewAsset",
                "parameters": [
                    {
                        "description": "JSON payload containing data for the new asset",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.NewAssets"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The asset was created successfully",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Failed to create the new asset due to validation errors or other issues",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/assets/disable": {
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Disables an existing asset based on the provided asset_id and asset_type.",
                "tags": [
                    "Assets"
                ],
                "summary": "Disable an existing asset",
                "operationId": "disableAnAsset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Asset ID to be disabled",
                        "name": "asset_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Type of the asset to be disabled",
                        "name": "asset_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The asset was disabled successfully",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Failed to disable the asset due to validation errors or other issues",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/assets/modify": {
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Modifies an existing asset based on the provided JSON payload.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "Modify an existing asset",
                "operationId": "modifyExistingAsset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Asset id to modify",
                        "name": "asset_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "JSON payload containing data for modifying the asset",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ModifyAssets"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The asset was updated successfully",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Failed to update the asset due to validation errors or other issues",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns a simple health check response indicating the status of the service.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health check"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.HeathCheckResponse"
                        }
                    }
                }
            }
        },
        "/property": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Add a new property to the property listings.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Property"
                ],
                "summary": "Add a new property",
                "operationId": "add-new-property",
                "parameters": [
                    {
                        "description": "New Property object to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.NewPropertyDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added the new property",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/property/sync": {
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Sync all the property to ES",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Property"
                ],
                "summary": "Sync all the property to ES",
                "operationId": "sync-properties",
                "responses": {
                    "200": {
                        "description": "Successfully",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "Handles the user sign-in process by validating input and authenticating the user via gRPC.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Handles user sign-in",
                "parameters": [
                    {
                        "description": "User sign-in data in JSON format",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.SignInDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "This endpoint allows users to sign up and create a new account. It expects a JSON payload containing user information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a new user account",
                "operationId": "signUp",
                "parameters": [
                    {
                        "description": "User information for registration",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.SignUpDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/proto.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Get the user profile for the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get user profile",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ModifyAssets": {
            "type": "object",
            "required": [
                "asset_type"
            ],
            "properties": {
                "asset_type": {
                    "type": "string",
                    "enum": [
                        "category",
                        "amenity"
                    ]
                },
                "icon_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.NewAssets": {
            "type": "object",
            "required": [
                "asset_type",
                "icon_url",
                "name"
            ],
            "properties": {
                "asset_type": {
                    "type": "string",
                    "enum": [
                        "category",
                        "amenity"
                    ]
                },
                "icon_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.NewDestinationDTO": {
            "type": "object",
            "required": [
                "city",
                "country",
                "lat",
                "long"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                },
                "country": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                },
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                }
            }
        },
        "dtos.NewPropertyDTO": {
            "type": "object",
            "required": [
                "amenity_id",
                "category_id",
                "description",
                "destination",
                "price",
                "title"
            ],
            "properties": {
                "amenity_id": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string",
                    "minLength": 3
                },
                "destination": {
                    "$ref": "#/definitions/dtos.NewDestinationDTO"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                }
            }
        },
        "dtos.SignInDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.SignUpDTO": {
            "type": "object",
            "required": [
                "email",
                "fullName",
                "password"
            ],
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "Male",
                        "Female"
                    ]
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "handlers.HeathCheckResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "proto.SignInResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "proto.SignUpResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Homify API",
	Description:      "This is a sample Swagger API for a Go Gin application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
