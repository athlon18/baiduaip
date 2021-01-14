package ocr

import (
	"net/url"
	"strings"
)

type generalOcr struct {
	Image        string `json:"image"`         //base64
	LanguageType string `json:"language_type"` //语言
}

type responseOcr struct {
	WordsResult    []WordsResult `json:"words_result"`
	LogID          int64         `json:"log_id"`
	WordsResultNum int           `json:"words_result_num"`
}
type WordsResult struct {
	Words string `json:"words"`
}

func generalBasic(image string, languageType string) (res responseOcr, err error) {
	values := url.Values{}
	values.Set("image", strings.Replace(image, "data:image/jpeg;base64,", "", -1))
	if languageType != "" {
		values.Set("language_type", languageType)
	}
	err = post(generalBasicUrl, values, &res)
	for index, item := range res.WordsResult {
		res.WordsResult[index].Words = strings.Replace(item.Words, " ", "", -1)
	}
	return res, err
}
