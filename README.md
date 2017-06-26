# secintro-samples
Code samples for the course Introduction to Information Security 

This repository contains Python and Golang sample code for the course. 
The code samples are intended to work as a skeleton for you own code. It can be daunting to start from scratch, particularly without much programming experience. 

Golang
------
The Go samples can be run in the Go playground at https://play.golang.org or compiled to a binary on a host that has Golang installed. You can get the Golang compiler from: https://golang.org/. Also the linux-ssh.cc.tut.fi SSH-machine has Golang installed. You should apply for an account through https://www.tut.fi/omatunnus (not a phishing link, trust me) 

Building and Running code with Go 
=================================
go build ex-1-4.go

./ex-1-4 

Go is very strict about a lot of things, such as unused variables. However, once you get it to compile, it is ready to ship for sure. 

Python
------

Because some of the students are more familiar with Python (or have personal issues with Golang) some Python examples are provided as well. If the exercise depends heavily on some library, or the dude making the exercise work has found good copypaste sources only in the other language, either python sample or Golang sample could be omitted. 

Building Python
===============
Python is interpreted, so only needs to be run with either 

python ex-1-4.py 

or 

./ex-1-4.py 

(provided that the python source is marked with executable bit (chmod +x ex-1-4.py)) 

Using Git 
=========
Git is a version control system. Using a version control system shows love and respect to you precious source code files. 
You can use this repository in two ways. 
1) Clone this repository on linux-ssh or some other host directly without forking. No upstream push with this origin. 

git clone https://github.com/joonakannisto/secintro-samples.git secintro 

If new code appears in the repository, you just run git pull to be in sync again. The course staff  should not be alter old files to prevent conflicts from happening during the course. 

2) If you have a Github account you may fork this repository and clone it from your own user page, (Not Recommended) 

If you clone the repository and push your local changes upstream, you need to make the repository private, or push the changes only after the week's session. You can also make a completely detached repo by changing the repository origin to an empty repository not linked to the original (security through obscurity). 
