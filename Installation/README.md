# How to Install Go on Windows?

Before, we start with the process of Installing Golang on our System. We must have first-hand knowledge of What the Go Language is and what it actually does? Go is an open-source and statically typed programming language developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009. It is also known as the Golang and it supports the procedural programming language.

Golang programs can be written on any plain text editor like notepad, notepad++ or anything of that sort. One can also use an online IDE for writing Golang codes or can even install one on their system to make it more feasible to write these codes. Using an IDE makes it easier to write Golang codes because IDEs provide a lot of features like intuitive code editor, debugger, compiler, etc.

## How to check the Preinstalled Go Language Version?

Before we begin with the installation of Go, it is good to check if it might be already installed on your System. To check if your device is preinstalled with Golang or not, just go to the Command line(For Windows, search for cmd in the Run dialog(Windows + R).

Now run the following command:

```
go version
```

If Golang is already installed, it will generate a message with all the details of the Golang’s version available, otherwise, if Golang is not installed then an error will arise stating Bad command or file name

## Downloading and Installing Go

Before starting with the installation process, you need to download it. For that, all versions of Go for Windows are available on [golang.org.](https://go.dev/dl/)

<p align = "center">
<img src = "https://user-images.githubusercontent.com/101946115/209450850-a4c69a1d-f511-4d0e-bb94-00d26dca21cc.png" height = 400 width = 600 />
  </p>

Download the Golang according to your system architecture and follow the further instructions for the installation of Golang.

<b>Step 1:</b> After downloading, unzip the downloaded archive file. After unzipping you will get a folder named go in the current directory.

<p align = "center">
<img src= "https://user-images.githubusercontent.com/101946115/209450903-d8553d10-f811-429b-be69-f18ebe61e9a9.png" />
</p>

<b>Step 2:</b> Now copy and paste the extracted folder wherever you want to install this. Here we are installing in C drive.

<b>Step 3:</b> Now set the environment variables. Right click on My PC and select Properties. Choose the Advanced System Settings from the left side and click on Environment Variables as shown in the below screenshots.

<p align = "center">
  <img src = "https://user-images.githubusercontent.com/101946115/209450941-fc6b95e8-2178-4061-ac5c-c43c058629f3.png" height = 400 width = 600/>
<br>
  <br>
  <br>
   <img src = "https://user-images.githubusercontent.com/101946115/209450944-a6c721c6-b4f5-4fa0-b7a6-5af203036710.png" height = 400 width = 600/>
</p>

<b>Step 4:</b> Click on Path from the system variables and then click Edit. Then Click New and then add the Path with bin directory where you have pasted the Go folder. Here we are editing the path C:\go\bin and click Ok as shown in the below screenshots.

<p align = "center">
<img src = "https://user-images.githubusercontent.com/101946115/209450983-6769b9d1-0082-4194-80d5-1896abd96e12.png" />
<br>
  <br>
  <br>
<img src = "https://user-images.githubusercontent.com/101946115/209450985-d37d9794-3093-49a3-8716-62bfa63ccf0e.png" />
  
  </p>


<b>Step 5:</b> Now create a new user variable which tells Go command where Golang libraries are present. For that click on New on User Variables as shown in the below screenshots.

<p align = "center">
  <img src = "https://user-images.githubusercontent.com/101946115/209451010-1482a206-183d-40b3-829d-16fd6a7029f8.png" height = 400 width = 600/>
</p>

Now fill the Variable name as GOROOT and Variable value is the path of your Golang folder. So here Variable Value is C:\go\. After Filling click OK.

<p align = "center">
  <img src = "https://user-images.githubusercontent.com/101946115/209451020-303445b1-d253-4f2d-8bf9-62bf955c2037.png" />
</p>

After that Click Ok on Environment Variables and your setup is completed. Now Let’s check the Golang version by using the command go version on command prompt.

<p align = "center">
  <img src = "https://user-images.githubusercontent.com/101946115/209451037-70ab26b5-d211-4390-8182-0a6b6e96c005.png" />
</p>

After completing the installation process, any IDE or text editor can be used to write Golang Codes and Run them on the IDE or the Command prompt with the use of command:

```
go run filename.go
```





