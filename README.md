# trigram

Package architecture:

Future improvements:
- Mention about memory usage in having map([2]string)int vs what was done
- Mention sharding that could be used for the trigrams
- Mention parallelism in the parsing
- For example, you could imagine a Slack bot that sends every message it sees to the POST /learn endpoint, while constantly generating text using the GET /generate endpoint. More "write intensive" app, so let's "prioritize" the channel that receives "save trigrams" requests

Algorithm:

the store runs on a different goroutine