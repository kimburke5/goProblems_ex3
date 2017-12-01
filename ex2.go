package main
 
 import "math/rand"
 import "time"
 import "fmt"
 import "regexp"

 var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
 }
 
 func ElizaResponse(input string) string {

	 if matched,_ := regexp.MatchString(`(?i).*\bfather\b.*`,input);matched{

		 return "why dont you tell me more about your father?"

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

 }