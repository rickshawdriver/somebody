package service

type ApiRuntime struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Url       []byte `json:"url"`
	Desc      string `json:"desc"`
	Method    string `json:"method"`
	ClusterId uint32 `json:"cluster_id"`
}

func (d *Dispatcher) addApi(api *ApiRuntime) error {
	d.Apis[api.ID] = api
	d.Router.Put(api.Url, api.ID)

	return nil
}

func (c *ApiRuntime) GetID() uint32 {
	return c.ID
}
