package site

type Site struct {
	Uri         string `json:"uri"`
	Urn         string `json:"urn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"` //joining , 加入域中  exiting,退出域中 normal ，正常  fault , 故障
}

type ListSiteResponse struct {
	Sites []Site `json:"sites"`
}
