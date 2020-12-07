FROM fedora:33
ADD bin/syncsets-server /opt/services/
ADD bin/syncsets-controllers /opt/services/
ENTRYPOINT ["/opt/services/syncsets-server", "--port", "7070"]
