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
        "/auth/change-password": {
            "post": {
                "description": "Меняет пароль по запросу пользователя",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Меняет пароль для пользователя",
                "parameters": [
                    {
                        "description": "Параметры смены пароля",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginInfo"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Проводит авторизацию на сервере и возвращает информацию о пользователе.",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизирует пользователя.",
                "parameters": [
                    {
                        "description": "Параметры аунтификации",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginInfo"
                        }
                    }
                }
            }
        },
        "/sample": {
            "get": {
                "description": "return sample data as json.",
                "produces": [
                    "application/json"
                ],
                "summary": "return sample data from server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Sample"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Sample": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "lorem": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "isRemember",
                "login",
                "password"
            ],
            "properties": {
                "isRemember": {
                    "type": "boolean"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "response.LoginInfo": {
            "type": "object",
            "required": [
                "token",
                "username"
            ],
            "properties": {
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "keeper cms server",
	Description:      "CMS с сервером на языке программирования GO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
