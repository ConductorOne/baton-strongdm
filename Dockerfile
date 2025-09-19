FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-strongdm"]
COPY baton-strongdm /