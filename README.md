# Compare the performance of json vs. toml as config format using golang

This small setup aims to compare the usage of json vs. toml as a config for go related projects. For json the go built-in
parser is used while for toml we use https://github.com/pelletier/go-toml.

## Setup 

```sh
go get https://github.com/pelletier/go-toml
```

## Run benchmark

```sh
go test -bench=. -benchmem -benchtime=10s
```

## Result

```sh
BenchmarkTomlConfig-8	  200000	     91145 ns/op	   12592 B/op	     349 allocs/op
BenchmarkJsonConfig-8	 1000000	     16829 ns/op	    1320 B/op	      16 allocs/op
```

From the performance point of view json clearly wins over toml.

## Further read suggestions

* http://everythingsysadmin.com/2015/06/toml-vs-json.html

