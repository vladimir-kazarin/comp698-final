# comp698-final

Why use git?
* It’s a distributed version control system which makes developers more independent from the central repository and provides better * reliability.
* It has fast-growing community of developers and extensive documentation.
* Has many GUI clients and supported by GitHub which can convenient for a small group of developers.
* Provides efficient branching and inexpensive merges.
* Git creates a local copy of the repository which allows developers to work offline.

**Why write an application in golang?**
* It's compiled into a single binary file
Compiling all dependencies into a one, single binary makes it easier to deploy the application. Otherwise you would have to install all dependencies on your server (like Python) for your app to run.
* Golang has static type system
In my opinion, languages with static type system for variables are much more safe and reliable because it prevents your app from wrong variable usage errors.
* Native tools for web development built-in
Golang has a lot of native tools for web development like http, json and html templating. 
* Golang is fast

**What purpose does Travis CI serve?**
Travis CI is a continuous integration service that allows testing projects on GitHub. Travis CI can be configured to make sure your applications builds and passes required tests.  

**What is the purpose of terraform?**
Terraform is a software system that allows configuring and building an infrastructure of computer systems through high-level code files. Terraform makes it easy to set up new virtual servers and deploy docker containers. 

**Why use virtualized resources?**
Virtual resources create a level of abstraction from hardware level. It saves a lot of time configuring new servers and environemnts by using pre-configured images. It also provides an easy way to mirror exisiting systems.

**Include a diagram of the architecture from laptop to GCP**
![Architecture diagram](/static/diagram.png)

**Why use bash commands vs clicking through UI?**
Bash commands can be harder to use, but they are highly efficient when it comes to repetitive actions. All bash commands that you use can be saved in a file and at end you have a script that will do the same operations without human effort. 
Also, sometimes you have access to more features through bash commands than GUI.

**Why docker vs installing application directly on the host?**
When installing an application directly, there is a chance that you will run into configuration issues when two applications use different versions of the same dependency. Docker containers on the other hand, contain all dependencies specifically for the application inside it and every container is isolated from others.

**What protections are there against accidentally introducing bad code into production?**
To decrease errors in production code the following protections are recommeneded to be used:
* Unit tests - tests against all functions in the code to check if the ystill return the desired output
* Separating development in branches, never direct commits to Master Protecting Master branch from any commits other than through a pull request.
* Deploying to staging first and testing the functionality there.
* Code reviews
