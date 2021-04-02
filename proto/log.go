package proto

type LogLevel int

const (
	LogLevelOff           LogLevel = 0
	LogLevelAdmin         LogLevel = 2
	LogLevelL1Data        LogLevel = 4
	LogLevelL1Request     LogLevel = 8
	LogLevelL1System      LogLevel = 16
	LogLevelL1Error       LogLevel = 32
	LogLevelL2Data        LogLevel = 64
	LogLevelL2Request     LogLevel = 128
	LogLevelL2System      LogLevel = 256
	LogLevelL2Error       LogLevel = 512
	LogLevelLookupData    LogLevel = 1024
	LogLevelLookupRequest LogLevel = 2048
	LogLevelLookupError   LogLevel = 4096
	LogLevelInformation   LogLevel = 8192
	LogLevelDebug         LogLevel = 16384
	LogLevelConnectivity  LogLevel = 32768
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelOff:
		return "Off"
	case LogLevelAdmin:
		return "Admin"
	case LogLevelL1Data:
		return "L1Data"
	case LogLevelL1Request:
		return "L1Request"
	case LogLevelL1System:
		return "L1System"
	case LogLevelL1Error:
		return "L1Error"
	case LogLevelL2Data:
		return "L2Data"
	case LogLevelL2Request:
		return "L2Request"
	case LogLevelL2System:
		return "L2System"
	case LogLevelL2Error:
		return "L2Error"
	case LogLevelLookupData:
		return "LookupData"
	case LogLevelLookupRequest:
		return "LookupRequest"
	case LogLevelLookupError:
		return "LookupError"
	case LogLevelInformation:
		return "Information"
	case LogLevelDebug:
		return "Debug"
	case LogLevelConnectivity:
		return "Connectivity"
	default:
		return "Off"
	}
}
