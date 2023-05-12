package vm

type Vm struct {
	Urn               string `json:"urn,omitempty,omitempty"`
	Uri               string `json:"uri,omitempty"`
	Uuid              string `json:"uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Arch              string `json:"arch,omitempty"`
	Description       string `json:"description,omitempty"`
	Group             string `json:"group,omitempty"`
	Location          string `json:"location,omitempty"`
	LocationName      string `json:"locationName,omitempty"`
	HostUrn           string `json:"hostUrn,omitempty"`
	Status            string `json:"status,omitempty"`
	PvDriverStatus    string `json:"pvDriverStatus,omitempty"`
	ToolInstallStatus string `json:"toolInstallStatus,omitempty"`
	CdRomStatus       string `json:"cdRomStatus,omitempty"`
	IsTemplate        bool   `json:"isTemplate,omitempty"`
	IsLinkClone       bool   `json:"isLinkClone,omitempty"`
	IsBindingHost     bool   `json:"isBindingHost,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	ToolsVersion      string `json:"toolsVersion,omitempty"`
	HostName          string `json:"hostName,omitempty"`
	ClusterName       string `json:"clusterName,omitempty"`
	HugePage          string `json:"hugePage,omitempty"`
	Idle              int    `json:"idle,omitempty"`
	VmType            int    `json:"vmType,omitempty"`
	DrStatus          int    `json:"drStatus,omitempty"`
	RpoStatus         int    `json:"rpoStatus,omitempty"`
	InitSyncStatus    int    `json:"initSyncStatus,omitempty"`
	VmConfig          Config `json:"vmConfig,omitempty"`
}

type Customization struct {
	OsType             string             `json:"osType,omitempty"`
	Hostname           string             `json:"hostname,omitempty"`
	IsUpdateVmPassword bool               `json:"isUpdateVmPassword,omitempty"`
	Password           string             `json:"password,omitempty"`
	NicSpecification   []NicSpecification `json:"nicSpecification,omitempty"`
}

type NicSpecification struct {
	SequenceNum int    `json:"sequenceNum,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Setdns      string `json:"setdns,omitempty"`
	Adddns      string `json:"adddns,omitempty"`
}

type ListVmResponse struct {
	Total int  `json:"total,omitempty"`
	Vms   []Vm `json:"vms,omitempty"`
}

type CloneVmRequest struct {
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	Group           string        `json:"group,omitempty"`
	Location        string        `json:"location,omitempty"`
	IsBindingHost   bool          `json:"isBindingHost,omitempty"`
	Config          Config        `json:"vmConfig,omitempty"`
	VmCustomization Customization `json:"vmCustomization,omitempty"`
}

type Config struct {
	Cpu    Cpu    `json:"cpu,omitempty"`
	Memory Memory `json:"memory,omitempty"`
	Disks  []Disk `json:"disks,omitempty"`
	Nics   []Nic  `json:"nics,omitempty"`
}
type Cpu struct {
	Quantity    int `json:"quantity,omitempty"`
	Reservation int `json:"reservation,omitempty"`
	Weight      int `json:"weight,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

type Memory struct {
	QuantityMB  int `json:"quantityMb,omitempty"`
	Reservation int `json:"reservation,omitempty"`
	Weight      int `json:"weight,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

type Disk struct {
	SequenceNum  int    `json:"sequenceNum,omitempty"`
	QuantityGB   int    `json:"quantityGb,omitempty"`
	IsDataCopy   bool   `json:"isDataCopy,omitempty"`
	DatastoreUrn string `json:"datastoreUrn,omitempty"`
	IsThin       bool   `json:"isThin,omitempty"`
}

type Nic struct {
	Name         string `json:"name,omitempty"`
	PortGroupUrn string `json:"portGroupUrn,omitempty"`
	Mac          string `json:"mac,omitempty"`
	Ip           string `json:"ip,omitempty"`
}

type CloneVmResponse struct {
	Urn     string `json:"urn,omitempty"`
	Uri     string `json:"uri,omitempty"`
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type DeleteVmResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type StartVmResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type StopVmResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type ImportTemplateRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
	VmConfig    Config   `json:"vmConfig,omitempty"`
	OsOptions   OsOption `json:"osOptions,omitempty"`
	Url         string   `json:"url,omitempty"`
	Protocol    string   `json:"protocol,omitempty"`
	IsTemplate  bool     `json:"isTemplate,omitempty"`
}

type OsOption struct {
	OsType      string `json:"osType,omitempty"`
	OsVersion   int    `json:"osVersion,omitempty"`
	GuestOSName string `json:"guestOsName,omitempty"`
}

type ImportTemplateResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

// "name": string,//如："vm132"
//    "portGroupUrn": string,//如：urn:sites:1:DVSwitchs:1:portgroup:1
//"mac":string //如："00:00:00:00:00:00"

type AddNicRequest struct {
	Name         string `json:"name ,omitempty"` //可选，长度【0,256】
	PortGroupUrn string `json:"portGroupUrn"`
	Mac          string `json:"mac,omitempty"`
	SequenceNum  int    `json:"sequenceNum,omitempty"` // 网卡对应的网线编号，0,15可选，不可重复
	VirtIo       int    `json:"virtio,omitempty"`      //可选。1：virtio（默认值） 2：e1000 (KVM未支持) 3：rtl8139 (KVM未支持) 4：vhost-user
	//portId	string	虚拟机交换机端口 ID，可选。
	PortId    string    `json:"portId,omitempty"`
	NicConfig NicConfig `json:"nicConfig,omitempty"`
}

type NicConfig struct {
	//vringbuf	integer	IO环大小，默认为256，取值范围为256、512、1024、2048、4096
	//queues	Integer	队列数，默认为1，取值范围为[1,8]
	Vringbuf int `json:"vringbuf,omitempty"`
	Queues   int `json:"queues,omitempty"`
}

//	{
//	   "urn":string, //如："urn:sites:1:vms:1:nics:1"
//	   "uri":string//如："/service/sites/1/vms/1/nics/1"
//	   " taskUrn": string,
//
// "taskUri": string
// }
type AddNicResponse struct {
	Urn     string `json:"urn,omitempty"`
	Uri     string `json:"uri,omitempty"`
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type DeleteNicResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type AddSecurityGroupRequest struct {
	//enableSecurityGroup	Boolean	可选，虚拟机网卡是否开启安全组：true表示开启，false表示未开启，默认为false。
	//securityGroupId	long	安全组Id，当enableSecurityGroup为true时有效。
	EnableSecurityGroup bool  `json:"enableSecurityGroup,omitempty"`
	SecurityGroupId     int64 `json:"securityGroupId,omitempty"`
}

type AddSecurityGroupResponse struct {
	//error	String	配置安全组返回的错误信息
	//result	boolean	配置安全组返回值
	//returnValue	boolean	配置安全组返回值
	//taskId	Int	配置安全组返回的taskId
	//uri	String	配置安全组任务对应的URI标识
	//urn	String	配置安全组任务对应的URN标识
	ErrMsg      string `json:"error"`
	Result      bool   `json:"result"`
	ReturnValue bool   `json:"returnValue"`
	TaskId      int64  `json:"taskId"`
	TaskUrn     string `json:"taskUrn,omitempty"`
	TaskUri     string `json:"taskUri,omitempty"`
}

type AttachVolumeRequest struct {
}

type AttachVolumeResponse struct {
}

type DetachVolume struct {
}
