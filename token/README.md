# Why is there this JWTS

In a certain company, due to the previous integration of OSS direct transmission, but when migrating to Google, the direct transmission function was not found yet. Therefore, we decided to write our own interface to provide the client with direct transmission

# Method List

~~~
BuildToken(aud, disk, bucket string) (*TokenResponse, error)
ParseToken(token string) (*FlysystemClaims, error)
WithTokenUploadMultipart(fs *flysystem.Flysystem, token, fileName string, file *multipart.FileHeader) error
WithTokenUploadReader(fs *flysystem.Flysystem, token, fileName string, reader io.Reader) error 
WithTokenUploadFilePath(fs *flysystem.Flysystem, token, fileName, filePath string) error
WithTokenUploadBase64(fs *flysystem.Flysystem, token, fileName, base64 string) error
UploadReader(fs *flysystem.Flysystem, disk, bucket, fileName string, reader io.Reader) error
UploadBase64(fs *flysystem.Flysystem, disk, bucket, fileName, base64Str string) error 
UploadByte(fs *flysystem.Flysystem, disk, bucket, fileName string, contents []byte) error
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

