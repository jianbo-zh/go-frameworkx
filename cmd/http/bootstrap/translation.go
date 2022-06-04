package bootstrap

import (
	goi18n "github.com/jianbo-zh/go-i18n"
	"golang.org/x/text/language"

	_ "goframeworkx/cmd/http/translations"
)

func initTranslation() {
	goi18n.Setup(language.Chinese)
}
