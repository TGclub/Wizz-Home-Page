package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
