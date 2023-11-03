package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"

	"github.com/kshedden/datareader"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type float64Values []float64

func (f float64Values) Len() int {
	return len(f)
}

func (f float64Values) Value(i int) float64 {
	return f[i]
}

func main() {

	//Read in MNIST_train from CSV
	mnist_train, err := base.ParseCSVToInstances("mnist_train.csv", true)
	if err != nil {
		fmt.Println("Error opening MNIST dataset:", err)
	}

	// Read Python values from a separate CSV file
	PyValues, err := readPyValues("pythonScores.csv")
	if err != nil {
		fmt.Println("Error reading actual values:", err)
		return
	}

	//Read R solitude values from separate CSV file
	RValues_Sol, err := readRValues_Sol("solitudeRScores.csv")
	if err != nil {
		fmt.Println("Error reading actual values:", err)
		return
	}

	//Read R isotree values from CSV
	RValues_Iso, err := readRValues_Iso("isotreeRScores.csv")
	if err != nil {
		fmt.Println("Error reading actual values:", err)
		return
	}

	//Create a new isolation forest model, fit it to training data and predict
	iforest := trees.NewIsolationForest(100, 100, 256)

	iforest.Fit(mnist_train)

	preds := iforest.Predict(mnist_train)

	//Create variables
	var avgScore float64
	var min float64
	min = 1

	//Get average prediction score and the minimum prediction score
	for i := 0; i < 60000; i++ {
		temp := preds[i]
		avgScore += temp
		if temp < min {
			min = temp
		}
	}
	fmt.Println("average score:", avgScore/60000)
	fmt.Println("minimum:", min)

	////////////////////////////
	//Write scores to CSV
	////////////////////////////

	// Create new file/writer
	file, err := os.Create("GoScores.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Column Header
	writer.Write([]string{"Go_Score"})
	writer.Flush()

	// Convert to string format and write all scores
	for i := 1000; i < 60000; i++ {
		predictionStr := strconv.FormatFloat(preds[i], 'f', -1, 64)
		writer.Write([]string{predictionStr})
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	/////////////////////////////////////////////////////////////
	//First Plot: Go Density Plot
	/////////////////////////////////////////////////////////////

	// Create a new plot
	p, err := plot.New()
	if err != nil {
		fmt.Println("Error creating plot:", err)
		return
	}

	// Create a histogram plotter
	vals := float64Values(preds)
	hist, err := plotter.NewHist(vals, 10)
	if err != nil {
		fmt.Println("Error creating histogram plotter:", err)
		return
	}

	// Add the histogram to the plot
	p.Add(hist)

	// Set title and labels
	p.Title.Text = "Density Plot"
	p.X.Label.Text = "Go Anomaly Score"
	p.Y.Label.Text = "Frequency"

	// Save the plot
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "density_plot.png"); err != nil {
		fmt.Println("Error saving plot:", err)
	}

	fmt.Println("Density plot saved as density_plot.png")

	//////////////////////////////////////////////////////////////////////
	//Second Plot: Python v Go Scores
	//////////////////////////////////////////////////////////////////////

	// Calculate the correlation between actual values and predictions
	correlation := stat.Correlation(preds, PyValues, nil)

	fmt.Printf("Correlation between preds and python values: %.2f\n", correlation)

	// Create a correlation plot
	p, err = plot.New()
	if err != nil {
		fmt.Println("Error creating plot:", err)
		return
	}

	//Title and Labels
	p.Title.Text = "Anomaly Score Correlation Plot"
	p.X.Label.Text = "Python"
	p.Y.Label.Text = "Go"

	// Create a new scatter plotter and set X and Y values
	pts := make(plotter.XYs, len(preds))
	for i, pred := range preds {
		pts[i].X = PyValues[i]
		pts[i].Y = pred
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		fmt.Println("Error creating scatter plot:", err)
		return
	}

	//Add to plot
	p.Add(s)

	//Save as a file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "go_python_corr.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}

	//////////////////////////////////////////////////////////////////////
	//Third Plot: R Solitude Scores v Go
	//////////////////////////////////////////////////////////////////////

	//Calculate correlation scores
	correlation2 := stat.Correlation(preds, RValues_Sol, nil)

	fmt.Printf("Correlation between preds and R solitude values: %.2f\n", correlation2)

	// Create a correlation plot
	p2, err := plot.New()
	if err != nil {
		fmt.Println("Error creating plot:", err)
		return
	}

	//Title and labels
	p2.Title.Text = "Anomaly Score Correlation Plot"
	p2.X.Label.Text = "R Solitude"
	p2.Y.Label.Text = "Go"

	// Create a new scatter plotter and set X and Y values
	pts2 := make(plotter.XYs, len(preds))
	for i, pred := range preds {
		pts2[i].X = RValues_Sol[i]
		pts2[i].Y = pred
	}

	s2, err := plotter.NewScatter(pts2)
	if err != nil {
		fmt.Println("Error creating scatter plot:", err)
		return
	}

	//Add to plot
	p2.Add(s2)

	//Save as a file
	if err := p2.Save(4*vg.Inch, 4*vg.Inch, "go_rsol_corr.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}

	//////////////////////////////////////////////////////////////////////
	//Fourth Plot: R Isotree Scores vs Go
	//////////////////////////////////////////////////////////////////////

	//Calculate correlation scores
	correlation3 := stat.Correlation(preds, RValues_Iso, nil)

	fmt.Printf("Correlation between preds and R isotree values: %.2f\n", correlation3)

	// Create a correlation plot
	p3, err := plot.New()
	if err != nil {
		fmt.Println("Error creating plot:", err)
		return
	}

	//Title and labels
	p3.Title.Text = "Anomaly Score Correlation Plot"
	p3.X.Label.Text = "R IsoTree"
	p3.Y.Label.Text = "Go"

	// Create a new scatter plotter and set X and Y values
	pts3 := make(plotter.XYs, len(preds))
	for i, pred := range preds {
		pts3[i].X = RValues_Iso[i]
		pts3[i].Y = pred
	}

	s3, err := plotter.NewScatter(pts3)
	if err != nil {
		fmt.Println("Error creating scatter plot:", err)
		return
	}

	//Add to plot
	p3.Add(s3)

	//Save as a file
	if err := p3.Save(4*vg.Inch, 4*vg.Inch, "go_riso_corr.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}

}

//////////////////
//Helper Functions
//////////////////

// Function to read Python values from a CSV file
func readPyValues(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := datareader.NewCSVReader(file)
	csvReader.HasHeader = true
	dataset, err := csvReader.Read(60000)
	if err != nil {
		return nil, err
	}

	x, m, _ := dataset[1].AsFloat64Slice()
	_ = m

	return x, nil
}

// Function to read R solitude values from a CSV file
func readRValues_Sol(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := datareader.NewCSVReader(file)
	csvReader.HasHeader = true
	dataset, err := csvReader.Read(60000)
	if err != nil {
		return nil, err
	}

	x, m, _ := dataset[2].AsFloat64Slice()
	_ = m

	return x, nil
}

// Function to read R isotree values from a CSV file
func readRValues_Iso(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := datareader.NewCSVReader(file)
	csvReader.HasHeader = true
	dataset, err := csvReader.Read(60000)
	if err != nil {
		return nil, err
	}

	x, m, _ := dataset[0].AsFloat64Slice()
	_ = m

	return x, nil
}
