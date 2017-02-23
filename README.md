# blob-store-explorer

## Contents

* [Overview](#overview)
* [Usage](#usage)
* [Package Management](#package-management)
* [Todo](#todo)

## Overview

The Blob Store explorer was written in Golang mostly because I wanted to learn it.

The tool is designed to query/explore the blob store directly. This can be used for development or support. Some things it should be able to do:

* find a count of soft-deleted blobs, verify the amount of storage that should be recovered after a compact.
* find blobs associated with a certain repository
* find blobs created by a specific user


## Usage

```
blob-store-explorer --c --filter deleted=true  ~/develop/sonatype/nexus-internal/target/sonatype-work/nexus3/blobs/default/
```

Example output:

```
Exploring /Users/ericdcobb/develop/sonatype/nexus-internal/target/sonatype-work/nexus3/blobs/default/Stats:
Total blobs: 6, Total size: 834, Soft Deleted: 6, Total Size Deleted 834
deleted = true
@BlobStore.created-by = admin
size = 40
@Bucket.repo-name = maven-releases
creationTime = 1486679665325
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml.sha1
@BlobStore.content-type = text/plain
sha1 = cbd5bce1c926e6b55b6b4037ce691b8f9e5dea0f

deleted = true
@BlobStore.created-by = admin
size = 40
@Bucket.repo-name = maven-releases
creationTime = 1486675424206
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml.sha1
@BlobStore.content-type = text/plain
sha1 = 35d39f8f5fade17cebd4474a07f3bdc28179bdac

deleted = true
@BlobStore.created-by = admin
size = 361
@Bucket.repo-name = maven-releases
creationTime = 1486679665310
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml
@BlobStore.content-type = application/xml
sha1 = 9a49697dae03eb74d05db06bc765fe050034fc60

deleted = true
@BlobStore.created-by = admin
size = 32
@Bucket.repo-name = maven-releases
creationTime = 1486675424219
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml.md5
@BlobStore.content-type = text/plain
sha1 = 4982969b96e1822a7afc32c500741c61c0a3d55a

deleted = true
@BlobStore.created-by = admin
size = 32
@Bucket.repo-name = maven-releases
creationTime = 1486679665338
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml.md5
@BlobStore.content-type = text/plain
sha1 = d18758e45b1557c218ebadd7455029c4cffe93fc

deleted = true
@BlobStore.created-by = admin
size = 329
@Bucket.repo-name = maven-releases
creationTime = 1486675424191
@BlobStore.blob-name = com/sonatype/training/nxs301/03-implicit-staging/maven-metadata.xml
@BlobStore.content-type = application/xml
sha1 = dbfa3838d7c51136f022e8c7698611daa23114d8
```
## Package Management

Dependencies are currently managed by [godep](https://github.com/tools/godep), which looks to be the [most popular](https://github.com/golang/go/wiki/PackageManagementTools) package management tool today. Interestingly, godep pulls dependencies into a 'vendor' folder, which is then managed in VC.

Check out [The Saga of Go Dependency Management](https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) for an interesting read. The upshot of all this is that the Go community recognizes the limitations inherent in "Just committing vendor" and has begun working on an official tool ([dep](https://github.com/golang/dep)) to reproduce builds, with committing vendor becoming optional.

## TODO:

- [ ] Tests!
- [x] some kind of package/ version management, you have to manually `go get` dependencies now
- [ ] better error handling
- [ ] output in various formats (json?)
