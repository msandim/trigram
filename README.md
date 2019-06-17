# Trigram

[![GoDoc](https://godoc.org/github.com/msandim/trigram?status.svg)](https://godoc.org/github.com/msandim/trigram)
[![GoCard](https://goreportcard.com/badge/github.com/msandim/trigram)](https://goreportcard.com/report/github.com/msandim/trigram)

This Golang written program generates random texts, based on texts it already learned. These generated texts maintain the "style" of the the texts used for learning.

## How does this work?

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
When generating a text, if the last 2 words we generated were `to` and `be`, we have in memory 2 possibilities for the third word: `or` or `that`. Note that we will have a 66% of choosing `or` as the next word, versus `that` which will a 33% chance, since `or` was 2 times more frequent than `that` in the `[to, be, _]` type of trigrams.

Current limitations:



Possible future improvements:
- Mention about memory usage in having map([2]string)int vs what was done
- Mention sharding that could be used for the trigrams
- Mention parallelism in the parsing
- For example, you could imagine a Slack bot that sends every message it sees to the POST /learn endpoint, while constantly generating text using the GET /generate endpoint. More "write intensive" app, so let's "prioritize" the channel that receives "save trigrams" requests

Algorithm:

the store runs on a different goroutine