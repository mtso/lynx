package lynx

import (
    "io"
    "os"
    "log"
    "bufio"
)

// STOLEN STRAIGHT FROM
// http://stackoverflow.com/a/9739903/2684355
func CopyFromTo(src, dest string) (err error) {
    // open input file
    fi, err := os.Open(src)
    if err != nil {
        return
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // open output file
    fo, err := os.Create(dest)
    if err != nil {
        return
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            return
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            return
        }
    }

    if err = w.Flush(); err != nil {
        return
    }

    return
}