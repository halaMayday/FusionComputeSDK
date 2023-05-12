package cluster

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"log"
	"testing"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://192.168.17.131:7443", "mlc", "1qaz@WSX")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()

	sm := site.NewManager(c)
	ss, err := sm.ListSite()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		cm := NewManager(c, s.Uri)
		cs, err := cm.ListCluster()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}
