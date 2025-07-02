# lucky-backend-day1

這是 Day 1 的 Go 語言學習模板，內容為簡易猜數字遊戲 CLI 範例。

## 環境需求

- Go 1.20 或以上版本

## 使用說明

1. 解壓縮本專案。
2. 在終端機執行：

```
go run main.go
```

3. 根據提示輸入 1 到 100 之間的數字，猜出正確數字。

## 後續建議

請依照速成課表，進行 Day 1 的學習與練習。


## 目前修改了 main.go 以下項目

1. 由於 rand.Seed() 已被拾棄, 調整為使用 r = rand.New() 的方法, 原因如下
```
在 Go 1.20 之後，rand.Seed() 這個方法已經被標記為 deprecated（已棄用）。這表示官方建議不要再使用這個方法，因為未來的 Go 版本可能會移除它，或是有更好的替代方案。

原因是：

math/rand 套件建議你建立自己的亂數產生器（rand.New(rand.NewSource(...))），而不是直接用全域的 rand.Seed() 和 rand.Intn()。
這樣可以避免多執行緒（goroutine）同時存取全域亂數產生器時出現問題，也讓程式更容易測試和維護。
簡單來說，deprecated 代表這個方法不建議再用，請改用新的寫法。例如：

這樣就不會用到已棄用的 rand.Seed() 了。
```

2. 測試過程看到了 bufio 套件的用法, 它是有緩衝區的I/O
```
套件是 Go 標準函式庫中的一個套件，主要用途是提供帶緩衝區的 I/O 操作。
它可以讓你更有效率地讀取或寫入資料，特別是處理大量或逐行的輸入輸出時。

常見用途包括：

逐行讀取檔案或標準輸入（如 bufio.NewScanner(os.Stdin)）
緩衝寫入檔案或網路連線
提供 ReadLine、ReadString、WriteString 等方便的方法
在你的程式中，bufio.NewScanner(os.Stdin) 用來逐行讀取使用者輸入。

scanner.Scan()

的用途是讀取下一行輸入，並判斷是否成功。

在你的程式中，scanner := bufio.NewScanner(os.Stdin) 建立了一個從標準輸入讀取的 Scanner。
每次呼叫 scanner.Scan()，它會嘗試讀取使用者輸入的一行文字：

如果成功讀到一行，回傳 true，你可以用 scanner.Text() 取得這一行內容。
如果遇到輸入結束（如 Ctrl+D/Ctrl+Z）或發生錯誤，回傳 false，for 迴圈就會結束。
簡單來說，scanner.Scan() 讓你可以逐行讀取輸入，常用於命令列互動或檔案讀取。
```

3. 接著因為看到了 os.Stdin 所以順便瞭解了 os.Stdout 及 fmt.Fprintln() 的用法, 也試著模擬 fmt.Fprintln() 第一個參數需要實作 Write interface{} 的部分

```
os.Stdin

是 Go 語言標準庫 os 套件中的一個變數，代表標準輸入（standard input）。

它的型別是 *os.File，通常對應於鍵盤輸入（在命令列執行時）。
你可以用它來讀取使用者輸入的資料，例如配合 bufio.NewScanner(os.Stdin) 來逐行讀取輸入。

簡單來說，os.Stdin 就是程式從外部（通常是使用者）接收資料的入口。

有的，Go 語言中也有 os.Stdout。

os.Stdout 是 os 套件提供的變數，代表標準輸出（standard output），型別同樣是 *os.File。
通常對應於螢幕輸出（命令列視窗），你可以用它來寫資料到螢幕，例如：

fmt.Fprintln(os.Stdout, "這是標準輸出")

這樣會直接把字串輸出到標準輸出裝置（通常是螢幕）。
同理，也有 os.Stderr 代表標準錯誤輸出。

fmt.Fprintln

的用法是：將資料格式化後寫入到指定的輸出目標。

語法如下：

fmt.Fprintln(輸出目標, 資料1, 資料2, ...)

第一個參數是 io.Writer 介面的物件，例如 os.Stdout、檔案、網路連線等。
後面的參數是你要輸出的內容，會自動加上換行字元。
os.Stdout 之所以可以傳入，是因為它實作了 io.Writer 介面，所以 fmt.Fprintln 可以直接把資料寫到標準輸出（螢幕）。

範例：

fmt.Fprintln(os.Stdout, "Hello, world!")

這會把 Hello, world! 輸出到螢幕。

os.Stdout

是一個 *os.File 型別的變數，而 *os.File 有實作 Go 的 io.Writer 介面。

io.Writer 介面定義

type Writer interface {
    Write(p []byte) (n int, err error)
}

只要型別有 Write([]byte) (int, error) 這個方法，就實作了 io.Writer。

*os.File 的實作
*os.File（包含 os.Stdout）有定義 Write 方法，能把資料寫到檔案、終端機或其他裝置。例如：

func (f *File) Write(b []byte) (n int, err error)

實際應用
因為 os.Stdout 實作了 io.Writer，所以你可以把它傳給像 fmt.Fprintln、io.Copy 這類需要 io.Writer 的函式，讓資料直接寫到標準輸出（螢幕）。

範例

fmt.Fprintln(os.Stdout, "Hello") // "Hello" 會被寫到螢幕

總結：
os.Stdout 之所以能當作 io.Writer 使用，是因為它的型別 *os.File 有實作 Write 方法，符合 io.Writer 介面規範。

os.Stdout

的型別是 *os.File，而 *os.File 有實作 io.Writer 介面。
io.Writer 介面只需要實作一個方法：

type Writer interface {
    Write(p []byte) (n int, err error)
}

*os.File 的 Write 方法大致內容如下（簡化說明）：

func (f *File) Write(b []byte) (n int, err error) {
    // 將 b 寫入檔案（或標準輸出、標準錯誤等）
    // 實際會呼叫底層作業系統的寫入功能
}

當你呼叫 fmt.Fprintln(os.Stdout, "hello") 時，Fprintln 會把資料轉成 bytes，然後呼叫 os.Stdout.Write()，資料就會被寫到螢幕。

總結：
os.Stdout 能當 io.Writer 用，是因為 *os.File 有實作 Write([]byte) (int, error) 方法，這個方法會把資料寫到標準輸出裝置。

```

```
func main() {
    ... (略)

    // 測試 Write 介面用 fmt.Fprintln
    test_obj := test_struct{Name: "測試物件"}
    fmt.Fprintln(&test_obj, "這是測試物件的輸出")
    fmt.Println("測試物件的 Name 欄位內容：", test_obj.Name)
}

type test_struct struct {
    // 這是一個測試用的結構體
    Name string
}

func (t *test_struct) Write(p []byte) (n int, err error) {
    // 實際的寫入邏輯會在這裡實現
    // 例如寫入到檔案或其他輸出目標
    t.Name = string(p) // 假設將寫入的資料存到 Name 欄位
    return len(p), nil // 假設寫入成功，返回寫入的字
}
```
```
O:\uk\tutorial\lucky-backend-day1>go run main.go
歡迎來到猜數字遊戲！請猜一個 1 到 100 的數字。
輸入你的猜測：^Z
這是標準輸出
測試物件的 Name 欄位內容： 這是測試物件的輸出
```