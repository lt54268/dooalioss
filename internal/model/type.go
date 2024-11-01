package model

// Config 用于存储配置信息
type Config struct {
	Port     string
	Region   string
	Endpoint string
	Bucket   string
}

type UploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	URL  string `json:"url,omitempty"`
}
