package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
)

func fibonacci() func(int) int {
  num1 := 0
  num2 := 1
  return func(x int) int {
    num3 := num1 + num2
    num1 = num2
    num2 = num3
    return num3
  }
}

func getFib(w http.ResponseWriter, r *http.Request) {
  num, err := strconv.Atoi(mux.Vars(r)["num"])
  if err != nil || num < 0{
    json.NewEncoder(w).Encode("Please add a positive number to your query such as localhost:8000/10")
  } else {
      sequence := []int{}
      f := fibonacci()
        for i := 0; i < num; i++ {
          if i <= 1 {
            sequence = append(sequence, i)
          } else {
            sequence = append(sequence, f(i))
          }
        }
      json.NewEncoder(w).Encode(sequence)
    }

  }

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", getFib).Methods("GET")
  router.HandleFunc("/{num}", getFib).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}
