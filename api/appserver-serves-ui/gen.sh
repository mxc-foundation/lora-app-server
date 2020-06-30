#!/usr/bin/env bash

GRPC_GW_PATH=$(go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

PROTOBUF_PATH=$(go list -f '{{ .Dir }}' github.com/golang/protobuf/ptypes)

# generate the gRPC code
protoc -I. -I.. -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --go_out=plugins=grpc:. \
  device.proto \
  application.proto \
  deviceQueue.proto \
  common.proto \
  user.proto \
  gateway.proto \
  organization.proto \
  profiles.proto \
  networkServer.proto \
  serviceProfile.proto \
  deviceProfile.proto \
  gatewayProfile.proto \
  multicastGroup.proto \
  fuotaDeployment.proto \
  internal.proto \
  serverInfo.proto \
  topup.proto \
  wallet.proto \
  withdraw.proto \
  settings.proto \
  staking.proto

# generate the JSON interface code
protoc -I. -I.. -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --grpc-gateway_out=logtostderr=true:. \
  device.proto \
  application.proto \
  deviceQueue.proto \
  common.proto \
  user.proto \
  gateway.proto \
  organization.proto \
  profiles.proto \
  networkServer.proto \
  serviceProfile.proto \
  deviceProfile.proto \
  gatewayProfile.proto \
  multicastGroup.proto \
  fuotaDeployment.proto \
  internal.proto \
  serverInfo.proto \
  topup.proto \
  wallet.proto \
  withdraw.proto \
  settings.proto \
  staking.proto

# generate the swagger definitions
protoc -I. -I.. -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --swagger_out=json_names_for_fields=true:./swagger \
  device.proto \
  application.proto \
  deviceQueue.proto \
  common.proto \
  user.proto \
  gateway.proto \
  organization.proto \
  profiles.proto \
  networkServer.proto \
  serviceProfile.proto \
  deviceProfile.proto \
  gatewayProfile.proto \
  multicastGroup.proto \
  fuotaDeployment.proto \
  internal.proto \
  serverInfo.proto \
  topup.proto \
  wallet.proto \
  withdraw.proto \
  settings.proto \
  staking.proto

# merge the swagger code into one file
#go run swagger/main.go swagger > ../static/swagger/api.swagger.json
