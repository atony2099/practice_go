package transport

type GetOder struct {
	Id string `json:"id,omitempty"`
}

type OderInfo struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}
