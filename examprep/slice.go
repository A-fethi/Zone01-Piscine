package piscine

func Slice(a []string, nbrs ...int) []string {
    length := len(a)
    
    if len(nbrs) == 0 {
        return a
    }

    start := nbrs[0]
    end := length

    if len(nbrs) > 1 {
        end = nbrs[1]
    }

    // Handle negative indices
    if start < 0 {
        start = length + start
    }
    if end < 0 {
        end = length + end
    }

    // Ensure start and end are within bounds
    if start < 0 {
        start = 0
    }
    if end > length {
        end = length
    }

    // If start is greater than end, return nil
    if start > end {
        return nil
    }

    return a[start:end]
}