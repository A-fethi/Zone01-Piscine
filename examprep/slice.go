package main

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}

	n := len(a)

	switch len(nbrs) {
	case 1:
		start := nbrs[0]

		if start < 0 {
			start += n
		}

		if start < 0 {
			start = 0
		}
		if start > n {
			start = n
		}

		return a[start:]

	case 2:
		start, end := nbrs[0], nbrs[1]

		if start < 0 {
			start += n
		}
		if end < 0 {
			end += n
		}

		if start < 0 {
			start = 0
		}
		if start > n {
			start = n
		}

		if end < 0 {
			end = 0
		}
		if end > n {
			end = n
		}

		if start > end {
			start, end = end, start
		}

		return a[start:end]

	default:
		return []string{}
	}
}

// func Slice(a []string, nbrs ...int) []string {
//     length := len(a)

//     if len(nbrs) == 0 {
//         return a
//     }

//     start := nbrs[0]
//     end := length

//     if len(nbrs) > 1 {
//         end = nbrs[1]
//     }

//     // Handle negative indices
//     if start < 0 {
//         start = length + start
//     }
//     if end < 0 {
//         end = length + end
//     }

//     // Ensure start and end are within bounds
//     if start < 0 {
//         start = 0
//     }
//     if end > length {
//         end = length
//     }

//     // If start is greater than end, return nil
//     if start > end {
//         return nil
//     }

//     return a[start:end]
// }
