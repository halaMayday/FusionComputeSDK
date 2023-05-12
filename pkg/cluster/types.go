package cluster

type Cluster struct {
	Name        string `json:"name"`
	Uri         string `json:"uri"`
	Urn         string `json:"urn"`
	Arch        string `json:"arch"`
	Description string `json:"description"`
	EnableHa    bool   `json:"isEnableHa"`
}

type ListClusterResponse struct {
	Clusters []Cluster `json:"clusters"`
}
