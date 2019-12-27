package pkg

type (
	ApiConfig struct {
		APIUser  string `yaml:"apiUser"`
		Username string `yaml:"username"`
		Token    string `yaml:"token"`
	}

	Config struct {
		API ApiConfig `yaml:"api"`
	}

	SyncRecordConfig struct {
		Type      string `yaml:"type"`
		Name      string `yaml:"name"`
		Address   string `yaml:"address"`
		TTL       int    `yaml:"ttl"`
		CreatePTR bool   `yaml:"createPTR"`
	}

	SyncConfig struct {
		Zone    string             `yaml:"zone"`
		Records []SyncRecordConfig `yaml:"records"`
	}
)
