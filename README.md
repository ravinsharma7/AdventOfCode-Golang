# AdventOfCode-Golang
AoC 2024

# Introduction
- The approach is as an idiot I can breakdown the concept part by part until chatgpt guides me to the correct answer.
- Compared to prompt engineering, where you create "fancy structured text" mindlessly to make the output follow a particular pattern or topic.
- I'm not going to do any fancy structuring of my text to chatgpt. Just applying "Socratic Method In Action" with the AI.
- My goal is not jump directly to a correct using chatGPT, but dance towards the correct answer. Is that what they called as RAG if you brute force the dance?(Ah whatever)
- Using golang code for 2 reasons. First golang has barely changed syntax over the decade. Second golang has probably one the best backward compatibility among the popular languages.(and only like 25 keywords, i think). Given this two factors, I'm guessing chatgpt should have a better result.
- Model used 4o

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
- Part 1 First try: Wrong. Correct third try:
  - I think it got confused the delimiter I was using to seperate the question and input, it generated a code that ran infinitely long(I suppose).
  - It forget about go code and generated python. Added "Generate go code" at the end. Delimeter using the new delimeter.
  - Same chat window
window.
- Part 2 First try: Wrong. Correct second try:
  - I think the ..., and # messed with the delimited and it got confused somehow. So I need to be aware of stop words and overload puntuations symbols.
  - I had to reformat the question to get it right.
  - Same chat window
