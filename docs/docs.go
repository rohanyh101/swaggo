// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "log in a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "user log in",
                "parameters": [
                    {
                        "description": "User log in request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"email or password is incorrect",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"email or password is incorrect",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "sign up a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "user sign up",
                "parameters": [
                    {
                        "description": "User sign up request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"email or password is incorrect",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"Error occurred while checking for email",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health check endpoint to see if the server is running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get a user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User profile",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"UnAuthenticated to access this resource",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error\": \"User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"Error occurred while fetching user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update a user profile",
                "parameters": [
                    {
                        "description": "User update request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"UnAuthenticated to access this resource",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error\": \"User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"Error occurred while updating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete a user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"UnAuthenticated to access this resource",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error\": \"User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"Error occurred while deleting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get all users (ADMIN privilege required)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "error\": \"no users available",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error\": \"UnAuthenticated to access this resource",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error\": \"Error occurred while listing users",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "role"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
