#!/bin/bash

terminate() {
        echo "Received -TERM. Sending -TERM to child process $PID..." >&2
        kill -TERM "$PID"
        sleep 60

        if ps -p "$PID" > /dev/null; then
                echo "Child process $PID did not terminate. Sending -KILL..." >&2
                kill -KILL "$PID"
        fi

        if [ -f "binary_name.txt" ]; then
          BINARY=$(cat binary_name.txt)
          if [ -n "$BINARY" ] && [ -f "$BINARY" ]; then
            rm "$BINARY"
          fi
        fi
        mv new_binary_name.txt binary_name.txt
        exit 0
}

trap terminate TERM

while true; do
        ./"$1" &>> app.log &
        PID=$!

        wait $PID
        echo "Application crashed with exit code $?. Restarting..." >&2

        mv my.log my-crashed.log
        mv gin.log gin-crashed.log
        sleep 5
done
