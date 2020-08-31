docker build -t y-test .
docker run -p 80:8080 -e SQLITE_PATH="/data/y-test.db" \
-e SERVER_PORT=80 \
-e JWT_SECRET="2zzFBsHBM7g3F2dmZJxZEK5ttYPYYG_iQwoSoIhXfoQ" \
-e STORAGE_ENDPOINT="127.0.0.1:9000" \
-e STORAGE_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE" \
-e STORAGE_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
-e STORAGE_USESSL=false \
-v ~/go/scr/y-test/:/data \
 y-test