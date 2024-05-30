package libc

// delslice removes the elements s[i:j] from s, returning the modified slice.
// delslice panics if j > len(s) or s[i:j] is not a valid slice of s.
// delslice is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// delslice zeroes the elements s[len(s)-(j-i):len(s)].
func delslice[S ~[]E, E any](s S, i, j int) S {
	_ = s[i:j:len(s)] // bounds check

	if i == j {
		return s
	}

	// oldlen := len(s)
	s = append(s[:i], s[j:]...)
	// clear(s[len(s):oldlen]) // zero/nil out the obsolete elements, for GC
	return s
}
