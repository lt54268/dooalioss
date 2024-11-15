# 阿里云OSS
运行：go run main.go

API文档：http://127.0.0.1:3030/swagger/index.html

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
### V1版本
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

### V2版本
http://127.0.0.1:3030/api/v2/list

参数（可选）：
prefix（返回的文件前缀，留空默认全部返回）

continuationToken（游标，列举时继续读取上次的标记）

maxKeys（每次返回的文件数量，默认一次返回1000条数据）

返回示例：
```
{
    "NextContinuationToken": "ChpBSeaooeWeiy_XXXXXXXXXXXXXXXXXXXXXXXXXX",
    "code": 200,
    "data": [
        {
            "key": "11.05会议纪要.docx",
            "content-length": 14567,
            "etag": "\"8AXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"",
            "last_modified": "2024-11-13T01:25:16Z"
        },
        {
            "key": "AI模型/测评报告.docx",
            "content-length": 16530,
            "etag": "\"1590XXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"",
            "last_modified": "2024-11-12T03:46:35Z"
        }
    ],
    "msg": "文件列表获取成功"
}
```

## 五、拷贝接口（POST）
http://127.0.0.1:3030/api/v1/copy

参数：srcBucket、srcObject、destBucket、destObject

返回示例：同一个桶不需要传destBucket参数
```
{
  "code": 200,
  "msg": "文件拷贝成功"
}
```

## 六、重命名接口（POST）
http://127.0.0.1:3030/api/v1/rename

参数：srcObject、destObject

返回示例：
```
{
  "code": 200,
  "msg": "文件'10.14会议纪要复制版.docx'重命名为'10.14会议纪要重命名版.docx'成功"
}
```

### 说明：
阿里云OSS对象存储，上传同名文件会自动覆盖旧文件
