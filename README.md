# pval = Pointer to Value
Since Go does not allow addressing a constant, and literal primitives are considered constants in this regard, it makes some code long-winded and noisy. I use this a lot in unit tests to keep them small and on-topic, but it's generally useful for areas of code where terseness is preferred over idiomatic grammar. I was annoyed, so I created this package.

Replace unnecessarily long code like the following:

```
var result *string

// ... do stuff

foo := "bar"
result = &foo
```

With this:

```
var result *string

// ... do stuff

result := pval.String("bar")
```

It's even worse if you happen to need a slice of pointers. Replace this eyesore:

```
var result []*string

// ... do stuff

foo1 := "bar1"
foo2 := "bar2"
foo3 := "bar3"
foo4 := "bar4"
result = append(result, &foo1, &foo2, &foo3, &foo4)
```

With this (the plural version):

```
var result []*string

// ... do stuff

result := pval.Strings("bar1", "bar2", "bar3", "bar4")
```

## License
Creative Commons CC0
![alt text](https://i.creativecommons.org/p/zero/1.0/88x31.png "Creative Commons CC0")

To the extent possible under law, Gerry Hernandez has waived all copyright and related or neighboring rights to pval.