<h1 class="libTop">cryptoserver_go</h1>

This is an implementation of the
[XLattice](https://jddixon/github.io/xlattice)
CryptoServer in the
[Go](https://golang.org)
programming language, Google's language for
system programming.

A CryptoServer serves static web pages from a
[BuildList](https://jddixon.github.io/buildList),
a cryptographically secure description of its content.  The BuildList
lists all of the pages in a website by name and then by the
[SHA](https://en/wikipedia.org/wiki/Secure_Hash_Algorithm)
hash of
the page.  The BuildList contains an indented list of the files holding
the content of each web page and the content key (SHA1 or SHA256 hash)
of that content.  The web pages are stored in a hierarchical file
system organized by content key instead of name.  In the current
implementation, this means that

* if the content hash of a page were `0x012345...def`,
* and the content key store is the directory `U`
* then the page would be stored at `U/01/23/012345...def`.

If the relative path to the page were `path/to/fileX` then the indented
list in the BuildList would contain

	path
	  ...
	  to
	    ...
	    fileX 012345...def

Given this path to the file, the system would first use the BuildList to
determine its content hash, and then use that content hash to obtain
the content (the web page itself) from the file system.  The CryptoServer
would have verified the correctness of the content key by hashing the
content during the file retrieval process.

CryptoServer web servers are generally organized in clusters (see
[xlCluster_go](https://jddixon.github.io/xlCluster_go),
with page content distributed across the clusters.  If the server
attempts to retrieve a web page (by content key) and the page is corrupt
or not present on the server, then the server will automatically retrieve
the file from another member of the cluster, reading it through its own disk
and in-memory caches, so that it will be available locally on future
requests.

At any time, the system administrator can change the website by first
adding the content of any altered web pages to one or more cluster
members and then by replacing
the BuildList.  This last action (replacing the BuildList)  will typically
invalidate a small fraction of the pages displayed.   Generally the first
fetch of an affected page will require that the content for the page be
retrieved from a peer; from that point there will almost always be a
local copy of the page and retrieval will be very fast.

## Motivation

The CryptoServer protects against the all-too-common style of web site
attack where a malicious hacker defaces the web site by replacing or
alterial legitimate content with the intention of confusing, misleading,
or deceiving the readaer.  In order to change what the CryptoServer
displays, it is necessary to alter the BuildList.  Doing this requires
access to the private key used to sign the BuildList -- you have to first
change the BuildList (adding your unauthorized content -- and then sign
the BuildList using the private key.  As this key need not (and should not)
be stored on the web server, this forces the hacker to comprovise not only
the web server but also the machine used by the designer, where the key
will have been used in preparing the BuildList.  The web server is necessarily
on an open Internet and so can be easily attacked.  The designer's machine
will normally be on a private subnet which is not publicly accessible.

To complete the attack, the hacker must

* gain access to the web designer's private network,
* locate the secret key,
* use that private key to generate a compromised BuildList
* introduce altered content into the system which distributes pages to the web servers
* and distribute the compromised BuildList to the web servers

This is a formidable list of tasks.

## Project Status

Pre-alpha.  Some Go code exists but nothing immediately useful to the
casual user.

There is a Java implementation of CryptoServer which
has been stable since 2012 or so.

