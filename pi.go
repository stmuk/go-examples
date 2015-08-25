package main

import (
    "fmt"
)

func main () {

    n = 100000;
    delta = 1.0 / n;
    sum = 0.0;
    i = 1;
    for ( i <= n) {
          x = ( i - 0.5 ) * delta;
            sum += 1.0 / ( 1.0 + x * x );
              i += 1;
          }
          pi = 4.0 * delta * sum;
          fmt.Println(pi);
}

