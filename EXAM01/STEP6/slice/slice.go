package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
    start, end := 0, len(a)

    if len(nbrs) >= 1 {
        start = nbrs[0]
    }

    if len(nbrs) == 2 {
        end = nbrs[1]
    }

    start, end = adjustIndices(start, end, len(a))

    if isValidIndices(start, end, len(a)) {
        return a[start:end]
    }

    return nil
}

func adjustIndices(start, end, length int) (int, int) {
    if start < 0 {
        start = length + start
    }

    if end < 0 {
        end = length + end
    }

    return start, end
}

func isValidIndices(start, end, length int) bool {
    return start >= 0 && start <= end && end <= length
}

func main() {
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}
