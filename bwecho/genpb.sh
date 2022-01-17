#!/bin/zsh

for dir in protos/*; do
  filename=$(basename "$dir")
  f_suffix=$(echo "$filename"|cut -d. -f2)
  if [ "$f_suffix" != "proto" ]; then
    continue
  fi
  protoc --go_out=protos --go-grpc_out=protos "$dir"

done