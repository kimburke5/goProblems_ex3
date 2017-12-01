
package main
 
 import "math/rand"
 import "time"
 import "fmt"
 import "regexp"
 import "strings"
	
func Reflect(input string) string {

	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	//reflected word list
	reflectedWords := [][]string{
		{`am`, `are`},
		{`was`, `were`},
		{`i`, `you`},
		{`i'd`, `you would`},
		{`i've`, `you have`},
		{`i'll`, `you will`},
		{`my`, `your`},
		{`are`, `am`},
		{`you've`, `I have`},
		{`you'll`, `I will`},
		{`your`, `my`},
		{`yours`, `mine`},
		{`you`, `me`},
		{`me`, `you`},
	}
   
	// Loop through each token and reflect it if there is a match
	for i, token := range tokens {
		for _, reflectedWords := range reflectedWords {
			if matched, _ := regexp.MatchString(reflectedWords[0], token); matched {
				tokens[i] = reflectedWords[1]
				break
			}
		}
	}
	return strings.Join(tokens, ``)
 }
 
 var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
 }
 
 func ElizaResponse(input string) string {

	 if matched,_ := regexp.MatchString(`(?i).*\bfather\b.*`,input);matched{
		 
		 return "Why don't you tell me more about your father?"
	 }

	 re := regexp.MustCompile(`(?i).*\bi am|i'm|im\b.*([^.?!]*)[.?!]?`)
	 
	 if matched := re.MatchString(input); matched {
		
		return re.ReplaceAllString(input, "How do you know you are$1")
	}

	return responses[rand.Intn(len(responses))]
 
 }
 
 func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println();

	fmt.Println("\nFather was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println();

	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println();

	fmt.Println("\nI'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
	fmt.Println();

	fmt.Println("\nMy grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println();

	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy.")) 
	fmt.Println();

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println();

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println();

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println();

 }