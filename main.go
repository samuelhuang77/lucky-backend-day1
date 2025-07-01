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
    rand.Seed(time.Now().UnixNano())
    secret := rand.Intn(100) + 1
    fmt.Println("歡迎來到猜數字遊戲！請猜一個 1 到 100 的數字。")

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("輸入你的猜測：")
        if !scanner.Scan() {
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
}
