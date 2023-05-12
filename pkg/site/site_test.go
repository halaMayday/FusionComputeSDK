package site

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
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
	m := NewManager(c)
	ss, err := m.ListSite()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ss)
}
