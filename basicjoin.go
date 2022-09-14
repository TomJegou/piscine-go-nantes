package piscine

func BasicJoin(elems []string) string {
	intermediate_slice := []byte{}
	for index_word := 0; index_word < len(elems); index_word++ {
		for i := 0; i < len(elems[index_word]); i++ {
			intermediate_slice = append(intermediate_slice, elems[index_word][i])
		}
	}
	return string(intermediate_slice)
}
