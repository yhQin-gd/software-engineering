// Package docs1 Code generated by swaggo/swag. DO NOT EDIT
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
        "/generate": {
            "post": {
                "description": "根据传入的参数生成图片并返回图片的URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片生成"
                ],
                "summary": "生成图片",
                "parameters": [
                    {
                        "description": "图片参数",
                        "name": "imageParaments",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/generate_s.ImageParaments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "tags": [
                    "登录注册"
                ],
                "summary": "用户注册",
                "description": "创建一个新的用户。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "用户注册信息",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserRegister"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "用户注册成功",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string",
                                    "example": "用户注册成功"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "请求数据无效",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string",
                                    "example": "请求数据无效"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "登录注册"
                ],
                "summary": "用户登录",
                "description": "验证用户身份并返回成功信息。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "用户登录信息",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string",
                                    "example": "登录成功"
                                }
                            }
                        }
                    },
                    "401": {
                        "description": "未授权 - 用户不存在或密码错误",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string",
                                    "example": "用户不存在"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "generate_s.ImageParaments": {
            "type": "object",
            "required": [
                "height",
                "prompt",
                "sampling_method",
                "seed",
                "steps",
                "width"
            ],
            "properties": {
                "height": {
                    "type": "integer",
                    "maximum": 1024,
                    "minimum": 128
                },
                "prompt": {
                    "type": "string"
                },
                "sampling_method": {
                    "type": "string",
                    "enum": [
                        "DDIM",
                        "PLMS",
                        "K-LMS"
                    ]
                },
                "seed": {
                    "type": "string"
                },
                "steps": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "width": {
                    "type": "integer",
                    "maximum": 1024,
                    "minimum": 128
                }
            }
        },
        "UserRegister": {
            "type": "object",
            "required": [
                "email",
                "user_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "description": "用户的邮箱",
                    "example": "user@example.com"
                },
                "user_name": {
                    "type": "string",
                    "description": "用户的昵称",
                    "example": "张三"
                },
                "password": {
                    "type": "string",
                    "description": "用户的密码",
                    "example": "mypassword"
                }
            }
        },
        "UserLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "description": "用户的邮箱",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "description": "用户的密码",
                    "example": "mypassword"
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
