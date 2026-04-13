package bob
import "strings" 
import "unicode"

func Hey(remark string) string {
    if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	}
    if hasNoLetter(remark) && isQuestion(remark){
        return "Sure."
    }
	 if isOnlyPunctuations(remark) {
		return "Sure."
	}
    
	if isTextUpper(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}
     if hasNoLetter(remark) == true {
		return "Whatever."
	}
	if isTextUpper(remark) {
		return "Whoa, chill out!"
	}
	if isQuestion(remark) {  
		return "Sure."
	}
    trimmed := strings.TrimSpace(remark)

    if strings.HasSuffix(trimmed, "?") {
		return "Sure." 
	}
	return "Whatever." 
}

func isQuestion(s string) bool {
	if strings.HasSuffix(s, "?") {
		return true
	}
	return false
}

func isTextUpper(s string) bool {
	for _, char := range s {
		if !unicode.IsUpper(char) && unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func isOnlyPunctuations(s string) bool {
	for _, char := range s {
		if !unicode.IsPunct(char) && char != ' ' {
			return false
		}
	}
	return true
}


func hasNoLetter(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
