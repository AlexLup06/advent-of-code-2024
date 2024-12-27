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

### Day 18
**Part 1:** Really easy, though at first I was extremely confused and I thought we walk through the map as the bytes fall. But it was actually just Dijkstra and then simulating Dijkstra for the more falling bytes.

**Part 2:** Really easy. Don't even know what to say. Took me 5 minutes or so

### Day 19
**Part 1:** Oh Boy, Oh Boy... I thought I am going to be the really smart one and wrote myself this beautiful graph that represents this language. Well it is slow as hell because it has to check every single damn option. Quickly deleted everything and just did some stupid old Stack. 100 fewer lines of code and did get me the answer within some milliseconds or so.

**Part 2:** This was a hell of a ride. Of course just brute forcing did not work. Not even close. So I got to thinking. Nothing came to my mind. Is there a really smart way of doing this!? I just thought of a cache: Instead of going down all the paths of the last 15 characters lets just calculate the score for the last 15 characters and cache that. Well building the cache was really quick. But then running the whole design up until the last 15 characters again took way too long! I had no clue what to do and started doing Day 20. Then at night it came to me. I will just start from the back and work my way up and all the previous answers will be my cache. This was soooo fast!! I did look into Reddit once but only saw that my idea with the cache was correct. 

### Day 20
**Part 1:** For this one I pretty quickly got a good idea right away. Just start from the back and do Dijkstra! So now I know for each tile how far its way from the 
end. If I now want to cheat and go from Tile A to tile B I just calculate the difference of the distance minus the cheated distance. Voilà you get the time saved

**Part 2:** I Applied the same idea for this Problem, but now the possible ways of cheating is adjusted in a way that from a tile you are allowed to cheat so that
|x|+|y|≤20 (Diamond shape). At first, I thought I need some kind of cache to safe already checked cheats. But this turned out to be really inefficient. So I removed
it. Then I faced a stupid bug where I double counted all cheats with y=0. In the end I got it, and I am pretty proud of the solution to this Problem.

### Day 21
**Part 1:**

**Part 2:**