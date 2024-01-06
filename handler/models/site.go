package models

type GetSiteRequest struct {
	Name                string `form:"name"`
	IsMaximumAccessTime bool   `form:"is_maximum_access_time"`
	IsMinimumAccessTime bool   `form:"is_minimum_access_time"`
}

type GetSiteResponse struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	AccessTime int64  `json:"access_time"`
}
