FROM ubuntu:20.04
ADD so/* /opt/dmdbms/bin/
RUN apt-get update && apt-get install -y --no-install-recommends unixodbc  unixodbc-dev && \
    mkdir -p /opt/dmdbms/bin && \
    apt-get clean && \
    rm -rf /var/cache/apt/* && \
    # odbcinst.ini 配置驱动路径
    echo '[DM7 ODBC DRIVER]' > /etc/odbcinst.ini && \
    echo 'Description = ODBC DRIVER FOR DM7' >> /etc/odbcinst.ini && \
    echo 'Driver = /opt/dmdbms/bin/libdodbc.so' >> /etc/odbcinst.ini && \
    # odbc.ini isql 使用
    echo '[dm]' > /etc/odbc.ini && \
    echo 'Description = DM ODBC DSN' >> /etc/odbc.ini && \
    echo 'Driver = DM7 ODBC DRIVER' >> /etc/odbc.ini && \
    echo 'SERVER = 192.168.17.187' >> /etc/odbc.ini && \
    echo 'UID = SYSDBA' >> /etc/odbc.ini && \
    echo 'PWD = SYSDBA' >> /etc/odbc.ini && \
    echo 'TCP_PORT = 5236' >> /etc/odbc.ini

ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/dmdbms/bin: