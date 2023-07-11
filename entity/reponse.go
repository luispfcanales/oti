package entity

type ReponseApi struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}
