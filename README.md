## 概述
Prometheus Nacos服务发现golang版本。

## JAVA版本
![java version](https://github.com/dtsola/devops-bootadmin-dashboard)

## 依赖
- go 1.15+
- nacos 2.x
- prometheus 2.38.0
- grafana 9.1.5

## 代码说明
### 核心接口
main.go:44 -> /api/prometheus/service/list

### prometheus 服务发现返回格式
```json
        [
                {
                    "targets": [ "<host>", ... ],
                    "labels": {
                      "<labelname>": "<labelvalue>", ...
                    }
                },
          ...
        ]
```

## 使用
### 配置prometheus
```yaml
  - job_name: 'service-monitoring-dev'
    metrics_path: '/actuator/prometheus'
    http_sd_configs:
    - url: 'http://localhost:6003/api/prometheus/service/list'
```

### 测试
#### 启动多个java应用
- 略
#### nacos界面
![](https://pan.bilnn.cn/api/v3/file/sourcejump/lAkO79Sb/PMuYDyFqMdA-yLOgTF28Rh-nVOpqdvsMUvId5ztuUHk*)
#### spring boot admin界面
![](https://pan.bilnn.cn/api/v3/file/sourcejump/6ea4vdHg/sb_f2Ek4o9sRwxbjcC-_hO1FN_KWUqgG0lEnmJ-JCQ8*)
#### prometheus界面
![](https://pan.bilnn.com/api/v3/file/sourcejump/GPoDyEiZ/kw8oDFOtigfwVF8t2hxne3O60Z9euae8PkjXXPf-Fx0*)
#### grafana界面
- dashboard id: 12856
![](https://pan.bilnn.cn/api/v3/file/sourcejump/mmkpobIw/LXRAcvAZkpzUZCeBsN1XpVNG0PyiwFPbVyJhGTqArxo*)

## 联系我
- E-Mail：dtsola@163.com
- QQ：550182738
