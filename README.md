- [一个 go 编写的启动器](#---go-------)
  - [改名](#--)
  - [编译](#--)
  - [配置](#--)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

## 一个 go 编写的启动器

### 改名

将 versioninfo.json.bk 更名为 versioninfo.json 并修改相关配置

图标配置`IconPath`

执行权限配置`ManifestPath`

> 举例添加管理员权限
>
> ManifestPath: "nac.mainfest"
>
> nac.mainfest 中内容为
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
> <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
> <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
>  <security>
>      <requestedPrivileges>
>          <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
>      </requestedPrivileges>
>  </security>
> </trustInfo>
> </assembly>
> ```

### 编译

```powershell
# 生成syso文件，使生成的exe带详细信息(目标为linux下可执行文件时可忽略此步骤)
go generate
# 生成不带控制台的启动程序
go build  -ldflags="-H windowsgui"
```

### 配置

启动配置依赖于`launcher.json`

```json
{
  "target": "ping", // 可以替换为java、./xxx.exe 等参数
  "workDir": ".", //工作目录
  "params": ["www.baidu.com", "-n", "2"] // 参数数组
}
```

程序支持将命令行参数透传添加到 params 后

示例：
配置修改为

```json
{
  "target": "ping",
  "workDir": ".",
  "params": ["www.baidu.com"]
}
```

编译后执行
`./launch.exe -n 2`和默认配置有相同效果
