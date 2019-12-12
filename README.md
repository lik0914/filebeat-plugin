# filebeat-plugin
filebeat plugin processor

# 版本信息
- os: mac 10.14.6
- golang: go1.12.7
- filebeat: 7.4.2

基于以上版本测试通过

# 运行
./filebeat-my -e -c filebeat_test.yml --plugin add_sample_data.so

# add_sample_data
- 基于URL维度计数，进行采样控制

- 采样率控制config.go中Sample

# 参考
https://github.com/zhonglongbo/filebeat_plugin
