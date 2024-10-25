LDFLAGS="-X 'main.buildDate=$(date)' -X 'main.gitHash=$(git rev-parse HEAD)' -X 'main.buildOn=$(go version)' -w -s "

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gml-windows.exe -trimpath -ldflags "${LDFLAGS}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gml-linux -trimpath -ldflags "${LDFLAGS}"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o gml-linux-arm64 -trimpath -ldflags "${LDFLAGS}"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gml-darwin -trimpath -ldflags "${LDFLAGS}"
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o gml-darwin-arm64 -trimpath -ldflags "${LDFLAGS}"

# sha256
sha256sum gml* > gml-sha256
cat gml-sha256

# chmod 
chmod +x gml-*

# gzip
gzip --best gml*