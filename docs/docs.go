// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-01-31 18:00:40.7676058 +0800 CST m=+0.085947701

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "117503445",
            "url": "https://github.com/117503445",
            "email": "t117503445@gmail.com"
        },
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/TGclub/Wizz-Home-Page/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "更改请求中的 Username 和 Password 进行登录。登陆成功以后，返回json中token字段比如说是\"token\":\"123\"，就在右上角Authorize按钮点一下，输入Bearer 123，大小写、空格敏感。然后就能使用需要身份验证的接口啦。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "身份验证"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录值",
                        "name": "loginVals",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"expire\":\"2020-02-05T23:11:41+08:00\",\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5MTU1MDEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDMxMDcwMX0.GWlmyTfCkXQYwgbtuTgVSTUSJXDcoDb_bptgRpt4HCU\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/members": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "获取所有成员",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "获取所有产品",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        },
        "/stories": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "获取所有历史事件",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Story"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "添加一个历史事件",
                "parameters": [
                    {
                        "description": "历史事件",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.PostStory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    }
                }
            }
        },
        "/stories/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "获取一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    },
                    "404": {
                        "description": "{\"message\":\"Story not found\"}",
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "更改一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "历史事件",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.PostStory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Story not found\"}",
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "删除一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"delete success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Story not found\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpModels.PostStory": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "为之诞生"
                },
                "storyDescribe": {
                    "type": "string",
                    "example": "11月,前身 TgClub 诞生"
                },
                "timeStamp": {
                    "type": "integer",
                    "example": 1580397149
                }
            }
        },
        "models.Member": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "个人简介",
                    "type": "string",
                    "example": "超级帅的前端dalao"
                },
                "id": {
                    "description": "member's id",
                    "type": "integer",
                    "example": 1
                },
                "memberType": {
                    "description": "成员类型,0 - 学生,1 - 导师",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ],
                    "example": 0
                },
                "name": {
                    "description": "成员姓名",
                    "type": "string",
                    "example": "tengye"
                },
                "omitempty": {
                    "description": "如果是老师,则有若干头衔",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "schoolYear": {
                    "description": "入学年份",
                    "type": "integer",
                    "example": 2017
                },
                "urlAvatar": {
                    "description": "成员头像图片的url",
                    "type": "string",
                    "example": "http://wuygewfuyd/weiug.jpg"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "项目介绍",
                    "type": "string",
                    "example": "大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天\u0026游戏两不误~"
                },
                "id": {
                    "description": "product's id",
                    "type": "integer",
                    "example": 1
                },
                "idBackGroundImg": {
                    "description": "项目背景图片的 Id",
                    "type": "integer"
                },
                "imgAvatar": {
                    "description": "项目图标的 Base64 字符串",
                    "type": "string"
                },
                "imgPartnerLogo": {
                    "description": "合作方Logo的 Base64 字符串",
                    "type": "string"
                },
                "imgProCode": {
                    "description": "二维码的 Base64 字符串",
                    "type": "string"
                },
                "imgScreenshot": {
                    "description": "项目截图的 Base64 字符串",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "littleDescribe": {
                    "description": "一句话介绍",
                    "type": "string",
                    "example": "一键约游戏"
                },
                "name": {
                    "description": "产品名称",
                    "type": "string",
                    "example": "一起来开黑"
                },
                "partner": {
                    "description": "合作方",
                    "type": "string",
                    "example": "腾讯"
                },
                "projectType": {
                    "description": "项目类型,0 - 校企合作,1 - 校园合作,2 - 校内自研",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ]
                },
                "urlAvatar": {
                    "description": "项目图标的 Url",
                    "type": "string"
                },
                "urlBackground": {
                    "description": "项目背景图片的 Url",
                    "type": "string"
                },
                "urlPartnerLogo": {
                    "description": "合作方Logo的 Url",
                    "type": "string"
                },
                "urlProCode": {
                    "description": "二维码的 Url",
                    "type": "string"
                },
                "urlScreenshot": {
                    "description": "项目截图的 Url",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Story": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID is story's id",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "为之诞生"
                },
                "storyDescribe": {
                    "type": "string",
                    "example": "11月,前身 TgClub 诞生"
                },
                "timeStamp": {
                    "type": "integer",
                    "example": 1580397149
                }
            }
        },
        "route.login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "ali.117503445.top:8080",
	BasePath:    "/api",
	Schemes:     []string{"http"},
	Title:       "Wizz-Home-Page API",
	Description: "Wizz's HomePage Backend",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
