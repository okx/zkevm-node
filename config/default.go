package config

// DefaultValues is the default configuration
const DefaultValues = `
IsTrustedSequencer = false
ForkUpgradeBatchNumber = 0
ForkUpgradeNewForkId = 0
Fork9UpgradeBatch = 0

[Log]
Environment = "development" # "production" or "development"
Level = "info"
Outputs = ["stderr"]

[State]
	[State.DB]
	User = "state_user"
	Password = "state_password"
	Name = "state_db"
	Host = "xlayer-state-db"
	Port = "5432"
	EnableLog = false	
	MaxConns = 200
	[State.Batch]
		[State.Batch.Constraints]
		MaxTxsPerBatch = 300
		MaxBatchBytesSize = 120000
		MaxCumulativeGasUsed = 1125899906842624
		MaxKeccakHashes = 2145
		MaxPoseidonHashes = 252357
		MaxPoseidonPaddings = 135191
		MaxMemAligns = 236585
		MaxArithmetics = 236585
		MaxBinaries = 473170
		MaxSteps = 7570538
		MaxSHA256Hashes = 1596

[Pool]
FreeClaimGasLimit = 150000
IntervalToRefreshBlockedAddresses = "5m"
EnableWhitelist = false
IntervalToRefreshWhiteAddresses = "1m"
IntervalToRefreshGasPrices = "5s"
MaxTxBytesSize=100132
MaxTxDataBytesSize=100000
DefaultMinGasPriceAllowed = 1000000000
MinAllowedGasPriceInterval = "5m"
PollMinAllowedGasPriceInterval = "15s"
AccountQueue = 64
GlobalQueue = 1024
FreeGasAddress = ["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"]
TxFeeCap = 1.0
    [Pool.EffectiveGasPrice]
	Enabled = false
	L1GasPriceFactor = 0.25
	ByteGasCost = 16
	ZeroByteGasCost = 4
	NetProfit = 1
	BreakEvenFactor = 1.1	
	FinalDeviationPct = 10
	EthTransferGasPrice = 0
	EthTransferL1GasPriceFactor = 0	
	L2GasPriceSuggesterFactor = 0.5
    [Pool.DB]
	User = "pool_user"
	Password = "pool_password"
	Name = "pool_db"
	Host = "xlayer-pool-db"
	Port = "5432"
	EnableLog = false
	MaxConns = 200

[Etherman]
URL = "http://localhost:8545"
ForkIDChunkSize = 20000
MultiGasProvider = false
	[Etherman.Etherscan]
		ApiKey = ""

[EthTxManager]
FrequencyToMonitorTxs = "1s"
WaitTxToBeMined = "2m"
ForcedGas = 0
GasPriceMarginFactor = 1
MaxGasPriceLimit = 0
	[EthTxManager.CustodialAssets]
		Enable = false
		URL = "http://localhost:8080"
		Symbol = 2882
		SequencerAddr = "0x1a13bddcc02d363366e04d4aa588d3c125b0ff6f"
		AggregatorAddr = "0x66e39a1e507af777e8c385e2d91559e20e306303"
		WaitResultTimeout = "2m"
		OperateTypeSeq = 3
		OperateTypeAgg = 4
		ProjectSymbol = 3011
		OperateSymbol = 2
		SysFrom = 3
		UserID = 0
		OperateAmount = 0
		RequestSignURI = "/priapi/v1/assetonchain/ecology/ecologyOperate"
		QuerySignURI = "/priapi/v1/assetonchain/ecology/querySignDataByOrderNo"

[RPC]
Host = "0.0.0.0"
Port = 8545
ReadTimeout = "60s"
WriteTimeout = "60s"
MaxRequestsPerIPAndSecond = 500
SequencerNodeURI = ""
EnableL2SuggestedGasPricePolling = true
BatchRequestsEnabled = false
BatchRequestsLimit = 20
MaxLogsCount = 10000
MaxLogsBlockRange = 10000
MaxNativeBlockHashBlockRange = 60000
EnableHttpLog = true
GasLimitFactor = 1
DisableAPIs = []
	[RPC.RateLimit]
		Enabled = false
		RateLimitApis = []
		RateLimitCount = 100
		RateLimitDuration = 1
		SpecialApis = []
	[RPC.WebSockets]
		Enabled = true
		Host = "0.0.0.0"
		Port = 8546
		ReadLimit = 104857600
	[RPC.DynamicGP]
        Enabled = false
        CongestionTxThreshold = 100
        CheckBatches = 5
        SampleNumber = 3
        Percentile = 70
        MaxPrice = 20000000000
        MinPrice = 2000000000
        UpdatePeriod = "10s"
	[RPC.ApiAuthentication]
		Enabled = false
		ApiKeys = []
	[RPC.ApiRelay]
		Enabled = false
		DestURI = "" 
		RPCs = []

[Synchronizer]
SyncInterval = "1s"
SyncChunkSize = 100
TrustedSequencerURL = "" # If it is empty or not specified, then the value is read from the smc
SyncBlockProtection = "safe" # latest, finalized, safe
L1SynchronizationMode = "sequential"
L1SyncCheckL2BlockHash = true
L1SyncCheckL2BlockNumberhModulus = 600
	[Synchronizer.L1BlockCheck]
		Enabled = true
		L1SafeBlockPoint = "finalized"
		L1SafeBlockOffset = 0
		ForceCheckBeforeStart = true
		PreCheckEnabled = true
		L1PreSafeBlockPoint = "safe"
		L1PreSafeBlockOffset = 0
	[Synchronizer.L1ParallelSynchronization]
		MaxClients = 10
		MaxPendingNoProcessedBlocks = 25
		RequestLastBlockPeriod = "5s"
		RequestLastBlockTimeout = "5s"
		RequestLastBlockMaxRetries = 3
		StatisticsPeriod = "5m"
		TimeoutMainLoop = "5m"
		RollupInfoRetriesSpacing= "5s"
		FallbackToSequentialModeOnSynchronized = false
		[Synchronizer.L1ParallelSynchronization.PerformanceWarning]
			AceptableInacctivityTime = "5s"
			ApplyAfterNumRollupReceived = 10
	[Synchronizer.L2Synchronization]
		Enabled = true
		AcceptEmptyClosedBatches = false
		ReprocessFullBatchOnClose = false
		CheckLastL2BlockHashOnCloseBatch = true

[Sequencer]
DeletePoolTxsL1BlockConfirmations = 100
DeletePoolTxsCheckInterval = "12h"
TxLifetimeCheckInterval = "10m"
TxLifetimeMax = "3h"
LoadPoolTxsCheckInterval = "500ms"
StateConsistencyCheckInterval = "5s"
	[Sequencer.Finalizer]
		NewTxsWaitInterval = "100ms"
		ForcedBatchesTimeout = "60s"
		ForcedBatchesL1BlockConfirmations = 64
		ForcedBatchesCheckInterval = "10s"
		L1InfoTreeL1BlockConfirmations = 64
		L1InfoTreeCheckInterval = "10s"
		BatchMaxDeltaTimestamp = "1800s"
		L2BlockMaxDeltaTimestamp = "3s"
		ResourceExhaustedMarginPct = 10
		StateRootSyncInterval = "3600s"
		FlushIdCheckInterval = "50ms"
		HaltOnBatchNumber = 0
		SequentialBatchSanityCheck = false
		SequentialProcessL2Block = false
		FullBatchSleepDuration = "0s"
	[Sequencer.Finalizer.Metrics]
		Interval = "60m"
		EnableLog = true
	[Sequencer.StreamServer]
		Port = 0
		Filename = ""
		Version = 0
		WriteTimeout = "5s"
		InactivityTimeout = "120s"
		InactivityCheckInterval = "5s"
		Enabled = false

[SequenceSender]
WaitPeriodSendSequence = "5s"
LastBatchVirtualizationTimeMaxWaitPeriod = "5s"
L1BlockTimestampMargin = "30s"
MaxTxSizeForL1 = 131072
MaxBatchesForL1 = 10
SequenceL1BlockConfirmations = 32
L2Coinbase = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
DAPermitApiPrivateKey = {Path = "/pk/sequencer.keystore", Password = "testonly"}
GasOffset = 80000

[Aggregator]
Host = "0.0.0.0"
Port = 50081
Parallel = false
ParaCount = 0
RetryTime = "5s"
VerifyProofInterval = "90s"
TxProfitabilityCheckerType = "acceptall"
TxProfitabilityMinReward = "1.1"
ProofStatePollingInterval = "5s"
CleanupLockedProofsInterval = "2m"
GeneratingProofCleanupThreshold = "10m"
GasOffset = 0
UpgradeEtrogBatchNumber = 0
BatchProofL1BlockConfirmations = 2
SettlementBackend = "l1"
AggLayerTxTimeout = "5m"
AggLayerURL = ""
SequencerPrivateKey = {}

[L2GasPriceSuggester]
Type = "follower"
UpdatePeriod = "10s"
Factor = 0.15
DefaultGasPriceWei = 2000000000
MaxGasPriceWei = 0
CleanHistoryPeriod = "1h"
CleanHistoryTimeRetention = "5m"

[MTClient]
URI = "xlayer-prover:50061"

[Executor]
URI = "xlayer-prover:50071"
MaxResourceExhaustedAttempts = 3
WaitOnResourceExhaustion = "1s"
MaxGRPCMessageSize = 100000000

[Metrics]
Host = "0.0.0.0"
Port = 9091
Enabled = false

[HashDB]
User = "prover_user"
Password = "prover_pass"
Name = "prover_db"
Host = "xlayer-state-db"
Port = "5432"
EnableLog = false
MaxConns = 200
`
