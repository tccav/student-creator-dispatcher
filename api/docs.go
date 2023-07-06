// Code generated by swaggo/swag. DO NOT EDIT.

package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/tccav",
            "email": "pedroyremolo@gmail.com"
        },
        "license": {
            "name": "No License",
            "url": "https://choosealicense.com/no-permission/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthcheck": {
            "get": {
                "tags": [
                    "Internal"
                ],
                "summary": "Check if service is healthy",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/students/students": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Register a student",
                "parameters": [
                    {
                        "description": "Student creation information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.StudentRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.StudentRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/students/students/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate a student",
                "parameters": [
                    {
                        "description": "Student credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.AuthenticateStudentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.AuthenticateStudentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/students/students/verify-auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verifies if Student Authentication is valid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_gateways_httpserver.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pkg_gateways_httpserver.AuthenticateStudentRequest": {
            "type": "object",
            "properties": {
                "secret": {
                    "type": "string",
                    "example": "celacanto-provoca-maremoto"
                },
                "student_id": {
                    "type": "string",
                    "example": "201210204310"
                }
            }
        },
        "pkg_gateways_httpserver.AuthenticateStudentResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string",
                    "format": "datetime",
                    "example": "2023-10-18T19:32:00.000Z"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
                },
                "token_id": {
                    "type": "string",
                    "format": "uuidv4",
                    "example": "1f6a4d3a-38c7-43fe-9790-2408fe595c93"
                }
            }
        },
        "pkg_gateways_httpserver.HTTPError": {
            "type": "object",
            "properties": {
                "err_code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "pkg_gateways_httpserver.StudentRegisterRequest": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "format": "date",
                    "example": "1990-10-18"
                },
                "course_id": {
                    "type": "string",
                    "format": "uuidv4",
                    "example": "1f6a4d3a-38c7-43fe-9790-2408fe595c93"
                },
                "cpf": {
                    "type": "string",
                    "example": "11111111030"
                },
                "email": {
                    "type": "string",
                    "format": "email",
                    "example": "jdoe@ol.com"
                },
                "id": {
                    "type": "string",
                    "example": "201210204310"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "secret": {
                    "type": "string",
                    "example": "celacanto provoca maremoto"
                }
            }
        },
        "pkg_gateways_httpserver.StudentRegisterResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "201210204310"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Identity Service API",
	Description:      "Service responsible for identity management of the Aluno Online's system.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
