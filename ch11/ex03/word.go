// Package word は、言葉遊び用のユーティリティを提供します。
package word

// IsPalindrome は s が前からでも後ろからでも同じように読めるかどうかを報告します。
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
