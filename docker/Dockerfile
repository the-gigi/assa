FROM debian:11-slim AS builder

ARG VERSION

COPY install.sh .
RUN ./install.sh $VERSION

FROM busybox
COPY --from=builder /usr/bin/kubectl /usr/bin/kubectl
CMD ["sleep", "1d"]
