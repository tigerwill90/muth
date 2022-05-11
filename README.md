# MutH - A thread safe, lock free, http.Handler hot updater

Note that this not make the underlying http.Handler implementation thread safe!

**Install**
````
go get -u github.com/tigerwill90/muth
````

### Usage
````go
mux := http.NewServeMux()
mux.HandleFunc("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("/foo/bar"))
})

h := muth.Handler(mux)
srv := &http.Server{
    Handler: h,
}
_ = srv.ListenAndServe()

// And some time later while the server is running, you can safely update your routing topology
mux := http.NewServeMux()
r.HandleFunc("/foo/bar/baz", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("/foo/bar/baz"))
})
h.Update(r)
````



