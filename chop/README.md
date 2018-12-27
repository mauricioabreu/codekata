# NOTES

There are my notes about the [kata02: karate chop](http://codekata.com/kata/kata02-karate-chop/)

I hope you find it useful. :-)

## 1st approach

My first approach is the easier to understand in my opinion.

We first the the lenght of the array. If it is empty we don't need to continue. 

The real job starts getting the elements in the left/right side of the array, sum the numbers and divides by 2 to get the middle (if it has a even number of elements, we use floor to adjust the number to the floor hehe)
```go
middle := int(math.Floor(float64((leftSide + rightSide) / 2)))
```

Now we have check if the middle number is lower or greater than the value we are passing in.
Let's have an example:

```array: [1 3 9 12 38 49 89]```

Let's find the position of the number `38`.

Middle position is `3` and the number is `12`.

Now we need to check if `12` is lower or greater than the value we are looking for. Why? Because with this we will know if it is located in the right or to the left. If the number is in the left side, we need to decrease the possibilites to find them in the whole array by decreasing the array closer to the middle of the array (to the right or to the left).

Tha being said, our new search will have less possibilities: `array: [38, 48, 89]`
And so on, until we find the position of our value.

This is how a `simple` binary search works. I don't know efficient this solution is. Next approaches will focus on different strategies.