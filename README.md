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

Example properties:
```
#Thu Feb 09 14:23:44 MST 2017
@BlobStore.created-by=admin
size=32
@Bucket.repo-name=maven-releases
creationTime=1486675424145
@BlobStore.content-type=text/plain
@BlobStore.blob-name=com/sonatype/training/nxs301/03-implicit-staging/1.24.0/03-implicit-staging-1.24.0.pom.md5
sha1=713a264172d968a64d784e3dd6f52c1653051bdf
```
