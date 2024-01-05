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
        "/admin/apkList": {
            "post": {
                "security": [
                    {
                        "admintoken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "app列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "管理员登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.AdminLoginParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/admin/release": {
            "post": {
                "security": [
                    {
                        "admintoken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "发布app",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/admin/upload": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "上传apk文件",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/replace/latest": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "版本信息"
                ],
                "summary": "最新版本信息",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.LatestParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/service/sendMail": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "邮件"
                ],
                "summary": "发送邮件",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.SendMail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/bindMail": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "绑定邮箱",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.BindMail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/cancelFollow": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "取消关注",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.Follow"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/editInfo": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "编辑信息",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.EditInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/editPwd": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.EditPwd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/fansList": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "粉丝列表",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.FollowList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/follow": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "关注",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.Follow"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/followList": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "关注列表",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.FollowList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.LoginParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/loginOut": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "退出登录接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.RegisterParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/retrievePwd": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "找回密码",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.RetrievePwd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        },
        "/user/userInfo": {
            "post": {
                "security": [
                    {
                        "Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户详情",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Res"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.Res": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "state": {
                    "type": "boolean"
                }
            }
        },
        "request.AdminLoginParams": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.BindMail": {
            "type": "object",
            "required": [
                "code",
                "mail",
                "source"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                }
            }
        },
        "request.EditInfo": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.EditPwd": {
            "type": "object",
            "required": [
                "newPassword",
                "newRpassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "newRpassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.Follow": {
            "type": "object",
            "required": [
                "followedPerson"
            ],
            "properties": {
                "followedPerson": {
                    "type": "integer"
                }
            }
        },
        "request.FollowList": {
            "type": "object",
            "required": [
                "page",
                "pageSize"
            ],
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "request.LatestParams": {
            "type": "object",
            "required": [
                "client_version"
            ],
            "properties": {
                "client_version": {
                    "type": "string"
                }
            }
        },
        "request.LoginParams": {
            "type": "object",
            "required": [
                "account",
                "source",
                "typeString"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                },
                "typeString": {
                    "type": "string"
                }
            }
        },
        "request.RegisterParams": {
            "type": "object",
            "required": [
                "source",
                "typeString"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "rpassword": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                },
                "typeString": {
                    "type": "string"
                }
            }
        },
        "request.RetrievePwd": {
            "type": "object",
            "required": [
                "code",
                "mail",
                "newPassword",
                "newRpassword",
                "source"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                },
                "newRpassword": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                }
            }
        },
        "request.SendMail": {
            "type": "object",
            "required": [
                "mail",
                "type"
            ],
            "properties": {
                "mail": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Token": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8888",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "简单的社交系统",
	Description:      "dj的go接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
