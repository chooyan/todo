package util

import (
        "os"
        "bufio"
        "encoding/json"
)

func Read() string {
        infilepath := GetFilePath()

        var fp *os.File
        var err error

        fp, err = os.Open(infilepath)
        if err != nil {
                panic(nil)
        }
        defer fp.Close()

        reader := bufio.NewReaderSize(fp, 4096)
        var line []byte

        line, _ = reader.ReadBytes('\n')
	return string(line)
}

func ReadObj() ScheduleList {
	jsonstr := Read()
        scheduleObjs := ScheduleList{}
        json.Unmarshal([]byte(jsonstr), &scheduleObjs)
	return scheduleObjs
}

func Write(scheduleJson string) {
        var writer *bufio.Writer

        filePath := GetFilePath()
	os.Remove(filePath)

        writeFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
        writer = bufio.NewWriter(writeFile)
        defer writeFile.Close()

        writer.Write([]byte(scheduleJson))
        writer.Flush()
}
