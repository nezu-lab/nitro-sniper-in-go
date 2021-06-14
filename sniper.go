package main
import (
"time"
"math/rand"
"os"
"net/http"
"github.com/fatih/color"
)

func check(code string) {
	file, err := os.Create("nitro.txt")
	if err != nil {return}
	defer file.Close()
	Red, Green := color.New(color.FgRed), color.New(color.FgGreen)
	r, err := http.Get("https://discord.com/api/v6/entitlements/gift-codes/"+code+"?with_application=false&with_subscription_plan=true")
	if err != nil {return}
	if r.StatusCode == 20 {
		Green.Println("[VALID] -> https://discord.gift/"+code)
		file.WriteString("https://discord.gift/"+code+"\n")

	} else {
		Red.Println("[INVALID] -> https://discord.gift/"+code)
	}
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
func nitro(n int) string {
        b := make([]rune, n)
        for i := range b {
                b[i] = chars[rand.Intn(len(chars))]
        }
        return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for true {
		check(nitro(16))
	}
}
