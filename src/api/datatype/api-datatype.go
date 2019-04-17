package datatype

const (
	ROLE_ROOT         = "ROOT"
	ROLE_ADMIN        = "ADMIN"
	ROLE_USER         = "USER"
	OP_TYPE_SET_SNMP  = "SET SNMP"
	OP_TYPE_SET_NTP   = "SET NTP"
	OP_TYPE_CREATE_USER   = "CREATE USER"
	OP_TYPE_DELETE_USER   = "DELETE USER"
	OP_TYPE_CREATE_VIEW   = "CREATE VIEW"
	OP_TYPE_DELETE_VIEW   = "DELETE VIEW"
	OP_TYPE_LOGIN   = "LOG IN"
	OP_TYPE_ADD_DISK   = "ADD DISK"
	OP_TYPE_CLEAR_DISK   = "CLEAR DISK"
	REQUEST_TOKEN_KEY = "TOKEN"
)

type QueryRequestNtp struct {
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
	Type       string `json:"type,omitempty"`
	NtpAddress string `json:"ntpaddress,omitempty"`
}
type QueryResponse struct {
	RetCode int         `json:"retCode"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type NtpRecord struct {
	NtpAddress string `json:"ntpAddress"`
}

