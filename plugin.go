package main

const (
	KUBECTL  = "/root/.kube/config"
	HELM_BIN = "/bin/helm"
)

type (
	// Config maps the params we need to run Helm
	Config struct {
		GitName    string `json:"git_name"`
		GitEmail   string `json:"git_email"`
		InputFile  string `json:"input_file"`
		OutputFile string `json:"output_file"`
	}
	// Plugin default
	Plugin struct {
		Config  Config
		command []string
	}
)

func (p *Plugin) Exec() error {

	return nil
}
