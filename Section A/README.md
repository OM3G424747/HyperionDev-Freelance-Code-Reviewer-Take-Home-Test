<H2>Section A - Option 1: Python Task</H2>

<H1>Compulsory Task 1 Review</H1>

<H2>Correctness</H2>
It seems several errors are preventing your code from running.
I'm going to cover each of them and suggest areas of improvement, but it would be a good idea to always test your code before submitting it and make sure the most recent working version is submitted.

<H4><ins>Indentation</ins></H4>
It seems your indentation is incorrect from lines 2 to 6, and then lines 8 and 10.
When running python code it's always important to ensure indentation is correct.
The easiest way to do this is by pressing "TAB" instead of pressing the spacebar multiple times.
If you are coding on a text editor like Nano on Linux, please make sure you configure it to add 4 spaces instead of 2.
If you're using Visual Studio Code, I would recommend using an extension like "<strong><a href="https://marketplace.visualstudio.com/items?itemName=oderwat.indent-rainbow">indent-rainbow</a></strong>", which will give you a clear visual indication of any incorrect indentation.

Your function declaration on line 2 should be indented by only 1 tab space, while the "results" variable below it should be indented by 2 tab spaces.
Remember the indentation in python helps indicate blocks of code, this means in this instance that the block of code belonging to your function would need to include the same level of indentation as the function declaration along with 1 additional level indentation.
So whenever you declare a class, function, loop, or condition, you'll need to include the indentation of the block of code that forms a part of.
Then you'll also need to indent the code that belongs to the declaration with an additional space.
Here's an example of a function that takes a string and then prints out each letter in a "for loop":

```python
def my_funct(some_string):
    # for loop indented by 1 TAB space to indicate it's a part of the function
    for letter in some_string:
        # print statement is indented 2 TAB spaces to indicate it's a part of the "for loop"
        print(letter)
```

As with my example, the print statement needs to include the indentation of the "for loop" (to show it's a part of the function), along with an additional indentation to indicate it forms a part of the "for loop".
Giving the print statement in my example the same level of indentation as the function declaration would mean it's outside of the function.

For this same reason, please remember to give your "return" statement the same indentation as your for loop in your code.
A return statement will exit your function, so if it's placed inside your for loop, it will result in your function exiting as soon as it reaches your return statement, which could mean that it could exit after just one iteration.

<H4><ins>Missing Variables</ins></H4>
It seems like you might have forgotten to pass "i" to the sorted function on line 5.


<H2>Efficiency</H2>
While your current submission can't be run, after correcting all the errors, it seems like you'll have a very simple yet elegant solution!
I can see what you were trying to accomplish, and you're on the right track here.

I like how you treat each string as a list of chars, by sorting it and using it as a key to creating a new list for any other string that matches it when sorted.
This shows you have a clear understanding of how strings work and you know how to manipulate them to accomplish your tasks!
Keep it up!


<H2>Style</H2>
When writing code in python, it's a good idea to use "<em>snake case</em>" to adhere to PEP8 standards when declaring variables and functions.

For example, instead of using this on line 3:
```python
def groupAnagrams(self, strs):
```

It would be better to use this instead:
```python
def group_anagrams(self, strs):
```

You can use tools like Flake8 to check your code and even enable linting for it.
Here's a brief guide on how to do this for Visual Studio Code:
<a href="https://code.visualstudio.com/docs/python/linting">Linting Python in Visual Studio Code</a>

This would also mean including linebreaks after your return statement on line 10.
While this won't affect if your code will run or not, it will help improve readability.
When you start working on larger projects with other developers, readability is vital, I highly recommend you try to adhere to PEP8 as best as possible.


<H2>Documentation</H2>
While this is a smaller project, it would still be a good idea to include a comment or two to indicate what actions you expect from certain lines of code.
Similarly, when declaring functions, it would be a good idea to make use of doc strings.
This will just help you get into the habit of documenting your code for when you do work on larger projects or projects with other developers.
Here is a guide on doc strings and how to declare them in python <a href="https://pandas.pydata.org/docs/development/contributing_docstring.html#:~:text=A%20Python%20docstring%20is%20a,html)%20documentation%20automatically%20from%20docstrings.">About docstrings and standards</a>

Similar to PEP8, this will also help make it easier for others to work and code alongside you.
Even if you're working alone on a project, these could still help you out tremendously in the long run.
You might know what your code is supposed to do now, but after working on the same project for a month or more, it could get difficult to remember what specific functions or blocks of code are supposed to do.
Adding comments and doc strings will help remind you of what you were trying to accomplish when you wrote those blocks and similarly if you're working with someone, it could help them understand what your code is supposed to do.


<H2>Outcome</H2>
While you will need to fix a few bugs and resubmit your code, I can see you've got the right idea and I know you'll be able to complete this assignment without any issues.
The top issue to focus on would be correct indentation and then just ensuring all variables are passed correctly where needed.
Please be sure to run your code before submitting it, to ensure it performs as expected.
You're doing well, keep it up, I'm looking forward to seeing more from you!

