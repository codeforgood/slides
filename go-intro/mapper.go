package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "sync"
  "time"
  )

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

// START MERGE
func merge(cs ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out := make(chan string, 10)
    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan string) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
// END MERGE

// START READ
func readFile(file string) <-chan string{
  out := make(chan string, 10)
  go func(){ // HL
    fmt.Println("Reading")
    f, err := os.Open(file)
    if err != nil {
      panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      out <- scanner.Text()
    }
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
    close(out)
  }() // HL
  return out
}
// END READ

// START FILTER
func filter(in <-chan string) <-chan string{
  out := make(chan string)
  go func(){
    fmt.Println("Filtering")
    for line := range in{
      if line[0:2] == "36" {
        out <- line
      }
    }
    close(out)
  }()
  return out
}
// END FILTER

// START WRITE
func writeFile(in <-chan string, file string){
  fmt.Println("writing")
  f, err := os.Create(file)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  w := bufio.NewWriter(f)
  for line := range in{
    fmt.Println(line)
    w.WriteString(line + "\n")
  }
  w.Flush()
}
// END WRITE

// START MAIN
func main(){
  defer timeTrack(time.Now(), "gofs")
  fmt.Println("Go - File Filter")
  input_chan := readFile("test.txt")
  out1 := filter(input_chan)
  out2 := filter(input_chan)
  out3 := filter(input_chan)
  out4 := filter(input_chan)
  out_chan := merge(out1, out2, out3, out4)
  writeFile(out_chan, "out.out")
  fmt.Println("Exciting from module")
}
// END MAIN
