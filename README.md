### Example

```go
ctx := context.Background()
endpointURL := "[s3-endpoint]"
accessKey := "[access-key]"
secretKey := "[secret-key]"
c := storage.NewClient(accessKey, secretKey, endpointURL)
c.Put(ctx, "bucket/path/to/file", []byte("sup"))
c.Get(ctx, "bucket/path/to/file")
```
