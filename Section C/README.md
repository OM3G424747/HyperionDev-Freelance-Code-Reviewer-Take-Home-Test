<H2>Section C - Option 4: International Standard Book Numbers</H2>

I decided to use <strong>GoLang</strong> for section C and chose to create Option 4: International Standard Book Numbers.
GoLang is similar to most C derivatives, including Java. 
However, my primary motivation for choosing GoLang was to help prove my versatility. Since I already completed the Software Engineering Bootcamp with HyperionDev, it's already obvious that I know how to code in Java, which is why I wanted to use GoLang instead.
All modules used are standard, however, go.mod files will need to be enabled


<H3>worst-case space complexity</H3>

The worst case space complexity would be calculated as: <strong>O(N)</strong>

This is because the complexity would vary depending on if an ISBN-10 or ISBN-13 is passed.
If an ISBN-13 is passed it would result in 3 additional values being stored


<H3>Build, test and run solution</H3>

To start, please follow the steps:

1) - Please download and install GoLang for your OS <a href="https://go.dev/dl/">HERE</a>
2) - If there's an option to add GoLang to your PATH environment variable, please ensure it's selected!
3) - After the installation has been completed, open a terminal and enter "go". If you are presented with a list of commands for go, it means the installation was successful.
4) - I've included a <strong>mod.bat</strong> file to automatically enable .mod files, please double-click the <strong>mod.bat</strong> file to run it and automatically enable mod files on windows. Alternatively please run the following command in your terminal if the bat file is not compatible with your OS: <strong>go env -w GO111MODULE=on</strong> 
5) - Next, in the terminal, navigate to the directory where the contents of the folder for this section are stored.
6) - In the terminal, run "go build main.go". This will compile the contents of this project and create a file you can run on your OS. For example, "main.exe" for Windows.
7) - Run the newly created file to use the coded solution for ISBM numbers

The test suit is stored in <strong>main_test.go</strong>. All the functions that return values are tested except for the main function and the "isbnCheck" function.
This is because both the main and "isbnCheck" function make use of all the other functions being tested in the test suit and no other functions are dependent on them to return a correct value.
To test the solution and run the test suit, please follow these steps:

1) - In the terminal, navigate to the directory where the contents of the folder for this section are stored.
2) - In the terminal, run "go test". If successful, it should display "PASS"


<H3>Using the ISBN solution</H3>

Simply enter an ISBN in the terminal window that was opened.
ISBN-13 codes that are entered will display either "Valid" or "Invalid".
While valid ISBN-10 numbers will display "Valid" along with an ISBN-13 version of the same number.
Invalid ISBN-10 numbers will only display "Invalid".

The solution is also able to handle white space and hyphens.
For example:
"0-9752298-0-X" and " 0    9752298  0 X ", will be handled as the same number.

