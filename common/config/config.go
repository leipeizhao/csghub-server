package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Saas          bool   `envconfig:"STARHUB_SERVER_SAAS" default:"true"`
	InstanceID    string `envconfig:"STARHUB_SERVER_INSTANCE_ID"`
	EnableSwagger bool   `envconfig:"STARHUB_SERVER_ENABLE_SWAGGER" default:"false"`
	// enable if you want to acess csghub through https, especially for space rproxy
	EnableHTTPS bool   `envconfig:"STARHUB_SERVER_ENABLE_HTTPS" default:"false"`
	APIToken    string `envconfig:"STARHUB_SERVER_API_TOKEN" default:"f3a7b9c1d6e5f8e2a1b5d4f9e6a2b8d7c3a4e2b1d9f6e7a8d2c5a7b4c1e3f5b8a1d4f9b7d6e2f8a5d3b1e7f9c6a8b2d1e4f7d5b6e9f2a4b3c8e1d7f995hd82hf"`

	APIServer struct {
		Port         int    `envconfig:"STARHUB_SERVER_SERVER_PORT" default:"8080"`
		PublicDomain string `envconfig:"STARHUB_SERVER_PUBLIC_DOMAIN" default:"localhost:8080"`
	}

	DocsHost string `envconfig:"STARHUB_SERVER_SERVER_DOCS_HOST" default:"http://localhost:6636"`

	Database struct {
		Driver   string `envconfig:"STARHUB_DATABASE_DRIVER" default:"pg"`
		DSN      string `envconfig:"STARHUB_DATABASE_DSN" default:"postgresql://postgres:postgres@localhost:5432/starhub_server?sslmode=disable"`
		TimeZone string `envconfig:"STARHUB_DATABASE_TIMEZONE" default:"Asia/Shanghai"`
	}

	Redis struct {
		Endpoint           string `envconfig:"STARHUB_SERVER_REDIS_ENDPOINT"              default:"localhost:6379"`
		MaxRetries         int    `envconfig:"STARHUB_SERVER_REDIS_MAX_RETRIES"           default:"3"`
		MinIdleConnections int    `envconfig:"STARHUB_SERVER_REDIS_MIN_IDLE_CONNECTIONS"  default:"0"`
		User               string `envconfig:"STARHUB_SERVER_REDIS_USER"`
		Password           string `envconfig:"STARHUB_SERVER_REDIS_PASSWORD"`
		SentinelMode       bool   `envconfig:"STARHUB_SERVER_REDIS_USE_SENTINEL"          default:"false"`
		SentinelMaster     string `envconfig:"STARHUB_SERVER_REDIS_SENTINEL_MASTER"`
		SentinelEndpoint   string `envconfig:"STARHUB_SERVER_REDIS_SENTINEL_ENDPOINT"`
	}

	GitServer struct {
		URL       string `envconfig:"STARHUB_SERVER_GITSERVER_URL"    default:"http://localhost:3000"`
		Type      string `envconfig:"STARHUB_SERVER_GITSERVER_TYPE"    default:"gitea"`
		Host      string `envconfig:"STARHUB_SERVER_GITSERVER_HOST"       default:"http://localhost:3000"`
		SecretKey string `envconfig:"STARHUB_SERVER_GITSERVER_SECRET_KEY" default:"619c849c49e03754454ccd4cda79a209ce0b30b3"`
		Username  string `envconfig:"STARHUB_SERVER_GITSERVER_USERNAME" default:"root"`
		Password  string `envconfig:"STARHUB_SERVER_GITSERVER_PASSWORD" default:"password123"`
	}

	MirrorServer struct {
		Enable    bool   `envconfig:"STARHUB_SERVER_MIRRORSERVER_ENABLE" default:"true"`
		URL       string `envconfig:"STARHUB_SERVER_MIRRORSERVER_URL"    default:"http://localhost:3001"`
		Type      string `envconfig:"STARHUB_SERVER_MIRRORSERVER_TYPE"    default:"gitea"`
		Host      string `envconfig:"STARHUB_SERVER_MIRRORSERVER_HOST"       default:"http://localhost:3001"`
		SecretKey string `envconfig:"STARHUB_SERVER_MIRRORSERVER_SECRET_KEY" default:"619c849c49e03754454ccd4cda79a209ce0b30b3"`
		Username  string `envconfig:"STARHUB_SERVER_MIRRORSERVER_USERNAME" default:"root"`
		Password  string `envconfig:"STARHUB_SERVER_MIRRORSERVER_PASSWORD" default:"password123"`
	}

	Frontend struct {
		URL string `envconfig:"STARHUB_SERVER_FRONTEND_URL" default:"https://portal-stg.opencsg.com"`
	}

	S3 struct {
		AccessKeyID     string `envconfig:"STARHUB_SERVER_S3_ACCESS_KEY_ID"`
		AccessKeySecret string `envconfig:"STARHUB_SERVER_S3_ACCESS_KEY_SECRET"`
		Region          string `envconfig:"STARHUB_SERVER_S3_REGION"`
		Endpoint        string `envconfig:"STARHUB_SERVER_S3_ENDPOINT" default:"oss-cn-beijing.aliyuncs.com"`
		Bucket          string `envconfig:"STARHUB_SERVER_S3_BUCKET" default:"opencsg-test"`
	}

	SensitiveCheck struct {
		Enable          bool   `envconfig:"STARHUB_SERVER_SENSITIVE_CHECK_ENABLE" default:"false"`
		AccessKeyID     string `envconfig:"STARHUB_SERVER_SENSITIVE_CHECK_ACCESS_KEY_ID"`
		AccessKeySecret string `envconfig:"STARHUB_SERVER_SENSITIVE_CHECK_ACCESS_KEY_SECRET"`
		Region          string `envconfig:"STARHUB_SERVER_SENSITIVE_CHECK_REGION"`
		Endpoint        string `envconfig:"STARHUB_SERVER_SENSITIVE_CHECK_ENDPOINT" default:"oss-cn-beijing.aliyuncs.com"`
	}

	JWT struct {
		SigningKey string `envconfig:"STARHUB_JWT_SIGNING_KEY" default:"signing-key"`
	}

	Inference struct {
		ServerAddr string `envconfig:"STARHUB_SERVER_INFERENCE_SERVER_ADDR" default:"http://localhost:8000"`
	}

	Space struct {
		BuilderEndpoint string `envconfig:"STARHUB_SERVER_SPACE_BUILDER_ENDPOINT" default:"http://localhost:8081"`
		// base url for space api running in k8s cluster
		RunnerEndpoint string `envconfig:"STARHUB_SERVER_SPACE_RUNNER_ENDPOINT" default:"http://localhost:8082"`

		// the internal root domain will be proxied to, should be internal access only
		InternalRootDomain string `envconfig:"STARHUB_SERVER_INTERNAL_ROOT_DOMAIN" default:"internal.example.com"`
		// the public root domain will be proxied from
		PublicRootDomain string `envconfig:"STARHUB_SERVER_PUBLIC_ROOT_DOMAIN" default:"public.example.com"`
		DockerRegBase    string `envconfig:"STARHUB_SERVER_DOCKER_REG_BASE" default:"registry.cn-beijing.aliyuncs.com/opencsg_public/"`
		ImagePullSecret  string `envconfig:"STARHUB_SERVER_DOCKER_IMAGE_PULL_SECRET" default:"opencsg-pull-secret"`
		// reverse proxy listening port
		RProxyServerPort int `envconfig:"STARHUB_SERVER_SPACE_RPROXY_SERVER_PORT" default:"8083"`
		// secret key for session encryption
		SessionSecretKey   string `envconfig:"STARHUB_SERVER_SPACE_SESSION_SECRET_KEY" default:"secret"`
		DeployTimeoutInMin int    `envconfig:"STARHUB_SERVER_SPACE_DEPLOY_TIMEOUT_IN_MINUTES" default:"30"`
		// gpu model label
		GPUModelLabel string `envconfig:"STARHUB_SERVER_GPU_MODEL_LABEL" default:"aliyun.accelerator/nvidia_name"`
	}

	Model struct {
		DeployTimeoutInMin int    `envconfig:"STARHUB_SERVER_MODEL_DEPLOY_TIMEOUT_IN_MINUTES" default:"30"`
		DownloadEndpoint   string `envconfig:"STARHUB_SERVER_MODEL_DOWNLOAD_ENDPOINT" default:"https://hub-stg.opencsg.com/"`
		DockerRegBase      string `envconfig:"STARHUB_SERVER_MODEL_DOCKER_REG_BASE" default:"opencsg-registry.cn-beijing.cr.aliyuncs.com/public/"`
	}

	Casdoor struct {
		ClientID         string `envconfig:"STARHUB_SERVER_CASDOOR_CLIENT_ID" default:"client_id"`
		ClientSecret     string `envconfig:"STARHUB_SERVER_CASDOOR_CLIENT_SECRET" default:"client_secret"`
		Endpoint         string `envconfig:"STARHUB_SERVER_CASDOOR_ENDPOINT" default:"http://localhost:80"`
		Certificate      string `envconfig:"STARHUB_SERVER_CASDOOR_CERTIFICATE" default:"/etc/casdoor/certificate.pem"`
		OrganizationName string `envconfig:"STARHUB_SERVER_CASDOOR_ORGANIZATION_NAME" default:"opencsg"`
		ApplicationName  string `envconfig:"STARHUB_SERVER_CASDOOR_APPLICATION_NAME" default:"opencsg"`
	}
}

func LoadConfig() (cfg *Config, err error) {
	cfg = new(Config)
	err = envconfig.Process("", cfg)
	if err != nil {
		return
	}

	return
}
