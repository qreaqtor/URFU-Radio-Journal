FROM golang:1.21-alpine AS build_stage
COPY . /go/src/radio-journal
WORKDIR /go/src/radio-journal
RUN go install ./cmd/radio-journal/

FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/src/radio-journal/$CONFIG_PATH /app_binary/$CONFIG_PATH
COPY --from=build_stage /go/bin/radio-journal /app_binary/
RUN chmod +x ./radio-journal
ENTRYPOINT ./radio-journal

EXPOSE 8080
CMD ./radio-journal
