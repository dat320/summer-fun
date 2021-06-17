# Multiple Choice Questions for Shell Commands

Answer the following questions by editing this file by replacing the `[ ]` for the correct answer with `[x]`.
Only one choice per question is correct.
Selecting more than one choice will result in zero points.
No other changes to the text should be made.

1. Which command lists the sizes of all the files and folders in top level of the current directory in human readable format?
   Note that if you are using max command line option for depth is `-d`.
   (Only Linux; macOS incompatible.)

    - [ ] a) `du -h -l .`
    - [ ] b) `du -h -–max-depth=1 .`
    - [ ] c) `du --max-depth=1`
    - [ ] d) `du -h -a --max-depth=1 .`

2. Which command continuously updates the displayed contents of a file named `file.txt` in real time, while its contents are being modified by some other process?

    - [ ] a) `cat file.txt`
    - [ ] b) `cat -f file.txt`
    - [ ] c) `tail -f file.txt`
    - [ ] d) `head -f file.txt`

3. Which command removes a non-empty directory called `temp_files`?

    - [ ] a) `rm temp_files`
    - [ ] b) `rm -r temp_files`
    - [ ] c) `rmdir temp_files`
    - [ ] d) `rem temp_files`

4. Which command prints the 10 most recent kernel messages?

    - [ ] a) `dmesg -k`
    - [ ] b) `dmesg -k | tail`
    - [ ] c) `dmesg | head`
    - [ ] d) `dmesg`

5. Which command is used to display the manual pages for the command cat?

    - [ ] a) `man cat`
    - [ ] b) `manual cat`
    - [ ] c) `? cat`
    - [ ] d) `guide cat`

6. Which command will show the first 10 lines of `readme.txt`?

    - [ ] a) `cat -10 readme.txt`
    - [ ] b) `less -10 readme.txt`
    - [ ] c) `tail readme.txt`
    - [ ] d) `head readme.txt`

7. Which command renames a file called `file1.txt` to `file2.txt`?

    - [ ] a) `mv file1.txt file2.txt`
    - [ ] b) `cp file1.txt file2.txt`
    - [ ] c) `ren file1.txt file2.txt`
    - [ ] d) `ren file2.txt file1.txt`

8. Which command will show all symbolic links?

    - [ ] a) `ls -l`
    - [ ] b) `ls -a`
    - [ ] c) `find . -type l -ls`
    - [ ] d) `find . -type f -ls`

9. Which command will display the contents of `readme.txt` with line numbers?

    - [ ] a) `cat readme.txt`
    - [ ] b) `cat -l readme.txt`
    - [ ] c) `cat -A readme.txt`
    - [ ] d) `cat -n readme.txt`

10. Which command will count only the number of lines in `readme.txt`?

    - [ ] a) `wc readme.txt`
    - [ ] b) `wc -l readme.txt`
    - [ ] c) `wc -m readme.txt`
    - [ ] d) `wc -n readme.txt`

11. Which command will display a list of currently logged in users on the system?

    - [ ] a) `whoami`
    - [ ] b) `who`
    - [ ] c) `top`
    - [ ] d) `ps -al`

12. Which command will remove the trailing new line from echoing hello?

    - [ ] a) `echo -n hello`
    - [ ] b) `echo \n hello`
    - [ ] c) `echo n hello`
    - [ ] d) `echo /n hello`

13. Which command is used to change your password?

    - [ ] a) `password`
    - [ ] b) `pwd`
    - [ ] c) `passwd`
    - [ ] d) `pw`

14. The command `history` will show the commands previously run in the terminal. For example

    ```text
    1052  ls -al
    1053  go test -v
    1054  pwd
    1055  cd ..
    1056  history
    ```

    How can you repeat command `ls -al`?
    - [ ] a) `repeat 1052`
    - [ ] b) `redo 1052`
    - [ ] c) `!1052`
    - [ ] d) `1052`

15. What does the `less` command do?

    - [ ] a) Displays the contents of a file in a manner that allows users to move forwards or backwards through the file.
    - [ ] b) Displays the contents of a file in a manner that allows users to move only forwards through the file.
    - [ ] c) Displays the contents of a file in a manner that allows users to move only backwards through the file.
    - [ ] d) Allows a user to edit a file.

16. How can you exit the `less` command?

    - [ ] a) Esc
    - [ ] b) q
    - [ ] c) z
    - [ ] d) x

17. What command will display the running processes of the current user?
    (Only Linux; macOS behaves differently.)

    - [ ] a) `ps -u <your user ID or user name>`
    - [ ] b) `ps -a`
    - [ ] c) `ps -e`
    - [ ] d) `ps`

18. What command can be used to find the process(es) consuming the most CPU?

    - [ ] a) `iostat`
    - [ ] b) `netstat`
    - [ ] c) `uptime`
    - [ ] d) `top`

19. What does the `screen` command do?

    - [ ] a) clears the screen
    - [ ] b) starts a virtual terminal
    - [ ] c) closes terminal
    - [ ] d) prints screen

20. Which command sorts `file.txt` according to descending numerical order?

    - [ ] a) `sort -r file.txt`
    - [ ] b) `sort file.txt`
    - [ ] c) `sort -n file.txt`
    - [ ] d) `sort -r -n file.txt`
