# Go Notes

[Huge Golang Guide](https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit?pli=1)

## Stack Overflow

1. [What does the . in a Go import statement do](https://stackoverflow.com/a/6478990)
    - Allows functions inside the package to be called directly, but is not recommended

1. [When should you use `errors.As` and when to use `errors.Is` for your own custom errors](https://stackoverflow.com/a/76918940)
    - `errors.Is` for `var` errors, `errors.As` for struct errors that has to be initialized (especially if it has custom fields)

1. [When to use Golang's default MUX versus doing your own](https://stackoverflow.com/a/30063908) 

1. [Should I use ServeMux or http directly in golang](https://stackoverflow.com/q/36248946)

## Good Articles

1. [init() in Go Programming](https://david-yappeter.medium.com/init-in-go-programming-31e2c2bc2371)
    - Read it

1. [Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/)
    - Read it

1. [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)

1. [Channel Axioms](https://dave.cheney.net/2014/03/19/channel-axioms)

## Good Videos

1. [Go concurrency patterns](https://youtu.be/f6kdp27TYZs?si=DqVBjbLDGg31j8XK)
    - Must watch

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

## Good Books

1. [100 Go Mistakes and How to Avoid Them](https://100go.co/)
