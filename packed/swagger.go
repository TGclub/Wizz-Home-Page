package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
