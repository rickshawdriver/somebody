package service

type Cluster struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (d *Dispatcher) addCluster(cluster *Cluster) error {
	d.Clusters[cluster.ID] = cluster

	return nil
}

func (c *Cluster) GetID() uint32 {
	return c.ID
}
