package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
