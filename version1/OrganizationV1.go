package version1

import (
	"time"
)

type OrganizationV1 struct {
	Id         string    `json:"id"`
	Code       string    `json:"code"`
	CreateTime time.Time `json:"create_time"`
	CreatorId  string    `json:"creator_id"`
	Deleted    bool      `json:"deleted"`
	Active     bool      `json:"active"`
	Version    int32     `json:"version"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`

	Center     interface{} `json:"center"`
	Radius     float32     `json:"float"`
	Geometry   interface{} `json:"geometry"`
	Boundaries interface{} `json:"boundaries"`

	Language string `json:"language"`
	Timezone string `json:"timezone"`
	Industry string `json:"industry"`
	OrgSize  int32  `json:"org_size"`
	Purpose  string `json:"purpose"`

	Params interface{} `json:"params"`
}

func EmptyOrganizationV1() *OrganizationV1 {
	return &OrganizationV1{}
}

func NewOrganizationV1(id string, code string, name string) *OrganizationV1 {
	return &OrganizationV1{
		Id:         id,
		Code:       code,
		Name:       name,
		CreateTime: time.Now(),
		Active:     true,
		Deleted:    false,
	}
}
