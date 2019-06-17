# Trigram

[![GoDoc](https://godoc.org/github.com/msandim/trigram?status.svg)](https://godoc.org/github.com/msandim/trigram)
[![GoCard](https://goreportcard.com/badge/github.com/msandim/trigram)](https://goreportcard.com/report/github.com/msandim/trigram)

This Golang written program generates random texts, based on texts it already learned. These generated texts maintain the "style" of the the texts used for learning.

## How can use this? 👩‍💻

Getting the program: 
- just clone this repository somewhere and `cd` into its main directory.

Running the program:
```
make run
```

Running the tests:
```
make test
```

Running the tests and looking at their coverage:
```
make cover
```

When running, the program will expose two endpoints:

`POST localhost:8080/learn` - feeds texts into the program.

`GET localhost:8080/generate` - generates texts based on the ones fed previously.

## How does this work? 🛠

This is achieved by structuring the learned texts in trigrams (groups of 3 words) and using its frequency to figure out the next word of the text.

Given the text `To be or, to be or, not to be, that is the question`, the following trigrams are generated and saved in memory:

```
[to, be, or]
[be, or, to]
[or, to, be]
[to, be, or]
[be, or, not]
[or, not, to]
[not, to, be]
[to, be, that]
[be, that, is]
[that, is, the]
[is, the, question]
```
When generating a text, if the last 2 words we generated were `to` and `be`, we have in memory 2 possibilities for the third word: `or` or `that`.
Note that we will have a 66% probability of choosing `or` as the next word, versus `that` which will a 33% chance, since `or` was 2 times more frequent than `that` in the `[to, be, _]` type of trigrams.

## How are these trigram beauties saved? 🧠

Currently memory usage is fairly optimized.
The trigrams are saved in a structure of type `map[string]map[string]map[string]int`, which represents a 3-dimensional map of frequencies of trigrams.

`ourMap["to"]["be"]["or"]` being 2, means that this sequence of words was found 2 times during past learning processes.

`ourMap["to"]["be"]` would return a map of possible endings for the trigram and their frequency. In the case of the example above, it would return `{"or": 2, "that": 1}`

## Awesome 😍, but what can it not do?

It can't serve coffee yet, unfortunately.

Also:
- It can only generate texts up to 100 words.
- It's only prepared to ignore the following characters in texts: `.,;!?`. If the text has portions like `stuff - me`, `-` will be considered a word. Sorry.

## But how could you improve this natural beauty? 🌳 

- Sharding when generating the texts: trigrams could be store in a "distributed" data store, each node responsible for a certain range of possible trigrams.
This would make the program more efficient both at learning (because we weren't be writing always on the same memory space/node, as different trigrams would be allocated to different nodes) and generating (because we wouldn't be needing to request access to the same memory node).
However it's important to say that this would increase the complexity of both operations slightly.

- Give priority somehow to `/learn` requests, as opposed to `/generate` requests.
Assuming this program would be used by a Slack bot, it makes sense the `/learn` endpoint will be a lot more used than the `/generate` one. It would be interesting if these operations could have a different priority in accessing the datastore (`/generate` would have a higher priority in this case).
