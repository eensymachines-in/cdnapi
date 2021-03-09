
FROM kneerunjun/gogingonic:latest 
# vanilla image of the go gin driver
COPY . .
RUN go mod download 
RUN go build -o cdnapi .