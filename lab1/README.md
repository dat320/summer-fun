# Lab 1: Introduction to Unix

| Lab 1: | Introduction to Unix |
| ---------------------    | --------------------- |
| Subject:                 | DAT320 Operating Systems and Systems Programming |
| Grading:                 | Pass/fail |
| Submission:              | Individually |

## Table of Contents

1. [Introduction](#introduction)
2. [The Linux Lab](#the-linux-lab)
3. [Task: Sign up for Unix Account - Do It Now](#task-sign-up-for-unix-account---do-it-now)
4. [The Missing Semester of Your CS Education by MIT](#the-missing-semester-of-your-cs-education-by-mit)
5. [Additional Resources and Tips](#additional-resources-and-tips)
6. [Task: Unix/Linux and Git Multiple-Choice Questions](#task-unixlinux-and-git-multiple-choice-questions)
7. [Remote Login with Secure SHell (SSH)](#remote-login-with-secure-shell-(ssh))
8. [Remote Access to the Linux Machines](#remote-access-to-the-linux-machines)
9. [Task: Logging Into the Linux Lab and Setting Up Password-less Logins](#task-logging-into-the-linux-lab-and-setting-up-password-less-logins)
10. [Submitting to QuickFeed](#submitting-to-quickfeed)

## Introduction

The overall aim of the labs in this course is to learn how to develop systems where you need to access operating system resources, and that may require some low-level tuning to obtain the desired performance.
We will do this through a series of lab assignments that will expose you to numerous developer tools, and by developing applications in the Go programming language.
We will not implement our own operating system, but in some of the assignments we try to emulate pieces of an operating system, such as memory management and scheduling.

In this first lab though, we will get started with some command-line tools, some of which are based on the Missing Semester course from MIT.

## The Linux Lab

Most lab assignments can be performed on your local machine.
If you already run Linux or macOS on your laptop, you should be ready to go.
Linux and macOS are to a large extent relatively similar at the command level.
If you are running Windows, please consult the instructions [here](https://github.com/dat320-2020/course-info/blob/master/setup-wsl.md).

However, we will also be using some machines in our Linux lab, named `pitter1` - `pitter20`,
through remote access; see below for more information about this.
All necessary software for this course should be installed on these machines.

Some of the assignments includes a networking part, requiring you to run your code on several machines.
This can be conveniently done using the machines mentioned above.
To be able to log into these machines you will need an account on the Unix system.

For more rapid testing, you may also run networking code from different ports on your local machine, i.e. `localhost`.

### Task: Sign up for Unix Account - Do It Now

Some of you may already have done this one:
You will need a Unix account to access machines in the Linux lab.
Get an account for UiS’ Unix system by following the instructions [here](http://user.ux.uis.no).
Be sure to read the section **Using the UNIX system**.

*It may take a while before you get access, but you can continue learning while you wait.*

## The Missing Semester of Your CS Education by MIT

Throughout this and other courses and as a software engineer, you will often need to use command-line tools to interact with computers.
Lack of knowledge of the available tools will lead to manually performing repetitive tasks or spending lots of time searching online for solutions.
For these reasons and more, we expect you to go through [The Missing Semester of Your CS Education](https://missing.csail.mit.edu/) from MIT (hereafter referred to as the Missing Semester).
You can read more about the motivation behind that course [here](https://missing.csail.mit.edu/about/).

Fun fact: One of the lecturers in this course is [Jon Gjengset](https://thesquareplanet.com).
He is a Norwegian PhD student at MIT and YouTuber, specializing in lengthy live-coding sessions about building stuff using the Rust programming language.

You should try to answer or at least understand the answers to the **Exercises** section at the end of each lecture.
Additionally we give a set of multiple choice questions below, which mostly correspond to lectures 1, 2, 4, 5, and 6.

### Additional Resources and Tips

- [UNIX Tutorial for Beginners](http://www.ee.surrey.ac.uk/Teaching/Unix/).
  Eight simple tutorials covering the basics of various Unix/Linux commands.
  You may use these as a reference if you struggle to answer some of the questions or want a more in-depth overview than that offered by the Missing Semester.

- [Unix/Linux Command Reference](https://files.fosswire.com/2007/08/fwunixref.pdf).
  A cheat sheet of several frequently used Unix/Linux commands.

- Remember that almost every Unix/Linux command has a manual page, or man page for short, which can be accessed with `man` command, e.g. `man ls` for the `ls` command.

- *Tip:* Use the `git help` command whenever in doubt about a Git command.
  It lets you read more about the functionality of each of Git's subcommands, e.g. by running `git help commit` for information about `git commit`, such as options, or `git help pull` for information about `git pull`.

- *Tip:* Navigating `man`, `less` and `git help` buffers: The buffers opened by the `man`, `less` and `git help` commands support vi(m)-like navigation.
  - You can move down by one line by pressing the `Down` arrow key or the `j` key, or up by one line by pressing the `Up` arrow key or the `k` key.
  - You can move up or down by one page by pressing the `PageUp` and `PageDown` keys.
    Alternatively you can press the `f` ("forward") or `b` ("backward") keys.
  - You can go to the start or end of the buffer by pressing the respective `Home` and `End` keys.
    Alternatively you can press the `g` or `G` keys for the same functionality.
    There are often examples at the end of man pages.
  - You can search for some text by pressing the `/` key.
    Press `n` to go to the next match, and `N` to go to the previous match.

### Task: Unix/Linux and Git Multiple-Choice Questions

Answer the questions inline in the markdown files, as explained in the heading of each file.

1. [Questions related to the Missing Semester](./missing_semester_questions.md) lectures 1, 2 and 4.
2. [Shell questions](shell_questions.md).
   Some of these commands may not be covered by the Missing Semester lectures.
   We recommend reading the relevant man pages and checking the other related resources mentioned above.
3. [Questions about Git](./git_questions.md) based on lecture 6 of the Missing Semester as well as some regularly used Git commands.
   *Hint:* Some of the questions may be heavily influenced by StackOverflow questions.

Note that, some commands behave differently on macOS and Linux, because they are based on different heritage.
Typically, macOS and Linux may sometimes use different flags for altering the behavior of a command.
We have made notes on these differences, where we are aware of them, but should you discover an incompatibility in these labs, please let us know.

Further, this lab was designed with the `bash` Unix shell, which is the default on Linux.
The default shell is `zsh` on macOS.
If you experience any issues related to running a different shell than `bash`, please try the same on Linux, and let us know.
To determine your shell, use the following command:

```console
echo $SHELL
```

## Remote Login with Secure SHell (SSH)

*Skip this part if you haven't got a Unix account password yet.*
*If so, come back and do it later, because you will need it to complete this lab.*

<!Here is a [SSH tutorial video](https://youtu.be/qik3HHZW6C0) illustrating the steps below (and a bit more).>

You can use `ssh` to log on to another machine without physically going to that
machine and login there. This makes it easy to run and test the example code
and your project later. To log onto a machine using ssh, open a terminal window
and type a command according to this template, and make sure to replace
username and hostname:

```console
ssh username@hostname
```

You first need to login to one of the jump hosts that are available for Internet access, for example:

```console
ssh nejm@ssh1.ux.uis.no 
ssh nejm@ssh2.ux.uis.no 
ssh nejm@ssh3.ux.uis.no 
ssh nejm@ssh4.ux.uis.no 
```
To ssh from outside the campus you will have to use two factor authentication as explained in [2FA UiS](https://foswiki.ux.uis.no/bin/view/Info/TwoFactorAuthentication).

Once you are in one of the jump hosts, you should be able to ssh into ```gorina1.ux.uis.no``` and perform your labs.


*The following may be skipped if you can login from the jump machines to the lab server without typing a password.*
*If so, your account was created with the appropriate ssh keys in your `authorized_keys` file.*
*(The following text is left in here for your information in case you want to configure your own machine's logins.)*

You can avoid having to type the password each time by generating a
public-private key-pair using the `ssh-keygen` command (see the man pages for
`ssh-keygen`). Type

```console
man ssh-keygen
```

and read the instructions. Then try running this command to generate your
key-pair; make sure that once asked to give a password, just press enter at the
password prompt. Once the key-pair have been generated, append the public-key
file (ends with .pub) to a file named `authorized_keys`.

If you have multiple keys in the latter file, make sure not to overwrite those
keys, and instead paste the new public-key at the end of your current file.
After having completed this process, try ssh to another machine and see whether
you have to type the password again.

Note that the security of this passphrase-less method of authenticating towards
a remote machine hinges on the protection of the private key file stored on
your client machine. Thus, it is actually recommended to create a key with a
passphrase, and instead use the `ssh-agent` command at startup, along with
`ssh-add` to add your key to this agent. Then, the `ssh`, `scp`, and other
ssh-based client commands can talk locally with the `ssh-agent`, and you as the
user only needs to type your password once. Please consult the `ssh-agent` and
`ssh-add` manual pages for additional details.

Another tip: If you are running from a laptop and wish to remain connected even
if you close the laptop-lid, you can check out the [mosh
command](http://mosh.mit.edu/).

### Remote Access to the Linux Machines

Due to UiS's firewall configuration you cannot access the machines in the Linux lab from outside UiS's network directly.
Information on how to do a remote login to the Linux machines from another network, e.g. from your home, can be found
[here](http://wiki.ux.uis.no/foswiki/Info/WebHome) and
[here](http://wiki.ux.uis.no/foswiki/Info/HvordanLoggeInnP%E5Unix-anlegget).

### Task: Logging Into the Linux Lab and Setting Up Password-less Logins

In this task you will log into the Linux Lab and set up an authorized key as described in the previous section.
Additionally, you will clone your `assignments` repository to the Linux lab and configure Git to use your SSH key for authentication.
To test these tasks, we have created an executable "token generator", which generates a unique token for each student indicating the number of checks that were successful, which will be checked by QuickFeed after being pushed to GitHub.

#### Generating a Key Pair on the Linux Lab

After having logged into the Linux Lab and set up SSH keys that are present in `$HOME/.ssh/authorized_keys` on the server, you should set up an additional key pair on the server.
Use `ssh-keygen` to generate a key pair on the Linux lab as described above.
The key pair should not require a passphrase, otherwise the token generator will fail.

NOTE: If you prefer to have a key pair with a passphrase, you can replace the key pair with a new one that is protected with a passphrase, after having passed the QuickFeed tests.

#### Setting up SSH Authentication on GitHub

After setting up a key pair on the server, follow the instructions for [Connecting to GitHub with SSH](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh) including the step [Testing your SSH connection](https://docs.github.com/en/github/authenticating-to-github/testing-your-ssh-connection).
Note that these guides provide a slight variation for Mac, Windows and Linux.
You can select your OS via a tab near the top of each article, and for operations on the Linux labs you should follow the instructions in the `Linux` tab.

#### Cloning the Repository to the Linux Lab

Next you will clone the `assignments` repository to the Linux lab servers, following the steps below.
As part of this process you must also generate a token, using the provided generator (as we explain below).
We recommend that you follow the steps below exactly.

In the following description, we refer to your GitHub user name as `username`.
*You must replace `username` with your actual GitHub user name.*
Perform the following steps to clone the repository:

```console
# change directory to $HOME
cd
```

In your browser, find the green `Code` menu on your `username-labs` repository on GitHub, select `Use SSH` and copy the URL in the `Clone with SSH` tab.
Below we use `git@github.com:dat320-2020/username-labs.git` as an example.

```console
# clone the Git repo with SSH into the $HOME/assignments directory
git clone git@github.com:dat320-2020/username-labs.git assignments
cd assignments
git remote add course-assignments git@github.com:dat320-2020/assignments.git
```

If you happened to clone the assignments repository using the HTTPS method by mistake, you can follow our [guide to configure Git to use SSH authentication](https://github.com/dat320-2020/course-info/blob/master/github-ssh.md) to fix it.

#### Generating a Token

After having completed the steps above you can generate your token.
The token generator will perform the following checks:

- Is the student logged in to one of the Linux lab computers?
- Are there public keys in `$HOME/.ssh/authorized_keys` to enable password-less login to the Unix lab?
- Is the assignments repository cloned to the Linux lab?
- Is SSH authentication used for Git, and can Git operations be performed without having to enter the password?

Perform the following steps to generate your token.
Navigate to the `lab1` directory:

```console
cd $HOME/assignments/lab1
```

Run the token generator:

```console
./generate_token
```

For each successful test you should see something like this:

```console
[v] Check passed.
```

Similarly failed checks provide some output briefly explaining what went wrong.

If the token was generated successfully you should see the following message:

```console
Token successfully generated and stored in $HOME/assignments/lab1/token.
You need to commit and push this directory so that QuickFeed can process it.
```

Navigate to this directory, then add, commit and push the token.

```console
cd $HOME/assignments/lab1/token
git add .
git commit -m "Submitted token"
git push
```

The test results should now show up in QuickFeed within a few minutes.

## Submitting to QuickFeed

Read our [lab submission guide](https://github.com/dat320-2020/course-info/blob/master/lab-submission.md) for more detailed instructions on how to submit your assignments to be evaluated by QuickFeed.
