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
        "/collection/": {
            "get": {
                "description": "Возвращает список всех коллекций на сервере.",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "collection"
                ],
                "summary": "Возвращает список всех коллекций на сервере",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CollectionListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Создаёт новую коллекцию на сервере.",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "collection"
                ],
                "summary": "Создаёт коллекцию на сервере.",
                "parameters": [
                    {
                        "description": "Параметры создания коллекции",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateCollectionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "patch": {
                "description": "Редактирует информацию о коллекции на сервере.",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "collection"
                ],
                "summary": "Редактирует коллекцию.",
                "parameters": [
                    {
                        "description": "Параметры создания коллекции",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateCollectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/collection/{id}": {
            "get": {
                "description": "Возвращает коллекцию на сервере, по id.",
                "produces": [
                    "aplication/json"
                ],
                "tags": [
                    "collection"
                ],
                "summary": "Возвращает коллекцию на сервере.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор коллекции",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.GetCollectionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateCollectionRequest": {
            "type": "object"
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
        "request.UpdateCollectionRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.CollectionListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CollectionValueResponse"
                    }
                }
            }
        },
        "response.CollectionValueResponse": {
            "type": "object",
            "properties": {
                "allowWrite": {
                    "type": "boolean"
                },
                "apiName": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.GetCollectionResponse": {
            "type": "object"
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
	Title:            "MinoraCMS server",
	Description:      "CMS с сервером на языке программирования GO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
