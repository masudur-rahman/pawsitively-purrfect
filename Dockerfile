FROM gcr.io/distroless/static

COPY bin/linux_arm64/pawsitively-purrfect /pawsitively-purrfect
COPY configs/.pawsitively-purrfect.yaml /configs/.pawsitively-purrfect.yaml

USER 65535:65535
ENV HOME /

ENTRYPOINT ["/pawsitively-purrfect"]
