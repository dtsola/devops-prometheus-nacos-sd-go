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
package main

import (
	"dtsola.com/oss/devops/devops-prometheus-nacos-sd-go/global"
	"dtsola.com/oss/devops/devops-prometheus-nacos-sd-go/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.uber.org/zap"
)

type ServiceTargets struct {
	Targets []string          `json:"Targets"`
	Labels  map[string]string `json:"Labels"`
}

func main() {

	initialize.InitLogger()
	initialize.InitConfig()

	r := gin.Default()

	/**
	 * Description: Get servie's infos from the registry center and
	 * conversion format adaptation the http_sd_configs
	 * *
	 */
	r.GET("/api/prometheus/service/list", func(c *gin.Context) {
		/*
		   Format Style :
		   [
		           {
		               "targets": [ "<host>", ... ],
		               "labels": {
		               "<labelname>": "<labelvalue>", ...
		           }
		           },
		     ...
		   ]
		*/
		// [{"Targets":["192.168.31.64:6002"],"Labels":{}}]
		var r []ServiceTargets = []ServiceTargets{}

		services, err := global.DiscoveryClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
			PageNo:   1,
			PageSize: 1000,
		})
		if err != nil {
			panic("GetAllService failed!")
		}
		//fmt.Printf("GetAllService result:%+v \n\n", services)
		for _, service := range services.Doms {
			instances, err := global.DiscoveryClient.SelectAllInstances(vo.SelectAllInstancesParam{
				ServiceName: service,
			})
			if err != nil {
				zap.S().Fatal("SelectAllInstances failed: " + err.Error())
				panic("SelectAllInstances failed!")
			}
			//fmt.Printf("SelectAllInstance result:%+v \n\n", instances)
			var tmp ServiceTargets = ServiceTargets{
				Targets: []string{},
				Labels:  make(map[string]string),
			}
			for _, instance := range instances {
				tmp.Targets = append(tmp.Targets, fmt.Sprintf("%s:%d", instance.Ip, instance.Port))
			}
			r = append(r, tmp)
		}
		////
		c.JSON(200, r)
	})

	r.Run(global.AppConfig.ServerPort)

}
