package mcast

// DistClient is a client within a distribution group
type DistClient struct {
	IPAddr   string `json:"ip_address"`
	DataType string `json:"data_type"`
}

// Distribution is a MusicCast distribution group
type Distribution struct {
	GroupID      string `json:"group_id"`
	GroupName    string `json:"group_name"`
	Role         string
	Status       string
	ServerZone   string       `json:"server_zone"`
	Clients      []DistClient `json:"client_list"`
	AudioDropout bool         `json:"audio_dropout"`
	// Unsupported fields: build_disable, mc_surround
}

// Distribution returns information about any distribution (Link group) the
// device belongs to
func (d Device) Distribution() (Distribution, error) {

	var dist Distribution

	err := d.GetRequest("dist/getDistributionInfo", &dist)
	if err != nil {
		return Distribution{}, err
	}

	return dist, nil
}
