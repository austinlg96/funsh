# CTF SSH Challanges

## Description:
A set of CTF challenges that require players to think outside of the box and script a connection to an SSH server to quickly process output.

## Simplest way to try it out:
  1. Clone the repository: `git clone [repository location]`
  2. Create a NOTICE.txt in the main directory (requires no contents): `touch NOTICE.txt` or `echo "" > NOTICE.txt`
  3. Compile the "shell" binary: `go build`
  4. Run the "shell": `./main` or `.\main.exe`
## Brief Build/Setup Logic:
  1. The "shell" that the user interacts with is built from `main.go` and `customfmt.go`
  2. SSH keys are generated for the host and the "user" (`jdoe`).
  3. A "hardened" docker container is built to try and prevent players from using an alternate/real shell or from being able to run arbitrary code. The container is built using:
     - The shell binary.
     - The SSH keys.
     - A "hardened" sshd_config.
  4. The container is started and port 22 on the host (or any other desired) is forwarded to port 22 on the container.
  5. The user key is distributed to contestants.

## Contents:
 - GoLang "Shell":
   - `customfmt.go`: Wrapper of the `fmt` package to allow for easy simultaneous logging and I/O. 
   - `main.go`: The "shell" program/challenge itself. 
     - `Challenge`: A single question. Has a question, an answer, and a cutoff/time limit.
     - `Level`: A group of questions and a message to be sent at the end.
     - `ResponseDetails`: Information about how a player responded to a question. 
     - `ask()`: Prints/logs the question from a challenge and generates a ResponseDetails.
     - `hex()`: A function that multiplies an integer by 6...except everything is a string. Awkward because `Challenge` was designed with strings in mind (for a few reasons). TODO: Expand types of challenges to support more than just strings.
     - `check()`: Determines which "mode" a player's response is valid for, if any. May seem backwards to determine what mode a player is in based on their response, but the goal was to force users to give different responses to the same questions based on the player's input. Currently supported, in order of priority:
       - slow: Response takes more than 30 seconds.
       - normal: Response is correct and was provided within the time specified.
       - hex: Response is 6 times what would normally be entered.  
     - `check_modes()`: Compares the two "modes" from a question or ensure that they are the same or one is empty/undefined.
     - `run_level()`: Uses `ask()` for each question in the `Level`, verifies that the modes are consistent with `check()` and then returns the mode that was acceptable.
     - `gen_level*()`: Generates a level at run-time so that responses can't be prepared ahead of time. TODO: There are probably better ways to do this with GoLang.
     - `main()`: 
       - Dump the `NOTICE` file (`notice.txt`) to `/dev/null` to require the compiler to actually include t in the binary.
       - Run each `Level` and confirm that modes are remaining consistent.
   - `NOTICE.txt`/`NOTICE.txt.example`: A notice that is contained in the output binary asking that a player who is able to obtain the binary or read arbitrary memory to reach out to the maintainers so that the environment can be futher hardened in the future.
 - Docker image:
   - `dockerfile`
   - `sshd_config`: A config for `sshd` that attempts to harden against the user being able to access a real shell on the system and instead be forced into communicating with the go "shell".
- Solution files:
   - `solver_base.py`: A script that can be used to build a solution to the challenge. Uses pure python + the host's default `ssh` handler. Provided to my colleagues who were testing the challenge to offset some of the difficulty with deadlocking when managing the asynchronous communication between processes. Alternative solutions (such as paramiko) would probably have been more straightforward.
   - `solution.py`: A solution that completes all three modes of the challenge. Naively uses `eval()` because I wrote the strings being parsed.
- Others:
   - `build.sh`: Helper script to rebuild the "shell" and keys then output a `tar`'d docker image.
   - `generation.py`: Helper script that was used to generate some Go code for static levels.

## License

[CC BY-SA](https://creativecommons.org/licenses/by-sa/4.0/)