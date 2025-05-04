package main

import (
    "fmt"
    "math/rand"
)

func rollDice(id int, ch chan int) {
    result := rand.Intn(6) + 1 // 1〜6の乱数を生成
    fmt.Printf("🎲 Dice %d rolled: %d\n", id, result)
    ch <- result
}

func main() {
    numDice := 5
    results := make(chan int, numDice)

    // 複数のサイコロを並行に振る
    for i := 1; i <= numDice; i++ {
        go rollDice(i, results)
    }

    // 結果を受信して合計を計算
    total := 0
    for i := 0; i < numDice; i++ {
        total += <-results
    }

    fmt.Printf("🎯 Total: %d\n", total)
}
