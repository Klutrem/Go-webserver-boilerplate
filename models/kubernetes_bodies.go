package models

type ChartBody struct {
	ChartPath   string `json:"chart_path"`
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
	Reponame    string `json:"repo_name"`
}

type RepositoryBody struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type ConfigMapBody struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Data      map[string]string `json:"env"`
}

type SecretBody struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Data      map[string]string `json:"env"`
}

type PV struct {
	Name    string `json:"name"`
	Storage string `json:"storage"`
	Path    string `json:"path"`
}

type PVC struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Storage   string `json:"storage"`
}

type Nodeport struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Port         int32  `json:"port"`
	RedirectPort int32  `json:"redirect_port"`
}

type Role struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Verbs     []string `json:"verbs"`
	Resources []string `json:"resourses"`
}

type RoleBinding struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	AccountName string `json:"account-name"`
	RoleName    string `json:"role-name"`
}

type ServiceAccount struct {
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	SecretNamespace string `json:"secret-namespace"`
	SecretName      string `json:"secret-name"`
}
