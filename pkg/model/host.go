package model

import (
	"errors"
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	"github.com/KubeOperator/KubeOperator/pkg/util/encrypt"
	uuid "github.com/satori/go.uuid"
)

const (
	Disconnect string = "DisConnect"
	SshError   string = "SshError"
)

type Host struct {
	common.BaseModel
	ID           string     `json:"id"`
	Name         string     `json:"name" gorm:"type:varchar(256);not null;unique"`
	Memory       int        `json:"memory" gorm:"type:int(64)"`
	CpuCore      int        `json:"cpuCore" gorm:"type:int(64)"`
	Os           string     `json:"os" gorm:"type:varchar(64)"`
	OsVersion    string     `json:"osVersion" gorm:"type:varchar(64)"`
	GpuNum       int        `json:"gpuNum" gorm:"type:int(64)"`
	GpuInfo      string     `json:"gpuInfo" gorm:"type:varchar(128)"`
	Ip           string     `json:"ip" gorm:"type:varchar(128);not null;unique"`
	Port         int        `json:"port" gorm:"type:varchar(64)"`
	CredentialID string     `json:"credentialId" gorm:"type:varchar(64)"`
	ClusterID    string     `json:"clusterId" gorm:"type:varchar(64)"`
	ZoneID       string     `json:"_"`
	Zone         Zone       `json:"_"  gorm:"save_associations:false" `
	Volumes      []Volume   `json:"volumes" gorm:"save_associations:false"`
	Credential   Credential `json:"_" gorm:"save_associations:false" `
	Cluster      Cluster    `json:"_" gorm:"save_associations:false" `
	Status       string     `json:"status" gorm:"type:varchar(64)"`
	Message      string     `json:"message" gorm:"type:text(65535)"`
}

func (h Host) GetHostPasswordAndPrivateKey() (string, []byte, error) {
	var err error = nil
	password := ""
	privateKey := []byte("")
	if "password" == h.Credential.Type {
		pwd, err := encrypt.StringDecrypt(h.Credential.Password)
		password = pwd
		if err != nil {
			return password, privateKey, err
		}
	}
	if "privateKey" == h.Credential.Type {
		privateKey = []byte(h.Credential.PrivateKey)
	}
	return password, privateKey, err
}

func (h *Host) BeforeCreate() error {
	h.ID = uuid.NewV4().String()
	return nil
}

func (h *Host) BeforeDelete() error {
	if h.ClusterID != "" {
		return errors.New("DELETE_HOST_FAILED")
	}
	return nil
}

func (h Host) TableName() string {
	return "ko_host"
}