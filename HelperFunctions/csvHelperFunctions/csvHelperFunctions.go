package csvHelperFunctions

import (
	"encoding/csv"
	"log"
	"math/big"
	"multihop_gobot/global"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

///// UNFILTERED POOLS

func WriteToCSV(filename string, queriedAddresses [][3]common.Address) {

	// convert queriedAddresses to strings
	var queriedAddressesSTR [][]string
	for i, _ := range queriedAddresses {
		tmp1 := queriedAddresses[i][0].String()
		tmp2 := queriedAddresses[i][1].String()
		tmp3 := queriedAddresses[i][2].String()
		tmp := []string{tmp1, tmp2, tmp3}
		queriedAddressesSTR = append(queriedAddressesSTR, tmp)
	}

	// Create file
    file, err := os.Create(filename)
    checkError("Cannot create file", err)
    defer file.Close()

	// Create csv object for serialising data into csv format
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range queriedAddressesSTR {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }
}


func LoadSavedQueriedPools(filename string) [][3]common.Address {
	f, err := os.Open(filename)
    if err != nil {
        log.Fatal("Unable to read input file " + filename, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filename, err)
    }

	// Parse the data into the common.Address format
	var queriedAddresses [][3]common.Address
	for i,_ := range data {
		tmp1 := common.HexToAddress(data[i][0])
		tmp2 := common.HexToAddress(data[i][1])
		tmp3 := common.HexToAddress(data[i][2])
		tmp := [3]common.Address{tmp1, tmp2, tmp3}
		queriedAddresses = append(queriedAddresses, tmp)
	}

    return queriedAddresses
}

///// FILTERED POOLS

func WriteToCSVFilteredPools(filename string, filteredPathsLength2 []global.PathsLength2 ) {

	// convert queriedAddresses to strings
	var filteredAddressesSTR [][]string

	for i,_ := range filteredPathsLength2{
		tmpPath := filteredPathsLength2[i]
		
		encodedPoolToSave := []string{tmpPath.Tuple1[0].String(), tmpPath.Tuple1[1].String(), tmpPath.Tuple1[2].String(),
										 tmpPath.Tuple2[0].String(), tmpPath.Tuple2[1].String(), tmpPath.Tuple2[2].String()}
		filteredAddressesSTR = append(filteredAddressesSTR, encodedPoolToSave)

	}

	// Create file
    file, err := os.Create(filename)
    checkError("Cannot create file", err)
    defer file.Close()

	// Create csv object for serialising data into csv format
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range filteredAddressesSTR {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }
}



func LoadSavedFilteredPaths(filename string) []global.PathsLength2 {

	f, err := os.Open(filename)
    if err != nil {
        log.Fatal("Unable to read input file " + filename, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filename, err)
    }

	// Parse the data into the common.Address format
	var filteredPaths2 []global.PathsLength2
	for i,_ := range data {
		token0_p1 := common.HexToAddress(data[i][0])
		token1_p1 := common.HexToAddress(data[i][1])
		pool_p1 := common.HexToAddress(data[i][2])
		token0_p2 := common.HexToAddress(data[i][3])
		token1_p2 := common.HexToAddress(data[i][4])
		pool_p2 := common.HexToAddress(data[i][5])
		
		Tuple1 := [3]common.Address{token0_p1, token1_p1, pool_p1}
		Tuple2 := [3]common.Address{token0_p2, token1_p2, pool_p2}
		Token1 := token0_p1
		Token2 := token1_p1
		ID := pool_p1.String() + pool_p2.String()

		tmpPath2 := global.PathsLength2{Tuple1: Tuple1, Tuple2: Tuple2, Token1: Token1, Token2: Token2, ID : ID}
		filteredPaths2 = append(filteredPaths2, tmpPath2)
	}

    return filteredPaths2
}


//// Uniswap V3

func LoadUniV3PoolAddresses(filename string) []common.Address {

	f, err := os.Open(filename)
    if err != nil {
        log.Fatal("Unable to read input file " + filename, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filename, err)
    }

	var dataAddress []common.Address
	for i,_ := range data{
		dataAddress = append(dataAddress, common.HexToAddress(data[i][0]))
	}
	return dataAddress

}

func SaveUniV3Structs(uniV3PoolStructs []global.UniV3Pool) {

	// convert queriedAddresses to strings
	var uniV3StructsStr [][]string
	for i, _ := range uniV3PoolStructs {
		tmp1 := uniV3PoolStructs[i].Token1.String()
		tmp2 := uniV3PoolStructs[i].Token2.String()
		tmp3 := uniV3PoolStructs[i].Pool.String()
		tmp4 := uniV3PoolStructs[i].Fee.String()
		tmp := []string{tmp1, tmp2, tmp3,tmp4}
		uniV3StructsStr = append(uniV3StructsStr, tmp)
	}

	// Create file
    file, err := os.Create(global.FilenameToSaveUniV3Structs)
    checkError("Cannot create file", err)
    defer file.Close()

	// Create csv object for serialising data into csv format
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range uniV3StructsStr {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }


}


func LoadUniV3Structs(filename string) []global.UniV3Pool {

	f, err := os.Open(filename)
    if err != nil {
        log.Fatal("Unable to read input file " + filename, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {        log.Fatal("Unable to parse file as CSV for " + filename, err)    }

	var dataStructs []global.UniV3Pool
	for i,_ := range data {
		tmpToken1 := common.HexToAddress(data[i][0])
		tmpToken2 := common.HexToAddress(data[i][1])
		tmpPool := common.HexToAddress(data[i][2])
		tmpFee, err := new(big.Int).SetString( data[i][3], 10)
		if !err {        panic("Unable to parse data in LoadUniV3Structs")    }
		tmpStruc := global.UniV3Pool{Token1 : tmpToken1, Token2 : tmpToken2, Pool: tmpPool, Fee: tmpFee}
		dataStructs = append(dataStructs, tmpStruc)

	}
	return dataStructs

}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
