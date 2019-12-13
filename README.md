# 说明
简单了解filebeat开发插件的流程，及插件工作的机制

# filebeat-plugin
filebeat plugin processor

# add_sample_data
- 基于URL维度计数，进行采样控制

- 采样率控制config.go中Sample

# add_sample_data使用场景
准确率要求不高的情况下，可以使用采样的机制，比如，通过日志，计算服务的可用性，基于统计计数的情况，会存在误差，服务波动情况下，可能导致抽样的数据都是正常的，异常的日志漏掉。

# 版本信息
- os: mac 10.14.6
- golang: go1.12.7
- filebeat: 7.4.2

基于以上版本测试通过

# 运行
./filebeat -e -c filebeat_test.yml --plugin add_sample_data.so

# 开发
- 下载beats源码 https://github.com/elastic/beats
- 可以选择filebeat对应版本的分支
- libbeat/processors下新建add_sample_data目录(下面有自带processor)
- 插件编译打包在当前目录下

# 打包
基于 https://github.com/elastic/beats 源码下进行编译打包

```$xslt
cd libbeat/processors/add_sample_data
go build -buildmode=plugin
```

# filebeat.yml配置
```$xslt
#================================ Processors =====================================
processors:
  - add_sample_data: ~
```

# 问题
1. 启动运行命令签名错误
> 由于Mac系统（csrutil）保护系统完整性导致的问题，运行 csrutil status 查看结果，如果是 enable，关闭disable即可。具体操作步骤自行百度
2. 启动 plugin was built with a different version of package errors
> plugin 编译问题导致，解决方法，基于beats对应版本源代码重新编译，执行filebeat下make命令，替换原有 filebeat 二进制文件

# 参考
https://github.com/zhonglongbo/filebeat_plugin
