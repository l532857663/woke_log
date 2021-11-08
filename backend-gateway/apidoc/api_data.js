define({ "api": [
  {
    "type": "get",
    "url": "/api/backend/goods/v1/list",
    "title": "获取商品列表数据",
    "name": "getGoodsListV1",
    "version": "1.0.0",
    "group": "goods",
    "description": "<p>根据条件获取商品列表数据</p>",
    "permission": [
      {
        "name": "已验证用户"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": true,
            "field": "X-Auth-Token",
            "description": "<p>（预留）用户Token</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "language",
            "description": "<p>客户端语言 EN-英文 CN-简体中文</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "entity_id",
            "description": "<p>Y星（URL参数）</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "请求数据格式:",
          "content": "?language=CN&entity_id=996",
          "type": "string"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>错误状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>错误信息</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>数据详情</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "data.timestamp",
            "description": "<p>记录获取时间</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "data.total_count",
            "description": "<p>记录总条数</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "data.tokens_have_balance",
            "description": "<p>数据详情</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.tokens_have_balance.symbol",
            "description": "<p>币种展示名称</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正确返回值:",
          "content": "HTTP/1.1 200 OK\n{\n    \"code\": 1000,\n    \"message\": \"\",\n    \"data\": {\n        \"timestamp\": 1608890232,\n        \"total_count\": 3,\n\t\t \"tokens_have_balance\": [\n              {\n                  \"symbol\": \"YTA\",\n              }\n\t\t ]\n    }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "server/backend-gateway/goods.go",
    "groupTitle": "goods",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "InvalidParams_1300",
            "description": "<p>无效参数</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "RecordNotFound_1301",
            "description": "<p>记录不存在</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "DataAlreadyExists_1302",
            "description": "<p>数据重复</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "InvalidKycStatus_1330",
            "description": "<p>身份认证状态错误</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "InvalidInviteCode_1331",
            "description": "<p>无效的邀请码</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "InvalidInviteAction_1332",
            "description": "<p>禁止自己邀请自己</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "DataAlreadyExists_1399",
            "description": "<p>服务内部错误</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "错误返回值:",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n     \"code\": 1301,\n     \"message\": \"record not found\"\n}",
          "type": "json"
        }
      ]
    }
  }
] });
