from subprocess import Popen, PIPE

from queue import Queue, Empty
from threading import Thread

# https://stackoverflow.com/questions/375427/a-non-blocking-read-on-a-subprocess-pipe-in-python

def enqueue_output(out, queue):
    try:   
        x = iter(out.readline, b'')
        for line in x:
            queue.put(line)
    except ValueError:
        print("FILE CLOSED")
    exit()

with Popen(["go","run","./main.go"], stdin=PIPE, stdout=PIPE, universal_newlines=True, text=True) as process:
    q = Queue()
    t = Thread(target=enqueue_output, args=(process.stdout, q))
    t.daemon = True
    t.start()

    line = ""
    while True:
        try:
            line = q.get_nowait()
        except Empty:
            t.join()
            t = Thread(target=enqueue_output, args=(process.stdout, q))
        else:
            print(line,end="")
            if line.startswith("Q:"):
                x=0
                exec(f"x={line[3:]}")
                process.stdin.writelines(("2","3","4"))
            line = ""