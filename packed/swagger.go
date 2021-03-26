package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
