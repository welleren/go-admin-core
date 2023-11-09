###

测试用例：
```
import (
	"fmt"
	"testing"
	
	"codeup.aliyun.com/64cb084d32e3d43cbddbe43d/go-admin-core/codeup.aliyun.com/64cb084d32e3d43cbddbe43d/go-admin-core/config"
	"codeup.aliyun.com/64cb084d32e3d43cbddbe43d/go-admin-core/codeup.aliyun.com/64cb084d32e3d43cbddbe43d/go-admin-core/config/source/file"
)

func TestApp(t *testing.T)  {
	c, err := config.NewConfig()
	if err != nil {
		t.Error(err)
	}
	err = c.Load(file.NewSource(file.WithPath("config/settings.yml")))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.Map())
}
```
