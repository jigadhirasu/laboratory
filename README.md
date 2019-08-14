

## protoc typescript

protoc -I=${INPUT_DIR} ${INPUT_FILE} --js_out=import_style=commonjs,binary:${OUTPUT_DIR} --grpc-web_out=import_style=typescript,mode=grpcweb:${OUTPUT_DIR}

## protoc golang

protoc --go_out=plugins=grpc:${OUTPUT_DIR} ${INPUT_FILE}