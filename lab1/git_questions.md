# Multiple Choice Questions for Git

Please consult [Lecture 5](https://missing.csail.mit.edu/2020/version-control/) from The Missing Semester, and the [Git book](https://git-scm.com/book/en/v2).

Answer the following questions by editing this file by replacing the `[ ]` for the correct answer with `[x]`.
Only one choice per question is correct.
Selecting more than one choice will result in zero points.
No other changes to the text should be made.

1. What does the following Git command do?
   Hint: See `git help commit` for options.

    ```console
    git commit -a -m "description of changes"
    ```

    - [ ] a) stages all changes to the repository and commits with the message "description of changes"
    - [ ] b) stages all changes to the repository and opens an editor to enter the commit message, which initially contains "description of changes"
    - [ ] c) stages changes to all files already tracked by Git which have been modified or deleted, and opens an editor to enter the commit message, which initially contains "description of changes"
    - [ ] d) stages changes to all files already tracked by Git which have been modified or deleted, and commits with the message "description of changes"

2. What does the following Git command do?

    ``` console
    git remote -v
    ```

    - [ ] a) shows a list of remote repositories along with their URLs
    - [ ] b) verifies that the remote repository exists
    - [ ] c) removes the reference to the remote repository and outputs all changes
    - [ ] d) sets the name of the remote repository to `v`

3. How would you revert *unstaged* (not yet added with `git add`) changes to `README.md` in the current branch?

    - [ ] a) `git reset HEAD README.md`
    - [ ] b) `git reset --hard README.md`
    - [ ] c) `git checkout -- README.md`
    - [ ] d) `git revert README.md`

4. Why do we add `course-assignments` as a remote repository in when setting up our Git for this course (command sequence listed below)?

    ```console
    git clone https://github.com/dat320-2020/username-labs assignments
    cd assignments
    git remote add course-assignments https://github.com/dat320-2020/assignments
    ```

    - [ ] a) so that we automatically pull changes from the `assignments` repo into the `username-labs` repo with `git pull` when they become available
    - [ ] b) so that our changes to the `username-labs` repo are also reflected in the `assignments` repo
    - [ ] c) so that we can pull updates from the `assignments` repo, such as new or updated lab assignments
    - [ ] d) so that the initial state of our `username-labs` repo will be the same as the state of the `assignments` repo

5. If the remote repositories are set up as in the following segment, how would you pull changes from the `dat320-2020/assignments` repo into the `dat320-2020/student-labs` repo you are working in?

    ```console
    course-assignments  git@github.com:dat320-2020/assignments.git (fetch)
    course-assignments  git@github.com:dat320-2020/assignments.git (push)
    origin              git@github.com:dat320-2020/student-labs.git (fetch)
    origin              git@github.com:dat320-2020/student-labs.git (push)
    ```

    - [ ] a) `git pull dat320-2020/assignments.git dat320-2020/student-labs.git`
    - [ ] b) `git pull course-assignments master`
    - [ ] c) `git pull master course-assignments`
    - [ ] d) `git pull course-assignments origin`
    - [ ] e) `git pull`

6. What does the following Git command do?

    ```console
    git checkout -b experimental
    ```

    - [ ] a) creates a new branch called `experimental` and switches to it
    - [ ] b) switches to the branch `experimental`, but only if it exists
    - [ ] c) resets any changes made to the `experimental` branch
    - [ ] d) opens a new window which displays the difference between the current and the `experimental` branches

7. If you had the following `.gitignore` file and file tree, and then ran `git add .`, which changes would be staged?
   Hint: Read about the [`.gitignore` file](https://git-scm.com/docs/gitignore) in the Git docs, specifically the *Pattern Format* section.

    ```console
    # file tree (from current path)
    go/
        .DS_Store
        main.go
    notes/
        personal-notes.md
        work-notes.md
    python/
        main.py
    .DS_Store

    # .gitignore content:
    /.DS_Store
    /notes
    !/notes/work-notes.md
    *.py
    ```

    - [ ] a) `go/.DS_Store`, `go/main.go` and `notes/work-notes.md`
    - [ ] b) `go/main.go`, `notes/work-notes.md` and `python/main.py`
    - [ ] c) `go/main.go` and `notes/work-notes.md`
    - [ ] d) `go/.DS_Store` and `go/main.go`

8. Which command lets you see the changes made to `file.txt` relative to the staging area (i.e. unstaged changes to `file.txt`)?

    - [ ] a) `git diff file.txt`
    - [ ] b) `git diff HEAD file.txt`
    - [ ] c) `git diff --cached file.txt`

9. Which command lets you see all changes made to `file.txt` since the latest commit?

    - [ ] a) `git diff file.txt`
    - [ ] b) `git diff HEAD file.txt`
    - [ ] c) `git diff --cached file.txt`

10. When you run `git pull` and get a merge conflict, which of the below approaches is most suited to resolving the merge conflict?

    - [ ] a) make a backup by copying the directory to another location, delete the original directory, clone the repository again, then manually copy changes into the newly cloned repository from the backup
    - [ ] b) make a backup by copying the directory to another location, revert the repository back to a previous commit without any conflicts, then manually copy changes into the repository while avoiding conflicts
    - [ ] c) fix the conflict by editing the conflicting files with a text editor, then remove the lines starting with `>>>>`, `====` and `<<<<` which were added by Git.
