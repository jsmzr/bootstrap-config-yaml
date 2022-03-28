# bootstrap/config-yaml

bootstrap 系列 yaml 配置适配器

## 使用说明

1. 通过 `import (_ "github.com/jsmzr/bootstrap-config-yaml/yaml)"` 导入适配器
2. 通过 `import "github.com/jsmzr/bootstrap-config/config"` 导入门面
3. 通过 `err := config.InitInstance("yaml", "resource/app.yaml")` 进行初始化
4. 代码中调用 `name, ok := config.Get("bootstrap.name")` 获取键值或 `err := config.Resolve("author", &author)` 解析配置到结构体

## 开发说明

bootstrap 系列是为了提供统一的门面和依赖管理而产生的，期望我们在项目中可以屏蔽底层的实现，可以通过切换适配器快速接入、切换到其他实现。
