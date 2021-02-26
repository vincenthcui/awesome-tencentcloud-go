package tmt

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextTranslateBatch = tencentcloud.Action{
		Service: "tmt",
		Version: "2018-03-21",
		Action:  "TextTranslateBatch",
	}
	TextTranslate = tencentcloud.Action{
		Service: "tmt",
		Version: "2018-03-21",
		Action:  "TextTranslate",
	}
	SpeechTranslate = tencentcloud.Action{
		Service: "tmt",
		Version: "2018-03-21",
		Action:  "SpeechTranslate",
	}
	LanguageDetect = tencentcloud.Action{
		Service: "tmt",
		Version: "2018-03-21",
		Action:  "LanguageDetect",
	}
	ImageTranslate = tencentcloud.Action{
		Service: "tmt",
		Version: "2018-03-21",
		Action:  "ImageTranslate",
	}
)
