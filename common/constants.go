package common

const (
	RewardMasterPercent        = 40
	RewardVoterPercent         = 30
	RewardFoundationPercent    = 30
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 21
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	MinGasPrice                = 2500

	// Maximum number of Masternodes Set to 21
)

var IsTestnet bool = false
