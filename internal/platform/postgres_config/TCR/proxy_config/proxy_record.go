package proxy_config

type ProxyRecord struct {
	ID   int64  `json:"id" db:"id"`
	UUID string `json:"uuid" db:"uuid"`
}
