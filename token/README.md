# Why is there this JWTS

In a certain company, due to the previous integration of OSS direct transmission, but when migrating to Google, the direct transmission function was not found yet. Therefore, we decided to write our own interface to provide the client with direct transmission

# Method List

~~~
BuildToken(aud, disk, bucket string) (*TokenResponse, error)
ParseToken(token string) (*FlysystemClaims, error)
WithTokenUploadMultipart(fs *flysystem.Flysystem, token, fileName string, file *multipart.FileHeader) (*Response, error) 
WithTokenUploadReader(fs *flysystem.Flysystem, token, fileName string, reader io.Reader) (*Response, error)
WithTokenUploadFilePath(fs *flysystem.Flysystem, token, fileName, filePath string) (*Response, error)
WithTokenUploadBase64(fs *flysystem.Flysystem, token, fileName, base64 string) (*Response, error)
WithTokenDelete(fs *flysystem.Flysystem, token, fileName string) (*Response, error)
Delete(fs *flysystem.Flysystem, disk, bucket, fileName string) (*Response, error) 
UploadBase64(fs *flysystem.Flysystem, disk, bucket, fileName, base64Str string) (*Response, error)
UploadReader(fs *flysystem.Flysystem, disk, bucket, fileName string, reader io.Reader) (*Response, error)
UploadByte(fs *flysystem.Flysystem, disk, bucket, fileName string, contents []byte) (*Response, error) 
~~~

# Code message

| code | Message               |
| ---- | --------------------- |
| 203  | Token is empty        |
| 403  | Token parsing failed  |
| 411  | file.Open()           |
| 305  | Driver  not found     |
| 204  | Write error           |
| 412  | Base64 parsing failed |
| 404  | gfs.Delete err        |

