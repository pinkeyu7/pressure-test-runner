# Pressure Test Runner

由於需要對 API 進行壓力測試，開發共用模板方便紀錄，並將測試結果儲存起來，並自動產生報告。

### 模擬兩種情境

1. 併發模式，使用 goroutine 將請求一次性發出，模擬併發情境
1. 序列模式，使用 for-loop 模擬 queuing 機制

### Environment

- `API_URL` 設定 API 位置
- `OUTPUT_FILE_PATH` 設定報告產生資料夾
- `TEST_TIMES` 設定執行次數

### 更改 API 請求及驗證方式

1. 調整 `api_request` 內的 function，以符合待測 API 的需求。
1. 在 `worker/dto.go` 中調整 `RequestDto` 內的 struct，並在 `main.go` 中新增測資。
1. 在 `worker/concurrence.go` 及 `worker/queue.go` 中將 `RequestDto` 內的參數交給 API。

### 執行 runner

```
go run main.go
```

執行後的 `各次測試結果` 以及 `結果綜合比較` 會在指定資料夾內。
