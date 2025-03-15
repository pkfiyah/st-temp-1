# League Backend Challenge - Dev Notes

The REPO can be cloned from github at: https://github.com/pkfiyah/st-temp-1
or downloaded as a .tar from: https://drive.google.com/file/d/1H0xFIuLX9ao-01lzh--TXnnH2oNUe2o4/view?usp=sharing
and ran in similar fashion to the initial description (from the base folder with go installed: `go run .`)

If using VSCode or another devcontainer capable IDE, it can also be launched as a devcontainer, and ran within the same command as above.

Tests should also be runnable from the base folder with `go test ./...`

Math commands ignore non-integer values, but continue with found integers

Middleware ensures properly sized array. Empty values are valid, just requires proper size based on delimiter

Endpoints can be hit with the following commands:  
curl -F 'file=@./testfiles/matrix.csv' "localhost:8080/echo"  
curl -F 'file=@./testfiles/matrix.csv' "localhost:8080/invert"  
curl -F 'file=@./testfiles/matrix.csv' "localhost:8080/flatten"  
curl -F 'file=@./testfiles/matrix.csv' "localhost:8080/sum"  
curl -F 'file=@./testfiles/matrix.csv' "localhost:8080/multiply"  

Does not account for Int overflow on large multiplication/summation.

# Everything below was provided by League via the e-mail 

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages
