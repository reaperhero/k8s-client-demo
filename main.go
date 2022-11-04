package main

import (
	"github.com/reaperhero/k8s-client-demo/common"
	"github.com/reaperhero/k8s-client-demo/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	var (
		restConf  *rest.Config
		crdClient *versioned.Clientset
		err       error
	)

	// 读取admin.conf, 生成客户端基本配置
	if restConf, err = common.GetRestConf(); err != nil {
		return
	}

	// 创建CRD的client
	if crdClient, err = versioned.NewForConfig(restConf); err != nil {
		return
	}

	// 获取CRD的nginx对象
	if _, err = crdClient.MycompanyV1().Nginxes("default").Get("my-nginx", v1.GetOptions{}); err != nil {
		return
	}

}
