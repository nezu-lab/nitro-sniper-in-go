package main
import (
"fmt"
"time"
"math/rand"
"os"
"strconv"
"github.com/fatih/color"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
func nitro(n int) string {
        b := make([]rune, n)
        for i := range b {
                b[i] = chars[rand.Intn(len(chars))]
        }
        return string(b)
}

func main() {
        var n int
        file, err := os.Create("nitro.txt")
        if err != nil {return}
        defer file.Close()
        c := color.New(color.FgCyan)
        c2 := color.New(color.FgMagenta)
        rand.Seed(time.Now().UnixNano())
        c.Print("How many nitro codes >>> ")
        fmt.Scanln(&n)
        for i := 1;i <= n;i++ {
                c2.Print("["+strconv.Itoa(i)+"] --> ")
                code := "https://discord.gg/"+nitro(16)
                color.Yellow(code)
                file.WriteString(code+"\n")
        }
        color.Green("all the codes were written in 'nitro.txt'")
}
