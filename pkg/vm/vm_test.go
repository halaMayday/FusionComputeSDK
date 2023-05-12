package vm

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/task"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
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
		cs, err := cm.ListVm(true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}

func TestManager_CloneVm(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()
	m := NewManager(c, "/service/sites/43BC08E8")
	ts, err := m.CloneVm("/service/sites/43BC08E8/vms/i-00000034", CloneVmRequest{
		Name:          "test-1",
		Description:   "test create vm",
		Location:      "urn:sites:43BC08E8:clusters:117",
		IsBindingHost: false,
		Config: Config{
			Cpu: Cpu{
				Quantity:    2,
				Reservation: 0,
			},
			Memory: Memory{
				QuantityMB:  2048,
				Reservation: 2048,
			},
			Disks: []Disk{
				{
					SequenceNum:  1,
					QuantityGB:   50,
					IsDataCopy:   true,
					DatastoreUrn: "urn:sites:43BC08E8:datastores:41",
					IsThin:       true,
				},
			},
			Nics: []Nic{
				{
					Name:         "vmnic1",
					PortGroupUrn: "urn:sites:43BC08E8:dvswitchs:1:portgroups:1",
				},
			},
		},
		VmCustomization: Customization{
			OsType:             "Linux",
			Hostname:           "test-1",
			IsUpdateVmPassword: false,
			NicSpecification: []NicSpecification{
				{
					SequenceNum: 1,
					Ip:          "100.199.10.88",
					Netmask:     "255.255.255.0",
					Gateway:     "100.199.10.1",
					Setdns:      "114.114.114.114",
					Adddns:      "8.8.8.8",
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("create vm %s", ts.Uri)
	fmt.Printf("task uri  %s", ts.TaskUri)

	tm := task.NewManager(c, "/service/sites/43BC08E8")
	for {
		tt, err := tm.Get(ts.TaskUri)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("task status %s", tt.Status)
		if tt.Status == "success" {
			break
		}
		time.Sleep(5 * time.Second)
	}

	resp, err := m.DeleteVm(ts.Uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete vm %s  response: %v", ts.Uri, resp)

}

func Test_manager_StartVm(t *testing.T) {
	type fields struct {
		client  client.FusionComputeClient
		siteUri string
	}
	type args struct {
		vmUri string
	}
	c := client.NewFusionComputeClient("https://192.168.17.131:7443", "mlc", "1qaz@WSX")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *StartVmResponse
		wantErr bool
	}{
		{name: "start", fields: fields{client: c, siteUri: ""}, args: args{"/service/sites/47520799/vms/i-00000061"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				client:  tt.fields.client,
				siteUri: tt.fields.siteUri,
			}
			got, err := m.StartVm(tt.args.vmUri)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartVm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartVm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_StopVm(t *testing.T) {
	type fields struct {
		client  client.FusionComputeClient
		siteUri string
	}
	type args struct {
		vmUri string
	}
	c := client.NewFusionComputeClient("https://192.168.17.131:7443", "mlc", "1qaz@WSX")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *StopVmResponse
		wantErr bool
	}{
		{name: "stop-test", fields: fields{client: c, siteUri: ""}, args: args{"/service/sites/47520799/vms/i-00000061"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				client:  tt.fields.client,
				siteUri: tt.fields.siteUri,
			}
			got, err := m.StopVm(tt.args.vmUri)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopVm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StopVm() got = %v, want %v", got, tt.want)
			}
		})
	}
}
