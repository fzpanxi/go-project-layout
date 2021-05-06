# Go项目目录标准 - Project Layout 最佳实践
## 介绍
这是Go应用程序项目的基础布局。不是官方团队定义的标准；无论是在经典项目还是在新兴的项目中，这都是Go生态系统中一组常见的项目布局模式。这其中有一些模式比另外的一些更受欢迎。它通过几个支撑目录为任何足够大规模的实际应用程序提供一些增强功能。

## 项目布局
每个公司都应当为不同的微服务项目建立个统一的支持库（基础库/基础微服务框架）和Application项目。基础库为独立项目，通常一个公司只有一个，同时按照功能目录来拆分会带来不少的管理工作，因此建议合并整合。

## 基础库

将基础库项目作为公司的标准库，因此应该只有一个。并且基础库应该具备以下特点：

- 简单：不过度设计，代码平实简单；
- 通用：通用业务开发所需要的基础库的功能；
- 高效：提高业务迭代的效率；
- 稳定：基础库可测试性高，覆盖率高，有线上实践安全可靠；
- 健壮：通过良好的基础库设计，减少错用；
- 高性能：性能高，但不特定为了性能做 hack 优化，引入 unsafe ；
- 扩展性：良好的接口设计，来扩展实现，或者通过新增基础库目录来扩展功能；
- 容错性：为失败设计，大量引入对 SRE 的理解，鲁棒性高；
- 工具链：包含大量工具链，比如辅助代码生成，lint 工具等等；

## Application 应用项目

如果你尝试学习 Go，或者你正在为自己建立一个 PoC 或一个玩具项目，这个项目布局是没啥必要的。从一些非常简单的事情开始（一个 main.go 文件绰绰有余）。当有更多的人参与这个项目时，你将需要更多的结构，包括需要一个 Toolkit 来方便生成项目的模板，尽可能大家统一的工程目录布局。

![avatar](https://storage.liaoxinma.com/go_project_layout.png)

一个典型目录结构可能是这样：
```python
application
├─ api
| ├─ demo
| | |____v1
| | |____errors
|____cmd
| |____demo
|____configs
|____internal
| |____conf
| |____domain
| | |____entity
| | |____event
| | |____repository
| | |____service
| |____infra
| |____server
| |____usecase
| | |____event
| | |____service
|____pkg
|____go.mod
|____go.sum
|____LICENSE
|____README.md
|____CHANGELOG.md
|____OWNERS.md
```
## 应用类型

为服务中的 application 服务类型主要分为5个类型：interface、service、job、task、admin，应用 cmd 目录负责程序的启动、关闭、初始化配置等。

- interface：对外BFF层服务，接受来自客户端的请求，暴露 HTTP/gRPC接口（聚合服务层）。

- service：对内分为服务，仅接受BFF层或者聚合服务层的请求，比如暴露 gRPC 接口只对聚合服务层开放（基础服务层）。

- job：流式任务处理服务，一般依赖MQ。

- task：定时任务处理，一般依赖于分布式任务调度系统。

- admin：区别于 service，它是一个单独的业务服务集合，应用于运营管理部门；通常此类服务集合数据权限更高，与 service 形成隔离带来更好的代码级别安全。

## Go应用目录

### `/cmd`

本项目的主干。

每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，`/cmd/myapp`)。

不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 `/pkg` 目录中。如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 `/internal` 目录中。你会惊讶于别人会怎么做，所以要明确你的意图!

通常有一个小的 `main` 函数，从 `/internal` 和 `/pkg` 目录导入和调用代码，除此之外没有别的东西。

### `/internal`

私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。有关更多细节，请参阅Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) 。注意，你并不局限于顶级 `internal` 目录。在项目树的任何级别上都可以有多个内部目录。

你可以选择向 internal 包中添加一些额外的结构，以分隔共享和非共享的内部代码。这不是必需的(特别是对于较小的项目)，但是最好有有可视化的线索来显示预期的包的用途。你的实际应用程序代码可以放在 `/internal/app` 目录下(例如 `/internal/app/myapp`)，这些应用程序共享的代码可以放在 `/internal/pkg` 目录下(例如 `/internal/pkg/myprivlib`)。

### `/pkg`

外部应用程序可以使用的库代码(例如 `/pkg/mypubliclib`)。其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思:-)注意，`internal` 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。`/pkg` 目录仍然是一种很好的方式，可以显式地表示该目录中的代码对于其他人来说是安全使用的好方法。由 Travis Jeffery  撰写的 [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 博客文章提供了 `pkg` 和 `internal` 目录的一个很好的概述，以及什么时候使用它们是有意义的。

## 通用应用目录

### `/configs`

配置文件、配置模板、默认配置信息。

### `/test`

​额外的外部测试应用程序和测试数据。你可以随时根据需求构造 /test 目录。对于较大的项目，有一个数据子目录是有意义的。例如，你可以使用 /test/data 或 /test/testdata (如果你需要忽略目录中的内容)。请注意，Go 还会忽略以 “.” 或 “_” 开头的目录或文件，因此在如何命名测试数据目录方面有更大的灵活性。

## 服务应用程序目录

### `/api`

API 协议定义 、user interfaces 实现；这一层可以直接实现 http handler，或把 http handler 封装到基础库的 http server 中。

### `/conf`

配置加载的helper

### `/domain`

类DDD领域层，包含entity、repository、service。

- entity：实体。

- event：类 domain event，用于处理事件相关的业务逻辑。

- repository：类 DDD 的repo interface。

- service：类domain service，负责业务逻辑具体实现，表达业务概念、业务规则。

### `/infra`

业务数据访问，包含 cache、db、MQ等封装，同时实现了 domain层 的 repository 接口。

### `/server`

为HTTP或gRPC实例的创建和配置，以及注册对应的 service、handler。

### `/usecase`

类ddd application层。

- 实现api定义的rpc接口；

- 提供 「粗粒度」 业务入口。

- 处理DTO到Domain entity的转换（DTO -> DO）；

- 对下层 domain 层进行协调，对业务逻辑进行编排；

- 该层还处理消息事件的发布/订阅；

- 该层不应该包含复杂的业务规则。

## 额外说明

您的项目目录里一般还会放置 README.md、CHANGELOG.md、OWNERS.md。
