package common

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	USD2RMB = 7
	USD     = 500 // $0.002 = 1 -> $1 = 500
	RMB     = USD / USD2RMB
)

var ModelRatio = map[string]float64{
	"gpt-4":                   15,
	"gpt-4-0314":              15,
	"gpt-4-0613":              15,
	"gpt-4-32k":               30,
	"gpt-4-32k-0314":          30,
	"gpt-4-32k-0613":          30,
	"gpt-4-1106-preview":      5,    // $0.01 / 1K tokens
	"gpt-4-0125-preview":      5,    // $0.01 / 1K tokens
	"gpt-4-turbo-preview":     5,    // $0.01 / 1K tokens
	"gpt-4-vision-preview":    5,    // $0.01 / 1K tokens
	"gpt-3.5-turbo":           0.75, // $0.0015 / 1K tokens
	"gpt-3.5-turbo-0301":      0.75,
	"gpt-3.5-turbo-0613":      0.75,
	"gpt-3.5-turbo-16k":       1.5, // $0.003 / 1K tokens
	"gpt-3.5-turbo-16k-0613":  1.5,
	"gpt-3.5-turbo-instruct":  0.75, // $0.0015 / 1K tokens
	"gpt-3.5-turbo-1106":      0.5,  // $0.001 / 1K tokens
	"gpt-3.5-turbo-0125":      0.25, // $0.0005 / 1K tokens
	"davinci-002":             1,    // $0.002 / 1K tokens
	"babbage-002":             0.2,  // $0.0004 / 1K tokens
	"text-ada-001":            0.2,
	"text-babbage-001":        0.25,
	"text-curie-001":          1,
	"text-davinci-002":        10,
	"text-davinci-003":        10,
	"text-davinci-edit-001":   10,
	"code-davinci-edit-001":   10,
	"whisper-1":               15,  // $0.006 / minute -> $0.006 / 150 words -> $0.006 / 200 tokens -> $0.03 / 1k tokens
	"tts-1":                   7.5, // $0.015 / 1K characters
	"tts-1-1106":              7.5,
	"tts-1-hd":                15, // $0.030 / 1K characters
	"tts-1-hd-1106":           15,
	"davinci":                 10,
	"curie":                   10,
	"babbage":                 10,
	"ada":                     10,
	"text-embedding-ada-002":  0.05,
	"text-embedding-3-small":  0.01,
	"text-embedding-3-large":  0.065,
	"text-search-ada-doc-001": 10,
	"text-moderation-stable":  0.1,
	"text-moderation-latest":  0.1,
	"dall-e-2":                8,  // $0.016 - $0.020 / image
	"dall-e-3":                20, // $0.040 - $0.120 / image
	// https://www.anthropic.com/api#pricing
	"claude-instant-1.2":       0.8 / 1000 * USD,
	"claude-2.0":               8.0 / 1000 * USD,
	"claude-2.1":               8.0 / 1000 * USD,
	"claude-3-haiku-20240229":  0.25 / 1000 * USD,
	"claude-3-sonnet-20240229": 3.0 / 1000 * USD,
	"claude-3-opus-20240229":   15.0 / 1000 * USD,
	// https://cloud.baidu.com/doc/WENXINWORKSHOP/s/hlrk4akp7
	"ERNIE-Bot":         0.8572,     // ￥0.012 / 1k tokens
	"ERNIE-Bot-turbo":   0.5715,     // ￥0.008 / 1k tokens
	"ERNIE-Bot-4":       0.12 * RMB, // ￥0.12 / 1k tokens
	"ERNIE-Bot-8k":      0.024 * RMB,
	"Embedding-V1":      0.1429, // ￥0.002 / 1k tokens
	"PaLM-2":            1,
	"gemini-pro":        1, // $0.00025 / 1k characters -> $0.001 / 1k tokens
	"gemini-pro-vision": 1, // $0.00025 / 1k characters -> $0.001 / 1k tokens
	// https://open.bigmodel.cn/pricing
	"glm-4":                     0.1 * RMB,
	"glm-4v":                    0.1 * RMB,
	"glm-3-turbo":               0.005 * RMB,
	"chatglm_turbo":             0.3572, // ￥0.005 / 1k tokens
	"chatglm_pro":               0.7143, // ￥0.01 / 1k tokens
	"chatglm_std":               0.3572, // ￥0.005 / 1k tokens
	"chatglm_lite":              0.1429, // ￥0.002 / 1k tokens
	"qwen-turbo":                0.5715, // ￥0.008 / 1k tokens  // https://help.aliyun.com/zh/dashscope/developer-reference/tongyi-thousand-questions-metering-and-billing
	"qwen-plus":                 1.4286, // ￥0.02 / 1k tokens
	"qwen-max":                  1.4286, // ￥0.02 / 1k tokens
	"qwen-max-longcontext":      1.4286, // ￥0.02 / 1k tokens
	"text-embedding-v1":         0.05,   // ￥0.0007 / 1k tokens
	"SparkDesk":                 1.2858, // ￥0.018 / 1k tokens
	"SparkDesk-v1.1":            1.2858, // ￥0.018 / 1k tokens
	"SparkDesk-v2.1":            1.2858, // ￥0.018 / 1k tokens
	"SparkDesk-v3.1":            1.2858, // ￥0.018 / 1k tokens
	"SparkDesk-v3.5":            1.2858, // ￥0.018 / 1k tokens
	"360GPT_S2_V9":              0.8572, // ¥0.012 / 1k tokens
	"embedding-bert-512-v1":     0.0715, // ¥0.001 / 1k tokens
	"embedding_s1_v1":           0.0715, // ¥0.001 / 1k tokens
	"semantic_similarity_s1_v1": 0.0715, // ¥0.001 / 1k tokens
	"hunyuan":                   7.143,  // ¥0.1 / 1k tokens  // https://cloud.tencent.com/document/product/1729/97731#e0e6be58-60c8-469f-bdeb-6c264ce3b4d0
	"ChatStd":                   0.01 * RMB,
	"ChatPro":                   0.1 * RMB,
	// https://platform.moonshot.cn/pricing
	"moonshot-v1-8k":   0.012 * RMB,
	"moonshot-v1-32k":  0.024 * RMB,
	"moonshot-v1-128k": 0.06 * RMB,
	// https://platform.baichuan-ai.com/price
	"Baichuan2-Turbo":      0.008 * RMB,
	"Baichuan2-Turbo-192k": 0.016 * RMB,
	"Baichuan2-53B":        0.02 * RMB,
	// https://api.minimax.chat/document/price
	"abab6-chat":    0.1 * RMB,
	"abab5.5-chat":  0.015 * RMB,
	"abab5.5s-chat": 0.005 * RMB,
	// https://docs.mistral.ai/platform/pricing/
	"open-mistral-7b":       0.25 / 1000 * USD,
	"open-mixtral-8x7b":     0.7 / 1000 * USD,
	"mistral-small-latest":  2.0 / 1000 * USD,
	"mistral-medium-latest": 2.7 / 1000 * USD,
	"mistral-large-latest":  8.0 / 1000 * USD,
	"mistral-embed":         0.1 / 1000 * USD}
var CompletionRatio = map[string]float64{}

var DalleSizeRatios = map[string]map[string]float64{
	"dall-e-2": {
		"256x256":   1,
		"512x512":   1.125,
		"1024x1024": 1.25,
	},
	"dall-e-3": {
		"1024x1024": 1,
		"1024x1792": 2,
		"1792x1024": 2,
	},
}

var DalleGenerationImageAmounts = map[string][2]int{
	"dall-e-2": {1, 10},
	"dall-e-3": {1, 1}, // OpenAI allows n=1 currently.
}

var DalleImagePromptLengthLimitations = map[string]int{
	"dall-e-2": 1000,
	"dall-e-3": 4000,
}

var ModelPrice = map[string]float64{
	"gpt-4-gizmo-*":      0.1,
	"mj_imagine":         0.1,
	"mj_variation":       0.1,
	"mj_reroll":          0.1,
	"mj_blend":           0.1,
	"mj_describe":        0.05,
	"mj_upscale":         0.05,
	"mj_action":          0.1,
	"mj_inpaint":         0.1,
	"mj_swapface":        0.1,
	"mj_shorten":         0.1,
	"mj_uploads":         0.05,
	"mj_turbo_imagine":   0.1,
	"mj_turbo_variation": 0.1,
	"mj_turbo_reroll":    0.1,
	"mj_turbo_blend":     0.1,
	"mj_turbo_describe":  0.1,
	"mj_turbo_upscale":   0.05,
	"mj_turbo_action":    0.1,
	"mj_turbo_inpaint":   0.1,
	"mj_turbo_swapface":  0.1,
	"mj_turbo_shorten":   0.1,
	"mj_turbo_uploads":   0.05,
	"mj_relax_imagine":   0.1,
	"mj_relax_variation": 0.1,
	"mj_relax_reroll":    0.1,
	"mj_relax_blend":     0.1,
	"mj_relax_describe":  0.1,
	"mj_relax_upscale":   0.05,
	"mj_relax_action":    0.1,
	"mj_relax_inpaint":   0.1,
	"mj_relax_swapface":  0.1,
	"mj_relax_shorten":   0.1,
	"mj_relax_uploads":   0.05,
	"midjourney":         0.1,
}

func ModelRatioJSONString() string {
	jsonBytes, err := json.Marshal(ModelRatio)
	if err != nil {
		SysError("error marshalling model ratio: " + err.Error())
	}
	return string(jsonBytes)
}

func UpdateModelRatioByJSONString(jsonStr string) error {
	ModelRatio = make(map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &ModelRatio)
}

func ModelRatio2JSONString() string {
	jsonBytes, err := json.Marshal(ModelPrice)
	if err != nil {
		SysError("error marshalling model ratio: " + err.Error())
	}
	return string(jsonBytes)
}

func UpdateModelRatio2ByJSONString(jsonStr string) error {
	ModelPrice = make(map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &ModelPrice)
}

func GetModelRatio(name string) float64 {
	if strings.HasPrefix(name, "gpt-4-gizmo") {
		name = "gpt-4-gizmo-*"
	}
	ratio, ok := ModelRatio[name]
	if !ok {
		SysError("model ratio not found: " + name)
		return 15
	}
	return ratio
}

func GetModelRatio2(name string) (float64, bool) {
	if strings.HasPrefix(name, "gpt-4-gizmo") {
		name = "gpt-4-gizmo-*"
	}
	ratio, ok := ModelPrice[name]
	if !ok {
		ratio, ok = ModelPrice["default"] // 尝试获取默认
	}
	return ratio, ok
}

func CompletionRatio2JSONString() string {
	jsonBytes, err := json.Marshal(CompletionRatio)
	if err != nil {
		SysError("error marshalling completion ratio: " + err.Error())
	}
	return string(jsonBytes)
}
func UpdateCompletionRatioByJSONString(jsonStr string) error {
	CompletionRatio = make(map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &CompletionRatio)
}

func GetCompletionRatio(name string) float64 {
	if ratio, ok := CompletionRatio[name]; ok {
		return ratio
	}
	if strings.HasPrefix(name, "gpt-3.5") {
		if strings.HasSuffix(name, "1106") {
			return 2
		}
		if strings.HasSuffix(name, "0125") {
			return 3
		}
		if name == "gpt-3.5-turbo" || name == "gpt-3.5-turbo-16k" {
			// TODO: clear this after 2023-12-11
			now := time.Now()
			// https://platform.openai.com/docs/models/continuous-model-upgrades
			// if after 2023-12-11, use 2
			if now.After(time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC)) {
				return 2
			}
		}
		return 1.333333
	}
	if strings.HasPrefix(name, "gpt-4") {
		if strings.HasSuffix(name, "preview") {
			return 3
		}
		return 2
	}
	if strings.HasPrefix(name, "claude-3") {
		return 5
	}
	if strings.HasPrefix(name, "claude-") {
		return 3
	}
	if strings.HasPrefix(name, "mistral-") {
		return 3
	}
	return 1
}
