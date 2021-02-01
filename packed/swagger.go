package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
