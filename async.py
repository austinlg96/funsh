import asyncio

async def main():
    proc = await asyncio.subprocess.create_subprocess_exec(
        "go","run","./main.go", stdin=asyncio.subprocess.PIPE, stdout=asyncio.subprocess.PIPE
    )
    while True:
        line = (await proc.stdout.readline()).decode()
        print(line,end="")
        if line.startswith("Q: "):
            x=eval(line[3:])
            print((await proc.stdout.read(3)).decode(), end ="")
            print(f"{x}")
            proc.stdin.write(f"{x}\n".encode())
        if line == "Flag:abcdefghijklmnop\n":
            break
    proc.stdin.write(("\n"*10).encode())
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
            proc.stdin.write(f"{x}\n".encode())
        # line = (asyncio.wait_for(proc.stdout.readline(), 0.5)).decode()

    proc.stdin.write(b"bob\n")
    print(await proc.stdout.read(1024))
    proc.stdin.write(b"alice\n")
    print(await proc.stdout.read(1024))
    proc.stdin.write(b"quit\n")
    await proc.wait()


asyncio.run(main())