package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
