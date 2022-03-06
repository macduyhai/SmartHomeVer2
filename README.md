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
[GIN-debug] GET    /ping                     --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).Ping-fm (4 handlers)
[GIN-debug] POST   /login                    --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).Login-fm (4 handlers)
[GIN-debug] POST   /api/v1/account           --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).CreateUser-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/download/:id/:name --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).Download-fm (4 handlers)
[GIN-debug] POST   /api/v1/device/add        --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).AddDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/list       --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).ListDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/delete     --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).DeleteDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/edit       --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).EditDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/upload     --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).Upload-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/push       --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).PushDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/device/getstatus  --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).GetstatusDevice-fm (5 handlers)
[GIN-debug] POST   /api/v1/media/add         --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).AddMedia-fm (5 handlers)
[GIN-debug] POST   /api/v1/media/list        --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).ListMedia-fm (5 handlers)
[GIN-debug] POST   /api/v1/media/delete      --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).DeleteMedia-fm (5 handlers)
[GIN-debug] POST   /api/v1/log               --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).CreateLog-fm (6 handlers)
[GIN-debug] GET    /api/v1/log               --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).GetLogs-fm (6 handlers)
[GIN-debug] GET    /api/v1/analysis/tag      --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).AnalysisByTag-fm (6 handlers)
[GIN-debug] GET    /api/v1/analysis/day      --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).AnalysisByDay-fm (6 handlers)
[GIN-debug] GET    /api/v1/average/day       --> github.com/macduyhai/SmartHomeVer2/controlers.(*Controller).GetAverageByDay-fm (6 handlers)
