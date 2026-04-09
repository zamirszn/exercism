package acronym
import "strings"


func Abbreviate(s string) string {
    if len(s) < 1{
        return s
    }
 	removeDash := strings.ReplaceAll(s, "-", " ")
    cleanedText := strings.Fields(strings.ReplaceAll(removeDash, "_", " "))
    
	abbr := []string{}
	for _, word := range cleanedText {
		firstChar := strings.ToUpper(string(word[0]))
		abbr = append(abbr, firstChar)
	}
	return strings.Join(abbr,"")
}


    
