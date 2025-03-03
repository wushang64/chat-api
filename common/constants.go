package common

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

var StartTime = time.Now().Unix() // unit: second
var Version = "v0.0.0"            // this hard coding will be replaced automatically when building, no need to manually change
var SystemName = "Chat API"
var SystemText = ""
var ServerAddress = "http://localhost:3000"
var PayAddress = ""
var EpayId = ""
var EpayKey = ""
var Price = 7.3
var RedempTionCount = 30
var Footer = ""
var Logo = ""
var TopUpLink = ""
var ChatLink = ""
var QuotaPerUnit = 500 * 1000.0 // $0.002 / 1K tokens
var DisplayInCurrencyEnabled = true
var DisplayTokenStatEnabled = true
var EmailNotificationsEnabled = true
var WxPusherNotificationsEnabled = true
var ModelRatioEnabled = true
var BillingByRequestEnabled = true
var GroupEnable = true
var LogContentEnabled = true
var Wx = true
var Zfb = true
var DrawingEnabled = true
var DataExportEnabled = true
var DataExportInterval = 5 // unit: minute
var MiniQuota = 1.0
var ProporTions = 10
var UserGroup = "default"
var VipUserGroup = "default"

// Any options with "Secret", "Token" in its key won't be return by GetOptions

var SessionSecret = uuid.New().String()

var OptionMap map[string]string
var OptionMapRWMutex sync.RWMutex

var ItemsPerPage = 10
var MaxRecentItems = 100

var PasswordLoginEnabled = true
var PasswordRegisterEnabled = true
var EmailVerificationEnabled = false
var GitHubOAuthEnabled = false
var WeChatAuthEnabled = false
var TurnstileCheckEnabled = false
var RegisterEnabled = true
var ApproximateTokenEnabled = false

var EmailDomainRestrictionEnabled = false
var EmailDomainWhitelist = []string{
	"gmail.com",
	"163.com",
	"126.com",
	"qq.com",
	"outlook.com",
	"hotmail.com",
	"icloud.com",
	"yahoo.com",
	"foxmail.com",
}

var DebugEnabled = os.Getenv("DEBUG") == "true"
var MemoryCacheEnabled = os.Getenv("MEMORY_CACHE_ENABLED") == "true"

var LogConsumeEnabled = true

var SMTPServer = ""
var SMTPPort = 587
var SMTPAccount = ""
var SMTPFrom = ""
var SMTPToken = ""

var AppToken = ""
var Uids = ""

var GitHubClientId = ""
var GitHubClientSecret = ""

var WeChatServerAddress = ""
var WeChatServerToken = ""
var WeChatAccountQRCodeImageURL = ""

var TurnstileSiteKey = ""
var TurnstileSecretKey = ""

var QuotaForNewUser = 0
var QuotaForInviter = 0
var QuotaForInvitee = 0
var ChannelDisableThreshold = 5.0
var AutomaticDisableChannelEnabled = false
var AutomaticEnableChannelEnabled = false
var TopupRatioEnabled = true
var TopupAmountEnabled = false
var QuotaRemindThreshold = 1000
var PreConsumedQuota = 500

var RetryTimes = 0

var RootUserEmail = ""
var GeminiSafetySetting = GetOrDefaultString("GEMINI_SAFETY_SETTING", "BLOCK_NONE")
var IsMasterNode = os.Getenv("NODE_TYPE") != "slave"

var requestInterval, _ = strconv.Atoi(os.Getenv("POLLING_INTERVAL"))
var RequestInterval = time.Duration(requestInterval) * time.Second

var SyncFrequency = GetOrDefault("SYNC_FREQUENCY", 60) // unit is second

var BatchUpdateEnabled = false
var BatchUpdateInterval = GetOrDefault("BATCH_UPDATE_INTERVAL", 5)

var RelayTimeout = GetOrDefault("RELAY_TIMEOUT", 0) // unit is second

const (
	RequestIdKey = "X-Oneapi-Request-Id"
)

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
	RoleRootUser   = 100
)

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

// All duration's unit is seconds
// Shouldn't larger then RateLimitKeyExpirationDuration
var (
	GlobalApiRateLimitNum            = GetOrDefault("GLOBAL_API_RATE_LIMIT", 180)
	GlobalApiRateLimitDuration int64 = 3 * 60

	GlobalWebRateLimitNum            = GetOrDefault("GLOBAL_WEB_RATE_LIMIT", 400)
	GlobalWebRateLimitDuration int64 = 3 * 400

	UploadRateLimitNum            = 10
	UploadRateLimitDuration int64 = 60

	DownloadRateLimitNum            = 10
	DownloadRateLimitDuration int64 = 60

	CriticalRateLimitNum            = 20
	CriticalRateLimitDuration int64 = 20 * 60
)

var RateLimitKeyExpirationDuration = 20 * time.Minute

const (
	UserStatusEnabled  = 1 // don't use 0, 0 is the default value!
	UserStatusDisabled = 2 // also don't use 0
)

const (
	TokenStatusEnabled   = 1 // don't use 0, 0 is the default value!
	TokenStatusDisabled  = 2 // also don't use 0
	TokenStatusExpired   = 3
	TokenStatusExhausted = 4
)

const (
	RedemptionCodeStatusEnabled  = 1 // don't use 0, 0 is the default value!
	RedemptionCodeStatusDisabled = 2 // also don't use 0
	RedemptionCodeStatusUsed     = 3 // also don't use 0
)

const (
	ChannelStatusUnknown          = 0
	ChannelStatusEnabled          = 1 // don't use 0, 0 is the default value!
	ChannelStatusManuallyDisabled = 2 // also don't use 0
	ChannelStatusAutoDisabled     = 3
)

const (
	ChannelTypeUnknown        = 0
	ChannelTypeOpenAI         = 1
	ChannelTypeAPI2D          = 2
	ChannelTypeAzure          = 3
	ChannelTypeCloseAI        = 4
	ChannelTypeOpenAISB       = 5
	ChannelTypeOpenAIMax      = 6
	ChannelTypeOhMyGPT        = 7
	ChannelTypeCustom         = 8
	ChannelTypeAILS           = 9
	ChannelTypeAIProxy        = 10
	ChannelTypePaLM           = 11
	ChannelTypeAPI2GPT        = 12
	ChannelTypeAIGC2D         = 13
	ChannelTypeAnthropic      = 14
	ChannelTypeBaidu          = 15
	ChannelTypeZhipu          = 16
	ChannelTypeAli            = 17
	ChannelTypeXunfei         = 18
	ChannelType360            = 19
	ChannelTypeOpenRouter     = 20
	ChannelTypeAIProxyLibrary = 21
	ChannelTypeFastGPT        = 22
	ChannelTypeTencent        = 23
	ChannelTypeGemini         = 24
	ChannelTypeChatBot        = 25
	ChannelTypeLobeChat       = 26
	ChannelTypeMoonshot       = 27
	ChannelTypeZhipu_v4       = 28
	ChannelTypeStability      = 29
)

var ChannelBaseURLs = []string{
	"",                                  // 0
	"https://api.openai.com",            // 1
	"https://oa.api2d.net",              // 2
	"",                                  // 3
	"https://api.closeai-proxy.xyz",     // 4
	"https://api.openai-sb.com",         // 5
	"https://api.openaimax.com",         // 6
	"https://api.ohmygpt.com",           // 7
	"",                                  // 8
	"https://api.caipacity.com",         // 9
	"https://api.aiproxy.io",            // 10
	"",                                  // 11
	"https://api.api2gpt.com",           // 12
	"https://api.aigc2d.com",            // 13
	"https://api.anthropic.com",         // 14
	"https://aip.baidubce.com",          // 15
	"https://open.bigmodel.cn",          // 16
	"https://dashscope.aliyuncs.com",    // 17
	"",                                  // 18
	"https://ai.360.cn",                 // 19
	"https://openrouter.ai/api",         // 20
	"https://api.aiproxy.io",            // 21
	"https://fastgpt.run/api/openapi",   // 22
	"https://hunyuan.cloud.tencent.com", //23
	"https://generativelanguage.googleapis.com", //24
	"",                         //25
	"",                         //26
	"https://api.moonshot.cn",  //27
	"https://open.bigmodel.cn", //28
	"https://api.stability.ai", //29
}

const (
	ConfigKeyPrefix = "cfg_"

	ConfigKeyAPIVersion = ConfigKeyPrefix + "api_version"
	ConfigKeyLibraryID  = ConfigKeyPrefix + "library_id"
	ConfigKeyPlugin     = ConfigKeyPrefix + "plugin"
)
