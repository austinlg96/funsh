import asyncio
import sys
import time

LEVEL1FLAG = "flag{level1flag}"
LEVEL2FLAG = "flag{level2flag}"
NORMALFLAG = "flag{normalflag}"
SLOWFLAG = "flag{slowflag}"
HEXFLAG = "flag{hexflag}"

async def main():
    proc = await asyncio.subprocess.create_subprocess_exec(
        "ssh", "sshuser@192.168.134.131", stdin=asyncio.subprocess.PIPE, stdout=asyncio.subprocess.PIPE
    )
    while True:
        line = (await proc.stdout.readline()).decode()
        print(line,end="")
        if line.startswith("Q: "):
            x=eval(line[3:])
            print((await proc.stdout.read(3)).decode(), end ="")
            print(f"{x}")
            time.sleep(5)
            proc.stdin.write(f"{x}\n".encode())
        if line == LEVEL1FLAG+"\n":
            break
    # proc.stdin.write(("\n"*10).encode())
    while True:
        line = (await proc.stdout.readline()).decode()
        print(line,end="")
        if line.startswith("Q: "):
            if "exit" in (line) or "new" in (line) or "die" in (line) or "return" in (line):
                line = "   " + (await proc.stdout.readline()).decode()
                print(line,end="")
            x=eval(line[3:])
            print((await proc.stdout.read(3)).decode(), end ="")
            print(f"{x}")
            time.sleep(5)
            proc.stdin.write(f"{x}\n".encode())
        # line = (asyncio.wait_for(proc.stdout.readline(), 0.5)).decode()

    proc.stdin.write(b"bob\n")
    print(await proc.stdout.read(1024))
    proc.stdin.write(b"alice\n")
    print(await proc.stdout.read(1024))
    proc.stdin.write(b"quit\n")
    await proc.wait()


asyncio.run(main())