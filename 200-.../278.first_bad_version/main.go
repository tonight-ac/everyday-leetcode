package main


var isBadVersion func (version int) bool

// 0 0 0 0 1 1
func firstBadVersion(n int) int {
	l, r := 1, n
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if isBadVersion(mid) {
			for isBadVersion(mid) {
				mid--
			}
			break
		} else {
			l = mid + 1
		}
	}

	return mid + 1
}

