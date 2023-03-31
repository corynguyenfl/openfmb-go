#!/bin/bash
protoc -I=coordination-messages/protos --go_out=. coordination-messages/protos/messages.proto