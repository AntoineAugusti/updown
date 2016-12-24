package updown

import (
	"net/http"
)

// NodeService interacts with the nodes section of the API
type NodeService struct {
	client *Client
}

// NodeDetails gives information about a node
type NodeDetails struct {
	IP          string `json:"ip,omitempty"`
	IP6         string `json:"ip6,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// IPs represents IP addresses in v4 or v6
type IPs []string

// Nodes represents multiple nodes
type Nodes map[string]NodeDetails

// List gets the nodes performing checks
func (s *NodeService) List() (Nodes, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "nodes", nil)
	if err != nil {
		return nil, nil, err
	}

	var res Nodes
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

// ListIPv4 gets the list of IPv4 performing checks
func (s *NodeService) ListIPv4() (IPs, *http.Response, error) {
	return s.genericIPList("4")
}

// ListIPv6 gets the list of IPv6 performing checks
func (s *NodeService) ListIPv6() (IPs, *http.Response, error) {
	return s.genericIPList("6")
}

// genericIPList get the list of IPv4 or IPv6 IPs performing checks
func (s *NodeService) genericIPList(version string) (IPs, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "nodes/ipv"+version, nil)
	if err != nil {
		return nil, nil, err
	}

	var res IPs
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}
