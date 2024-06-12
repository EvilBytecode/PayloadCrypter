package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/EvilBytecode/GolangStyle/pkg"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf16"
	"time"
)

func obfps1payload(scriptContent string, n int) string {
	rand.Seed(time.Now().UnixNano())
	encodedScript := base64.StdEncoding.EncodeToString([]byte(scriptContent))
	for i := 1; i <= n; i++ {
		gostyle.Write(fmt.Sprintf("Obfuscation Layer %d/%d\n", i, n), gostyle.YELLOW_TO_RED, true)
		randVar1 := rndstr(8)
		randVar2 := rndstr(8)
		scriptContent = fmt.Sprintf(`# obf by codepulze - evilbytecode

$%s = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String(@"
%s
"@))
iex $%s
`, randVar1, encodedScript, randVar2)
		encodedScript = base64.StdEncoding.EncodeToString([]byte(scriptContent))
	}
	return scriptContent
}

func getinput(prompt string) string {
	gostyle.Write(prompt, gostyle.YELLOW_TO_RED, false)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func rndstr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CryptPSPayload() {
	rand.Seed(time.Now().UnixNano())

	var psPath string
	gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Enter the path to the PowerShell script: ", gostyle.YELLOW_TO_RED, false)
	fmt.Scanln(&psPath)

	psContent, err := ioutil.ReadFile(psPath)
	if err != nil {
		gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error reading PowerShell script: %s\n", err), gostyle.RED_TO_YELLOW, false)
		return
	}

	layerInput := getinput("[@Admin:~/CodepulzeCrypter$] > Enter the number of obfuscation layers: ")
	numLayers, err := strconv.Atoi(layerInput)
	if err != nil {
		gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Invalid input. Please enter a valid number.\n", gostyle.RED_TO_YELLOW, false)
		return
	}

	obfuscatedContent := obfps1payload(string(psContent), numLayers)

	outputDir := filepath.Join("Built", "PowerShellPayloads")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error creating output directory: %s\n", err), gostyle.RED_TO_YELLOW, false)
		return
	}

	outputPath := filepath.Join(outputDir, "CodepulzeObfuscated.ps1")
	err = ioutil.WriteFile(outputPath, []byte(obfuscatedContent), 0644)
	if err != nil {
		gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error writing obfuscated script: %s\n", err), gostyle.RED_TO_YELLOW, false)
		return
	}

	gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Obfuscated script saved as %s\n", outputPath), gostyle.YELLOW_TO_RED, false)
}

func Obfuscate(code, alphabet string, setsNum, minNameLength int) string {
	Il := []string{"I", "l"}
	mixedAlphabet := shuffleString(alphabet)

	sets := make([]string, 0)
	alphabets := make([][]rune, 0)

	for x := 0; x < setsNum; x++ {
		name := generateSetName(Il, minNameLength, sets)
		sets = append(sets, name)
	}

	chunkSize := (len(mixedAlphabet) + setsNum - 1) / setsNum
	for _, chunk := range chunkString(mixedAlphabet, chunkSize) {
		alphabets = append(alphabets, []rune(chunk))
	}

	var newScript strings.Builder
	newScript.WriteString("@echo off\n")

	for idx, set := range sets {
		newScript.WriteString(fmt.Sprintf("Set %s=%s\n", set, string(alphabets[idx])))
	}
	newScript.WriteString("cls\n\n")

	for _, c := range code {
		replaced := false
		for idx, alphabet := range alphabets {
			if strings.ContainsRune(string(alphabet), c) {
				newScript.WriteString(fmt.Sprintf("%%%s:~%d,1%%", sets[idx], strings.IndexRune(string(alphabet), c)))
				replaced = true
				break
			}
		}
		if !replaced {
			newScript.WriteRune(c)
		}
	}

	return newScript.String()
}

func shuffleString(s string) string {
	r := []rune(s)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	return string(r)
}

func generateSetName(Il []string, minLength int, sets []string) string {
	rand.Seed(time.Now().UnixNano())
	var name string
	for {
		name = ""
		for i := 0; i < minLength; i++ {
			name += Il[rand.Intn(len(Il))]
		}
		found := false
		for _, set := range sets {
			if set == name {
				found = true
				break
			}
		}
		if !found {
			break
		}
		minLength++
	}
	return name
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func BatchObf() {
	var batchFilePath string
	fmt.Print("Enter the path to the batch file (.bat): ")
	fmt.Scanln(&batchFilePath)

	fileContent, err := ioutil.ReadFile(batchFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	obfuscatedScript := Obfuscate(string(fileContent), alphabet, 8, 5)

	outputDir := "Built/BatchfilePaylods"
	err = ioutil.WriteFile(outputDir+"/CodepulzeObfuscated.bat", []byte(obfuscatedScript), 0644)
	if err != nil {
		fmt.Println("Error writing obfuscated script:", err)
		return
	}

	fmt.Printf("Obfuscated script saved as %s/CodepulzeObfuscated.bat\n", outputDir)
}



func main() {
	gostyle.Init()
	gostyle.ClearConsole()
	gostyle.Write(`

 ____  ____  ____  _____ ____  _     _     ____  _____   _____  ____  ____  _     ____ 
/   _\/  _ \/  _ \/  __//  __\/ \ /\/ \   /_   \/  __/  /__ __\/  _ \/  _ \/ \   / ___\
|  /  | / \|| | \||  \  |  \/|| | ||| |    /   /|  \      / \  | / \|| / \|| |   |    \
|  \__| \_/|| |_/||  /_ |  __/| \_/|| |_/\/   /_|  /_     | |  | \_/|| \_/|| |_/\\___ |
\____/\____/\____/\____\\_/   \____/\____/\____/\____\    \_/  \____/\____/\____/\____/  
   
Author : Evilbytecode & Codepulze
   
   `, gostyle.YELLOW_TO_RED, false)
   gostyle.Write("[1] Powershell to Batchfile", gostyle.YELLOW_TO_RED, false)
   gostyle.Write("[2] Python Payload Crypter", gostyle.YELLOW_TO_RED, false)
   gostyle.Write("[3] Powershell Payload Crypter", gostyle.YELLOW_TO_RED, false)
   gostyle.Write("[4] JavaScript Payload Crypter", gostyle.YELLOW_TO_RED, false)
   gostyle.Write("[5] Batchfile Payload Crypter", gostyle.YELLOW_TO_RED, false)
	gostyle.Write("[@Admin:~/CodepulzeCrypter$] >", gostyle.YELLOW_TO_RED, false)

	var optionStr string
	fmt.Scanln(&optionStr)
	optionStr = strings.TrimSpace(optionStr)

	switch optionStr {
	case "1":
		ConvertPSToBAT()
	case "2":
		CryptPythonPayload()
	case "3":
		CryptPSPayload()
	case "4":
		CryptJSPayload()
	case "5":
		BatchObf() 
	default:
		gostyle.Write("Invalid option selected\n", gostyle.RED_TO_YELLOW, false)
	}
}

func ConvertPSToBAT() {
	reader := bufio.NewReader(os.Stdin)
	gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Enter the full path of the PowerShell script file: ", gostyle.YELLOW_TO_RED, false)
	hehe, _ := reader.ReadString('\n')
	hehe = strings.TrimSpace(hehe)
	hehe = strings.TrimRight(hehe, "\r\n")

	monkicont, _ := os.ReadFile(hehe)

	u16byt := utf16.Encode([]rune(string(monkicont)))
	u16l := make([]byte, len(u16byt)*2)
	for i, r := range u16byt {
		u16l[i*2] = byte(r)
		u16l[i*2+1] = byte(r >> 8)
	}

	monkcoded := base64.StdEncoding.EncodeToString(u16l)

	monkiii := "Built/BatchfilePaylods"
	monkgong := "CodepulzeBuiltPS2BAT.bat"
	opat := filepath.Join(monkiii, monkgong)

	os.MkdirAll(monkiii, 0755)

	monkbat, _ := os.Create(opat)
	defer monkbat.Close()

	monkbat.WriteString("@echo off\n")
	monkbat.WriteString(fmt.Sprintf("powershell.exe -NoExit -encodedCommand %s\nexit", monkcoded))

	gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Converted: %s\n", opat), gostyle.YELLOW_TO_RED, false)
}

func CryptPythonPayload() {
    rand.Seed(time.Now().UnixNano())

    var pythonPath string
    gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Enter the path to the Python script: ", gostyle.YELLOW_TO_RED, false)
    fmt.Scanln(&pythonPath)

    if _, err := os.Stat(pythonPath); os.IsNotExist(err) {
        gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Invalid path.\n", gostyle.RED_TO_YELLOW, false)
        return
    }

    pythonContent, err := ioutil.ReadFile(pythonPath)
    if err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error reading Python script: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    layerInput := getinput("[@Admin:~/CodepulzeCrypter$] > Enter the number of obfuscation layers: ")
    numLayers, err := strconv.Atoi(layerInput)
    if err != nil {
        gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Invalid input. Please enter a valid number.\n", gostyle.RED_TO_YELLOW, false)
        return
    }

    obfuscatedContent := string(pythonContent)
    for i := 1; i <= numLayers; i++ {
        randVar := rndstr(8)
        obfuscatedContent = fmt.Sprintf("import base64\n%s='%s'\n%s=%s.replace('*','')\nexec(base64.b64decode(%s))",
            randVar, base64.StdEncoding.EncodeToString([]byte(obfuscatedContent)),
            randVar, randVar, randVar)
    }

    outputPath := filepath.Join("Built", "PythonPayloads", "CodepulzeObfuscated.py")
    if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error creating output directory: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    if err := ioutil.WriteFile(outputPath, []byte(obfuscatedContent), 0644); err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error writing obfuscated Python script: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Obfuscated script saved as %s\n", outputPath), gostyle.YELLOW_TO_RED, false)
}


func CryptJSPayload() {
    rand.Seed(time.Now().UnixNano())

    var jsPath string
    gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Enter the path to the JavaScript file: ", gostyle.YELLOW_TO_RED, false)
    fmt.Scanln(&jsPath)

    if _, err := os.Stat(jsPath); os.IsNotExist(err) {
        gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Invalid path.\n", gostyle.RED_TO_YELLOW, false)
        return
    }

    jsContent, err := ioutil.ReadFile(jsPath)
    if err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error reading JavaScript file: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    layerInput := getinput("[@Admin:~/CodepulzeCrypter$] > Enter the number of obfuscation layers: ")
    numLayers, err := strconv.Atoi(layerInput)
    if err != nil {
        gostyle.Write("[@Admin:~/CodepulzeCrypter$] > Invalid input. Please enter a valid number.\n", gostyle.RED_TO_YELLOW, false)
        return
    }

    obfuscatedContent := string(jsContent)
    for i := 1; i <= numLayers; i++ {
        obfuscatedContent = fmt.Sprintf("eval(atob('%s'))", base64.StdEncoding.EncodeToString([]byte(obfuscatedContent)))
    }

    outputPath := filepath.Join("Built", "JavaScriptPayloads", "CodepulzeObfuscated.js")
    if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error creating output directory: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    if err := ioutil.WriteFile(outputPath, []byte(obfuscatedContent), 0644); err != nil {
        gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Error writing obfuscated JavaScript file: %s\n", err), gostyle.RED_TO_YELLOW, false)
        return
    }

    gostyle.Write(fmt.Sprintf("[@Admin:~/CodepulzeCrypter$] > Obfuscated script saved as %s\n", outputPath), gostyle.YELLOW_TO_RED, false)
}
