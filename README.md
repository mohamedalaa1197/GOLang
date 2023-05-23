# GOLang

To pass Value by reference in Go, use ```&``` operator 
``` 
    var answer string
		fmt.Scan(&answer)
```

To define more than one variable in the same lin ==> 

```
question, question2 := record[0], record[1] 
```

To ignore a variable u can use the magic key **_**, here u ignored the record[1]

```
question, _ := record[0], record[1] 
```

To loop through a list or array of items

```
for i, record := range records {
		// define more than one variable in the same line
		question, _ := record[0], record[1]
		fmt.Println("%d. %s?\n", i+1, question)
		//to pass value with pointer (reference), we use the &
		var answer string
		fmt.Scan(&answer)
	}

```

i => indicates the index of the array, and record indicates the record itself

in Go each function, can have more than one return, and we need to handel this return (and listen if there is an error), as there is no throwh exception in Go.
Like here we are handling the exception.
```	
f, err := os.Open(problemsFilename)
	// in go we should handle the error in each functio, as there is no throw try/catch checkes in Go, if we didn't handle this then a panic will happen
	// and may cos the program to stop.
	if err != nil {
		fmt.Printf("failed to open the file %s", err)
		return
	}
```

also for any IO operation, we need to handle closing the connection after u finish, So Go is giving us a magic key word to handle this once the operation is finished 
like the example here 
```
f, err := os.Open(problemsFilename)
	// in go we should handle the error in each functio, as there is no throw try/catch checkes in Go, if we didn't handle this then a panic will happen
	// and may cos the program to stop.
	if err != nil {
		fmt.Printf("failed to open the file %s", err)
		return
	}
	// to make sure after we finish dealing with the file we close the file, and release it back to the IO system, also useful when error happen as it will also run if an error happen
	defer f.Close()

```

```Char``` type in Go ==> ```rune```


another type use for If 

```
	if length := getLength(email); length < 1{
		fmt.Printf("it is one")
	}

```
