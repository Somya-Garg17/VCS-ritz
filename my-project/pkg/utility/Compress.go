package utility

import (
    "bytes"
    "compress/zlib"
    "io/ioutil"
    "io"
)

// Compress compresses the input data using zlib
func Compress(data []byte) ([]byte, error) {
    var b bytes.Buffer
    w := zlib.NewWriter(&b)
    if _, err := w.Write(data); err != nil {
        return nil, err
    }
    if err := w.Close(); err != nil {
        return nil, err
    }
    return ioutil.ReadAll(&b)
}

// Decompress decompresses the input data using zlib
func Decompress(data []byte) ([]byte, error) {
    r, err := zlib.NewReader(bytes.NewReader(data))
    if err != nil {
        return nil, err
    }
    defer r.Close()
    var out bytes.Buffer
    _, err = io.Copy(&out, r)
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(&out)
}

