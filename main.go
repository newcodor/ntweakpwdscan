package main

import (
	"flag"
	"fmt"
	"hashcompare/commonutils"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	hashFilePath     string = "hash.txt"
	passwordFilePath string = "password.txt"
	outpath          string = "result.txt"
	outType          string = "xlsx"
	help             bool
	columns          []string               = []string{"username", "password", "ntlm_hash"}
	columnLength     map[string]int         = map[string]int{"username": 16, "password": 10, "ntlm_hash": 25}
	headStyle        *commonutils.XlsxStyle = commonutils.NewXlsxStyle("center", "00A2A5A1", "Verdana", 13)
	cellCenterStyle  *commonutils.XlsxStyle = commonutils.NewXlsxStyle("center", "FFFFFFFF", "宋体", 11)
	cellLeftStyle    *commonutils.XlsxStyle = commonutils.NewXlsxStyle("left", "FFFFFFFF", "宋体", 11)
	column_tags      []commonutils.XlsxCol  = []commonutils.XlsxCol{commonutils.XlsxCol{0, "Index", 5.5, headStyle.Style, cellCenterStyle.Style}, commonutils.XlsxCol{1, "username", 12.0, headStyle.Style, cellLeftStyle.Style}, commonutils.XlsxCol{2, "password", 10, headStyle.Style, cellCenterStyle.Style}, commonutils.XlsxCol{3, "ntlm_hash", 25, headStyle.Style, cellLeftStyle.Style}}
)

func main() {
	flag.StringVar(&hashFilePath, "f", "hash.txt", "hash file path")
	flag.StringVar(&passwordFilePath, "pass", "password.txt", "password file path")
	flag.StringVar(&outType, "oT", "xlsx", "out file type:txt,csv,xlsx")
	flag.BoolVar(&help, "h", false, "show help")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if !commonutils.IsFileExist(hashFilePath) {
		fmt.Printf("%s is not exist!\n", hashFilePath)
		os.Exit(0)
	}

	if !commonutils.IsFileExist(passwordFilePath) {
		fmt.Printf("%s is not exist!\n", passwordFilePath)
		os.Exit(0)
	}
	hashList := commonutils.ReadFileLines(hashFilePath)
	passwords := commonutils.ReadFileLines(passwordFilePath)
	fmt.Println("hash file: " + hashFilePath)
	fmt.Printf("password file: %s\n", passwordFilePath)
	var passwd_hash_list map[string]string = make(map[string]string, len(passwords))
	for _, v := range passwords {
		passwd_hash_list[commonutils.FromASCIIStringToHex(v)] = v
	}
	// for k,v :=range passwd_hash_list {
	// 	fmt.Printf("%s,%s\n",k,v)
	// }
	var result_list []interface{}
	for i := 0; i < len(hashList); i++ {
		v := hashList[i]
		if strings.HasPrefix(v, "NTLM : ") && len(v) != 7 {
			hash := strings.TrimSpace(strings.Split(v, ":")[1])
			password, exits := passwd_hash_list[hash]
			if exits {
				username := strings.TrimSpace(strings.Split(hashList[i-2], ":")[1])
				fmt.Printf("%s,%s,%s\n", username, password, hash)
				res := make(map[string]string, 3)
				res["username"] = username
				res["password"] = password
				res["ntlm_hash"] = hash
				result_list = append(result_list, res)
			}
		}
	}
	// for _,v :=range hashList {

	// }
	fmt.Printf("\nDone!\nfind accounts with weakpassword: %v\n", len(result_list))
	if len(result_list) > 0 {
		i := 1
		for _, v := range result_list {
			v.(map[string]string)["Index"] = strconv.Itoa(i)
			i++
		}
		outpath = time.Now().Format("accounts_2006-01-02_150405") + "." + outType
		isSaved := false
		if outType == "xlsx" {
			isSaved = commonutils.WriteFileLinesToExcel(column_tags, outpath, result_list)
		} else if outType == "csv" {
			isSaved = commonutils.WriteFileLinesBySplitChar(columns, outpath, result_list, ",")
		} else {
			isSaved = commonutils.WriteFileLinesBySplitChar(columns, outpath, result_list, " ")
		}
		if isSaved {
			fmt.Printf("save to file %s ……\n", outpath)
		} else {
			fmt.Println("save result failed!")
		}
	}
}
