package utility

import (
    "bytes"
    "compress/zlib"
    "io/ioutil"
)

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

