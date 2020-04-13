package service

type apiRuntime struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Method    string `json:"method"`
	Status    Status `json:"status"`
	ClusterId uint32 `json:"cluster_id"`
}
