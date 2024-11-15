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
        "/api/v1/copy": {
            "post": {
                "description": "根据源存储桶和对象，将文件拷贝到目标存储桶和对象",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "拷贝文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "源存储桶名称",
                        "name": "srcBucket",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "源对象名称（文件名）",
                        "name": "srcObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "目标存储桶名称（可选）",
                        "name": "destBucket",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "目标对象名称（文件名）",
                        "name": "destObject",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件拷贝成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "缺少必需的参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件拷贝失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/delete": {
            "delete": {
                "description": "根据文件名删除文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件名",
                        "name": "objectName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件删除成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "object_Name 参数缺失",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件删除失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/download": {
            "get": {
                "description": "根据文件名下载文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "下载文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件名",
                        "name": "objectName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件数据流",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "object_Name 参数缺失",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件下载失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/list": {
            "get": {
                "description": "返回存储中的所有文件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "获取文件列表",
                "responses": {
                    "200": {
                        "description": "文件列表获取成功",
                        "schema": {
                            "$ref": "#/definitions/model.FileInfo"
                        }
                    },
                    "500": {
                        "description": "文件列表获取失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/logs": {
            "get": {
                "description": "根据条件查询日志",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "日志管理"
                ],
                "summary": "查询日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询条件",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "开始时间，Unix 时间戳",
                        "name": "startTime",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "结束时间，Unix 时间戳",
                        "name": "endTime",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "返回日志条数上限",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页起始位置",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "日志查询成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "日志查询失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/rename": {
            "post": {
                "description": "根据源对象名称将文件重命名为目标对象名称",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "重命名文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "源对象名称（当前文件名）",
                        "name": "srcObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "目标对象名称（新的文件名）",
                        "name": "destObject",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件重命名成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "缺少必需的参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件重命名失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/upload": {
            "post": {
                "description": "处理文件上传请求",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "上传成功",
                        "schema": {
                            "$ref": "#/definitions/model.UploadResponse"
                        }
                    },
                    "400": {
                        "description": "文件解析失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "上传失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v2/list": {
            "get": {
                "description": "获取指定目录下的文件列表，并支持分页查询，返回文件列表及下一页的分页标记（Continuation Token）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "获取文件列表（V2版本）",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件前缀",
                        "name": "prefix",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "分页标记，继续上次查询的位置",
                        "name": "continuationToken",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页返回的文件数，最大值为1000，默认为1000",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件列表获取成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid limit parameter",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "获取文件列表失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.FileInfo": {
            "type": "object",
            "properties": {
                "content-length": {
                    "type": "integer"
                },
                "etag": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "last_modified": {
                    "type": "string"
                }
            }
        },
        "model.UploadResponse": {
            "type": "object",
            "properties": {
                "content-length": {
                    "type": "integer"
                },
                "etag": {
                    "type": "string"
                },
                "last-modified": {
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
