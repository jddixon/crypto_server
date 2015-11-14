# crypto_server

This is an implementation of the [XLattice](http://www.xlattice.org)
CryptoServer in the Go programming language.

A CryptoServer servers static web pages from a
[BuildList](https://jddixon.github.io/buildList),
a cryptographically secure description of its content.  The BuildList
lists all of the pages in a website by name and then by the SHA hash of
the page.  The BuildList contains an indented list of the files holding
the content of each web page and the content key (SHA1 or SHA256 hash)
of that content.  The web pages are stored in a hierarchical file
system organized by content key instead of name.  In the current
implementation, this means that if the content hash of a page were
`0x012345...def`, then the page would be stored at `U/01/23/45...def`.
If the relative path to the page were `path/to/fileX` then the indented
list would contain
	path
	  ...
	  to
	    ...
	    fileX 012345...def

Given this file name, the system would first use the BuildList to
determine its content hash, and then use that content hash to obtain
the content (the web page itself) from the file system.  The CryptoServer
would have verified the correctness of the content key by hashing the
content during the file retrieval process.

CryptoServer web servers are generally organized in clusters (see
[xlCluster_go](https://jddixon.github.io/xlCluster_go),
with page content distributed across the clusters.  If the server
attempts to retrieve a web page (by content key) and the page is corrupt
or not present on the server, then the server will automatically retrieve
it from another member of the cluster, reading it through its own disk
and in-memory caches, so that it will be available locally on future
requests.

At any time, the system administrator can change the website by first
adding the content of any altered web pages to one or more cluster
members and then by replacing
the build list.  This last action (replacing the build list)  will usually
invalidate a small fraction of the pages displayed.   Generally the first
fetch of an affected page will required that the content for the page be
retrieved from a peer; from that point there will almost always be a
local copy of the page and retrieval will be very fast.

## Project Status

Pre-alpha.  Some code exists but nothing immediately useful to the
casual user.

## On-line Documentation

More information on the **crypto_server** project can be found
[here](https://jddixon.github.io/crypto_server)
