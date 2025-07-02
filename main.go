package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secret := r.Intn(100) + 1
	fmt.Println("歡迎來到猜數字遊戲！請猜一個 1 到 100 的數字。")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("輸入你的猜測：")
		if !scanner.Scan() { // Ctrl+C 或 Ctrl+Z 會導致 Scan 返回 false
			break
		}
		input := strings.TrimSpace(scanner.Text())
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("請輸入有效的整數！")
			continue
		}

		if guess < secret {
			fmt.Println("太小了，再試試看！")
		} else if guess > secret {
			fmt.Println("太大了，再試試看！")
		} else {
			fmt.Println("恭喜你猜對了！遊戲結束。")
			break
		}
	}

	// 這邊測試一下 fmt.Fprintln 用標準輸出 os.Stdout
	fmt.Fprintln(os.Stdout, "這是標準輸出")

    // 測試 Writer 介面用 fmt.Fprintln
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