# trigram

- Mention about memory usage in having map([2]string)int
- Mention sharding that could be used for the trigrams
- Mention parallelism in the parsing
- For example, you could imagine a Slack bot that sends every message it sees to the POST /learn endpoint, while constantly generating text using the GET /generate endpoint. More "write intensive" app, so let's "prioritize" the channel that receives "generate text" requests