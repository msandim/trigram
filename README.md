# Trigram

[![GoDoc](https://godoc.org/github.com/msandim/trigram?status.svg)](https://godoc.org/github.com/msandim/trigram)
[![GoCard](https://goreportcard.com/badge/github.com/msandim/trigram)](https://goreportcard.com/report/github.com/msandim/trigram)

This Golang written program generates random texts, based on texts it already learned. These generated texts maintain the "style" of the the texts used for learning.

## How does this work? 🛠

This is achieved by structuring the learned texts in trigrams (groups of 3 words) and using its frequency to write the next word of the text.

Given the text `To be or, to be or, not to be, that is the question`, the following trigrams are generated and saved in memory:

```
[to, be, or]
[be, or, to]
[to, be, or]
[be, or, not]
[or, not, to]
[not, to, be]
[to, be, that]
[be, that, is]
[that, is, the]
[is, the, question]
```
When generating a text, if the last 2 words we generated were `to` and `be`, we have in memory 2 possibilities for the third word: `or` or `that`. Note that we will have a 66% probability of choosing `or` as the next word, versus `that` which will a 33% chance, since `or` was 2 times more frequent than `that` in the `[to, be, _]` type of trigrams.

## How are these trigram beauties saved?

Currently memory usage is fairly optimized.
The trigrams are saved in a structure of type `map[string]map[string]map[string]int`, which represents a 3-dimensional map of frequencies of trigrams.

`ourMap["to"]["be"]["or"]` being 2, means that this sequence of words was found 3 times during past learning processes.

`ourMap["to"]["be"]` would return a map of possible endings for the trigram and their frequency. In the case of the example above, it would return `{"or": 2, "that": 1}`

## Awesome 😍, but what can it not do?

It can't serve coffee yet, unfortunately.

Also:
- It can only generate texts up to 100 words.
- It's prepared to ignore the following characters in texts: `.,;!?`. If the text has portions like `stuff - me`, `-` will be considered a word. Sorry.

## But how could you improve this natural beauty? 🌳 

- Currently memory usage is fairly optimized.
The trigrams are saved in a structure of type `map[string]map[string]map[string]int`, which represents a 3-dimensional map of frequencies of trigrams. `ourMap["a"]["b"]["c"]` being 3, means that this sequence of words was found 3 times during past learning processes.

Possible future improvements:
- Mention about memory usage in having map([2]string)int vs what was done
- Mention sharding that could be used for the trigrams
- Mention parallelism in the parsing
- For example, you could imagine a Slack bot that sends every message it sees to the POST /learn endpoint, while constantly generating text using the GET /generate endpoint. More "write intensive" app, so let's "prioritize" the channel that receives "save trigrams" requests

Algorithm:

the store runs on a different goroutine