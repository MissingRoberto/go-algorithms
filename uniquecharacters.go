package algorithms

// O(N) space,  O(N) in memory

func IsUnique(s string) bool {
	found := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if found[s[i]] {
			return false
		}
		found[s[i]] = true
	}
	return true
}
