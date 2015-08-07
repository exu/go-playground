package main

import "testing"
import "text/template"
import "bytes"
import "fmt"
import "log"

// Parallel benchmark for text/template.Template.Execute on a single object.

func main() {
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		i := 1
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer
			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()
				err := templ.Execute(&buf, "World")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(buf.Bytes()), i)
				i++
			}
		})
	})
}
