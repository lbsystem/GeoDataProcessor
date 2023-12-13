# geo

`geo` 是一个命令行工具，用于处理和输出特定地理数据文件中的IP地址或网站信息。它支持从GeoIP或GeoSite格式的文件中提取数据，并按国家或网站类别进行过滤。

## 安装

[提供安装指南，例如从源代码编译或直接下载编译好的二进制文件等]

## 命令行参数

- `-f`: 指定输入文件的路径。必须提供。
- `-m`: 模式选择，可以是`ip`或`site`。默认为`ip`。
- `-c`: 当模式为`ip`时，指定国家代码（例如`cn`）。默认为`cn`。
- `-s`: 当模式为`site`时，指定网站类别（如`gfw`、`google`等）。默认为`gfw`。
- `-v`: 指定IP版本，可以是`all`、`4`或`6`。默认为`all`。

## 使用示例

### 处理GeoIP数据

- 处理指定国家（如中国）的所有IP地址：
```geo -f path/to/geoip.dat -m ip -c cn```



- 仅处理IPv4地址：
```geo -f path/to/geoip.dat -m ip -c cn -v 4```
除了可以指定国家还有各种模式
cloudflare
cloudfront
facebook
fastly
google
netflix
telegram
twitter


- 仅处理IPv6地址：
```geo -f path/to/geoip.dat -m ip -c cn -v 6```



### 处理GeoSite数据

- 处理特定类别（如`gfw`）的网站信息：
```geo -f path/to/geosite.dat -m site -s gfw```


## 输出

程序将处理后的数据以文本形式输出到标准输出（stdout）。在IP模式下，输出格式为IP地址/子网掩码；在网站模式下，输出为网站地址。

## 错误处理

如果发生错误（如文件无法加载、格式错误等），程序将输出错误信息并退出。

## 注意事项
可以在https://github.com/Loyalsoldier/v2ray-rules-dat 获得最新的geo 文件

- 确保提供的文件路径和格式正确。
- 使用合适的权限运行程序，以避免文件访问权限问题。

