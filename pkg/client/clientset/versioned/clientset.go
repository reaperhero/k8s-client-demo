package versioned

import (
	mycompanyv1 "github.com/reaperhero/k8s-client-demo/pkg/client/clientset/versioned/typed/nginx_controller/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	mycompanyV1 *mycompanyv1.MycompanyV1Client
}

// MycompanyV1 retrieves the MycompanyV1Client
func (c *Clientset) MycompanyV1() mycompanyv1.MycompanyV1Interface {
	return c.mycompanyV1
}

// Deprecated: Mycompany retrieves the default version of MycompanyClient.
// Please explicitly pick a version.
func (c *Clientset) Mycompany() mycompanyv1.MycompanyV1Interface {
	return c.mycompanyV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.mycompanyV1, err = mycompanyv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
