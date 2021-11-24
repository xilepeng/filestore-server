package main

import (
	"LeiliNetdisk/store/ceph"
	"fmt"
	"os"

	"gopkg.in/amz.v1/s3"
)

func main() {
	bucket := ceph.GetCephBucket("userfile")

	d, _ := bucket.Get("/ceph/7c2bfbc5301dff40d8bd6d61fc134034fb31af48")
	tmpFile, _ := os.Create("/tmp/test_file")
	tmpFile.Write(d)
	return

	// // 创建一个新的bucket
	err := bucket.PutBucket(s3.PublicRead)
	fmt.Printf("create bucket err: %v\n", err)

	// 查询这个bucket下面指定条件的object keys
	res, _ := bucket.List("", "", "", 100)
	fmt.Printf("object keys: %+v\n", res)

	// // 新上传一个对象
	err = bucket.Put("/testupload/a.txt", []byte("just for test"), "octet-stream", s3.PublicRead)
	fmt.Printf("upload err: %+v\n", err)

	// // 查询这个bucket下面指定条件的object keys
	res, err = bucket.List("", "", "", 100)
	fmt.Printf("object keys: %+v\n", res)
}
