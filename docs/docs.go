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
        "/auth/refresh": {
            "post": {
                "description": "Обновление токенов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновление токенов",
                "operationId": "auth-refresh",
                "parameters": [
                    {
                        "description": "Рефреш токен",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        },
        "/auth/sign_in": {
            "post": {
                "description": "Вход для всех пользователей по логину и паролю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Вход пользователей",
                "operationId": "auth-sign-in",
                "parameters": [
                    {
                        "description": "Логин и пароль",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SignInOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        },
        "/auth/sign_up": {
            "post": {
                "description": "Регистрация пользователя по логину и паролю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "operationId": "auth-sign-up",
                "parameters": [
                    {
                        "description": "Логин, пароль, роль",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        },
        "/goods": {
            "get": {
                "description": "Возвращает список всех товаров",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Получение списка товаров",
                "operationId": "get-all-goods",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Good"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавление товара",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Добавление товара",
                "operationId": "add-good",
                "parameters": [
                    {
                        "description": "Товар",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Good"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        },
        "/goods/{good_id}": {
            "get": {
                "description": "Возвращает товар по его айди",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Получение товара по его айди",
                "operationId": "get-good-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Good"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновление данных о товаре",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Обновление данных о товаре",
                "operationId": "update-good",
                "parameters": [
                    {
                        "description": "Товар",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Good"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление товара по его айди",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Удаление товара по его айди",
                "operationId": "delete-good",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Role": {
            "type": "string",
            "enum": [
                "none",
                "admin",
                "user"
            ],
            "x-enum-varnames": [
                "NoneRole",
                "AdminRole",
                "UserRole"
            ]
        },
        "domain.Tokens": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "dto.Good": {
            "type": "object",
            "required": [
                "desc",
                "measure_unit",
                "name",
                "price"
            ],
            "properties": {
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "measure_unit": {
                    "type": "string",
                    "enum": [
                        "METER",
                        "KILOGRAM",
                        "LITER",
                        "PIECE"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock_quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.RefreshInput": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.SignInInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.SignInOutput": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.SignUpInput": {
            "type": "object",
            "required": [
                "password",
                "role",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Body": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.IDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
