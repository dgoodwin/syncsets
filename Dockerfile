FROM fedora:33
ADD bin/syncsets-api /opt/services/
ADD bin/syncsets-controllers /opt/services/
ENTRYPOINT ["/opt/services/syncsets-api"]
