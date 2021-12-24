# Mars Rover

## Build and Run the Image


Build image from current directory:

```
docker build -t marsrover . 
```

Run image interactively:

```
docker run -i marsrover
```

## Rover Instructions

Enter grid size (m x n). Leave space between values, example format [m n]:

```
4 8
```

Add as many robots as like on a new line after hitting return:

```
(2, 3, E) LFRFF
(0, 2, N) FFLRFF
(2, 3, N) FLLFR
(1, 0, S) FFRLF

```

Hit return as a blank entry after the last robot to see results:

```
(4, 4, E)
(0, 4, N) LOST
(2, 3, W)
(1, 0, S) LOST
```

## Running Unit Tests

Run the test shell file to run unit tests.

```
./test.sh
```