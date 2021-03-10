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
