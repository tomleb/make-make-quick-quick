module github.com/tomleb/make-make-quick-quick/local-replace/foo

go 1.24.2

// replace github.com/tomleb/make-make-quick-quick/local-replace/bar => ../bar

require github.com/tomleb/make-make-quick-quick/local-replace/bar v0.0.0-20250503041349-5bf5bd4edbb7
