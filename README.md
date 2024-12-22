# advent-of-code-2024
Advent of Code 2024

After Day 13 I decided to upload everything to GitHub. I am using Go which I want to learn with the intent to write Applications 
with it. Go is really nice.! I have literally never used it before, just did the [Tour of Go](https://go.dev/tour/welcome/1).
After those 13 days I have to say that Go is one of my favorite languages! It feels like I have already used it for many years. 
It is really refreshing

I did not measure my time up until now. Most of them took me 30-120 Minutes. Most of the time I spent on bug fixing and googling some Go Libraries. I really had no issues besides: 
- **Day 6 Part 2:** I had issues defining a loop. Then I had some issues with brute forcing... It took way too long. I then narrowed down the 
search space → Only checking the positions to place obstacles which the guard could hit
- **Day 11 Part 2:** Of course brute forcing does not work. After some thinking I got it. From the beginning I was thinking of some way to generalize it. In
the end I found a pretty neat solution myself without googling
- **Day 12 Part 2:** Took me some time to define Corners, though I quickly figured out that the number of edges in a Polygon is the number of corners. Part 1 was really quick. Coding just took some time 

## Diary

### Day 14
**Part 1:** It was pretty easy. No brute force or greedy algorithm simulating 100 Seconds. Just some simple math which came into my head right away. Still took me 20 min because I am slow as fuck at coding and I forgot about the "-" in the regex and only had positive velocities.

**Part 2:** This was really messed up. My first though was that this Christmas tree is full page. So I checked for the symmetry of the whole map. Nothing... So I was lost. The task had no clues or anything. I went on Reddit and saw some post about statistics. I did not read more into but got some ideas. I calculated the distributions of the robots after x seconds on the map in grid. I first did a grid of squares of size 10x10 which was 100 squares in the grid. The calculated the average and distribution and checked for at least 1/2 squares to be outside the 3*StandardDeveation. Found nothing for 100_000 Seconds. So I changed the square size to 20x20 and voilà, I found the beautiful tree! Took me an extra 90 minutes though haha

### Day 15
**Part 1:** I basically got the Idea on how to do it right away and solved it withing 30 minutes. Had practically no error but some typos or just little mistakes

**Part 2:** For this I also got the idea right away. Still took some time to code up and code some little silly mistakes like not including the last entry in an array and stuff. But the idea worked out right away after getting the little bugs working. For both parts in total it took me like 2 hours and 15 minutes. Was also my first time doing pointers in Go, so that also took some time away to google and try out in the Go Playground. This was a nice puzzle!

### Day 16
**Part 1:** It was pretty easy. Go it working really quickly

**Part 2:** This was difficult. I made the crucial mistake to turn and to go in the same step. So I was not able to compare the values with each other.
This took me probably like 10 hours or so. Really stupid. This was again a perfect example of simple but not easy. But I did it without googling or using any 
other libraries.

### Day 17
**Part 1:** The first part was really easy. I already built a whole computer class because I thought the second part building upon it. I got it first try and quickly moved on to the second part.

**Part 2:** At first I really had no idea what to do. But then I started looking into the program that is running and tried to figure out what it is doing. So I did a full program analysis and quickly found out the output is calculated like this:
```
	output_k = {
		b = A[A.len-3*(k+1),A.len-3*k-1] ^ 011
		return A[A.len-3*(k+1),A.len-3*k-1] ^ A[A.len-3*(k+1)-1-b,A.len-3*k-1-b] ^ 110
	}
```
After some time I wrote down the code and ran it: Fuck it only tries one path and does not stop if the output does not match the wanted output (the program itself).
I really did not know how to easily code this up but then *The Stack* came into my mind and with some copy pasted really simple Stack implementation in go I got it to work within 15 minutes. It prints out all possible starting values for Register A but I just need to manually choose the smallest one.