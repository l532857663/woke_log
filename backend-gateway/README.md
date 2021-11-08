## 框架选择
1. 日志系统:seelog
2. rpc框架:gin
3. DB SDK：gorm

#### 外部依赖包管理
- 在需要使用新的第三方包时，请在本项目目录下使用`go get github.com/xxxxx`的方式来获取第三方包(go mod会自动维护go.mod)
- 本项目使用本地依赖包管理的形式(-mod=vendor)进行编译，在获取第三方包后，请使用`go mod vendor`加入到本地包管理，项目组所有成员需严格遵守此项约定，防止在拉取代码之后编译时还需要自动下载依赖包

#### 接口文档
- 安装apidoc库
   ```shell script
   npm install --registry=https://registry.npm.taobao.org
   npm config set strict-ssl false
   npm install -g apidoc
   ```
- 生成api-doc
   ```shell script
   make apidoc
   ```
- 接口文档维护时需要注意以下原则：
  > 1. 新增加接口版本必须跟随最新的产品版本
  > 2. 参数必须进行说明并且给出是否是可选参数以及其他属性
  > 3. 调用示例地址统一用`DEV`内测地址

#### 单元测试
- 安装gomock库
   ```shell script
   go get github.com/golang/mock/gomock
   go get github.com/golang/mock/mockgen
   ```
- 本项目已集成gomock框架，并已经注入到service层，在接口开发的过程中，请先使用mock代替数据库dao层，以便最快速度实现业务功能接口，加速前后端开发的整体进度
- go test与gomock可无缝对接使用，所有的业务层代码请务必同时维护test的单元测试代码

#### TODO
- 命令行支持指定应用配置文件和指定日志配置文件

#### gin BaiscAuth鉴权校验方法
- 请求方式
    > 1. 浏览器访问，弹出对话框，输入用户名和密码
    > 2. curl -XGET username:password@localhost:4396/api/backend/backstage/admin/v1/dbconf/reload
    > 3. 添加请求头 Authorization: Basic dXNlcjE6Y2VzaGk=('username:password'字符串的base64加密)
