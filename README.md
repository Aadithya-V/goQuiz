# goQuiz
A timed concurrent command line quiz application.

Execute the goQuiz.exe file from the release in a terminal. Currently compiled only for windows(x86). 
Use the source code provided to compile to any target architecture.

Use the flag -h to view the help information of the flags.

The problem set is provided in a .csv file. 
Use -csv string to provide the input file:
        a csv file in the form of question, answer (default "problems.csv").

The timer for the quiz runs concurrently. 
Use  -limit int to change the time limit:
        the time limit for the quiz in seconds (default 30).
