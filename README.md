# giligili
giligili是一个基于gin、gorm、redis、docker的视频网站后端项目。

项目结构为：

giligili

- api（实现handler）
- cache（实现redis缓存）
- config（使用viper初始化配置）
- middleware（中间件，包括session、auth）
- model（模型）
- server（路由）
- service（服务层）
- utils（工具包，统一返回体与错误枚举）
  - errmsg



实现或即将实现的功能有：

- 用户创建登录
- 视频投稿、查看视频、视频列表以及视频的删改
- redis实现日排行榜
- 阿里云oss对象存储
- docker镜像打包
- 一些权限认证和跨域
