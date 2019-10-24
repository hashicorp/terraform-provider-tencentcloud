package tencentcloud

/*
 all gate way types
 https://cloud.tencent.com/document/api/215/15824#Route
*/
const GATE_WAY_TYPE_CVM = "CVM"
const GATE_WAY_TYPE_VPN = "VPN"
const GATE_WAY_TYPE_DIRECTCONNECT = "DIRECTCONNECT"
const GATE_WAY_TYPE_SSLVPN = "SSLVPN"
const GATE_WAY_TYPE_NAT = "NAT"
const GATE_WAY_TYPE_NORMAL_CVM = "NORMAL_CVM"
const GATE_WAY_TYPE_EIP = "EIP"
const GATE_WAY_TYPE_CCN = "CCN"

var ALL_GATE_WAY_TYPES = []string{GATE_WAY_TYPE_CVM,
	GATE_WAY_TYPE_VPN,
	GATE_WAY_TYPE_DIRECTCONNECT,
	GATE_WAY_TYPE_SSLVPN,
	GATE_WAY_TYPE_NAT,
	GATE_WAY_TYPE_NORMAL_CVM,
	GATE_WAY_TYPE_EIP,
	GATE_WAY_TYPE_CCN,
}

/*
EIP
*/
const (
	EIP_STATUS_CREATING  = "CREATING"
	EIP_STATUS_BINDING   = "BINDING"
	EIP_STATUS_BIND      = "BIND"
	EIP_STATUS_UNBINDING = "UNBINDING"
	EIP_STATUS_UNBIND    = "UNBIND"
	EIP_STATUS_OFFLINING = "OFFLINING"
	EIP_STATUS_BIND_ENI  = "BIND_ENI"

	EIP_TYPE_EIP     = "EIP"
	EIP_TYPE_ANYCAST = "AnycastEIP"

	EIP_ANYCAST_ZONE_GLOBAL   = "ANYCAST_ZONE_GLOBAL"
	EIP_ANYCAST_ZONE_OVERSEAS = "ANYCAST_ZONE_OVERSEAS"

	EIP_INTERNET_PROVIDER_BGP  = "BGP"
	EIP_INTERNET_PROVIDER_CMCC = "CMCC"
	EIP_INTERNET_PROVIDER_CTCC = "CTCC"
	EIP_INTERNET_PROVIDER_CUCC = "CUCC"
)

var EIP_INTERNET_PROVIDER = []string{
	EIP_INTERNET_PROVIDER_BGP,
	EIP_INTERNET_PROVIDER_CMCC,
	EIP_INTERNET_PROVIDER_CTCC,
	EIP_INTERNET_PROVIDER_CUCC,
}

var EIP_TYPE = []string{
	EIP_TYPE_EIP,
	EIP_TYPE_ANYCAST,
}

var EIP_ANYCAST_ZONE = []string{
	EIP_ANYCAST_ZONE_GLOBAL,
	EIP_ANYCAST_ZONE_OVERSEAS,
}

// ENI
const (
	ENI_DESCRIBE_LIMIT = 100
)

const (
	ENI_STATE_PENDING   = "PENDING"
	ENI_STATE_AVAILABLE = "AVAILABLE"
	ENI_STATE_ATTACHING = "ATTACHING"
	ENI_STATE_DETACHING = "DETACHING"
	ENI_STATE_DELETING  = "DELETING"
)

const (
	ENI_IP_PENDING   = "PENDING"
	ENI_IP_AVAILABLE = "AVAILABLE"
	ENI_IP_ATTACHING = "ATTACHING"
	ENI_IP_DETACHING = "DETACHING"
	ENI_IP_DELETING  = "DELETING"
)

const (
	VPCNotFound = "ResourceNotFound"
)
