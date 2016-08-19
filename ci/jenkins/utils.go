package main

import (
    "io"
    "os"
    "fmt"
)

func Copy(src, dst string) (int64, error) {
    src_file, err := os.Open(src)
    if err != nil {
        return 0, err
    }
    defer src_file.Close()

    src_file_stat, err := src_file.Stat()
    if err != nil {
        return 0, err
    }

    if !src_file_stat.Mode().IsRegular() {
        return 0, fmt.Errorf("%s is not a regular file", src)
    }

    dst_file, err := os.Create(dst)
    if err != nil {
        return 0, err
    }
    defer dst_file.Close()
    return io.Copy(dst_file, src_file)
}
