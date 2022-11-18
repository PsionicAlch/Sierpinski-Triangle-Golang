# Sierpinski Triangle Golang

This programming challenge was inpsired by this [TikTok video](https://vm.tiktok.com/ZMNNk5PW7/?k=1).

![Image of completed Sierpinski Triangle](https://github.com/Wolfman13/Sierpinski-Triangle-Golang/blob/master/sierpinski.png)

The idea is to draw a Sierpinski triangle using the following algorithm:

1. Start with three points. The points of a triangle.
2. Choose a random point inside the triangle and draw it.
3. Find the midpoint between the random point and a random point of the triangle and draw it.
4. Repeat step 3 by using the mid point as the new random point.

If done enough times a fractal triangle will begin to form.

## Running the application.

To run the application, you just need to have golang installed (I originally wrote and tested this application using Go 1.19.3) and then you can run the following command.

go run . [width of the image] [height of the image] [amount of pixels]

For this project I had some difficulties with using iterations to generate the image. I am not sure what the issue is but I have a feeling it had something to do with random number generation. To get around this problem I decided to just keep track of how many unique pixels coordinates there are. So the amount of pixels is the amount of pixels you would like to have filled in. The more pixels you use the clearer the image will become but the longer it will take to generate since the program will have to randomly stumble upon each unique coordinate. The image above had a width and height of 501, and 23k unique pixels. I found that after 22k the progress slows down a lot.

## Similar projects:

- [Python](https://github.com/Wolfman13/Sierpinski-Triangle-Python)
