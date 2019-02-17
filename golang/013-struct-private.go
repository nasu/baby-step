package main

type S struct {
    Id   int
    name string
}

func main() {
    //_ = S{1} // too few values
    _ = S{1, ""}
    _ = S{Id: 1}
}
