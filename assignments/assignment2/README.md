# Assignment: TCP Congestion Control and Bufferbloat

This is not a group assignment. You are not allowed to copy or look at code
from other students. However, you are welcome to discuss the assignments with
other students without sharing code.

## Getting Started

Let's get started!

## For existing VirtualBox and Vagrant users

If you already installed vagrant and VirtualBox, go to the CS513 folder. You can likely skip most of PART A below.
```
$ git clone https://github.com/bcarpent/CS513.git
$ cd assignments
```
 Pull the latest update from Github.
```
$ git pull
```

Now you need to run vagrant up and ssh again, as the Vagrantfile has been updated for Jupyter support.
```
$ vagrant up
$ vagrant ssh
```

## PART A: Set Up Virtual Machine

The first part of this assignment is to set up the virtual machine (VM) you
will use for the rest of the course. This will make it easy to install all
dependencies for the programming assignments, saving you the tedium of
installing individual packages and ensuring your development environment is
correct.

### Step 1: Install Vagrant

Vagrant is a tool for automatically configuring a VM using instructions given
in a single "Vagrantfile."

**macOS & Windows:** You need to install Vagrant using the correct download
link for your computer here: https://www.vagrantup.com/downloads.html.

**Windows only**: You will be asked to restart your computer at the end of the
installation. Click Yes to do so right away, or restart manually later, but
don't forget to do so or Vagrant will not work!

**Linux:** First, make sure your package installer is up to date by running the
command `sudo apt-get update`. To install Vagrant, you must have the "Universe"
repository on your computer; run `sudo apt-add-repository universe` to add it.
Finally, run `sudo apt-get install vagrant` to install vagrant.

### Step 2: Install VirtualBox

VirtualBox is a VM provider (hypervisor).

**macOS & Windows:** You need to install VirtualBox using the correct download
link for your computer here: https://www.virtualbox.org/wiki/Downloads. The
links are under the heading "VirtualBox 6.x.x platform packages."

**For macOS Big Sur:** After installation, you need to go to 
`System Preferences  > Security & Privacy` 
and allow system software updates from Oracle (there might be a prompt about 
this during installation). Then, restart your mac.

**Windows only:** Use all the default installation settings, but you can
uncheck the "Start Oracle VirtualBox 6.x.x after installation" checkbox.

**Linux:** Run the command `sudo apt-get install virtualbox`. If `virtualbox` is 
not already added to your package repository (if apt-get prompts "virtualbox has 
no installation candidate"), please follow the instructions 
here: https://www.virtualbox.org/wiki/Linux_Downloads

**Note:** This will also install the VirtualBox application on your computer,
but you should never need to run it, though it may be helpful (see Step 6).

### Step 3: Install Git (and SSH-capable terminal on Windows)

Git is a distributed version control system.

**macOS & Windows:** You need to install Git using the correct download link
for your computer here: https://git-scm.com/downloads.

**macOS only:** Once you have opened the .dmg installation file, you will see a
Finder window including a .pkg file, which is the installer. Opening this
normally may give you a prompt saying it can't be opened because it is from an
unidentified developer. To override this protection, instead right-click on
thet .pkg file and select "Open". This will show a prompt asking you if you are
sure you want to open it. Select "Yes". This will take you to the
(straightforward) installation.

**Windows only:** You will be given many options to choose from during the
installation; using all the defaults will be sufficient for this course (you
can uncheck "View release notes" at the end). The installation includes an
SSH-capable Bash terminal usually located at `C:\Program
Files\Git\bin\bash.exe`. You should use this as your terminal in this class,
unless you prefer another SSH-capable terminal (the command prompt will not
work). Feel free to create a shortcut to it; copying and pasting the executable
somewhere else will not work, however.

**Linux:** `sudo apt-get install git`.

### Step 4: Install X Server

You will need an X Server to input commands to the virtual machine.

**macOS:** Install [XQuartz](https://www.xquartz.org/). You will need to log
out and log back in to complete the installation (as mentioned by the prompt at
the end).

**Windows:** Install
[Xming](https://sourceforge.net/projects/xming/files/Xming/6.9.0.31/Xming-6-9-0-31-setup.exe/download).
Use default options and uncheck "Launch Xming" at the end.

**Linux:** The X server is pre-installed!

### Step 5: Clone course Git repository

Open your terminal (use the one mentioned in step 3 if using Windows) and `cd`
to wherever you want to keep files for this course on your computer.  

Run `git clone https://github.com/bcarpent/CS513` to
download the course files from GitHub.

`cd CS513/assignments` to enter the course assignment directory.

### Step 6: Provision virtual machine using Vagrant

From the `assignments` directory you just entered, run the command  `vagrant
up` to start the VM and provision it according to the Vagrantfile. You will
likely have to wait several minutes. You may see warnings/errors in red, such
as "default: dpkg-preconfigure: unable to re-open stdin: No such file or 
directory", but you shouldn't worry about them.

**Note 1**: The following commands will allow you to stop the VM at any point
(such as when you are done working on an assignment for the day):
* `vagrant suspend` will save the state of the VM and stop it.
* `vagrant halt` will gracefully shutdown the VM operating system and power
  down the VM.
* `vagrant destroy` will remove all traces of the VM from your system. If you
  have important files saved on the VM (like your assignment solutions) **DO
  NOT** use this command.

Additionally, the command `vagrant status` will allow you to check the status
of your machine in case you are unsure (e.g. running, powered off, saved...).
You must be in some subdirectory of the directory containing the Vagrantfile to
use any of the commands above, otherwise Vagrant will not know which VM you are
referring to.

**Note 2**: The VirtualBox application that was installed in Step 2 provides a
visual interface as an alternative to these commands, where you can see the
status of your VM and power it on/off or save its state. It is not recommended
to use it, however, since it is not integrated with Vagrant, and typing
commands should be no slower. It is also not an alternative to the initial
`vagrant up` since this creates the VM.

### Step 7: Test SSH to VPN

Run `vagrant ssh` from your terminal. This is the command you will use every
time you want to access the VM. If it works, your terminal prompt will change
to `vagrant@cs513:~$`. All further commands will execute on the VM. You can
then run `cd /vagrant` to get to the course directory that's shared between
your regular OS and the VM.

Vagrant is especially useful because of this shared directory structure.  You
don't need to copy files to and from the VM. Any file or directory in the
`assignments` directory where the `Vagrantfile` is located is automatically
shared between your computer and the virtual machine. This means you can use
your IDE of choice from outside the VM to write your code (but will still have
to build and run within the VM).

The command `logout` will stop the SSH connection at any point.

### Extra Note for Windows users

Line endings are symbolized differently in DOS (Windows) and Unix
(Linux/MacOS). In the former, they are represented by a carriage return and
line feed (CRLF, or "\r\n"), and in the latter, just a line feed (LF, or "\n").
Given that you ran `git pull` from Windows, git detects your operating system
and adds carriage returns to files when downloading. This can lead to parsing
problems within the VM, which runs Ubuntu (Unix). Fortunately, this only seems
to affect the shell scripts (\*.sh files) we wrote for testing. The
`Vagrantfile` is set to automatically convert all files back to Unix format, so
**you shouldn't have to worry about this**. **However**, if you want to
write/edit shell scripts to help yourself with testing, or if you encounter
this problem with some other type of file, use the pre-installed program
`dos2unix`. Run `dos2unix [file]` to convert it to Unix format (before
editing/running in VM), and run `unix2dos [file]` to convert it to DOS format
(before editing on Windows). A good hint that you need to do this when running
from the VM is some error message involving `^M` (carriage return). A good hint
you need to do this when editing on Windows is the lack of new lines. Remember,
doing this should only be necessary if you want to edit shell scripts.

## PART B: Starting up the Jupyter Notebook Server

In the VM, run the command `sudo jupyter notebook &`. This will
start a new Jupyter notebook server in the background. Even though it is
running in the background, it will sometimes print informative messages to the
terminal. You can press Enter each time you get a message to get the shell
prompt back. To shut down the notebook, run `fg` then press Control-C twice
(once to get the confirmation message, another time to skip confirmation).

While the notebook is running, on your host machine, open up your browser and
type `localhost:8888` in the address bar. This should take you to the Jupyter
notebook file selection window.  Juypter notebook is actually running on port
8888 on your vagrant VM, but you can access it through your host machine
browser because the port is being forwarded between the VM and the host
machine.  

In the file selection window, enter the `assignment2` directory and then open
`Assignment2_Notebook.ipynb`. This will open a notebook with the instructions
for the rest of the assignment.  Work through this notebook from top to bottom
and complete the sections marked "TODO."

**Remember to "Save and Checkpoint" (from the "File" menu) before you leave the
notebook or close your tab.**  

## Jupyter Notebook

Jupyter Notebook (formerly called iPython Notebook) is a browser-based IDE with
a cell-based editor.

Every cell in a notebook can contain either code or text ("Markdown"). Begin
editing a cell by double-clicking it. You can execute the code in a cell (or
typeset the text) by pressing `shift-enter` with the cell selected.  Global
variables and functions are retained across cells. Save your work with the
"Save and Checkpoint" option in the "File" menu. If your code hangs, you can
interrupt it with the "Interrupt" option in the "Kernel" menu.  You can also
clear all variables and reset the environment with the "Restart" option in the
"Kernel" menu.

The "Help" menu contains many additional resources about Jupyter notebooks
(including a user interface tour, useful keyboard shortcuts, and links to
tutorials).

## Submission

Submit your completed `Assignment2_Notebook.ipynb` file on Canvas along with screenshots of all graphs presented.

Remember to put your name and id in the marked location at the top of the file.
