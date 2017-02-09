# blob-store-explorer

## Overview

* Golang (why? because I wanted to learn it)
* used to query the blobstore directly
* useful as a development tool
** How many blobs are in this directory?
** how many of them are soft-deleted?
** whats the total size of soft-deleted blobs?
** what is the total size of non-deleted blobs?
* useful for troubleshooting
** blobs created by this person

## Goals

```
bse /path/to/blob/store --metrics
```

Produces

```json
{
  "total-blobs" : 1234,
  "soft-deleted" : 345,
  "total-size" : 135123,
  "total-size-deleted" : 23543
}
```
