package word

import (
	"math/rand"
	"testing"
	"time"

	word "github.com/mikan/go-training-course/ch11/ex03"
)

// randomPalindrome は、擬似乱数生成器 rng から長さと内容が計算された回文を返します。
func randomPalindrome(rng *rand.Rand) string {
	src := []rune("abcdefghijklmnopqrstuvwxyz,. ")
	n := rng.Intn(25) // 24 までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// r := rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン
		r := src[rng.Intn(len(src))]
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// 疑似乱数生成器を初期化する。
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
