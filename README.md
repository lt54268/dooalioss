# 阿里云OSS
运行：go run main.go

## 一、上传接口（POST）
http://127.0.0.1:3030/api/v1/upload

Body：file

返回示例：
```
{
    "code": 200,
    "data": {
        "content-length": 670197,
        "etag": "\"FDCXXXXXXXXXXXXXXXXXXXXXXXXXX\"",
        "last-modified": "2024-11-08T01:59:32Z"
    },
    "msg": "上传成功"
}
```

## 二、下载接口（GET）
http://127.0.0.1:3030/api/v1/download

参数：objectName

返回示例：返回文件，浏览器自动跳转下载

## 三、删除接口（DELETE）
http://127.0.0.1:3030/api/v1/delete

参数：objectName

返回示例：
```
{
    "code": 200,
    "msg": "文件删除成功"
}
```

## 四、获取文件列表接口（GET）
http://127.0.0.1:3030/api/v1/list

参数：无

返回示例：
```
{
    "code": 200,
    "data": [
        {
            "key": "10.14会议纪要.docx",
            "content-length": 13769,
            "etag": "\"850XXXXXXXXXXXXXXXXXXXXXXXXXXXX\"",
            "last_modified": "2024-11-07T07:55:07Z"
        },
        {
            "key": "test1/",
            "content-length": 0,
            "etag": "\"D41XXXXXXXXXXXXXXXXXXXXXXXXXXXX\"",
            "last_modified": "2024-11-07T08:34:18Z"
        }
    ],
    "msg": "文件列表获取成功"
}
```

### 说明：
阿里云OSS对象存储，上传同名文件会自动覆盖旧文件
