// this combines string manipulation, looping,slices,function declarationand unit testing in go

package main

import (
"strings"
"testing"
)

//ReverseWords take a string as input and returns the string with each word reversed
func ReverseWords(input string) string {
words := strings.Fields(input)
reversed := make ([]string,
  len(words))
for i, word := range words {
reversedWord := ""
for _, char := range word {
reversedWord = string(char) +
 reversedWord
}
reversed[i] = reversedWord
}
return strings.Join(reversed, "")
}
//TestReverseWords tests the ReverseWords function
func TestReverseWords(t *testing.T) {
testCases := []struct {
  input          string
  expectedOutput string
  }{
{"Hello, World!","olleH, !dlroW"},
{"Go Programming", "oGgmimmarrgorp"},
{"I love Go!", "I evol !oG"},
}

for _, tc := range testCases {
output :=ReverseWords(tc.input)
if output != tc.expectedOutput {
t.Errorf("ReverseWords(%s) =%s, expected %s",tc.input,output,tc.expectedOutput)
}
}
}
