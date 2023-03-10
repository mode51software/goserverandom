# goserverandom

Return random bytes to an HTTP request:

```
GET /genrandomfile/<number of bytes>
```

Chunk size set to 128 bytes between rand reads.

Max size defaults to 100GB. Should include a property that configures a max limit.


