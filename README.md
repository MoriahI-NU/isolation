# Isolation Trees
## A Comparison Between Go, Python, and R

### The Setup
- summarize the work on Go package selection, implementation, and testing. 

For this experiment, I had to research various go packages that would allow me to implement an isolation forest with the MNIST dataset. I looked at a few that were specifically for iforests like "go_iforest" (2021), "eif" (2020), and "isolationforest" (2017). Although go_iforest seemed the most promising, I had issues trying to import the package so I looked elsewhere, to the "trees" package from sjwhitworth/golearn. I had already used other packages from this developer and had confidence in their performance and ease of use.  

Once I had chosen a package, the implementation and testing were quite easy. All I required from the "trees" package were three short functions to produce my anomaly scores. I saved the Go-generated anomaly scores in a csv file (GoScores.csv) and visualized them with a plot (density_plot.png). I then compared my Go scores with Python and R generated anomaly scores, creating plots and tables to see how similar the results from each program are. (See below for more on this.)

### Package Used
- What led to the selection of one package over others?  

I used the sjwhitworth/golearn/trees package to implement an isolation forest model. I chose this one specifically because of multiple reasons. First off, I have worked with other sjwhitworth/golearn packages and have found them to be well documented, popular to use (which means plenty of examples), and relatively current (this package was published December 28, 2022). 

- How easy was it to clone the package repository, implement the code, and set up the tests?  

Using this package was a smooth and easy process. I had no issues importing it. With the functions provided, I was able to complete the test process in a mere three lines of code - one to create a new isolation forest instance, one to fit the model to the training data, and lastly one to predict scores. It makes the process incredibly short, concise, and easy to read.


### The Data
I was unable to upload one of my datasets to this repo due to GitHub's restrictions on file sizes. So I have decided to list all the datasets I used here. Feel free to go and check them out at their original sources:

MNIST datasets in CSV from Kaggle:
https://www.kaggle.com/datasets/oddrationale/mnist-in-csv?select=mnist_train.csv

- mnist_train.csv
- mnist_test.csv

Python and R score datasets (used for correlation plots):
https://github.com/ThomasWMiller/jump-start-mnist-iforest/tree/main

- pythonScores.csv
- solitudeRScores.csv
- isotreeRScores.csv

### Results
- Were the tests successful?  
Yes, the tests completed successfully and anomaly scores from the Go isolation forests were written out to the "GoScores.csv" file in this repo. You can also see a density plot of these scores in "density_plot.png".

- To what extent were Go results similar to results from R and/or Python?  
The Go results were quite similar to Python. I had my application print correlation scores rather than write them to a csv. Correlation scores with my Go application are as follows:

#### Go Correlations
|Language    |Score|
|------------|-----|
|Python      |0.95 |
|R (solitude)|0.64 |
|R (isotree) |0.75 |

You will also find correlation plots in the files that end with"..._corr.png". I'd like to compare these scores with correlcation scores between Python and R. Below is a table showing correlation scores obtained by https://github.com/ThomasWMiller/jump-start-mnist-iforest/tree/main

#### Python Correlations
|Language    |Score|
|------------|-----|
|R (solitude)|0.58 |
|R (isotree) |0.72 |

Both Python and Go share similar correlation scores with R. (Although Go has just a touch more in common with R scores) One important thing to keep in mind, however, is that hyperparameter tuning may affect these scores. Python and Go are so close (0.95) that with a little tweaking of the parameters, we could potentially get them to output the exact same anomaly results.

### Insights
- How difficult will it be for the firm to implement Go-based isolation forests as part of its data processing pipeline?

Based on my experience with the sjwhitworth/golearn/trees package, I don't think the firm will have much trouble at all. There are plenty of examples out there, the functions are easy to learn, and (as far as isolation forests go) the code required is short and clean. It performs about on par with Python and the learning curve is very small.