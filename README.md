# go-crawler
A distributed crawler based on Golang.

---

Details:

[https://blog.csdn.net/chao2016/article/details/81697353](https://blog.csdn.net/chao2016/article/details/81697353)


Tree:

```
go-crawler
├── README.md
├── engine
│   ├── concurrent.go (update)
│   ├── simple.go (update)
│   └── types.go
├── fetcher
│   └── fetcher.go (update)
├── main.go (update)
├── model
│   └── profile.go
├── scheduler
│   └── simple.go (update)
└── zhenai
    └── parser
        ├── city.go
        ├── citylist.go
        ├── citylist_test.go
        ├── citylist_test_data.html
        ├── profile.go
        ├── profile_test.go
        └── profile_test_data.html
```

Download:

```
git clone git@github.com:chao2015/go-crawler.git
```

or download the previous version via [the release page](https://github.com/chao2015/go-crawler/releases)


Run:

```
mv go-crawler/ $GOPATH/src/
cd $GOPATH/src/go-crawler/
go run main.go 
```
Have fun! ^_^
