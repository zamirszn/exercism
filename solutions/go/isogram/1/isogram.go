package isogram
import "unicode"


func IsIsogram(word string) bool {
	wordMap := make(map[rune]bool)
    for _, char := range word{
        if unicode.IsLetter(char){
		r := unicode.ToLower(char)
		if wordMap[r]{
			return false
		}
		wordMap[r] = true
		}
		
    }
return true
}
