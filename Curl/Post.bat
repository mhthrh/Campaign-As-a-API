
@REM :start
@REM :: Post Add Score
@REM 	curl -Lvso /dev/null -d  "@Score.json" -X POST http://localhost:8585/scores
@REM 	curl GET http://localhost:8585/Score/49385234
@REM goto start

curl -Lvso /dev/null -d  "@Score.json" -X POST http://localhost:8585/scores