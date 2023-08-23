package conf

func newConfig() *Config {
	return &Config{}
}

type Config struct {
	FEISHU_BOT_URL   string `env:"FEISHU_BOT_URL"`
	DINGDING_BOT_URL string `env:"DINGDING_BOT_URL"`
	WECHAT_BOT_URL   string `env:"WECHAT_BOT_URL"`

	DEPLOY_CLUSTER_ID string `env:"DEPLOY_CLUSTER_ID"`
	DEPLOY_ID         string `env:"DEPLOY_ID"`
	BUILD_ID          string `env:"BUILD_ID"`
	MCENTER_BUILD_ID  string `env:"MCENTER_BUILD_ID"`
	SERVICE_ID        string `env:"SERVICE_ID"`
	PIPELINE_TASK_ID  string `env:"PIPELINE_TASK_ID"`

	DEPLOY_JOB_ID     string `env:"DEPLOY_JOB_ID"`
	BUILD_JOB_ID      string `env:"BUILD_JOB_ID"`
	CICD_PIPELINE_ID  string `env:"CICD_PIPELINE_ID"`
	MFLOW_PIPELINE_ID string `env:"MFLOW_PIPELINE_ID"`

	DEVCLOUD_DEPLOY_APPROVAL_ID string `env:"DEVCLOUD_DEPLOY_APPROVAL_ID"`

	MCENTER_BUILD_TASK_ID     string `env:"MCENTER_BUILD_TASK_ID"`
	MCENTER_BUILD_TASK_TOKEN  string `env:"MCENTER_BUILD_TASK_TOKEN"`
	MCENTER_DEPLOY_TASK_ID    string `env:"MCENTER_DEPLOY_TASK_ID"`
	MCENTER_DEPLOY_TASK_TOKEN string `env:"MCENTER_DEPLOY_TASK_TOKEN"`
	MCENTER_DEPLOY_ID         string `env:"MCENTER_DEPLOY_ID"`
	MCENTER_SERVICE_ID        string `env:"MCENTER_SERVICE_ID"`
	MCENTER_GRPC_ADDRESS      string `env:"MCENTER_GRPC_ADDRESS"`
	MCENTER_CLINET_ID         string `env:"MCENTER_CLINET_ID"`
	MCENTER_CLIENT_SECRET     string `env:"MCENTER_CLIENT_SECRET"`
}