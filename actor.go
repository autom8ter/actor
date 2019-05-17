//go:generate godocdown -o README.md

package actor

/*
1. Add Action
addAction (string) returning error
This function accepts a json serialized string of the form below and maintains an average time
for each action. 3 sample inputs:
1) {"action":"jump", "time":100}
2) {"action":"run", "time":75}
3) {"action":"jump", "time":200}
Assume that an end user will be making concurrent calls into this function.
2. Statistics
getStats () returning string
Write a second function that accepts no input and returns a serialized json array of the average
time for each action that has been provided to the addAction function. Output after the 3
sample calls above would be:
[
{"action":"jump", "avg":150},
{"action":"run", "avg":75}
]
Assume that an end user will be making concurrent calls into all functions.
*/

type Actor interface {
	addAction(string) error
	getStats() string
}

type ActionStore struct {
	Actions map[string]error
}
