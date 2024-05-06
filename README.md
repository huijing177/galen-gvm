# GVM
参考gin-vue-main项目，用于练手

主要学习gin、gorm框架，以及web端相关内容开发

有一些修改，例如：只支持少数数据库-mysql和pgsql
oss部分也只支持aliyun和aws

[参考链接](https://github.com/flipped-aurora/gin-vue-admin)


## 已完成功能列表
- 使用`viper`+`mapstructure`读取配置文件，并能实时更新配置
- 使用`gorm`链接数据库
- 使用`zap`构建日志模块
- 使用`swag`构建`swagger`页面
- 路由基础分组完成

## TODO
- 用户注册，登陆，修改密码，修改权限
    使用到验证码、redis、JWT
- 黑名单判断，加入以及删除
- 中间件，其他接口需要验证登陆信息


## 防爆
1. 可无限刷新验证码图片----> captcha 接口
2. 同一IP，调用登录接口，失败n次后，界面开始显示验证码输入栏，此时登录需要验证码才能登录