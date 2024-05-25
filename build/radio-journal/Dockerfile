FROM golang:1.21-alpine AS build_stage
COPY ../../ /go/src/radio_journal
WORKDIR /go/src/radio_journal
RUN go install .

FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/radio_journal /app_binary/
RUN chmod +x ./radio_journal
EXPOSE 8080/tcp
ENTRYPOINT ./radio_journal

EXPOSE 8080/tcp
CMD ./radio_journal
