import asyncio
import time

printlines = True

def log(data: str) -> None:
    with open("shell_log.txt", "a") as o:
        o.write(data)

def skip(line: str) -> bool:
    criteria = ("exit","throw","die","return")
    for item in criteria:
        if item in line:
            return True
    return False

def process_stdout(stdout: str, mode: str) -> str: 
    """Should accept multi-line STDOUT from the challenge ending with an answer prompt ("A: ") and then determine the appropriate response."""
    usenextline = False
    stdin = ""
    if printlines:
        print(stdout, end="")
    for line in stdout.split("\n"):
        if skip(line):
            usenextline = True
            continue
        elif usenextline:
            usenextline = False
            response = eval(line)
        elif line.startswith("Q: "):
            response = eval(line[3:])
        else:
            continue
        if mode =="hex":
            response = response * 6
        if mode == "slow":
            time.sleep(30)
        stdin = f"{response}\n"
    return stdin

def stdin_write(proc, stdin: str) -> None:
    """Writes a string to the stdin of the proces and logs with the log() function."""
    proc.stdin.write(stdin.encode())
    log(stdin)
    if printlines:
        print(stdin,end="")

async def _get_question(proc) -> str:
    """Gets all data from the stdout of the SSH session until either the next answer prompt ("A: ") or the SSH process is killed."""
    try:
        stdout = await proc.stdout.readuntil(b"A: ")
    except asyncio.IncompleteReadError as e:
        stdout = e.partial
    return stdout.decode()

async def get_question(proc) -> str:
    # GOAL: Get stdout from the SSH process without hanging/blocking.
    if proc.returncode is not None:
        return None
    try:
        # Start an async task that tries to get a question from the challenge.
        stdout_task = asyncio.create_task(_get_question(proc))

        # If no data is sent within 10 seconds, assume that something is wrong and throw asyncio.TimeoutError.
        stdout = await asyncio.wait_for(asyncio.shield(stdout_task),timeout=10)
    except asyncio.TimeoutError:
        # Kill the SSH process so that asyncio.IncompleteReadError 
        proc.kill()
        stdout = await stdout_task
    
    return stdout

async def main(mode:str="normal"):

    # Open SSH session to the challenge.
    proc = await asyncio.subprocess.create_subprocess_exec(
        "ssh", "jdoe@192.168.42.183", "-i", "./keys/user_key", stdin=asyncio.subprocess.PIPE, stdout=asyncio.subprocess.PIPE
    )

    question = 1
    try:
        # While the SSH process is still running.
        while proc.returncode is None:

                # Get data from the SSH process
                stdout = await get_question(proc)

                # Record that data
                log(stdout)

                # Determine the response for the data
                response = process_stdout(stdout, mode)

                # stdin the response for the data.
                stdin_write(proc, response)
                if question == 10:
                    stdin_write(proc, "\n"*10)
                question += 1
    except BrokenPipeError:
        pass

if __name__ == "__main__":
    asyncio.run(main())
    asyncio.run(main("hex"))
    asyncio.run(main("slow"))
