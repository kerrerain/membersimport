package common

import (
	"time"
)

const SUPERCOOP_MAIL_SUFFIX string = "@supercoop.fr"
const EXPORT_FOLDER string = "exports/"

func ExportFileName(name string) string {
	return EXPORT_FOLDER + name + timestamp() + ".csv"
}

func timestamp() string {
	return time.Now().Format("20060102150405")
}
