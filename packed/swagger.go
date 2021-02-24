package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
