# AdventOfCode-Golang
AoC 2024 : https://adventofcode.com/2024

# Introduction
- The approach is as an idiot I can breakdown the concept part by part until chatgpt guides me to the correct answer.
- Compared to prompt engineering, where you create "fancy structured text" mindlessly to make the output follow a particular pattern or topic.
- I'm not going to do any fancy structuring of my text to chatgpt. Just applying "Socratic Method In Action" with the AI.
  > Scratch that. I have not choice to reformat the questions sometime. Because the input might clash with my delimiter between the question and input when I send text to chatgpt. But only this.
- My goal is not jump directly to a correct using chatGPT, but dance towards the correct answer. Is that what they called as RAG if you brute force the dance?(Ah whatever)
- Using golang code for 2 reasons. First golang has barely changed syntax over the decade. Second golang has probably one the best backward compatibility among the popular languages.(and only like 25 keywords, i think). Given this two factors, I'm guessing chatgpt should have a better result.
- Model used 4o or o1

# Results
- Mad that ChatGPT can generate the code for the correct answer. I'm commiting the right answer from chatgpt for reference later. See how it goes.

### Day 1
- Part 1 First try: Correct
  - Direct Paste of Question+Input. No additional text aside appending "Generate go code" 
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 2
- Part 1 First try: Correct  
  - Direct Paste of Question+Input. No additional text. But in the same chat window.
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 3
- Part 1 First try: Correct  
  - Direct Paste of Question+Input. No additional text. But in the same chat window.
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 4
- Part 1 First try: Correct  
  - Direct Paste of Question+Input. No additional text. But in the same chat window.
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 5
- Part 1 First try: Correct  
  - Direct Paste of Question+Input. I had to split the rules and updates inputs by myself since the resulting code assumed it wasy already split. It starting to lose context from Day 1, no additional text aside appending "Generate go code". But in the same chat window.
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 6
- Part 1 First try Wrong. Correct third try:
  - I think it got confused that the delimiter I was using was part of the question and input somehow, it generated a code that ran infinitely long(I suppose).
  - It forgot about go code and generated python. Added "Generate go code" at the end. Delimeter using the new delimeter.
    ```
    INPUT:
    
    #...#...
      
    ME:
    Generate go code
    ```
  - Same chat window
window.
- Part 2 First try: Wrong. Correct second try:
  - I think the `...`, and `#` messed with the delimiter and it got confused somehow. So I need to be aware of stop words and overloaded puntuations symbols as delimiter and input.
  - I had to reformat the question to get it right.
  - Same chat window

### Day 7
- Part 1 First try: Correct  
  - Direct Paste of Question+Input. Added text "Generate go code". But in the same chat window.
- Part 2 First try: Correct  
  - Direct Paste of Question without the same Input. No additional text. But in the same chat window.

### Day 8
- Part 1 first try: Correct(o1) / Wrong(4o)
  - I reformat the question because the delimiter might clash.
  - 4o gets the code wrong. Tried like 3-4 times.
  - o1 got it right the first time. Kinda, it generated code with unused variable in it, this trigger compiler error, so I had to manually change into _
- Part 2 first try: Correct(o1)
  - Only tried with o1. 4o didnt bother
  - Got right the first time. Kinda, it generated code with unused variable in it, this trigger compiler error, so I had to manually change into _
 
### Day 9
- Part 1 first try: Correct(o1)
  - Only tried with o1. 4o didnt bother
  - I reformat the question because the delimiter might clash.
  - o1 got it right the first time. 
- Part 2 first try: Correct(o1)
  - Only tried with o1. 4o didnt bother
  - Got right the first time.

### Day 10
- Part 1 first try: Correct(o1)
  - Only tried with o1-mini. 
  - I reformat the question because the delimiter might clash.
- Part 2 first try: Correct(o1)
  - Only tried with o1-mini.
  - Got right the first time.
