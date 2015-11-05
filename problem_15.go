/* I did this on paper, but can adapt to golang.
What I did was, use n choose k formula. for the example, we have
4 choose 2. there are four moves to get to the bottom, and we need 
2 moves right (or down). This choose is the same as this formula:
n! / k!(n-k)! . plugging in for 40 and 20, we get 40!/ 20!20!.
