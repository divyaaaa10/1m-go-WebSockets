# 1M Go Web Sockets
## Introduction 
### web_server 
This consists of the server and client code. The client code establishes 1m persistent connections in a span 30 minutes
### storing_data 
The client code here establishes 10 connections and sends the data of size 1Kb per client through the web sockets to the server, the server stores the data.
### data_500
The client code establishes 1 million connections as well as send data of size 500kb in the span of 30 minutes . However the data is not stored at the server.
