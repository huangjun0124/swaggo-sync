### 一、说明
此目录仅供`swago`命令使用，因为其解析其他的`pkg`文件会报错。
+ `swag init --parseDependency`会连带项目使用到的第三方库一起解析，造成莫名其妙的`parse comment error`
+ 降级到`1.6.7`版本，会导致生成的`swagger`文档，字段没有注释
+ 所以折中采用这个方案

### 二、使用
#### 1、代码使用
+ 方法上添加的`swag`的注释里面，`dto`写成`swag_dto.`下面的
```go
// @Param UserLogInReq body swag_dto.UserLogInReq true "参数"
// @Success 200  {object} swag_dto.UserLogInResp
```
+ `go`不需要引用`swag_dto`这个包

#### 2、`makefile`新增
```makefile
swag:
	cd cmd/server/interface/rest/swag_dto/swagsync && ./swagsync
	cd cmd/server && swag init
```
+ 自动根据配置复制代码实际使用的文件到`swag_dto`下