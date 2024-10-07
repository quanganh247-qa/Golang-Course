package main

import "fmt"

// The addEmailsToQueue function tries to send all the emails to the channel emailsToSend synchronously, 
// but the channel is unbuffered. This means that it blocks on the first emailsToSend <- email because there's no goroutine actively receiving from the channel.
// Only after addEmailsToQueue finishes and returns the channel to the caller, the sendEmails function begins to receive from it.
// But by that point, the sending has already stopped, resulting in a deadlock.

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string)
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}
