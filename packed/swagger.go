package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
