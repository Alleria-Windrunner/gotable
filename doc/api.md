


# API
This section describes the gotable APIs.




## github.com/Alleria-Windrunner/gotable

### Create a simple table (recommended)
```go
func Create(columns ...string) (*table.Table, error)
```

```go
table, err := gotable.Create("China", "US", "French")
```

### Create a simple table from struct

```go
func CreateByStruct(v interface{}) (*table.Table, error)
```

### Create a safe table

```go
func CreateSafeTable(columns ...string) (*table.SafeTable, error)
```

### Specify default value

The ```gotable.Default``` constant replaces the default value stored in column. Refer to Set Default Values in the Demo 
section for more information.

```go
gotable.Default
```

### Load data from file

Currently，csv and json file are supported.
```go
func Read(path string) (*table.Table, error)
```

### Color control

The following constants are used in conjunction with the ```*table.SetColumnColor``` method to change the column color.
#### display type
Default display
```go
gotable.TerminalDefault
```
Fonts are highlighted
```go
gotable.Highlight
```
Underline
```go
gotable.Underline
```
Font flash
```go
gotable.Flash
```
#### color
```go
gotable.Black
gotable.Red
gotable.Green
gotable.Yellow
gotable.Blue
gotable.Purple
gotable.Cyan
gotable.Write
```

Do not set the background color
```go
gotable.NoneBackground
```





## APIs for simple table type(*table.Table)

### Clear data
The clear method is used to clear all data in the table, include columns, rows and parts.
```go
func (tb *Table) Clear()
```



### Get table type

The type method returns a type of table.
```go
func (tb *Table) Type() string
```

### Add part (recommended)
Add a new part with new columns to the table. Similar to Creat(). See the Demo section for more information.
```go
func (tb *Table) AddPart(columns ...string) error
```

```go
table.AddPart("name", "salary")
```

### Add row (recommended)
Add a row to the table. Support Map and Slice. See the Demo section for more information.

 ```go
func (tb *Table) AddPNRow(partNumber int, row interface{}) error
```
```go
func (tb *Table) AddRow(row interface{}) error
```
```go
// Use map
	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["French"] = "Paris"
	err = table.AddRow(row)
// Use Slice
	row2 := []string{"Yinchuan", "Los Angeles", "Orleans"}
	err = table.AddRow(row2)
```
### Add a list of rows

Method ```AddRows``` add a list of rows. It returns a slice that consists of adding failed rows.

```go
func (tb *Table) AddRows(rows []map[string]string) []map[string]string
```



### Add column

Add a new column to table,can assign partNumber or last part in default
```go
func (tb *Table) AddPNColumn(partNumber int, column string) error
```

```go
func (tb *Table) AddColumn(column string) error
```

### Clear rows

Clear rows in selected part, or the last one in default

```go
func (tb *Table) DelPNRows(partNumber int) error
```

```go
func (tb *Table) DelRows() error
```

### Set Max Length of column

Set the max length of the selected column in selected part
```go
func (tb *Table) SetColumnMaxLength(partNumber int, column string, maxlength int) error
```

### Adapt Column Length (recommended)

Auto adapt the length of selected column in selected part according to a longer one
```go
func (tb *Table) AdaptColLen(longPN int, shortPN int, adCol string) error
```
```go
table.AdaptColLen(0, 1, "salary")
```

### Get Column Length

Get the selected column length
```go
func (tb *Table) GetPNColumnLen(pn int, column string) (int, error)
```

### Get part length
```go
func (tb *Table) PartLength() int
```

### Get columns
```go
func (tb *Table) GetPNColumns(partNumber int) []string
```
```go
func (tb *Table) GetColumns() []string
```

### Get values map

Use table method ```GetValues``` to get the map that save values.
```go
func (tb *Table) GetPNValues(partNumber int) []map[string]string
```

```go
func (tb *Table) GetValues() []map[string]string
```

### Print table

```*table``` implements ```fmt.Stringer``` interface, so you can use the ```fmt.Print```, ```fmt.Printf``` functions 
and so on to print the contents of the table instance.
```go
func (tb *Table) String() string
```



### Set default value

By default, the default value for all columns is an empty string.

```go
func (b *base) SetDefault(column string, defaultValue string)
```



### Drop default value

```go
func (b *base) DropDefault(column string)
```



### Get default value

Use table method ```GetDefault``` to get default value of column. If column does not exist in the table, the method returns an empty string.

```go
func (b *base) GetDefault(column string) string
```



### Get default map

Use table method ```GetDefaults``` to get default map of head.

```go
func (b *base) GetDefaults() map[string]string
```



### Arrange: center, align left or align right (recommended)

By default, the table is centered. You can set a header to be left 
aligned or right aligned. See the next section for more details on how 
to use it.

```go
func (tb *Table) PNAlign(partNumber int, column string, mode int)
```
```go
func (tb *Table) Align(column string, mode int)
```



### Check empty

Use table method ```Empty``` to check if the table is empty.

```go
func (tb *Table) Empty() bool
```

### Check value exists

```go
func (tb *Table) PNExist(partNumber int, value map[string]string) bool
```

```go
func (tb *Table) Exist(value map[string]string) bool
```



### Get table length

```go
func (tb *Table) Length() int
```



### To JSON string

Use table method ```JSON``` to convert the table to JSON format.
The argument ```indent``` indicates the number of indents.
If the argument ```indent``` is less than or equal to 0, then the ```JSON``` method unindents.
```go
func (tb *Table) JSON(indent int) (string, error)
```



### To XML string

Use table method ```XML``` to convert the table to XML format.
The argument ```indent``` indicates the number of indents.
If the argument ```indent``` is less than or equal to 0, then the ```XML``` method unindents.
```go
func (tb *Table) XML(indent int) string
```



### Save the table data to a JSON file

Use table method ```ToJsonFile``` to save the table data to a JSON file.
```go
func (tb *Table) ToJsonFile(path string, indent int) error
```



### Save the table data to a CSV file

Use table method ```ToCSVFile``` to save the table data to a CSV file.
```go
func (tb *Table) ToCSVFile(path string) error
```



### Close border

Use table method ```CloseBorder``` to close table border.
```go
func (tb *Table) CloseBorder()
```

### Open border

Use table method ```OpenBorder``` to open table border. By default, the border property is turned on.
```go
func (tb *Table) OpenBorder()
```

### Set border
0:" " 1:"-" 2:"=" 3:"~" 4:"+"
```go
func (tb *Table) SetBorder(border int8)
```

### Set the title line
Set the title line to hide or not in the part, 
```go
func (tb *Table) SetPNTitleLine(partNumber int, value int8) error
```

```go
func (tb *Table) SetTitleLine(value int8) error
```

### Set Title Hide
Set the title (first line) to hide or not in the part
```go
func (tb *Table) SetPNTitleHide(partNumber int, ishide bool) error
```

```go
func (tb *Table) SetTitleHide(ishide bool) error
```


### Has column

Table method ```HasColumn``` determine whether the column is included.
```go
func (tb *Table) HasPNColumn(partNumber int, column string) bool
```
```go
func (tb *Table) HasColumn(column string) bool
```



### Check whether the columns of the two tables are the same

Table method ```EqualColumns``` is used to check whether the columns of two tables are the same. This method returns
true if the columns of the two tables are identical (length, name, order, alignment, default), and false otherwise.
```go
func (tb *Table) EqualColumns(other *Table) bool
```



### Set column color

Table method ```SetColumnColor``` is used to set the color of a specific column. The first parameter specifies the name 
of the column to be modified. The second parameter indicates the type of font to display. Refer to the Color control 
section in this document for more information. The third and fourth parameters specify the font and background color.
```go
func (tb *Table) SetColumnColor(columnName string, display, fount, background int)
```



### Custom ending string

By default, a new blank line will print after table printing. You can designate your ending string by reset
```table.End```.



### Is simple table

Method IsSimpleTable is used to check whether the table type is simple table.

```go
func (b *base) IsSimpleTable() bool
```



### Is safe table

Method IsSafeTable is used to check whether the table type is safe table.

```go
func(b *base) IsSafeTable() bool
```



## APIs for safe table type(*table.SafeTable)

### Get table type
The type method returns a type of table.
```go
func (tb *Table) Type() string
```



### Is simple table

Method IsSimpleTable is used to check whether the table type is simple table.

```go
func (b *base) IsSimpleTable() bool
```



### Is safe table

Method IsSafeTable is used to check whether the table type is safe table.

```go
func(b *base) IsSafeTable() bool
```



### Add row

Add a row to the safe table. Only support Map. See the Demo section for more information.
```go
func (s *SafeTable) AddRow(row interface{}) error
```



### Add a list of rows

Method ```AddRows``` add a list of rows. It returns a slice that consists of adding failed rows.

```go
func (s *SafeTable) AddRows(rows []map[string]string) []map[string]string
```





### Custom ending string

By default, a new blank line will print after table printing. You can designate your ending string by reset 
```table.End```.



### Add column

```go
func (s *SafeTable) AddColumn(column string) error
```



### Set default value

By default, the default value for all columns is an empty string.

```go
func (b *base) SetDefault(column string, defaultValue string)
```



### Drop default value

```go
func (b *base) DropDefault(column string)
```



### Get default map

Use table method ```GetDefaults``` to get default map of head.

```go
func (b *base) GetDefaults() map[string]string
```



### Get table length

```go
func (s *SafeTable) Length() int
```





