package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
