package main

func isMatch(s string, p string) bool {
	tokens := []byte{}
	for i := 0; i < len(p); i++ {
		if i >= 1 && p[i-1] == '*' && tokens[len(tokens)-1] == '*' {
			continue
		}
		tokens = append(tokens, p[i])
	}
	currTok := 0
	prevTok := 0
	for i := 0; i < len(s); {
		if tokens[currTok] == '*' {
			prevTok = currTok
			currTok += 1
		}
		if 
	}
	return true
}

func main() {

}
