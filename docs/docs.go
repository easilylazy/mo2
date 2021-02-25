// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/accounts": {
            "post": {
                "description": "add by json account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Add an account",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            }
        },
        "/api/accounts/addUser": {
            "post": {
                "description": "为新用户创建信息，加入数据库",
                "produces": [
                    "application/json"
                ],
                "summary": "新增用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/accounts/login": {
            "post": {
                "description": "login by json model.LoginAccount and set cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "login an account",
                "parameters": [
                    {
                        "description": "login account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginUserInfo"
                        }
                    }
                }
            }
        },
        "/api/accounts/logout": {
            "get": {
                "description": "logout and delete cookies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "logout",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/accounts/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a account",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/addBlogs2Categories": {
            "post": {
                "description": "blogs 与 categories皆为id列表，方便批量操作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "add blogs to chosen categories",
                "parameters": [
                    {
                        "description": "dto.AddBlogs2Categories",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddBlogs2Categories"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.QueryBlog"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/addCategory": {
            "post": {
                "description": "add by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "upsert category",
                "parameters": [
                    {
                        "description": "Add category",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    }
                }
            }
        },
        "/api/blogs/addCategory2Category": {
            "post": {
                "description": "category为model.Category(若id存在，直接存放；否则新建) parent category 为id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "add category to parent category",
                "parameters": [
                    {
                        "description": "category info and parent id",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddCategory2Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    }
                }
            }
        },
        "/api/blogs/addCategory2User": {
            "post": {
                "description": "user 与 category 皆为id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "add category to user",
                "parameters": [
                    {
                        "description": "category id and user id",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddCategory2User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AddCategory2User"
                        }
                    }
                }
            }
        },
        "/api/blogs/find/id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "find Blog by id",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "bool true",
                        "name": "draft",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string xxxxxxxx",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Blog"
                        }
                    }
                }
            }
        },
        "/api/blogs/find/own": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "find Blog",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "bool true",
                        "name": "draft",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 0",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 5",
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
                                "$ref": "#/definitions/dto.QueryBlogs"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/find/userId": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "find Blog by userid",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "bool true",
                        "name": "draft",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string xxxxxxxx",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 0",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 5",
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
                                "$ref": "#/definitions/dto.QueryBlogs"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/findAllCategories": {
            "get": {
                "description": "若id为空，返回所有categories；若id不为空，返回该id的category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "find categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string ObjectID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Category"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/findCategoriesByUserId": {
            "get": {
                "description": "return (main category)个人的主存档 于前端不可见，用于后端存储",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "find categories by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string ObjectID",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.Category"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/findCategoryByUserId": {
            "get": {
                "description": "return (main category)个人的主存档 于前端不可见，用于后端存储",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "find category by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string ObjectID",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Category"
                        }
                    }
                }
            }
        },
        "/api/blogs/publish": {
            "post": {
                "description": "add by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "Publish Blog",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "bool true",
                        "name": "draft",
                        "in": "query"
                    },
                    {
                        "description": "Add blog",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Blog"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Blog"
                        }
                    }
                }
            }
        },
        "/api/blogs/query": {
            "get": {
                "description": "find",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "find all Blogs",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "bool true",
                        "name": "draft",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 0",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int 5",
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
                                "$ref": "#/definitions/dto.QueryBlog"
                            }
                        }
                    }
                }
            }
        },
        "/api/img/{filename}": {
            "get": {
                "description": "add by json",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "img"
                ],
                "summary": "Gen img token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file name",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ImgUploadToken"
                        }
                    }
                }
            }
        },
        "/api/logs": {
            "get": {
                "description": "get by check cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginUserInfo"
                        }
                    }
                }
            }
        },
        "/sayHello": {
            "get": {
                "description": "say something",
                "produces": [
                    "application/json"
                ],
                "summary": "simple test",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddBlogs2Categories": {
            "type": "object",
            "properties": {
                "blog_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "category_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.AddCategory2Category": {
            "type": "object",
            "properties": {
                "category_id": {
                    "$ref": "#/definitions/model.Category"
                },
                "parent_id": {
                    "type": "string",
                    "example": "xxxxxxx"
                }
            }
        },
        "dto.AddCategory2User": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string",
                    "example": "xxxxxxx"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.ImgUploadToken": {
            "type": "object",
            "properties": {
                "file_key": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.LoginUserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@qq.com"
                },
                "id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "name": {
                    "type": "string",
                    "example": "account name"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "ordinaryUser"
                    ]
                }
            }
        },
        "dto.QueryBlog": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Category"
                    }
                },
                "cover": {
                    "type": "string",
                    "example": "https://xxx/xxx"
                },
                "description": {
                    "type": "string",
                    "example": "mouse ❤ monkey"
                },
                "entityInfo": {
                    "$ref": "#/definitions/model.Entity"
                },
                "id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "keyWords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "xxx",
                        "xxx"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "mouse ❤ monkey"
                }
            }
        },
        "dto.QueryBlogs": {
            "type": "object"
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@qq.com"
                },
                "entityInfo": {
                    "$ref": "#/definitions/model.Entity"
                },
                "hashedPassword": {
                    "type": "string",
                    "example": "$2a$10$rXMPcOyfgdU6y5n3pkYQAukc3avJE9CLsx1v0Kn99GKV1NpREvN2i"
                },
                "id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "infos": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "'avatar'": " 'www.avatar.com/account_name'",
                        "'site'": "'www.limfx.com'"
                    }
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "ordinaryUser"
                    ]
                },
                "userName": {
                    "type": "string",
                    "example": "account name"
                }
            }
        },
        "model.AddAccount": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@qq.com"
                },
                "password": {
                    "type": "string",
                    "example": "p@ssword"
                },
                "userName": {
                    "type": "string",
                    "example": "account name"
                }
            }
        },
        "model.Blog": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Category"
                    }
                },
                "content": {
                    "type": "string",
                    "example": "xxxx\nxxxx"
                },
                "cover": {
                    "type": "string",
                    "example": "https://xxx/xxx"
                },
                "description": {
                    "type": "string",
                    "example": "mouse ❤ monkey"
                },
                "entityInfo": {
                    "$ref": "#/definitions/model.Entity"
                },
                "id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxx=="
                },
                "keyWords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "xxx",
                        "xxx"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "mouse ❤ monkey"
                }
            }
        },
        "model.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxxx=="
                },
                "name": {
                    "type": "string",
                    "example": "records"
                },
                "parent_id": {
                    "type": "string",
                    "example": "xxxxxxxxxxxxxx=="
                }
            }
        },
        "model.Entity": {
            "type": "object",
            "properties": {
                "createTime": {
                    "type": "string",
                    "example": "2020-10-1"
                },
                "updateTime": {
                    "type": "string",
                    "example": "2020-10-1"
                }
            }
        },
        "model.LoginAccount": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "p@ssword"
                },
                "userNameOrEmail": {
                    "type": "string",
                    "example": "account name/email@qq.com"
                }
            }
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
	Host:        "localhost:5000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Mo2",
	Description: "This is a Mo2 server.",
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