import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    a, b := nums1, nums2
    total := len(nums1) + len(nums2)
    half := (total + 1) / 2

    if len(b) < len(a) {
        a, b = b, a
    }

    low, hi := 0, len(a) - 1

    for  {
        pointerA := (low + hi) >> 1
        pointerB := half - pointerA - 2

        var ALeft, ARight float64
        var BLeft, BRight float64

        if pointerA >= 0 {
            ALeft = float64(a[pointerA])
        } else {
            ALeft = math.Inf(-1)
        }

        if (pointerA + 1) < len(a) {
            ARight = float64(a[pointerA+1])
        } else {
            ARight = math.Inf(1)
        }

        if pointerB >= 0 {
            BLeft = float64(b[pointerB])
        } else {
            BLeft = math.Inf(-1)
        }

        if (pointerB + 1) < len(b) {
            BRight = float64(b[pointerB+1])
        } else {
            BRight = math.Inf(1)
        }

        // Correct partition
        if ALeft <= BRight && BLeft <= ARight {
            if total % 2 == 1 {
                return max(ALeft, BLeft)
            } else {
                return (max(ALeft, BLeft) + min(ARight, BRight)) / 2
            }
        } else if ALeft > BRight {
            hi = pointerA - 1
        } else {
            low = pointerA + 1
        }
    }

    return 0
}

func max(a,b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func min(a,b float64) float64 {
    if a < b {
        return a
    }
    return b
}
