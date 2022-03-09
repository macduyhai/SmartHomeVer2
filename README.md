# Smart QC Monitor

Build a environment for IoT(Embedded soft ware, Back end data base, Mobile App)

<h1>Build server</h1>
<p>How to run this example</p>

<p>step 1: create volume<br>
    $ docker volume create smarthome_db_data<br>
step 2: run two containers with docker compose<br>
    $ docker-compose up<br>
step 3: data base migration
    - sudo docker exec -it mysql-qc-monitor /bin/bash
    - login mysql container with user: root, password: bW90aGVyIGZ1Y2tlciBub29i<br>
    - create mysql table by copy file scripts/log.sql and scripts/user.sql, paste to mysql terminal<br>
</p>
<p>List API support </p>
[GIN-debug] <b> GET </b>    /ping                  <br>
[GIN-debug] <b>POST </b>  /login                   <br>
[GIN-debug] <b>POST </b>  /api/v1/account          <br>
[GIN-debug] <b>GET  </b> /api/v1/device/download/:id/:name <br>
[GIN-debug] <b>POST </b>  /api/v1/device/add        <br>
[GIN-debug] <b>POST </b>  /api/v1/device/list       <br>
[GIN-debug] <b>POST </b>  /api/v1/device/delete     <br>
[GIN-debug] <b>POST </b> /api/v1/device/edit       <br>
[GIN-debug] <b>POST </b>  /api/v1/device/upload    <br>
[GIN-debug] <b>POST </b> /api/v1/device/push       <br>
[GIN-debug] <b>POST  </b> /api/v1/device/getstatus  <br>
[GIN-debug] <b>POST  </b> /api/v1/media/add         <br>
[GIN-debug] <b>POST  </b> /api/v1/media/list        <br>
[GIN-debug] <b>POST  </b> /api/v1/media/delete      <br>
[GIN-debug] <b>POST  </b> /api/v1/log              <br>
[GIN-debug] <b>GET   </b> /api/v1/log              <br>
[GIN-debug] <b>GET   </b> /api/v1/analysis/tag     <br>
[GIN-debug] <b>GET   </b> /api/v1/analysis/day     <br>
[GIN-debug] <b>GET   </b> /api/v1/average/day      <br>
