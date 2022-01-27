# Links R Us

## Prerequisites
### Run
* Go
* CockroachDB
    * brew install cockroachdb/tap/cockroach
    * cockroach start-single-node --insecure --http-port=26256 --host=localhost
    * cockroach sql --insecure -e 'CREATE DATABASE linkgraph;'
* Elasticsearch
    * java11
    * brew tap elastic/tap
    * brew install elastic/tap/elasticsearch-full
    * to start: elasticsearch
* Dependencies
    * start new module: go mod init github.com/mebobby2/linksrus
    * add new example: go get github.com/google/uuid
    * to install missing dependencies: go build or go test
* DB Migrations
    * brew install golang-migrate
    * migrate create -ext sql -dir linkgraph/store/cdb/migrations -seq create_edges_table
    * make run-cdb-migrations

### Test
* export CDB_DSN='postgresql://root@localhost:26257/linkgraph?sslmode=disable'
* make run-cdb-migrations
* export ES_NODES='http://localhost:9200'
* make test
## Quick Notes
* A facade is a software design pattern that abstracts the complexity of one or more software components behind a simple interface.
In the context of microservice-based design, the facade pattern allows us to transparently compose or aggregate data across multiple, specialized microservices while providing a simple API for the facade clients to access it.
* Most cloud providers run an internal metadata service that each computing node can query to obtain information about itself. This service is typically accessed via a link-local address such as 169.254.169.254 and nodes can perform simple HTTP GET requests to retrieve the information they are interested in.
* Introduced in Go 1.11, Go Modules aim to replace Go Workspaces and the GOPATH. It is recommended to start using modules

## Notes
### Software Enginering Roles
1. software engineer (SWE)
2. software development engineer in test (SDET)
3. site reliability engineer (SRE)
    * SREs spend approximately 50% of their time developing software and the other 50% dealing with ops-related aspects such as support tickets, responding to alerts, being on-call, upgrading systems, DR scenarios etc
    * It is in the best interests of SREs to increase the stability and reliability of the services they operate
    * The basic mantra of SREs is to eliminate potential sources of human errors by automating repeated tasks. One example of this philosophy is the use of a Continuous Deployment (CD) pipeline
4. release engineer (RE)
    * collaborates with all the engineering teams to define and document all the required steps and processes for building and releasing code to production
    * Goal: reproducible builds
    * Ensuring that builds are both repeatable and hermetic: changes to external dependencies (for example, third-party libraries) between builds of the same software version should have no effect on the artifacts that are produced by each build
5. System Architect
    * While software engineering teams focus on building the various components of the system, the architect is the one person who sees the big picture: what components comprise the system, how each component must be implemented, and how all the components fit and interact with each other.
    * In smaller companies, the role of the architect is usually fulfilled by one of the senior engineers. In larger companies, the architect is a distinct role that's filled by someone with both a solid technical background and strong analytical and communication skills.

### Software Development Models
1. Waterfall
2. Iterative enhancement
    * an attempt to improve on some of the caveats of the waterfall model
    * recognizing that requirements may potentially change for long-running projects, the model advocates executing a set of evolution cycles or iterations, with each one being allocated a fixed amount of time out of the project's time budget
    * the iterative model has exerted quite a bit of influence in the evolution of most of the contemporary software development models
3. Spiral
    * combines the ideas and concepts from the waterfall and iterative models with a risk assessment and analysis process
4. Agile
    * a sort of umbrella term that encompasses not only a set of frameworks but also a fairly long list of best practices for software development.
    * Lean
    * Scrum
    * Kanban
    * DevOps
        * basic premise behind the DevOps model is that each engineering team owns the services they build
        * DevOps advocates tend to gravitate toward two different models: CAMS & the three ways model
        * CAMS - Culture, Automation, Measurement, and Sharing
        * three ways - Systems thinking and workflow optimization, Amplifying feedback loops, and Culture of continuous experimentation and learning

### Engineering Principles
#### Dependency inversion principle (DIP)
* High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.
* The DIP essentially summarizes all the other principles we've discussed so far. If you have been applying the rest of the SOLID principles to your code base, you will find that it already adheres to the preceding definition!
* The introduction and use of interfaces aids in decoupling high-level and low-level modules. The open/closed principle ensures that interfaces themselves are immutable but does not preclude us from coming up with any number of alternative implementations (the details bit in the preceding definition) that satisfy an implicit or explicit interface. At the same time, the LSP guarantees that we can rely on the established abstractions while also having the flexibility to swap the underlying implementation at compile time or even runtime without worrying about breaking our applications.

### SOLID
* SRP - Group structs and functions based on their purpose and organize them into packages with clear logical boundaries.
* Open/Closed - Use composition and embedding of simple types to construct more complex types that still retain the same implicit interface as the types they consist of.
* LSP - Avoid unnecessary coupling by using interfaces rather than concrete types to define the contract between packages.
* ISP - Make sure your function or method signatures only depend on the behaviors they need and nothing more; use the smallest possible interface to describe function/method arguments and avoid coupling to the implementation details of concrete types.
* DIP - Use the appropriate level of abstraction when designing your code to decouple high-level and low-level modules while at the same time ensuring that the implementation details rely on the abstractions and not the other way round.

#### Applying SOLID principles
Even though we analyzed the SOLID principles through the eyes of a Go engineer, the principles themselves have a much wider scope and can also be applied to system design in general. For instance, in a microservice-based deployment, you should be aiming to build and deploy services with a single purpose (SRP) that communicate through clearly defined contracts and boundaries (ISP).

### Golang Conventions
* in contrast to other programming languages whose standard libraries usually come with utility libraries or packages with generic-sounding names such as common or util, Go is quite opinionated against this practice. This is actually justified from the SOLID principles' point of view as those packages are more likely to be violating the SRP versus aptly named packages whose name enforces a logical boundary for their contents. To add to this, as the number of published Go packages grows over time, searching for and locating packages with generic-sounding names will become more and more difficult.
* Accept interfaces, return structs.
    * states that we should always try to return concrete types rather than interfaces. This advice actually makes sense: as a package consumer, if I am calling a function that creates a type, Foo, I am probably interested in calling one or more methods that are specific to that type. If the NewFoo function returns an interface, the client code would have to manually cast it to Foo so that it can invoke the Foo-specific methods; this would defeat the purpose of returning an interface in the first place.

### Version Management
#### Single repository – multiple branches
A much better approach would be to still use a single repository but maintain a different branch (in Git terminology) for each major package version, extra feature, or development branches for ongoing work. If we were to apply this approach to the case of the weather package that we discussed before, our repository would normally contain the following branches:
* v1: This is the branch where the released 1.x.y line of the weather package is located.
* v2: Another branch for the 2.x.y release of the weather package.
* develop: Code in development branches is generally considered to be work in progress, and therefore unstable for use. Eventually, once the code stabilizes, it will be merged back into one or more of the stable release branches.

Similar to the versioned folder approach, the multibranch approach also ensures that the tip or head of each release branch contains the latest release version for a package; however, it is sometimes useful to be able to refer to an older semantic version of the package. A typical use case for this is repeatable builds, where we always want to compile against a specific version of the package and not the latest, albeit stable, version from a particular package line.

To satisfy the preceding requirement, we can exploit the VCS's capability to tag each release so we can easily locate it in the future without having to scan the commit history.

e.g.
| Name      | Type |
| ----------- | ----------- |
| v1.0.10      | Tag       |
| v1.1.9   | Tag        |
| v1   | Branch        |
| v3~dev   | Branch        |

To make a project compatible with the gopkg.in service, you need to make sure that either your branches or your tags match the expected patterns that gopkg.in looks for: vx, vx.y, vx.y.z, and so on.

#### Vendoring
The fact that services such as gopkg.in always redirect the go get tool to
the latest available major version for a given version selector is, technically speaking, a show-stopper for engineering teams that endeavor to set up a development pipeline that guarantees repeatable builds.

However, what if I told you that there is actually a way to retain the benefits of lazy package resolution and at the same time have the flexibility to pin down package versions for each build? The mechanism that will assist us in this matter is called vendoring.

In the context of Go programming, we refer to vendoring as
the process where immutable snapshots (also known as vendored dependencies) for all nodes in the import graph of a Go application get created. The vendored dependencies are used instead of the original imported packages whenever a Go application is compiled.

First and foremost, the key promise of vendoring is nothing other than the capability to run reproducible builds. Many customers, especially larger corporations, tend to stick to stable or LTS releases for the software they deploy and forego upgrading their systems unless it's absolutely necessary.

Being able to check out the exact software version that a customer uses and generate a bit-for-bit identical binary for use in a test environment is an invaluable tool for any field engineer attempting to diagnose and reproduce bugs that the customers are facing.

Another benefit of vendoring is that it serves as a safety net in case an upstream dependency suddenly disappears from the place where it is hosted (for example, a GitHub or GitLab repository), thereby breaking builds for software that depends on it. If you are thinking that this is a highly unlikely scenario, let me take you back to 2016 and share an interesting engineering horror story from the world of Node.js!

One common problem across engineering teams is that in spite of the fact that engineers are keen on vendoring their dependencies, they often forget to periodically refresh them. As I argued in a previous section, all code can contain potential security bugs. It is therefore likely that some security bugs (perhaps from a transitive dependency of an imported package) will eventually end up in production.

Security-related or not, when bugs are reported to the package maintainers, a fix is usually promptly released and the package version is incremented accordingly (that is, if a package is using semantic versioning). As large-scale projects tend to import a large volume of packages, it is not feasible to monitor each imported package's repository for security fixes. Even if this was possible, we couldn't realistically do this for their transitive dependencies. As a result, production code can remain unpatched for a long time even though the affected upstream packages have already been patched.

#### The dep tool
The Go team—being well aware that having several competing tools for managing dependencies could result in the fragmentation of the Go ecosystem and encumber the growth of the Go community—decided to assemble a committee and produce an official specification document detailing the way to move forward regarding Go package dependency management. The dep tool is the first tool that conforms to the published specification.

The output of the dep constraints solver is the highest possible supported version
across all dependencies. The dep tool creates two text-based files in the project's root folder that the user must commit to their VCS: Gopkg.toml and Gopkg.lock. To speed up CI builds, users may also optionally commit the populated vendor folder to version control. Alternatively, assuming that both Gopkg.toml and Gopkg.lock are available, a prebuild hook can populate the vendor folder on the fly by running dep ensure -vendor-only.

The Gopkg.toml file serves as a manifest for controlling the dep tool's behavior. The dep init invocation will analyze the import graph of the project and produce a Gopkg.toml file with an initial set of constraints. From that point on, whenever a constraint needs to be updated (usually to bump the minimum supported version), users need to manually modify the generated Gopkg.toml file.

By committing the Gopkg.lock file to the VCS, the dep support in Go 1.9+ guarantees that we can produce repeatable builds, provided, of course, that all referenced dependencies remain available.

#### Go Modules
One limitation of the dep tool is that it does not let us use multiple major versions of a package in our projects, as each path to an imported package must be unique.

To begin using Go modules, we first need to declare a new Go module by running go mod init parser in the folder where the preceding example is located. This will generate a file called go.mod.

If you list the contents of the project's folder, you will notice a new file called go.sum. This file stores the cryptographic hashes of the dependencies that have been downloaded and serves as a safeguard for ensuring that the contents of the packages have not been modified between builds (that is, a package maintainer force-pushed some changes, overwriting the previous version); a very useful feature when aiming for repeatable builds.

The go.mod and go.sum files serve the same purpose as
the Gopkg.toml and Gopkg.lock files used by the dep tool, and they also need to be committed to your version control system.

### Testing
#### Timing attacks exampe
* Start by providing answers following the $a pattern and measuring the time it takes to get a response. The $ symbol is a placeholder for all possible alphanumeric characters. In essence, we try combinations such as aa, ba, and so on.
* Once we have identified an operation that takes longer than the rest, we can assume that that particular value of $ (say, 4) is the expected first character of the CAPTCHA answer! The reason this takes longer is that the string comparison code matched the first character and then tried matching the next character instead of immediately returning it, like it would if there was a mismatch.
* Continue the same process of providing answers but this time using the 4$a pattern and keep extending the pattern until the expected CAPTCHA answer can be recovered.

#### Mocks, Stubs, Spies
* stubs are devoid of any logic; they just return a canned answer.
* A spy is nothing more than a stub that keeps a detailed log of all the methods that are invoked on it. For each method invocation, the spy records the arguments that were provided by the caller and makes them available for inspection by the test code.
* mocks as stubs on steroids! Contrary to the fixed behavior exhibited by stubs, mocks allow us to specify, in a declarative way, not only the list of calls that the mock is expected to receive but also their order and expected argument values. In addition, mocks allow us to specify different return values for each method invocation, depending on the argument tuple provided by the method caller.

#### Testing in Prod
* If you are working with a microservice architecture, you can engineer your services so that they do not talk to other services directly but rather to do so via a local proxy that is deployed in tandem with each service as a sidecar process. This pattern is known as
the ambassador pattern and opens up the possibility of implementing a wide range of really cool tricks for testing in production. Such as facilitating dark launch of a new service version and using the proxy to divert test traffic to the new service under test.

In my view, if your system is built in such a way that you can easily introduce one of these patterns to facilitate live testing, you should definitely go for it. After all, there is only so much data that you can collect when running in an isolated environment whose load and traffic profiles don't really align with the ones of your production systems.

#### Smoke Tests
When it comes to execution, smoke tests are the exact antithesis of functional tests. While functional tests are allowed to execute for long periods of time, smoke tests must execute as quickly as possible. As a result, smoke tests are crafted so as to exercise specific, albeit limited, flows in the user-facing parts of a system that are deemed critical for the system's operation. For example, smoke tests for a social network application would verify the following:
* A user can login with a valid username and password
* Clicking the like button on a post increases the like counter for that post
* Deleting a contact removes them from the user's friends list
* Clicking the logout button signs the user out of the service

#### Chaos Testing
You might be wondering: but, if some failures are statistically unlikely to occur, how can we trigger them in the first place? The only way to do this is to engineer our systems in such a way that failure can be injected on demand. In the Functional tests part deux – testing in production! section, we talked about the ambassador pattern, which can help us achieve exactly that.

The ambassador pattern decouples service discovery and communication from the actual service implementation. This is achieved with the help of a sidecar process that gets deployed with each service and acts as a proxy.

The sidecar proxy service can be used for other purposes, such as conditionally routing traffic based on tags or headers, acting as a circuit breaker, bifurcating traffic to perform A/B testing, logging requests, enforcing security rules, or to inject artificial failures into the system.

From a chaos engineering perspective, the sidecar proxy is an easy avenue for introducing failures. Let's look at some examples of how we can exploit the proxy to inject failure into the system:
* Instruct the proxy to delay outgoing requests or wait before returning upstream responses to the service that initiated the request. This is an effective way to model latency. If we opt not to use fixed intervals but to randomize them, we can inject jitter into intra-service communication.
* Configure the proxy to drop outgoing requests with probability P. This emulates a degraded network connection.
* Configure the proxy for a single service to drop all outgoing traffic to another service. At the same time, all the other service proxies are set up to forward traffic as usual. This emulates a network partition.

### User Stories
One final thing that I would like to point out is that, while each story will record a need for a particular feature, all of them will be completely devoid of any sort of implementation detail. This is quite intentional, and congruent with the recommended practices when working with any Agile framework. Our goal is to defer any technical implementation decisions up to the last possible moment. If we were to decide up-front about how we are going to implement each user story, we would be placing unnecessary constraints on our development process, hence limiting our flexibility and the amount of work we can achieve given a particular time budget.

### Databases
#### Key Value Stores
* memcached, AWS DynamoDB, LevelDB, and SSD-optimized RocksDB
* The data access patterns (insertions, deletions, and lookups) that are used by key-value stores make data partitioning across multiple nodes much easier compared to other database technologies. This property allows key-value stores to scale horizontally so as to accommodate increased traffic demand.
* Use cases
    * Caches! We can use a key-value store as a general-purpose cache for all sorts of things. We could, for instance, cache web pages for a CDN service or store the results of frequently used database queries to reduce the response time for a web application.
    * A distributed store for session data: Imagine for a moment that we operate a high-traffic website. To handle the traffic, we would normally spin up a bunch of backend servers and place them behind a load balancer. Unless our load balancer had built-in support for sticky sessions (always sending requests from the same user to the same backend server), each request would be handled by a different backend server. This could cause issues with stateful applications as they require access to the session data associated with each user. If we tagged each user request with a unique per-user ID, we could use that as a key and retrieve the session data from a key-value store.
    * A storage layer for a database system. The properties of key-value stores make them a very attractive low-level primitive for implementing more sophisticated types of databases. For example, relational databases such as CockroachDB [5] and NoSQL databases such as Apache Cassandra [2] are prime examples of systems built on top of key-value stores.
* The main caveat of key-value stores is that we cannot efficiently search within the stored data without introducing some kind of auxiliary data structure to facilitate the role of an index.

#### Relational
* The standardized way to access and query relational databases is via the use of an English- like structured query language (SQL), which is actually a subset of various domain-specific languages:
    * A data definition language, which includes commands for managing the database schema; for example, creating, altering, or dropping tables, indexes, and constraints
    * A data manipulation language, which supports a versatile set of commands for inserting, deleting, and, of course, querying the database contents
    * A data control language, which provides a streamlined way to control the level of access that individual users have to the database
    * A transaction control language, which allows database users to start, commit, or abort database transactions
* In terms of performance, relational databases such as PostgreSQL [18] and MySQL [17] are generally easy to scale vertically. Switching to a beefier CPU and/or adding more memory to your database server is more or less a standard operating procedure for increasing the queries per second (QPS) or transactions per second (TPS) that the DB can handle. On the other hand, scaling relational databases horizontally is much harder and typically depends on the type of workload you have.
    * For write-heavy workloads, we usually resort to techniques such as data sharding. Data sharding allows us to split (partition) the contents of one or more tables into multiple database nodes. This partitioning is achieved by means of a per-row shard key, which dictates which node is responsible for storing each row of the table. One caveat of this approach is that it introduces additional complexity at query time. While writes are quite efficient, reads are not trivial as the database might need to query each individual node and then aggregate the results together in order to answer even a simple query such as SELECT COUNT(*) FROM X.
    * On the other hand, if our workloads are read-heavy, horizontal scaling is usually achieved by spinning up read-replicas, which mirror updates to one or more primary nodes. Writes are always routed to the primary nodes while reads are handled by the read-replicas (ideally) or even by the primaries if the read-replicas cannot be reached.
* While relational databases are a great fit for transactional workloads and complex queries, they are not the best tool for querying hierarchical data with arbitrary nesting or for modeling graph-like structures.

#### NoSQL
* To achieve a performance boost over relational databases, NoSQL databases have to sacrifice something! Being distributed systems, NoSQL databases must adhere to the rules of the CAP theorem
* Let's briefly analyze the behavior as to how each of the CAP configurations reacts in the presence of errors:
    * Consistency – Partition (CP) tolerance: Distributed systems in this category typically use a voting protocol to ensure that the majority of nodes agree that they have the most recent version of the stored data; in other words, they reach a quorum. This allows the system to recover from network partitioning events. However, if not enough nodes are available to reach quorum, the system will return an error to clients as data consistency is preferred over availability.
    * Availability – Partition (AP) tolerance: This class of distributed systems favors availability over consistency. Even in the case of a network partition, an AP system will try to process read requests, although stale data may be returned to the clients.
    * Consistency – Availability (CA): In practice, all distributed systems are, to some extent, affected by network partitions. Therefore, a pure CA type of system is not really feasible unless, of course, we are talking about a single-node system. We could probably classify a single-node deployment of a traditional relational database as a CA system.
* At the end of the day, the choice of an appropriate NoSQL solution largely depends on your particular use case. What happens, though, if the use case requires all three of these properties? Are we simply out of luck?
    * Fortunately, over the years, several NoSQL solutions (for example, Cassandra [2]) have evolved support for what is now referred to as tunable consistency. Tunable consistency allows clients to specify their desired level of consistency on a per-query basis. For example, when creating a new user account, we would typically opt for strong consistency semantics. On the other hand, when querying the number of views of a popular video, we could dial down the desired level of consistency and settle for an approximate, eventually-consistent, value.

#### Document database
* Document databases are specialized NoSQL databases that store, index, and query complex and possibly deeply nested document-like objects.
* MongoDB, CouchDB, and Elasticsearch

### Miscellaneous
#### What is Link-local addressing?
A PC automatically acquires a 169.254.x.x/16 address if it does not receive an IP address from a DHCP server. If you disable the DHCP server on your home or lab network and issue the "ipconfig/release" and "ipconfig/renew" commands, your PC will receive a 169.254.x.x address.

The purpose of these self-assigned link-local addresses is to facilitate communication with other hosts within the subnet even in the absence of external address configuration (via manual input or DHCP). Unlike in IPv6, implementation of IPv4 link-local addresses is recommended only in the absence of a normal, routable address. Hosts pseudo-randomly generate the last two octets of the address to mitigate address conflicts. Because of the broadcast nature of some local networking protocols (for example, Microsoft's NetBIOS), hosts may be able to detect one another even without any preexisting knowledge of the address scheme.

https://networkengineering.stackexchange.com/questions/24749/what-is-link-local-addressing

#### CQRS
The CQRS pattern belongs to the family of architectural patterns. The key idea behind CQRS is to separate the write and read models exposed by a particular component so they can be optimized in isolation. Commands refer to operations that mutate the state of the model, whereas queries retrieve and return the current model state.

This separation allows us to execute different business logic paths for reads and writes, and, in effect, enables us to implement complex access patterns. For example, writes could be a synchronous process whereas reads might be asynchronous and provide a limited view over the data.

As another example, the component could utilize separate data stores for writes and reads. Writes would eventually trickle into the read store but perhaps the read store data could also be augmented with external data obtained from other downstream components.

### Why CockroachDB
* It can easily scale horizontally just by increasing the number of nodes available to the cluster. CockroachDB clusters can automatically rebalance and heal themselves when nodes appear or go down.
* CockroachDB is fully ACID-compliant and supports distributed SQL transactions.
* The SQL flavor supported by CockroachDB is compatible with the PostgreSQL syntax.
* CockroachDB implements the PostgreSQL wire protocol; this means that we do not require a specialized driver package to connect to the database but can simply use the battle-tested pure-Go Postgres package to connect to the database.

## Issues
### Elasticsearch start failure
* sudo mkdir /usr/local/var/lib/elasticsearch
* sudo chown -R bobbylei /usr/local/var/lib
## Source Code
https://github.com/PacktPublishing/Hands-On-Software-Engineering-with-Golang
## Upto
Page 238

An in-memory Indexer implementation using bleve

Before that: set up the base test package https://github.com/PacktPublishing/Hands-On-Software-Engineering-with-Golang/tree/master/Chapter06/textindexer/index/indextest
