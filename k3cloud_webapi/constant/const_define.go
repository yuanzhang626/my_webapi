package constant

// InvokeMethod 定义调用方法的枚举类型
type InvokeMethod int

const (
	// SYNC 同步调用方法
	SYNC InvokeMethod = iota + 1
	// ASYNC 异步调用方法
	ASYNC
	// QUERY 查询调用方法
	QUERY
)

// QueryState 定义查询状态的枚举类型
type QueryState int

const (
	// Pending 查询状态：待处理
	Pending QueryState = iota
	// Running 查询状态：运行中
	Running
	// Complete 查询状态：已完成
	Complete
)

const (
	// BeginMethod_Header 查询模式：开始方法头
	BeginMethod_Header = "beginmethod"
	// BeginMethod_Method 查询模式：开始查询实现方法
	BeginMethod_Method = "BeginQueryImpl"
	// QueryMethod_Header 查询模式：查询方法头
	QueryMethod_Header = "querymethod"
	// QueryMethod_Method 查询模式：查询异步结果方法
	QueryMethod_Method = "QueryAsyncResult"
)
