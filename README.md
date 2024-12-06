# Go Notes

[Huge Golang Guide](https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit?pli=1)

## Useful Repos

1. [golangci-lint](https://github.com/golangci/golangci-lint)
1. [golangci-lint-config-examples](https://github.com/ccoVeille/golangci-lint-config-examples)
1. [gci](https://github.com/daixiang0/gci)

        gci write --skip-generated -s standard -s default -s localmodule .

1. [gofumpt](https://github.com/mvdan/gofumpt)
1. [mingo](https://github.com/bobg/mingo.git)

## Stack Overflow || Reddit

1. [What does the . in a Go import statement do](https://stackoverflow.com/a/6478990)
    - Allows functions inside the package to be called directly, but is not recommended

1. [When should you use `errors.As` and when to use `errors.Is` for your own custom errors](https://stackoverflow.com/a/76918940)
    - `errors.Is` for `var` errors, `errors.As` for struct errors that has to be initialized (especially if it has custom fields)

1. [When to use Golang's default MUX versus doing your own](https://stackoverflow.com/a/30063908)

1. [Should I use ServeMux or http directly in golang](https://stackoverflow.com/q/36248946)

1. [When to use and not to use generics](https://www.reddit.com/r/golang/comments/17evm2i/comment/k65wps6/)

1. [How to test a unexported (private) function in go (golang)?](https://stackoverflow.com/a/60813569), the [article](https://medium.com/@robiplus/golang-trick-export-for-test-aa16cbd7b8cd)

1. [How to add new methods to an existing type in Go?](https://stackoverflow.com/a/43507669)

1. [Go constants beyond basics](https://www.reddit.com/r/golang/comments/1gll74w/go_constants_beyond_basics/)

1. [Why to not choose Ginkgo](https://www.reddit.com/r/golang/comments/1azj63h/comment/ks1srp2/)

1. [Return zero value for generic types](https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type/70589302#70589302)

1. [Truncating a string in templates](https://stackoverflow.com/questions/23466497/how-to-truncate-a-string-in-a-golang-template/36093426#36093426)

1. [Slices behavior](https://stackoverflow.com/questions/39993688/are-slices-passed-by-value)

## Good Articles

1. [init() in Go Programming](https://david-yappeter.medium.com/init-in-go-programming-31e2c2bc2371)

1. [Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/)

1. [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)

1. [Channel Axioms](https://dave.cheney.net/2014/03/19/channel-axioms)

1. [Proper HTTP Shutdown in Go](https://dev.to/mokiat/proper-http-shutdown-in-go-3fji)

1. [Golang Trick: Export unexport method for test](https://medium.com/@robiplus/golang-trick-export-for-test-aa16cbd7b8cd)

1. [Mastering Go Structs: 7 Advanced Techniques for Efficient Code](https://blog.stackademic.com/mastering-go-structs-7-advanced-techniques-for-efficient-code-fe71f3b56613)

1. [Understanding Go's UTF-8 Support](https://ashwingopalsamy.hashnode.dev/understanding-gos-utf-8-support)

1. [Google style guide assertion libraries](https://google.github.io/styleguide/go/decisions.html#assertion-libraries)

1. [Understanding Singleflight in Go](https://www.codingexplorations.com/blog/understanding-singleflight-in-golang-a-solution-for-eliminating-redundant-work)

1. [Go Singleflight panic and GoExit behaviours](https://victoriametrics.com/blog/go-singleflight/) 

## Bites from the Golang Manual (aka. RTFM)

1. [nil_error behaviors](https://go.dev/doc/faq#nil_error)

1. [Constants](https://go.dev/ref/spec#Constants)

1. [Go 1.23 timer changes](https://tip.golang.org/doc/go1.23#timer-changes)

1. [Go subtests](https://go.dev/blog/subtests)

1. [Go slice tricks](https://go.dev/wiki/SliceTricks)

## Good Videos

1. [Go concurrency patterns](https://youtu.be/f6kdp27TYZs?si=DqVBjbLDGg31j8XK)

1. [7 Deadly Mistakes Beginner Go Developers Make (and how to fix them)](https://www.youtube.com/watch?v=biGr232TBwc)
    - You can label loops apparently
    - map, slice, values in an interface are not addressable
    - Pointer receivers can only accept pointers, while value receivers can accept both values and pointers

1. [Golang struct configuration pattern](https://www.youtube.com/watch?v=MDy7JQN5MN4)
    - Use a struct to pass optional values
    - Allow the constructing function to take a list of functions that can modify the options to allow for extensibility

1. [GopherCon 2016: Francesc Campoy - Understanding nil](https://www.youtube.com/watch?v=ynoY2xz-F8s)
    - 16:30
    - Use nil to disable channels
    - Nil is valid as a read only empty map

1. [The standard library now has all you need for advanced routing in Go](https://www.youtube.com/watch?v=H7tbjKFSg58)

1. [GothamGo 2018 - Things in Go I Never Use by Mat Ryer](https://www.youtube.com/watch?v=5DVV36uqQ4E)
    - Lazy initialization (20:54)
        - ```sync.Once```: Guarantees thing inside code block will be called exactly once
    - You can pass {the struct you want to work on} to {a method called on the struct using dot, ex. ```Struct.method(actualObject)```}, don't do this but its interesting (22:20)

1. [Golang UK Conference 2016 - Dave Cheney - SOLID Go Design](https://www.youtube.com/watch?v=zzAdEt3xZ1M)
    - Single Responsibility Principle
        - Structure functions and types into packages that exhibit natural cohesion
        - Avoid `shared`, `utils`, which causes the packages to have multiple responsibilities
    - [Open / Closed Principle](https://www.freecodecamp.org/news/open-closed-principle-solid-architecture-concept-explained/)
        - Open for extension, closed for modification
        - Use embedding rather than inheritanca
    - Liskov Substitution Principle
        - Implicit interfaces, use small interfaces and express the dependencies between packages as interfaces
        - Require no more, promise no less
    - Interface Segregation Principle
        - Clients should not be forced to depend on methods they don't use
        - Ex. A `Save` function should only need a `Writer`, instead of `Read` or `File`, or even `Close`, therefore it can work with network writer, file writers etc
    - Dependency Inversion Principle
        - High level modules should not depend on low level modules, both should depend on abstractions
        - Abstractions should not depend on details. Details should depend on abstractions
        - Import graph should be acyclic, push responsibility as high as possible in the import tree

1. [Golang UK Conference 2016 - Mat Ryer - Idiomatic Go Tricks](https://www.youtube.com/watch?v=yeetIgNeIkc)
    - 9:16: ```defer log.Println("------")``` even if exits abnormally this code will print
    - defer can also be used for teardown functions for any setup / timers that need to stop at end of the line
    - 21:18: [Retry code](https://github.com/matryer/try/blob/master/try.go)
    - 22:10: Empty structs to group methods methods don't capture the receiver, so the variable need not expose the private type
    - 23:40: Semaphore code, limit hom many stuff is executed at once

1. [How To Build A Chat And Data Feed With WebSockets In Golang?](https://www.youtube.com/watch?v=JuUAEYLkGbM)

## Good Books

1. [100 Go Mistakes and How to Avoid Them](https://100go.co/)

## Notes

1. `runtime.NumCPU`
2. `runtime.GOMAXPROCS()`

## Others

1. [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)

1. [What is the difference between grpc and websocket?](https://stackoverflow.com/a/48113832)

1. [Will WebSocket survive HTTP/2](https://www.infoq.com/articles/websocket-and-http2-coexist/)

1. [gRPC vs WebSockets](https://ably.com/topic/grpc-vs-websocket)

1. [Docker stop containers](https://spacelift.io/blog/docker-stop-container)

1. [Docker distroless Docker exec failed: No such file or directory when CGO is enabled](https://stackoverflow.com/a/72727838)

1. [Golang's null safety proposal discussion](https://github.com/golang/go/issues/49202)

1. [Postgres don't do this](https://wiki.postgresql.org/wiki/Don%27t_Do_This)

## GraphQL

1. [Fragment, minimizes token usage as well](https://graphql.org/learn/queries/#fragments)

1. [Aliases](https://graphql.org/learn/queries/#aliases)

1. [Meta-fields, ex. __typename](https://graphql.org/learn/queries/#meta-fields)

1. [Default variables](https://graphql.org/learn/queries/#meta-fields)

1. [Interface types](https://graphql.org/learn/schema/#interface-types)

1. [Inline fragments](https://graphql.org/learn/queries/#inline-fragments)
    - It allows to show different fields depending on what the runtime type for the interface is

1. [Union types, because why not](https://graphql.org/learn/schema/#union-types)

1. [Input object types](https://graphql.org/learn/schema/#input-object-types)
    - Use objects for input

1. [GraphQL check complexity](https://gqlgen.com/reference/complexity/)

## Fuck Yaml Specifically

1. [noyaml website](https://noyaml.com/)
    - this is good

1. [Yaml: Probably not so great after all](https://www.arp242.net/yaml-config.html)

1. [Valueable lessons in over engineering](https://tinyurl.com/lessons-in-over-engineering)

1. [Yaml test matrix](https://matrix.yaml.info/valid.html)

1. [Yaml strips out underscore](https://drupal.stackexchange.com/questions/238639/why-does-yaml-parser-strippes-out-underscore)
