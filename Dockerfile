FROM fedora:33
ADD bin/syncsets-api /opt/services/
ENTRYPOINT ["/opt/services/syncsets-api"]
