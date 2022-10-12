# ginkgo_keyauth

用户中心

## Mongodb

1. 如何连接MongoDB, 以及MongoDB的配置
2. 基于MongoDB的 CRUD
    + bson struct tag: 用于MongoDB 保存时，完成Object -> 
    + "_id": 内置的BSON TAG, 类似于MySQL主键, _: 代表的倒序索引, 不用额外多创建一个Index, 最好沿用 
    + 通过DB对象的Collection对象来 进行CRUD操作
    + 

## 对接注册中心
主题: 服务发现
服务注册
   1. 添加注册中心的配置
   2. 然后初始化全局的 注册中心客户端实例 rpc.C()
   3. GRPC服务启动时 调用注册中心的客户端把当前GRPC监听的地址注册过去
   4. 当GRPC服务Stop时,注销把注册中心的实例下线掉

服务解锁(GRPC Client)
   1. 通过GRPC的NamedResolver来进行服务的发现
   2. 也是加载注册中心的GRPC客户端: rpc.C() --? , 因为Mcenter提供的Resolver需要依赖注册中心的客户端来进行服务实例的搜索;
   3. 在服务启动初始化时候,就已经完成了以上步骤;
   4. GRPC客户端配置注册中心的访问凭证, 已经需要访问的服务的名称, Resovler就能完成 服务名称 ---> 地址的解析

准备完成Keyauth的客户端, CMDB就可以通过该客户端来和Keyauth进行交互
```shell
// keyauth 客户端
// 需要配置注册中心的地址
// 获取注册中心的客户端，使用注册中心的客户端 查询 keyauth的地址
func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")

	// 把Mcenter的配置传递给Keyauth的客户端
	c, err := rpc.NewClient(conf)

	// 使用SDK 调用Keyauth进行 凭证的校验
	// c.Token().ValidateToken()

	if should.NoError(err) {
		resp, err := c.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest("yTGTAj3fnPWqXIEkuicr57bf1"),
		)
		should.NoError(err)
		fmt.Println(resp)
	}
}
```

cmdb 如何使用Keyauth的客户端进行Token的校验,需要一个HTTP的认证中间件
因为 Keyauth 是用户中心, HTTP的权限中间件需要和 keyatuh 交互,依赖 keyauth 的SDK, 因此这个中间件由 keyauth 提供

## 架构图

## 项目说明

```
├── protocol                       # 脚手架功能: rpc / http 功能加载
│   ├── grpc.go              
│   └── http.go    
├── client                         # 脚手架功能: grpc 客户端实现 
│   ├── client.go              
│   └── config.go    
├── cmd                            # 脚手架功能: 处理程序启停参数，加载系统配置文件
│   ├── root.go             
│   └── start.go                
├── conf                           # 脚手架功能: 配置文件加载
│   ├── config.go                  # 配置文件定义
│   ├── load.go                    # 不同的配置加载方式
│   └── log.go                     # 日志配置文件
├── dist                           # 脚手架功能: 构建产物
├── etc                            # 配置文件
│   ├── xxx.env
│   └── xxx.toml
├── apps                            # 具体业务场景的领域包
│   ├── all
│   │   |-- grpc.go                # 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序。  
│   │   |-- http.go                # 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载。                    
│   │   └── internal.go            #  注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖。 
│   ├── book                       # 具体业务场景领域服务 book
│   │   ├── http                   # http 
│   │   │    ├── book.go           # book 服务的http方法实现，请求参数处理、权限处理、数据响应等 
│   │   │    └── http.go           # 领域模块内的 http 路由处理，向系统层注册http服务
│   │   ├── impl                   # rpc
│   │   │    ├── book.go          # book 服务的rpc方法实现，请求参数处理、权限处理、数据响应等 
│   │   │    └── impl.go           # 领域模块内的 rpc 服务注册 ，向系统层注册rpc服务
│   │   ├──  pb                    # protobuf 定义
│   │   │     └── book.proto       # book proto 定义文件
│   │   ├── app.go                 # book app 只定义扩展
│   │   ├── book.pb.go             # protobuf 生成的文件
│   │   └── book_grpc.pb.go        # pb/book.proto 生成方法定义
├── version                        # 程序版本信息
│   └── version.go                    
├── README.md                    
├── main.go                        # Go程序唯一入口
├── Makefile                       # make 命令定义
└── go.mod                         # go mod 依赖定义
```


## 快速开发
make脚手架
```sh
➜  keyauth-g7 git:(master) ✗ make help
dep                            Get the dependencies
lint                           Lint Golang files
vet                            Run go vet
test                           Run unittests
test-coverage                  Run tests with coverage
build                          Local build
linux                          Linux build
run                            Run Server
clean                          Remove previous build
help                           Display this help screen
```

1. 使用安装依赖的Protobuf库(文件)
```sh
# 把依赖的probuf文件复制到/usr/local/include

# 创建protobuf文件目录
$ make -pv /usr/local/include/github.com/infraboard/mcube/pb

# 找到最新的mcube protobuf文件
$ ls `go env GOPATH`/pkg/mod/github.com/infraboard/

# 复制到/usr/local/include
$ cp -rf pb  /usr/local/include/github.com/infraboard/mcube/pb
```

2. 添加配置文件(默认读取位置: etc/keyauth-g7.toml)
```sh
$ 编辑样例配置文件 etc/keyauth-g7.toml.book
$ mv etc/keyauth-g7.toml.book etc/keyauth-g7.toml
```

3. 启动服务
```sh
# 编译protobuf文件, 生成代码
$ make gen
# 如果是MySQL, 执行SQL语句(docs/schema/tables.sql)
$ make init
# 下载项目的依赖
$ make dep
# 运行程序
$ make run
```

## 相关文档