#### DDL创建
```sql
// 用户表创建
  CREATE TABLE `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
    `addr` varchar(255) NOT NULL DEFAULT '' COMMENT '住址',
    `age` smallint(4) NOT NULL DEFAULT '0' COMMENT '年龄',
    `birth` varchar(100) NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '生日',
    `sex` smallint(4) NOT NULL DEFAULT '0' COMMENT '性别',
    `update_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    `create_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表'


  接口文档
  环境配置：
   base_url: http://localhost:8080

   管理登录接口
      地址：/login
      类型：post
      参数：
      username string 必填 用户名
      password string 必填 密码

      返回：
      {
        "errno": 0,
        "errmsg": 0,
        "data": "",
        "trace_id": "xxxxxxx"
      }
    
    管理退出接口
      地址：/loginout
      类型：get
      参数：

      返回：
      {
        "errno": 0,
        "errmsg": 0,
        "data": "",
        "trace_id": "xxxxxxx"
      }
    
    用户列表接口
      地址：/user/listpage
      类型：post
      参数：
      page string 必填 用户名
      name string 选填 用户名

      返回：
      {
        "errno": 0,
        "errmsg": 0,
        "data": {
          "list": [
            {
              "id"
            }
          ]
        },
        "trace_id": "xxxxxxx"
      }
```
