sh start-business.sh : 用于启动业务服务

controllers: 
      login_cor.go: 处理登录请求的控制器。
      register_cor.go: 处理注册请求的控制器。
      uploadCor.go: 处理文件上传请求的控制器。
model: 
      rdbConnect.go: 连接到关系数据库的代码。
      sqlConnect.go: 连接到 SQL 数据库的代码。
      user.go:  User 模型，它表示应用程序的用户。
routes: 
      login_router.go: 定义与登录相关请求的路由。
      pop_router.go: 定义其他一些功能的路由。
      refister_success.go: 定义注册成功页面的路由。
      register_router.go: 定义与注册相关请求的路由。
      upload_router.go: 定义与文件上传相关请求的路由。
static: 
      Iryh.pdf: 传入的单文件pdf。
templates: 
      login.html: 登录页面的模板。
      pop-ups.html: 一些弹出窗口的模板。
      register_success.html: 注册成功页面的模板。
      register.html: 注册页面的模板。
      result.html: 一些结果页面的模板。
      upload.html: 文件上传页面的模板。
tmp: 临时文件。
tool: 
      tools.go: 公共方法。
go.mod: 依赖项。
go.sum: 依赖项的校验和文件。
main.go: 应用程序的入口。
runner.conf: 是应用程序的配置文件。




sh start-pdf.sh 用于启动pdf工具服务

pdf-tool-service 
├── controllers 
│   └── massageCor 
│       ├──massage.go 消息控制器
├── model 
│   └──pdf.go pdf数据结构模型
|   └──rdbConnect.go redis数据库连接
├── router 
│   └──message.go 路由
├── services
│   └── pdf_process.go  pdf工具服务
├── static
│   ├── even_pages 
│   │   ├── even_pages.pdf 生成的偶数页pdf
│   │—— odd_pages
|   |   |——odd_pages.pdf 生成的奇数页pdf
│   └── Iryh.pdf 传入的pdf
├── tmp
│   └── runner-build.exe
|── go.mod
│── go.sum
│──main.go
├── runner.conf
