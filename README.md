# Learning Go
This is the repository for the LinkedIn Learning course Learning Go. The full course is available from [LinkedIn Learning][lil-course-url].

![Learning Go][lil-thumbnail-url] 
What is Go? Go is a next-generation, open-source programming language created by Google for building systems, web, and other applications. This course is designed to help developers get started with Go, covering its core language elements and syntax. David Gassner introduces tools and skills used in a Go workflow—including Go Playground, an online tool that takes Go development off the desktop. He also covers basic programming tasks: managing values, using math operators, storing values as complex types, and managing program flow. Plus, learn how to create reusable Go code, read and write files, and make simple web requests.

## Learning Objectives
- Installing Go development tools and Visual Studio Code
- Exploring variables, constants, and types
- Storing ordered and unordered values
- Grouping related values in structs
- Programming conditional logic and loops
- Defining and calling functions
- Handling errors
- Working with files

## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go development tools](https://golang.org/dl/)
	- [Visual Studio Code](https://code.visualstudio.com/download)
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. As you work through the course, commit your changes in each branch to your local repository before checking out the next branch. Optionally, you can fork the exercise files repository to your own GitHub account and commit your changes to the forked version. A demonstration of how to work with GitHub in Visual Studio Code is available in [this video][E&P: include url for video 01_08].

### Instructor

**David Gassner**

_Author of 60+ video-based training courses for software developers_

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/david-gassner?u=104).

[lil-course-url]: https://www.linkedin.com/learning/learning-go-8399317
[lil-thumbnail-url]: https://cdn.lynda.com/course/2875237/2875237-1616177158000-16x9.jpg
