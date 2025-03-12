*以下是可能的需要准备工作，熟悉者可忽略*

## 安装redis并启动服务

`yum install redis`

`systemctl start redis`

## redis-cli使用

终端输入redis-cli进入 exit退出

HSET、HGET、HGETALL用于设置和查看哈希表数据
（需要设置全局容量才能正常购票）

`HSET ticket_hash_key ticket_total_nums 150 ticket_sold_nums 0`

----

## 安装GVM管理go版本

下载脚本，自动执行安装

`bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)`

## go安装使用
*要编译安装的话，可能要先下载1.4版本(1.5+版本使用go编译)*

*建议直接下载二进制版本（参数-B 或 --binary）*

`gvm install go<version>`

`gvm use go<verison>`

## GVM查看可用go版本

可下载的：`gvm listall`

已安装的：`gvm list`


### 其余命令查看readme文档(/root/.gvm/README.md)
----

### GVM与vscode插件
vscode的go插件不能识别GVM的go环境，需要手动配置插件的root和path（建议使用最新版go，运行和插件的版本可以不同）

Open User Settings(JSON)打开配置文件settings.json
添加类似下面的配置
```
"go.goroot": "/Users/xxx/.gvm/gos/go1.24.1",
"go.gopath": "/Users/xxx/.gvm/pkgsets/go1.24.1/global",
```
