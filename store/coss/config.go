package coss

const (
	ENDPOINT_COS = "%s.cos.%s.myqcloud.com"
	ENDPOINT_S3  = "s3.%s.amazonaws.com"
	ENDPOINT_OSS = "oss-%s.aliyuncs.com"
	ENDPOINT_OBS = "obs.%s.myhuaweicloud.com"
)

type Config struct {
	Source       string `json:"source"        toml:"source"        yaml:"source"` // oss / cos / obs / s3
	Region       string `json:"region"        toml:"region"        yaml:"region"`
	Bucket       string `json:"bucket"        toml:"bucket"        yaml:"bucket"`
	AccessId     string `json:"access_id"     toml:"access_id"     yaml:"access_id"`
	AccessSecret string `json:"access_secret" toml:"access_secret" yaml:"access_secret"`
	Domain       string `json:"domain"        toml:"domain"        yaml:"domain"`

	Endpoint string `json:"-" toml:"-" yaml:"-"`
}

var endpoints = map[string]string{
	COS: ENDPOINT_COS,
	S3:  ENDPOINT_S3,
	OSS: ENDPOINT_OSS,
	OBS: ENDPOINT_OBS,
}
