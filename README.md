# DLL-call-in-Golang

How to call Windows DLL using LoadLibrary in Golang

First, of course you need to create a simple DLL.  Follow the steps by PauloGP's github :
https://gist.github.com/paulogp/1345006

I also put the .zip file that contains the project code for my DLL.  You just need to open the project and build it to create the .DLL.  
For simplicity, i also attach together the 64-bit square.dll which created earlier.

Make sure you build for 64-bit if your Golang is running in 64-bit machine.  
Once you have the .DLL, copy it and paste under the same folder where your Golang code will reside.

In this example, i create a simple Golang function in which it will load the DLL and finally call the function
inside the DLL.  Here, the function name is "dll_math_example".  It is as simple function where it takes 
4 input parameters of type int to be added together.  

In line 20, u can see i used "syscall.Syscall9(uintptr(proc), 0, 2, 2, 2, 2, 0, 0, 0, 0, 0)".  This will invoke 
the function with 4 inputs. So, the operation will be 2+2+2+2, yields a result of 8.


