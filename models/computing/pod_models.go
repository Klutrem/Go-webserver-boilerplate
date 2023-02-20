package models

type PodBody struct {
	Name          string   `json:"name"`
	Namespace     string   `json:"namespace"`
	ContainerName string   `json:"container_name"`
	Image         string   `json:"image"`
	Command       []string `json:"command"`
	ConfigmapName string   `json:"configmap_name"`
	SecretName    string   `json:"secret_name"`
	Port          int32    `json:"port"`
	ClaimName     string   `json:"claim_name"`
	VolumeName    string   `json:"volume_name"`
	MountPath     string   `json:"mountpath"`
}

type CPU struct {
	Cores   int `json:"cores"`
	Sockets int `json:"sockets"`
	Threads int `json:"threads"`
}

type VMBody struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	CPU       CPU    `json:"cpu"`
	Memory    string `json:"memory"`
}
