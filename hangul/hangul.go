package hangul

var (
    start = rune(44032) // UTF-8  "가"
    end   = rune(55204) // UTF-8  "힝"
)

func HasConsonantSuffix(s string) bool {
    numEnd := 28
    result := false
    for _, r := range s {
        if start <= r && r < end {
            index := int(r-start)
            result = index%numEnd !=0
        }

    }
    return result

}
