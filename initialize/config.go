/*
 * Copyright 2022 Dtsola.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package initialize

import (
	"dtsola.com/oss/devops/devops-prometheus-nacos-sd-go/config"
	"dtsola.com/oss/devops/devops-prometheus-nacos-sd-go/global"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func InitConfig() {

	configPath := ""
	configEnv := os.Getenv("GO_ENV")
	switch configEnv {
	case "prod":
		configPath = "./config-prod.yaml"
	default:
		configPath = "./config-dev.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	var appConfig config.AppConfig

	if err := v.Unmarshal(&appConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", appConfig)

	//
	//从nacos中读取配置信息
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(appConfig.NacosHost, appConfig.NacosPort, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(appConfig.NacosNamespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	zap.S().Info("nacos初始化成功.")
	//
	global.DiscoveryClient = client
	global.AppConfig = appConfig
	zap.S().Info("nacos注册成功.")
	///

}
